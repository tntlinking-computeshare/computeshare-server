package biz

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/samber/lo"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	queue "github.com/mohaijiang/computeshare-server/api/queue/v1"
)

type NetworkMapping struct {
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// gateway id
	FkGatewayID uuid.UUID `json:"fk_gateway_id,omitempty"`
	// computer_id
	FkComputerID uuid.UUID `json:"fk_computer_id,omitempty"`
	// 映射到网关的端口号
	GatewayPort int `json:"gateway_port,omitempty"`
	// 需要映射的虚拟机端口号
	ComputerPort int `json:"computer_port,omitempty"`
	//  0 待开始 1 进行中 2 已完成，3 失败
	Status int `json:"status,omitempty"`
}

type NetworkMappingCreate struct {
	Name         string
	ComputerId   uuid.UUID
	ComputerPort int32
}

type NetworkMappingRepo interface {
	CreateNetworkMapping(ctx context.Context, entity *NetworkMapping) error
	GetNetworkMapping(ctx context.Context, id uuid.UUID) (*NetworkMapping, error)
	DeleteNetworkMapping(ctx context.Context, id uuid.UUID) error
	PageNetworkMappingByComputerID(ctx context.Context, computerId uuid.UUID, page int32, size int32) ([]*NetworkMapping, int32, error)
	UpdateNetworkMapping(ctx context.Context, entity *NetworkMapping) error
}

type Gateway struct {
	ID   uuid.UUID
	Name string
	IP   string
	Port int
}

type GatewayRepo interface {
	ListGateway(ctx context.Context) ([]*Gateway, error)
	GetGateway(ctx context.Context, id uuid.UUID) (*Gateway, error)
	// FindInstanceSuitableGateway 查询实例需要链接到的gateway
	// 1). 判断此实例有无配置端口映射，如果配置，直接使用此端口映射对应的gatewayId
	// 2). 如果此实例未进行端口映射， 选择一个gateway进行端口映射
	FindInstanceSuitableGateway(ctx context.Context, instanceId uuid.UUID) (*Gateway, error)
}

type GatewayPort struct {
	ID          uuid.UUID
	FkGatewayID uuid.UUID
	Port        int64
	IsUse       bool
}

type GatewayPortCount struct {
	FkGatewayID uuid.UUID `json:"fk_gateway_id" db:"fk_gateway_id"`
	Count       int32     `json:"count" db:"count"`
}

type GatewayPortRepo interface {
	CountGatewayPortByIsUsed(ctx context.Context, isUsed bool) ([]*GatewayPortCount, error)
	GetGatewayPortFirstByNotUsed(ctx context.Context, gatewayID uuid.UUID) (*GatewayPort, error)
	Update(ctx context.Context, gp *GatewayPort) error
	GetGatewayPortByGatewayIdAndPort(ctx context.Context, id uuid.UUID, port int) (*GatewayPort, error)
}

type NetworkMappingUseCase struct {
	repo            NetworkMappingRepo
	gatewayRepo     GatewayRepo
	gatewayPortRepo GatewayPortRepo
	taskRepo        TaskRepo
	ciu             *ComputeInstanceUsercase
	log             *log.Helper
}

func NewNetworkMappingUseCase(repo NetworkMappingRepo,
	gatewayRepo GatewayRepo,
	gatewayPortRepo GatewayPortRepo,
	taskRepo TaskRepo,
	ciu *ComputeInstanceUsercase,
	logger log.Logger) *NetworkMappingUseCase {
	return &NetworkMappingUseCase{
		repo:            repo,
		gatewayRepo:     gatewayRepo,
		gatewayPortRepo: gatewayPortRepo,
		ciu:             ciu,
		taskRepo:        taskRepo,
		log:             log.NewHelper(logger),
	}
}

