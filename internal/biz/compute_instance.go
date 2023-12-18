package biz

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	queue "github.com/mohaijiang/computeshare-server/api/queue/v1"
	"github.com/mohaijiang/computeshare-server/internal/global"
	"github.com/mohaijiang/computeshare-server/internal/global/consts"
	"time"
)

type ComputeSpecRepo interface {
	List(ctx context.Context) ([]*ComputeSpec, error)
	Get(ctx context.Context, id int32) (*ComputeSpec, error)
}

type ComputeInstanceRepo interface {
	List(ctx context.Context, owner string) ([]*ComputeInstance, error)
	ListByPeerId(ctx context.Context, peerId string) ([]*ComputeInstance, error)
	ListAll(ctx context.Context) ([]*ComputeInstance, error)
	Create(ctx context.Context, instance *ComputeInstance) error
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, id uuid.UUID, instance *ComputeInstance) error
	UpdateStatus(ctx context.Context, id uuid.UUID, status consts.InstanceStatus) error
	SetInstanceExpiration(ctx context.Context) error
	Get(ctx context.Context, id uuid.UUID) (*ComputeInstance, error)
	SaveInstanceStats(ctx context.Context, id uuid.UUID, rdbInstance *ComputeInstanceRds) error
	GetInstanceStats(ctx context.Context, id uuid.UUID) ([]*ComputeInstanceRds, error)
}

type ComputeImageRepo interface {
	List(ctx context.Context) ([]*ComputeImage, error)
	Get(ctx context.Context, id int32) (*ComputeImage, error)
}

type ComputeInstanceUsercase struct {
	specRepo           ComputeSpecRepo
	instanceRepo       ComputeInstanceRepo
	imageRepo          ComputeImageRepo
	agentRepo          AgentRepo
	taskRepo           TaskRepo
	gatewayRepo        GatewayRepo
	gatewayPortRepo    GatewayPortRepo
	networkMappingRepo NetworkMappingRepo
	p2pClient          *P2pClient
	log                *log.Helper
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
	logger log.Logger) *ComputeInstanceUsercase {
	return &ComputeInstanceUsercase{
		specRepo:           specRepo,
		instanceRepo:       instanceRepo,
		imageRepo:          imageRepo,
		agentRepo:          agentRepo,
		taskRepo:           taskRepo,
		gatewayRepo:        gatewayRepo,
		gatewayPortRepo:    gatewayPortRepo,
		networkMappingRepo: networkMappingRepo,
		log:                log.NewHelper(logger),
	}
}

func (uc *ComputeInstanceUsercase) ListComputeSpec(ctx context.Context) ([]*ComputeSpec, error) {
	return uc.specRepo.List(ctx)
}

func (uc *ComputeInstanceUsercase) ListComputeImage(ctx context.Context) ([]*ComputeImage, error) {
	return uc.imageRepo.List(ctx)
}

func (uc *ComputeInstanceUsercase) Create(ctx context.Context, cic *ComputeInstanceCreate) (*ComputeInstance, error) {

	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("cannot get user ID")
	}

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

	instance := &ComputeInstance{
		Owner:          claim.UserID,
		Name:           cic.Name,
		Core:           computeSpec.Core,
		Memory:         computeSpec.Memory,
		Port:           fmt.Sprintf("%d", computeImage.Port),
		Image:          fmt.Sprintf("%s:%s", computeImage.Image, computeImage.Tag),
		ExpirationTime: time.Now().AddDate(0, int(cic.Duration), 0),
		AgentId:        agent.ID.String(),
		Status:         consts.InstanceStatusCreating,
		VncIP:          gw.InternalIP,
		VncPort:        gp.Port,
	}

	err = uc.instanceRepo.Create(ctx, instance)
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

	err = uc.SendTaskQueue(ctx, instance, queue.TaskCmd_VM_DELETE, nil)
	if err != nil {
		return err
	}

	return uc.instanceRepo.UpdateStatus(ctx, instance.ID, consts.InstanceStatusDeleting)
}

func (uc *ComputeInstanceUsercase) ListComputeInstance(ctx context.Context, owner string) ([]*ComputeInstance, error) {
	list, err := uc.instanceRepo.List(ctx, owner)
	for _, ins := range list {
		ins.Stats, _ = uc.instanceRepo.GetInstanceStats(ctx, ins.ID)
	}
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
func (uc *ComputeInstanceUsercase) GetVncConsole(ctx context.Context, instanceId uuid.UUID) (string, error) {
	instance, err := uc.Get(ctx, instanceId)
	if err != nil {
		return "", err
	}

	// 查询实例需要链接到的gateway
	// 1). 判断此实例有无配置端口映射，如果配置，直接使用此端口映射对应的gatewayId
	// 2). 如果此实例未进行端口映射， 选择一个gateway进行端口映射
	gateway, err := uc.gatewayRepo.FindInstanceSuitableGateway(ctx, instanceId)
	if err != nil {
		return "", err
	}

	gp, err := uc.gatewayPortRepo.GetGatewayPortFirstByNotUsed(ctx, gateway.ID)
	taskParam := queue.NatNetworkMappingTaskParamVO{
		Id:           uuid.NewString(),
		Name:         instance.Name,
		InstanceId:   instance.ID.String(),
		InstancePort: 0,
		RemotePort:   gp.Port,
		GatewayId:    gateway.ID.String(),
		GatewayIp:    gateway.IP,
		GatewayPort:  gateway.Port,
	}
	paramData, err := json.Marshal(taskParam)
	if err != nil {
		return "", err
	}
	param := string(paramData)
	task := Task{
		AgentID:    instance.AgentId,
		Cmd:        queue.TaskCmd_VM_VNC_CONNECT,
		Params:     &param,
		Status:     queue.TaskStatus_CREATED,
		CreateTime: time.Now(),
	}
	err = uc.taskRepo.CreateTask(ctx, &task)

	gp.IsUse = true
	err = uc.gatewayPortRepo.Update(ctx, gp)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("http://%s:%d/vnc.html", gateway.IP, gp.Port), nil

}

// SyncContainerOverdue 同步资源实例的过期状态
func (uc *ComputeInstanceUsercase) SyncContainerOverdue() {
	ctx := context.Background()
	_ = uc.instanceRepo.SetInstanceExpiration(ctx)
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
