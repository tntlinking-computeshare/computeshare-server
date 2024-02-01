package biz

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	queue "github.com/mohaijiang/computeshare-server/api/queue/v1"
	"github.com/mohaijiang/computeshare-server/internal/global"
	"github.com/mohaijiang/computeshare-server/internal/global/consts"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

type ComputeSpecRepo interface {
	List(ctx context.Context) ([]*ComputeSpec, error)
	Get(ctx context.Context, id int32) (*ComputeSpec, error)
	GetSpecPrice(ctx context.Context, id int32) (*ComputeSpecPrice, error)
}

type ComputeInstanceRepo interface {
	List(ctx context.Context, owner string) ([]*ComputeInstance, error)
	ListByStatus(ctx context.Context, owner string, status consts.InstanceStatus) ([]*ComputeInstance, error)
	ListByAgentId(ctx context.Context, agentId string) ([]*ComputeInstance, error)
	ListAll(ctx context.Context) ([]*ComputeInstance, error)
	Create(ctx context.Context, instance *ComputeInstance) error
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, id uuid.UUID, instance *ComputeInstance) error
	UpdateStatus(ctx context.Context, id uuid.UUID, status consts.InstanceStatus) error
	SetInstanceExpiration(ctx context.Context) error
	Get(ctx context.Context, id uuid.UUID) (*ComputeInstance, error)
	SaveInstanceStats(context.Context, uuid.UUID, []*ComputeInstanceRds) error
	GetInstanceStats(ctx context.Context, id uuid.UUID) ([]*ComputeInstanceRds, error)
	ListExpiration(ctx context.Context) ([]*ComputeInstance, error)
	IfNeedSyncInstanceStats(ctx context.Context, id uuid.UUID) bool
	ListByOrderDue3Day(ctx context.Context) []*ComputeInstance
}

type ComputeImageRepo interface {
	List(ctx context.Context) ([]*ComputeImage, error)
	Get(ctx context.Context, id int32) (*ComputeImage, error)
}

type ComputeInstanceUsercase struct {
	specRepo              ComputeSpecRepo
	instanceRepo          ComputeInstanceRepo
	imageRepo             ComputeImageRepo
	agentRepo             AgentRepo
	taskRepo              TaskRepo
	gatewayRepo           GatewayRepo
	gatewayPortRepo       GatewayPortRepo
	networkMappingRepo    NetworkMappingRepo
	cycleRepo             CycleRepo
	cycleOrderRepo        CycleOrderRepo
	cycleTransactionRepo  CycleTransactionRepo
	cycleRenewalRepo      CycleRenewalRepo
	networkMappingUseCase *NetworkMappingUseCase
	smsUseCase            *SmsUseCase
	log                   *log.Helper
}

func NewComputeInstanceUsercase(
	specRepo ComputeSpecRepo,
	instanceRepo ComputeInstanceRepo,
	imageRepo ComputeImageRepo,
	agentRepo AgentRepo,
	taskRepo TaskRepo,
	gatewayRepo GatewayRepo,
	gatewayPortRepo GatewayPortRepo,
	networkMappingRepo NetworkMappingRepo,
	cycleRepo CycleRepo,
	cycleOrderRepo CycleOrderRepo,
	cycleTransactionRepo CycleTransactionRepo,
	cycleRenewalRepo CycleRenewalRepo,
	networkMappingUseCase *NetworkMappingUseCase,
	smsUseCase *SmsUseCase,
	logger log.Logger) *ComputeInstanceUsercase {
	return &ComputeInstanceUsercase{
		specRepo:              specRepo,
		instanceRepo:          instanceRepo,
		imageRepo:             imageRepo,
		agentRepo:             agentRepo,
		taskRepo:              taskRepo,
		gatewayRepo:           gatewayRepo,
		gatewayPortRepo:       gatewayPortRepo,
		networkMappingRepo:    networkMappingRepo,
		cycleRepo:             cycleRepo,
		cycleOrderRepo:        cycleOrderRepo,
		cycleTransactionRepo:  cycleTransactionRepo,
		cycleRenewalRepo:      cycleRenewalRepo,
		networkMappingUseCase: networkMappingUseCase,
		smsUseCase:            smsUseCase,
		log:                   log.NewHelper(logger),
	}
}

func (uc *ComputeInstanceUsercase) ListComputeSpec(ctx context.Context) ([]*ComputeSpec, error) {
	return uc.specRepo.List(ctx)
}

