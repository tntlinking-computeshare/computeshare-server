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

type StorageProvider struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// agent 节点ID
	AgentID uuid.UUID `json:"agent_id,omitempty"`
	// 提供状态： 0：未运行，1：启动中，2： 启动失败，3： 运行中，4： 运行失败
	Status consts.StorageProviderStatus `json:"status,omitempty"`
	// 存储节点master http地址
	MasterServer string `json:"master_server,omitempty"`
	// 存储volume的nat 映射IP
	PublicIP string `json:"public_ip,omitempty"`
	// 存储节点volume的http nat映射端口
	PublicPort int32 `json:"public_port,omitempty"`
	// 存储节点volume的grpc nat映射端口
	GrpcPort int32 `json:"grpc_port,omitempty"`
	// 创建时间
	CreatedTime time.Time `json:"created_time,omitempty"`
}

type StorageProviderRepo interface {
	Create(ctx context.Context, sp *StorageProvider) (*StorageProvider, error)
	List(ctx context.Context) ([]*StorageProvider, error)
	Get(ctx context.Context, id uuid.UUID) (*StorageProvider, error)
	Delete(ctx context.Context, id uuid.UUID) error
	QueryByAgentId(ctx context.Context, id uuid.UUID) (*StorageProvider, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status consts.StorageProviderStatus) error
}

func NewStorageProviderUseCase(
	logger log.Logger,
	storageProviderRepo StorageProviderRepo,
	agentRepo AgentRepo,
	gatewayPortRepo GatewayPortRepo,
	networkMappingRepo NetworkMappingRepo,
	gatewayRepo GatewayRepo,
	taskRepo TaskRepo,
	networkMappingUseCase *NetworkMappingUseCase,
) *StorageProviderUseCase {
	return &StorageProviderUseCase{
		log:                   log.NewHelper(logger),
		storageProviderRepo:   storageProviderRepo,
		agentRepo:             agentRepo,
		gatewayRepo:           gatewayRepo,
		gatewayPortRepo:       gatewayPortRepo,
		networkMappingRepo:    networkMappingRepo,
		networkMappingUseCase: networkMappingUseCase,
		taskRepo:              taskRepo,
	}
}

type StorageProviderUseCase struct {
	log                   *log.Helper
	storageProviderRepo   StorageProviderRepo
	agentRepo             AgentRepo
	gatewayRepo           GatewayRepo
	gatewayPortRepo       GatewayPortRepo
	taskRepo              TaskRepo
	networkMappingRepo    NetworkMappingRepo
	networkMappingUseCase *NetworkMappingUseCase
}

func (c *StorageProviderUseCase) createNetworkMappingPort(ctx context.Context, agentId uuid.UUID, networkMappingName string) (int32, error) {

	claim, ok := global.FromContext(ctx)
	if !ok {
		return 0, errors.New("unauthorize")
	}

	// 查询下一个可用的gateway 端口
	gp, err := c.networkMappingUseCase.GetNextGatewayPort(ctx, agentId)
	if err != nil {
		return 0, err
	}
	gateway, err := c.gatewayRepo.GetGateway(ctx, gp.FkGatewayID)
	if err != nil {
		return 0, err
	}
	// 保存数据库
	// 进行网络映射转换
	nm := NetworkMapping{
		ID:       uuid.UUID{},
		Name:     networkMappingName,
		Protocol: "TCP",
		// gateway id
		FkGatewayID: gp.FkGatewayID,
		// computer_id
		FkComputerID: uuid.Nil,
		// 映射到网关的端口号
		GatewayPort: gp.Port,
		// 需要映射的虚拟机端口号
		ComputerPort: gp.Port,
		//  0 待开始 1 进行中 2 已完成，3 失败
		Status:    0,
		UserId:    claim.GetUserId(),
		GatewayIP: gateway.IP,
	}
	err = c.networkMappingRepo.CreateNetworkMapping(ctx, &nm)
	if err != nil {
		return 0, err
	}
	nptp := &queue.NatNetworkMappingTaskParamVO{
		Id:           nm.ID.String(),
		Name:         nm.Name,
		InstanceId:   "",
		InstancePort: gp.Port,
		RemotePort:   gp.Port,
		GatewayId:    gp.FkGatewayID.String(),
		GatewayIp:    gateway.IP,
		GatewayPort:  gateway.Port,
	}
	paramData, err := json.Marshal(nptp)
	if err != nil {
		return 0, err
	}
	param := string(paramData)
	task := &Task{
		AgentID: agentId.String(),
		Cmd:     queue.TaskCmd_NAT_PROXY_CREATE,
		// 执行参数，nat 网络类型对应 NatProxyCreateVO, 虚拟机类型对应 ComputeInstanceTaskParamVO
		Params: &param,
		//   CREATED = 0; //创建
		Status: queue.TaskStatus_CREATED,
		// CreateTime holds the value of the "create_time" field.
		CreateTime: time.Now(),
	}
	err = c.taskRepo.CreateTask(ctx, task)
	if err != nil {
		return 0, err
	}

	// 标记此gateway 被使用
	gp.IsUse = true
	err = c.gatewayPortRepo.Update(ctx, gp)
	if err != nil {
		return 0, err
	}
	return gp.Port, nil
}