func (m *NetworkMappingUseCase) CreateNetworkMapping(ctx context.Context, nmc *NetworkMappingCreate) (*NetworkMapping, error) {
	// 查看当前 gatewayID
	gpcList, err := m.gatewayPortRepo.CountGatewayPortByIsUsed(ctx, false)
	if err != nil {
		return nil, err
	}

	maxItem := lo.MaxBy(gpcList, func(item *GatewayPortCount, max *GatewayPortCount) bool {
		return item.Count > max.Count
	})

	if maxItem == nil {
		return nil, fmt.Errorf("无可用 Gateway")
	}
	if maxItem.Count <= 1 {
		return nil, fmt.Errorf("无可用端口")
	}
	// 查询当前 gateway 的空余端口并进行分配
	gp, err := m.gatewayPortRepo.GetGatewayPortFirstByNotUsed(ctx, maxItem.FkGatewayID)
	if err != nil {
		return nil, err
	}
	fkGatewayId, err := uuid.Parse(gp.FkGatewayID)
	if err != nil {
		return nil, err
	}
	gatewayPort := gp.Port

	// 保存数据库
	// 进行网络映射转换
	nm := NetworkMapping{
		ID:   uuid.UUID{},
		Name: nmc.Name,
		// gateway id
		FkGatewayID: gp.FkGatewayID,
		// computer_id
		FkComputerID: nmc.ComputerId,
		// 映射到网关的端口号
		GatewayPort: int(gatewayPort),
		// 需要映射的虚拟机端口号
		ComputerPort: int(nmc.ComputerPort),
		//  0 待开始 1 进行中 2 已完成，3 失败
		Status: 0,
	}
	err = m.repo.CreateNetworkMapping(ctx, &nm)
	if err != nil {
		return nil, err
	}
	// 通过 computerID 得到 agentID
	ci, err := m.ciu.Get(ctx, nmc.ComputerId)
	if err != nil {
		return nil, err
	}
	// 发送任务给 agentID
	// 构建任务参数

	gateway, err := m.gatewayRepo.GetGateway(ctx, gp.FkGatewayID)
	if err != nil {
		return nil, err
	}

	nptp := &queue.NatNetworkMappingTaskParamVO{
		Id:           nm.ID.String(),
		Name:         nm.Name,
		InstanceName: ci.Name,
		InstancePort: int64(nm.ComputerPort),
		RemotePort:   int64(nm.GatewayPort),
		GatewayId:    nm.FkGatewayID.String(),
		GatewayIp:    gateway.IP,
		GatewayPort:  int64(gateway.Port),
	}
	paramData, err := json.Marshal(nptp)
	if err != nil {
		return nil, err
	}
	param := string(paramData)
	task := &Task{
		AgentID: ci.AgentId,
		//   NAT_PROXY_CREATE = 6; // nat 代理创建
		Cmd: queue.TaskCmd_NAT_PROXY_CREATE,
		// 执行参数，nat 网络类型对应 NatProxyCreateVO, 虚拟机类型对应 ComputeInstanceTaskParamVO
		Params: &param,
		//   CREATED = 0; //创建
		Status: queue.TaskStatus_CREATED,
		// CreateTime holds the value of the "create_time" field.
		CreateTime: time.Now(),
	}
	err = m.taskRepo.CreateTask(ctx, task)
	if err != nil {
		return nil, err
	}

	gp.IsUse = true
	err = m.gatewayPortRepo.Update(ctx, gp)
	return &nm, err
}

func (m *NetworkMappingUseCase) PageNetworkMapping(ctx context.Context, computerId uuid.UUID, page int32, size int32) ([]*NetworkMapping, int32, error) {
	m.log.WithContext(ctx).Infof("PageNetorkMapping %s %d %d", computerId, page, size)
	return m.repo.PageNetworkMappingByComputerID(ctx, computerId, page, size)
}

func (m *NetworkMappingUseCase) GetNetworkMapping(ctx context.Context, id uuid.UUID) (*NetworkMapping, error) {
	m.log.WithContext(ctx).Infof("GetNetorkMapping %s", id)
	return m.repo.GetNetworkMapping(ctx, id)
}

func (m *NetworkMappingUseCase) DeleteNetworkMapping(ctx context.Context, id uuid.UUID) error {
	m.log.WithContext(ctx).Infof("DeleteNetworkMapping %s", id)

	np, err := m.GetNetworkMapping(ctx, id)
	if err != nil {
		return err
	}

	gp, err := m.gatewayPortRepo.GetGatewayPortByGatewayIdAndPort(ctx, np.FkGatewayID, np.GatewayPort)
	if err != nil {
		return err
	}

	instance, err := m.ciu.Get(ctx, np.FkComputerID)
	if err != nil {
		return err
	}
	fkGatewayId := gp.FkGatewayID
	gateway, err := m.gatewayRepo.GetGateway(ctx, fkGatewayId)
	if err != nil {
		return err
	}

	nptp := &queue.NatNetworkMappingTaskParamVO{
		Id:           np.ID.String(),
		Name:         np.Name,
		InstanceName: instance.Name,
		InstancePort: int64(np.ComputerPort),
		RemotePort:   int64(np.GatewayPort),
		GatewayId:    gp.FkGatewayID.String(),
		GatewayIp:    gateway.IP,
		GatewayPort:  int64(gateway.Port),
	}
	paramData, err := json.Marshal(nptp)
	if err != nil {
		return err
	}
	param := string(paramData)
	task := &Task{
		AgentID: instance.AgentId,
		Cmd:     queue.TaskCmd_NAT_PROXY_DELETE,
		// 执行参数，nat 网络类型对应 NatProxyCreateVO, 虚拟机类型对应 ComputeInstanceTaskParamVO
		Params: &param,
		//   CREATED = 0; //创建
		Status: queue.TaskStatus_CREATED,
		// CreateTime holds the value of the "create_time" field.
		CreateTime: time.Now(),
	}
	err = m.taskRepo.CreateTask(ctx, task)
	if err != nil {
		return err
	}

	gp.IsUse = false

	err = m.gatewayPortRepo.Update(ctx, gp)
	if err != nil {
		return err
	}

	return m.repo.DeleteNetworkMapping(ctx, id)
}

func (m *NetworkMappingUseCase) UpdateNetorkMapping(ctx context.Context, id uuid.UUID, status int) error {
	m.log.WithContext(ctx).Infof("UpdateNetorkMapping %s", id, status)
	nm, err := m.repo.GetNetworkMapping(ctx, id)
	if err != nil {
		return err
	}
	nm.Status = status
	return m.repo.UpdateNetworkMapping(ctx, nm)
}