func (uc *ComputeInstanceUsercase) GetComputeSpecPrice(ctx context.Context, specId int32) (*ComputeSpecPrice, error) {
	return uc.specRepo.GetSpecPrice(ctx, specId)
}

func (uc *ComputeInstanceUsercase) ListComputeImage(ctx context.Context) ([]*ComputeImage, error) {
	return uc.imageRepo.List(ctx)
}

func (uc *ComputeInstanceUsercase) Create(ctx context.Context, cic *ComputeInstanceCreate) (*ComputeInstance, error) {

	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New(400, "unauthorized", "无权限")
	}

	userId := claim.GetUserId()

	computeSpec, err := uc.specRepo.Get(ctx, cic.SpecId)
	if err != nil {
		return nil, err
	}

	computeImage, err := uc.imageRepo.Get(ctx, cic.ImageId)
	if err != nil {
		return nil, err
	}

	// 选择一个agent节点进行通信
	agent, err := uc.agentRepo.FindOneActiveAgent(ctx, computeSpec.Core, computeSpec.Memory)
	if err != nil {
		return nil, err
	}

	// 查询价格
	specPrice, err := uc.specRepo.GetSpecPrice(ctx, cic.SpecId)
	if err != nil {
		return nil, err
	}

	gatewayId, err := uc.networkMappingRepo.QueryGatewayIdByAgentId(ctx, agent.ID)
	if err != nil {
		return nil, err
	}
	gw, err := uc.gatewayRepo.GetGateway(ctx, gatewayId)
	if err != nil {
		return nil, err
	}
	gp, err := uc.gatewayPortRepo.GetGatewayPortFirstByNotUsedAndIsPublic(ctx, gatewayId, false)
	if err != nil {
		return nil, err
	}

	var dockerComposeDecode string
	if cic.DockerCompose != "" {
		data, err := base64.StdEncoding.DecodeString(cic.DockerCompose)
		if err != nil {
			dockerComposeDecode = string(data)
		}
	}
	instance := &ComputeInstance{
		Owner:          claim.UserID,
		Name:           cic.Name,
		Core:           computeSpec.Core,
		Memory:         computeSpec.Memory,
		Port:           fmt.Sprintf("%d", computeImage.Port),
		Image:          fmt.Sprintf("%s:%s", computeImage.Image, computeImage.Tag),
		ImageId:        computeImage.ID,
		ExpirationTime: time.Now().AddDate(0, 0, int(specPrice.Day)),
		AgentId:        agent.ID.String(),
		Status:         consts.InstanceStatusCreating,
		VncIP:          gw.InternalIP,
		VncPort:        gp.Port,
		DockerCompose:  dockerComposeDecode,
		CreateTime:     time.Now(),
	}

	err = uc.instanceRepo.Create(ctx, instance)
	if err != nil {
		return nil, err
	}

	// 创建续费管理
	renewalTime := instance.ExpirationTime.AddDate(0, 0, -9)
	renewalTime = time.Date(renewalTime.Year(), renewalTime.Month(), renewalTime.Day(), 23, 0, 0, 0, renewalTime.Location())

	fmt.Println("=============")
	fmt.Println("instanceId:", instance.ID.String())
	fmt.Println("=============")

	cycle, err := uc.cycleRepo.FindByUserID(ctx, userId)

	if err != nil || cycle.Cycle.LessThan(decimal.NewFromFloat32(specPrice.Price)) {
		_ = uc.smsUseCase.InsufficientBalance(instance.Name, decimal.NewFromFloat32(specPrice.Price), userId)
		return nil, errors.New(400, "insufficient balance", "Cycle不足，请先充值再试！")
	}

	// 扣余额
	cycle.Cycle = cycle.Cycle.Sub(decimal.NewFromFloat32(specPrice.Price))
	cycle.UpdateTime = time.Now()
	err = uc.cycleRepo.Update(ctx, cycle)
	if err != nil {
		return nil, err
	}

	// 创建订单
	var orderNo string
	// 判断订单号不重复
	var exists bool
	for {
		orderNo = NewOrderNo()
		exists = uc.cycleOrderRepo.CheckOrderNoExists(ctx, orderNo)
		if !exists {
			break
		}
	}

	resourceId := instance.ID.String()

	cycleOrder := &CycleOrder{
		OrderNo:     orderNo,
		FkUserID:    userId,
		ProductName: string(consts.RentingCloudServers),
		ProductDesc: fmt.Sprintf("%s | %d核%dGB | %s | %d天", cic.Name, computeSpec.Core, computeSpec.Memory, computeImage.GetImageTag(), specPrice.Day),
		Symbol:      "-",
		Cycle:       float64(specPrice.Price),
		ResourceId:  &resourceId,
		CreateTime:  time.Now(),
	}
	cycleOrder, err = uc.cycleOrderRepo.Create(ctx, cycleOrder)
	if err != nil {
		return nil, err
	}

	// 创建交易流水
	balance, _ := cycle.Cycle.Float64()
	cycleTransaction := &CycleTransaction{
		FkCycleID:         cycle.ID,
		FkUserID:          userId,
		FkCycleOrderID:    cycleOrder.ID,
		FkCycleRechargeID: uuid.Nil,
		Operation:         cycleOrder.ProductName,
		Symbol:            cycleOrder.Symbol,
		Cycle:             cycleOrder.Cycle,
		Balance:           balance,
		OperationTime:     time.Now(),
	}

	cycleTransaction, err = uc.cycleTransactionRepo.Create(ctx, cycleTransaction)
	if err != nil {
		return nil, err
	}

	renewal := &CycleRenewal{
		FkUserID:     userId,
		ResourceID:   instance.ID,
		ResourceType: int(consts.RenewalResourceType_Resource),
		ProductName:  cycleOrder.ProductName,
		ProductDesc:  cycleOrder.ProductDesc,
		State:        int8(consts.RenewalState_IN_SERVICE),
		ExtendDay:    int8(specPrice.Day),
		ExtendPrice:  float64(specPrice.Price),
		DueTime:      &instance.ExpirationTime,
		RenewalTime:  &renewalTime,
		AutoRenewal:  true,
	}
	renewal, err = uc.cycleRenewalRepo.Create(ctx, renewal)
	if err != nil {
		return nil, err
	}

	err = uc.SendTaskQueue(ctx, instance, queue.TaskCmd_VM_CREATE, func() InstanceCreateParam {
		return InstanceCreateParam{
			PublicKey:      cic.PublicKey,
			Password:       cic.Password,
			GatewayIP:      gw.IP,
			GatewayPort:    gw.Port,
			VncConnectIP:   gw.InternalIP,
			VncConnectPort: gp.Port,
			DockerCompose:  cic.DockerCompose,
		}
	})
	if err != nil {
		return nil, err
	}

	gp.IsUse = true

	err = uc.gatewayPortRepo.Update(ctx, gp)
	if err != nil {
		return nil, err
	}

	// 发送扣款短信
	_ = uc.smsUseCase.ChargingSuccess(instance.Name, decimal.NewFromFloat32(specPrice.Price), userId)

	return instance, err
}