func (c *StorageProviderUseCase) deleteNetworkMappingPort(ctx context.Context, agentId uuid.UUID, ip string, port int32) error {
	_, ok := global.FromContext(ctx)
	if !ok {
		return errors.New("unauthorize")
	}

	nm, err := c.networkMappingRepo.GetNetworkMappingByPublicIpdAndPort(ctx, ip, port)

	if err != nil {
		return err
	}

	// 查询下一个可用的gateway 端口
	gatewayId, err := c.networkMappingRepo.QueryGatewayIdByAgentId(ctx, agentId)
	gateway, err := c.gatewayRepo.GetGateway(ctx, gatewayId)
	gp, err := c.gatewayPortRepo.GetGatewayPortByGatewayIdAndPort(ctx, gatewayId, port)

	nptp := &queue.NatNetworkMappingTaskParamVO{
		Id:           nm.ID.String(),
		Name:         nm.Name,
		InstanceId:   "",
		InstancePort: gp.Port,
		RemotePort:   gp.Port,
		GatewayId:    gp.FkGatewayID.String(),
		GatewayIp:    gateway.IP,
		GatewayPort:  gateway.Port,
	}
	paramData, err := json.Marshal(nptp)
	if err != nil {
		return err
	}
	param := string(paramData)
	task := &Task{
		AgentID: agentId.String(),
		Cmd:     queue.TaskCmd_NAT_PROXY_DELETE,
		// 执行参数，nat 网络类型对应 NatProxyCreateVO, 虚拟机类型对应 ComputeInstanceTaskParamVO
		Params: &param,
		//   CREATED = 0; //创建
		Status: queue.TaskStatus_CREATED,
		// CreateTime holds the value of the "create_time" field.
		CreateTime: time.Now(),
	}
	err = c.taskRepo.CreateTask(ctx, task)
	if err != nil {
		return err
	}

	// 标记此gateway 被使用
	gp.IsUse = false
	err = c.gatewayPortRepo.Update(ctx, gp)
	if err != nil {
		return err
	}

	return c.networkMappingRepo.DeleteNetworkMapping(ctx, nm.ID)

}

