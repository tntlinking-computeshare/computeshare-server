package biz

import (
	"context"
	"fmt"
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
}

type GatewayPort struct {
	ID          uuid.UUID
	FkGatewayID string
	Port        int64
	IsUse       bool
}

type GatewayPortCount struct {
	FkGatewayID string `json:"fk_gateway_id" db:"fk_gateway_id"`
	Count       int32  `json:"count" db:"count"`
}

type GatewayPortRepo interface {
	CountGatewayPortByIsUsed(ctx context.Context, isUsed bool) ([]*GatewayPortCount, error)
	GetGatwayPortFirstByNotUsed(ctx context.Context, gatewayID string) (*GatewayPort, error)
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
	maxItem := findMax(gpcList, func(item *GatewayPortCount) int32 {
		return item.Count
	})
	if maxItem == nil {
		return nil, fmt.Errorf("无可用 Gateway")
	}
	if maxItem.Count <= 1 {
		return nil, fmt.Errorf("无可用端口")
	}
	// 查询当前 gateway 的空余端口并进行分配
	gp, err := m.gatewayPortRepo.GetGatwayPortFirstByNotUsed(ctx, maxItem.FkGatewayID)
	if err != nil {
		return nil, err
	}
	fkGatewayId, err := uuid.Parse(gp.FkGatewayID)
	if err != nil {
		return nil, err
	}
	gatewayPort := gp.Port
	// TODO:更新 gateway 端口变为已使用
	// 保存数据库
	// 进行网络映射转换
	nm := NetworkMapping{
		ID:   uuid.UUID{},
		Name: nmc.Name,
		// gateway id
		FkGatewayID: fkGatewayId,
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
	nptp := &queue.NatNetworkMappingTaskParamVO{
		Id:           nm.ID.String(),
		Name:         nm.Name,
		InstanceName: ci.Name,
		InstancePort: int64(nm.ComputerPort),
		RemotePort:   int64(nm.GatewayPort),
		GatewayId:    nm.FkGatewayID.String(),
	}
	param := nptp.String()
	task := &Task{
		AgentID: ci.PeerID,
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
	return &nm, nil
}

func (m *NetworkMappingUseCase) PageNetorkMapping(ctx context.Context, computerId uuid.UUID, page int32, size int32) ([]*NetworkMapping, int32, error) {
	m.log.WithContext(ctx).Infof("PageNetorkMapping %s %d %d", computerId, page, size)
	return m.repo.PageNetworkMappingByComputerID(ctx, computerId, page, size)
}

func (m *NetworkMappingUseCase) GetNetorkMapping(ctx context.Context, id uuid.UUID) (*NetworkMapping, error) {
	m.log.WithContext(ctx).Infof("GetNetorkMapping %s", id)
	return m.repo.GetNetworkMapping(ctx, id)
}

func (m *NetworkMappingUseCase) DeleteNetorkMapping(ctx context.Context, id uuid.UUID) error {
	m.log.WithContext(ctx).Infof("DeleteNetworkMapping %s", id)
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

// findMax 函数用于查找属性值最大的列表对象
func findMax(list []*GatewayPortCount, getValue func(*GatewayPortCount) int32) *GatewayPortCount {
	if len(list) == 0 {
		// 处理空列表的情况
		return nil
	}

	maxItem := list[0]
	maxValue := getValue(maxItem)

	for i := 1; i < len(list); i++ {
		currentValue := getValue(list[i])
		if currentValue > maxValue {
			maxValue = currentValue
			maxItem = list[i]
		}
	}

	return maxItem
}