func (uc *ComputeInstanceUsercase) SendTaskQueue(ctx context.Context, instance *ComputeInstance, cmd queue.TaskCmd, publicKeyAndPassword func() InstanceCreateParam) error {

	taskParam := queue.ComputeInstanceTaskParamVO{
		Id:         instance.ID.String(),
		Name:       instance.Name,
		Cpu:        instance.GetCore(),
		Memory:     instance.GetMemory(),
		Image:      instance.Image,
		InstanceId: instance.ID.String(),
	}
	if publicKeyAndPassword != nil {
		instanceCreateParam := publicKeyAndPassword()
		taskParam.PublicKey = instanceCreateParam.PublicKey
		taskParam.Password = instanceCreateParam.Password
		taskParam.GatewayIp = instanceCreateParam.GatewayIP
		taskParam.GatewayPort = instanceCreateParam.GatewayPort
		taskParam.VncConnectIp = instanceCreateParam.VncConnectIP
		taskParam.VncConnectPort = instanceCreateParam.VncConnectPort
		taskParam.DockerCompose = instanceCreateParam.DockerCompose
	}
	paramData, err := json.Marshal(&taskParam)
	if err != nil {
		return err
	}
	param := string(paramData)
	task := Task{
		AgentID:    instance.AgentId,
		Cmd:        cmd,
		Params:     &param,
		Status:     queue.TaskStatus_CREATED,
		CreateTime: time.Now(),
	}
	err = uc.taskRepo.CreateTask(ctx, &task)
	return err
}