func (c *StorageProviderUseCase) CreateStorageProvider(ctx context.Context, agentId uuid.UUID) (*StorageProvider, error) {
	agent, err := c.agentRepo.GetAgent(ctx, agentId)
	if err != nil {
		return nil, err
	}
	// 判断是否已经创建过
	sp, err := c.storageProviderRepo.QueryByAgentId(ctx, agentId)

	if sp != nil {
		return sp, errors.New("StorageProvider Exists")
	}

	//创建http 端口映射
	publicHttpPort, err := c.createNetworkMappingPort(ctx, agentId, fmt.Sprintf("WEED_VOLUME_HTTP_%s", agentId.String()))
	if err != nil {
		return nil, err
	}

	//创建grpc 端口映射
	publicGrpcPort, err := c.createNetworkMappingPort(ctx, agentId, fmt.Sprintf("WEED_VOLUME_GRPC_%s", agentId.String()))
	if err != nil {
		return nil, err
	}
	sp = &StorageProvider{
		AgentID:      agent.ID,
		Status:       consts.StorageProviderStatus_NOT_RUN,
		MasterServer: "computeshare.newtouch.com:9333",
		PublicIP:     "computeshare.newtouch.com",
		PublicPort:   publicHttpPort,
		GrpcPort:     publicGrpcPort,
		CreatedTime:  time.Now(),
	}
	sp, err = c.storageProviderRepo.Create(ctx, sp)

	if err != nil {
		return nil, err
	}

	sstp := &queue.StorageSetupTaskParamVO{
		Id:           sp.ID.String(),
		MasterServer: sp.MasterServer,
		PublicIp:     sp.PublicIP,
		PublicPort:   sp.PublicPort,
		GrpcPort:     sp.GrpcPort,
	}
	paramData, err := json.Marshal(sstp)
	if err != nil {
		return nil, err
	}
	param := string(paramData)

	// 下发Storage 任务
	task := &Task{
		AgentID: agentId.String(),
		Cmd:     queue.TaskCmd_STORAGE_CREATE,
		// 执行参数，nat 网络类型对应 NatProxyCreateVO, 虚拟机类型对应 ComputeInstanceTaskParamVO
		Params: &param,
		//   CREATED = 0; //创建
		Status: queue.TaskStatus_CREATED,
		// CreateTime holds the value of the "create_time" field.
		CreateTime: time.Now(),
	}
	err = c.taskRepo.CreateTask(ctx, task)
	if err != nil {
		return nil, err
	}

	return sp, err
}

func (c *StorageProviderUseCase) DeleteStorageProvider(ctx context.Context, id uuid.UUID) error {
	sp, err := c.GetStorageProvider(ctx, id)
	if err != nil {
		return err
	}

	err = c.deleteNetworkMappingPort(ctx, sp.AgentID, sp.PublicIP, sp.PublicPort)
	if err != nil {
		return err
	}

	err = c.deleteNetworkMappingPort(ctx, sp.AgentID, sp.PublicIP, sp.GrpcPort)
	if err != nil {
		return err
	}
	sstp := &queue.StorageSetupTaskParamVO{
		Id:           sp.ID.String(),
		MasterServer: sp.MasterServer,
		PublicIp:     sp.PublicIP,
		PublicPort:   sp.PublicPort,
		GrpcPort:     sp.GrpcPort,
	}
	paramData, err := json.Marshal(sstp)
	if err != nil {
		return err
	}
	param := string(paramData)

	// 下发Storage 任务
	task := &Task{
		AgentID: sp.AgentID.String(),
		Cmd:     queue.TaskCmd_STORAGE_DELETE,
		// 执行参数，nat 网络类型对应 NatProxyCreateVO, 虚拟机类型对应 ComputeInstanceTaskParamVO
		Params: &param,
		//   CREATED = 0; //创建
		Status: queue.TaskStatus_CREATED,
		// CreateTime holds the value of the "create_time" field.
		CreateTime: time.Now(),
	}
	err = c.taskRepo.CreateTask(ctx, task)
	if err != nil {
		return err
	}

	return c.storageProviderRepo.Delete(ctx, id)
}

func (c *StorageProviderUseCase) GetStorageProvider(ctx context.Context, id uuid.UUID) (*StorageProvider, error) {
	return c.storageProviderRepo.Get(ctx, id)
}

func (c *StorageProviderUseCase) ListStorageProvider(ctx context.Context) ([]*StorageProvider, error) {

	return c.storageProviderRepo.List(ctx)
}

func (c *StorageProviderUseCase) UpdateStorageProviderStatus(ctx context.Context, id uuid.UUID, status consts.StorageProviderStatus) error {
	return c.storageProviderRepo.UpdateStatus(ctx, id, status)
}