func (uc *ComputeInstanceUsercase) Delete(ctx context.Context, id uuid.UUID) error {
	instance, err := uc.Get(ctx, id)
	if err != nil {
		return err
	}

	mappings, _ := uc.networkMappingUseCase.ListByComputeInstanceId(ctx, id)
	if len(mappings) > 0 {
		for _, mapping := range mappings {
			_ = uc.networkMappingUseCase.DeleteNetworkMapping(ctx, mapping.ID)
		}
	}

	err = uc.SendTaskQueue(ctx, instance, queue.TaskCmd_VM_DELETE, nil)
	if err != nil {
		return err
	}

	return uc.instanceRepo.UpdateStatus(ctx, instance.ID, consts.InstanceStatusDeleting)
}

func (uc *ComputeInstanceUsercase) ListComputeInstance(ctx context.Context, owner string) ([]*ComputeInstance, error) {
	list, err := uc.instanceRepo.List(ctx, owner)
	for _, ins := range list {
		ins.Stats, _ = uc.GetInstanceStats(ctx, ins.ID)
	}
	return list, err
}

func (uc *ComputeInstanceUsercase) ListComputeInstanceByStatus(ctx context.Context, owner string, status int32) ([]*ComputeInstance, error) {
	list, err := uc.instanceRepo.ListByStatus(ctx, owner, consts.InstanceStatus(status))
	return list, err
}

func (uc *ComputeInstanceUsercase) Get(ctx context.Context, id uuid.UUID) (*ComputeInstance, error) {
	return uc.instanceRepo.Get(ctx, id)
}

func (uc *ComputeInstanceUsercase) Start(ctx context.Context, id uuid.UUID) error {
	instance, err := uc.Get(ctx, id)
	if err != nil {
		return err
	}

	instance.Status = consts.InstanceStatusStarting

	err = uc.instanceRepo.Update(ctx, instance.ID, instance)

	if err != nil {
		uc.log.Error("创建容器部署指令失败")
		uc.log.Error(err)
		return err
	}

	err = uc.SendTaskQueue(ctx, instance, queue.TaskCmd_VM_START, nil)
	if err != nil {
		return err
	}

	return nil

}

func (uc *ComputeInstanceUsercase) Stop(ctx context.Context, id uuid.UUID) error {
	instance, err := uc.Get(ctx, id)
	if err != nil {
		return err
	}

	instance.Status = consts.InstanceStatusClosing

	err = uc.instanceRepo.Update(ctx, instance.ID, instance)

	if err != nil {
		uc.log.Error("创建容器部署指令失败")
		uc.log.Error(err)
		return err
	}

	err = uc.SendTaskQueue(ctx, instance, queue.TaskCmd_VM_SHUTDOWN, nil)
	if err != nil {
		return err
	}

	return nil
}

// Terminal Deprecate
func (uc *ComputeInstanceUsercase) GetVncConsole(ctx context.Context, instanceId uuid.UUID, userId string) (string, error) {
	instance, err := uc.Get(ctx, instanceId)
	if err != nil {
		return "", err
	}
	if instance.Owner != userId {
		return "", errors.New(400, "unauthorized", "无权限")
	}
	return fmt.Sprintf("ws://%s:%d/websockify", instance.VncIP, instance.VncPort), err
}

// SyncContainerOverdue 同步资源实例的过期状态
func (uc *ComputeInstanceUsercase) SyncContainerOverdue() {
	ctx := context.Background()
	uc.log.Info("查询过期实例")
	expirationList, err := uc.instanceRepo.ListExpiration(ctx)
	if err != nil {
		fmt.Println("查询过期实例失败")
		return
	}

	for _, instance := range expirationList {
		if instance.Status == consts.InstanceStatusRunning {
			// 停止实例
			_ = uc.Stop(ctx, instance.ID)
		}

		err := uc.instanceRepo.UpdateStatus(ctx, instance.ID, consts.InstanceStatusExpire)
		if err != nil {
			uc.log.Error("更新实例状态失败： ", err)
			break
		}

		cycleRenew, err := uc.cycleRenewalRepo.QueryByResourceId(ctx, instance.ID)
		if err != nil {
			uc.log.Error("更新实例状态失败： ", err)
			return
		}
		cycleRenew.State = int8(consts.RenewalState_STOP)
		err = uc.cycleRenewalRepo.Update(ctx, cycleRenew.ID, cycleRenew)
		if err != nil {
			uc.log.Error("更新实例状态失败： ", err)
			return
		}
	}
}

func (uc *ComputeInstanceUsercase) Reboot(ctx context.Context, instanceId uuid.UUID) error {
	instance, err := uc.Get(ctx, instanceId)
	if err != nil {
		return err
	}

	instance.Status = consts.InstanceStatusRestarting

	err = uc.instanceRepo.Update(ctx, instance.ID, instance)

	if err != nil {
		uc.log.Error("创建容器重启指令失败")
		uc.log.Error(err)
		return err
	}

	err = uc.SendTaskQueue(ctx, instance, queue.TaskCmd_VM_RESTART, nil)
	if err != nil {
		return err
	}

	return nil
}

func (uc *ComputeInstanceUsercase) Recreate(ctx context.Context, instanceId uuid.UUID, param *ComputeInstanceCreate) error {
	instance, err := uc.Get(ctx, instanceId)
	if err != nil {
		return err
	}

	computeImage, err := uc.imageRepo.Get(ctx, param.ImageId)
	if err != nil {
		return err
	}
	instance.Image = fmt.Sprintf("%s:%s", computeImage.Image, computeImage.Tag)
	instance.DockerCompose = param.DockerCompose

	err = uc.instanceRepo.Update(ctx, instance.ID, instance)

	if err != nil {
		uc.log.Error("重建虚拟机指令失败")
		uc.log.Error(err)
		return err
	}

	gateways, err := uc.gatewayRepo.ListGateway(ctx)

	if err != nil {
		uc.log.Error("重建虚拟机指令失败")
		uc.log.Error(err)
		return err
	}

	var g *Gateway
	for _, ga := range gateways {
		if ga.InternalIP == instance.VncIP {
			g = ga
		}
	}
	if g == nil {
		uc.log.Error("重建虚拟机指令失败")
		uc.log.Error(err)
		return errors.New(400, "RECREATE_FAIL", "Gateway cannot found")
	}

	err = uc.SendTaskQueue(ctx, instance, queue.TaskCmd_VM_RECREATE, func() InstanceCreateParam {
		return InstanceCreateParam{
			DockerCompose:  param.DockerCompose,
			Password:       param.Password,
			PublicKey:      param.PublicKey,
			GatewayIP:      g.IP,
			GatewayPort:    g.Port,
			VncConnectIP:   instance.VncIP,
			VncConnectPort: instance.VncPort,
		}
	})
	if err != nil {
		return err
	}

	return nil
}

func (uc *ComputeInstanceUsercase) GetInstanceStats(ctx context.Context, instanceId uuid.UUID) ([]*ComputeInstanceRds, error) {
	go func() {
		ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
		need := uc.instanceRepo.IfNeedSyncInstanceStats(ctx, instanceId)
		if need {
			data, err := uc.GetLast24HInstanceStats(ctx, instanceId.String())
			if err != nil {
				uc.log.Debug("GetLast24HInstanceStats Err: ", err.Error())
				return
			}
			err = uc.instanceRepo.SaveInstanceStats(ctx, instanceId, data)
			uc.log.Error(err)
		}
	}()
	return uc.instanceRepo.GetInstanceStats(ctx, instanceId)
}

func (uc *ComputeInstanceUsercase) GetLast24HInstanceStats(_ context.Context, instanceId string) ([]*ComputeInstanceRds, error) {
	step := 600
	to := time.Now().Round(10 * time.Minute)
	from := to.AddDate(0, 0, -1)
	cpuData, err := uc.PrometheusQuery(getCpuUsageQuery(instanceId), step, from.Unix(), to.Unix())
	if err != nil {
		return nil, err
	}

	memoryData, err := uc.PrometheusQuery(getMemoryUsageQuery(instanceId), step, from.Unix(), to.Unix())
	if err != nil {
		return nil, err
	}

	m := make(map[float64]*ComputeInstanceRds)

	if len(cpuData.Data.Result) == 0 || len(memoryData.Data.Result) == 0 {
		return []*ComputeInstanceRds{}, nil
	}

	for _, v := range cpuData.Data.Result[0].Values {
		timestamp := v[0].(float64)
		statsTime := time.UnixMilli(int64(timestamp * 1000))
		value := v[1].(string)
		cpuUsage, err := strconv.ParseFloat(value, 32)
		if value == "NaN" || err != nil {
			cpuUsage = 0
		}
		m[timestamp] = &ComputeInstanceRds{
			ID:          uuid.New().String(),
			CpuUsage:    float32(cpuUsage),
			MemoryUsage: 0,
			StatsTime:   statsTime,
		}
	}

	for _, v := range memoryData.Data.Result[0].Values {
		timestamp := v[0].(float64)
		value := v[1].(string)
		memoryUsage, err := strconv.ParseFloat(value, 32)
		if value == "NaN" || err != nil {
			memoryUsage = 0
		}
		if cis, ok := m[timestamp]; ok {
			cis.MemoryUsage = float32(memoryUsage)
		} else {
			m[timestamp] = &ComputeInstanceRds{
				ID:          uuid.New().String(),
				CpuUsage:    0,
				MemoryUsage: float32(memoryUsage),
				StatsTime:   time.UnixMilli(int64(timestamp * 1000)),
			}
		}

	}

	result := lo.Values(m)
	sort.Slice(result, func(i, j int) bool {
		return result[i].StatsTime.Before(result[j].StatsTime)
	})
	return result, err
}

func getCpuUsageQuery(instanceId string) string {
	return fmt.Sprintf("(sum by(instance) (irate(node_cpu_seconds_total{instance=\"%s\",job=\"node\", mode!=\"idle\"}[10m15s])) / on(instance) group_left sum by (instance)((irate(node_cpu_seconds_total{instance=\"%s\",job=\"node\"}[10m15s])))) * 100", instanceId, instanceId)
}
func getMemoryUsageQuery(instanceId string) string {
	return fmt.Sprintf("100 - ((avg_over_time(node_memory_MemAvailable_bytes{instance=\"%s\",job=\"node\"}[10m15s]) * 100) / avg_over_time(node_memory_MemTotal_bytes{instance=\"%s\",job=\"node\"}[10m15s]))", instanceId, instanceId)
}

func (uc *ComputeInstanceUsercase) PrometheusQuery(expr string, step int, from, to int64) (*PrometheusQueryResult, error) {
	// Prometheus 查询 API 地址
	prometheusURL := "http://61.172.179.73:9090/api/v1/query_range"

	// PromQL 查询语句
	// cpu (sum by(instance) (irate(node_cpu_seconds_total{instance="76fe0a88-1960-4966-9beb-41b1c1251595",job="node", mode!="idle"}[10m15s])) / on(instance) group_left sum by (instance)((irate(node_cpu_seconds_total{instance="76fe0a88-1960-4966-9beb-41b1c1251595",job="node"}[10m15s])))) * 100
	// memory  100 - ((avg_over_time(node_memory_MemAvailable_bytes{instance="76fe0a88-1960-4966-9beb-41b1c1251595",job="node"}[10m15s]) * 100) / avg_over_time(node_memory_MemTotal_bytes{instance="76fe0a88-1960-4966-9beb-41b1c1251595",job="node"}[10m15s]))

	// 构建查询参数
	params := fmt.Sprintf("query=%s", url.QueryEscape(expr))
	queryURL := fmt.Sprintf("%s?%s&start=%d&end=%d&step=%d", prometheusURL, params, from, to, step)

	//queryURL = fmt.Sprintf("%s?%s", prometheusURL, "query=%28sum%28increase%28node_cpu_seconds_total%7Bmode%3D%27system%27%2Cinstance%3D%2276fe0a88-1960-4966-9beb-41b1c1251595%22%7D%5B10m%5D%29%29by%28instance%29%29+%2F+%28sum%28increase%28node_cpu_seconds_total%5B10m%5D%29%29by%28instance%29%29++*100%0A&start=1704437735.053&end=1704441335.053&step=14")
	// 发起 GET 请求
	response, err := http.Get(queryURL)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return nil, err
	}
	defer response.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	var queryResult PrometheusQueryResult
	if err := json.Unmarshal(body, &queryResult); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}
	return &queryResult, err
}

func (uc *ComputeInstanceUsercase) Rename(ctx context.Context, id uuid.UUID, name string) error {
	instance, err := uc.Get(ctx, id)
	if err != nil {
		return err
	}

	instance.Name = name

	return uc.instanceRepo.Update(ctx, id, instance)
}

func (uc *ComputeInstanceUsercase) NotificationOverDue(ctx context.Context) {
	items := uc.instanceRepo.ListByOrderDue3Day(ctx)

	for _, item := range items {
		userId, err := uuid.Parse(item.Owner)
		if err != nil {
			continue
		}
		_ = uc.smsUseCase.ResourceBecomeDue(item.Name, userId)
	}
}
