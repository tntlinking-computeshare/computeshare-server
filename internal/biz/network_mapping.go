package biz

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mohaijiang/computeshare-server/internal/global"
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
	// 协议
	Protocol string `json:"protocol"`
	// gateway id
	FkGatewayID uuid.UUID `json:"fk_gateway_id,omitempty"`
	// computer_id
	FkComputerID uuid.UUID `json:"fk_computer_id,omitempty"`
	// computer name
	ComputerInstanceName string `json:"computer_instance_name"`
	// 映射到网关的端口号
	GatewayPort int32 `json:"gateway_port,omitempty"`
	// 需要映射的虚拟机端口号
	ComputerPort int32 `json:"computer_port,omitempty"`
	//  0 待开始 1 进行中 2 已完成，3 失败
	Status int `json:"status,omitempty"`
	// 用户id
	UserId uuid.UUID `json:"user_id"`
	// gateway ip
	GatewayIP string `json:"gateway_ip"`
}

type NetworkMappingCreate struct {
	Name         string
	Protocol     string
	ComputerId   uuid.UUID
	ComputerPort int32
}

type NextNetworkMappingInfo struct {
	PublicIP   string
	PublicPort int32
}

type NetworkMappingRepo interface {
	CreateNetworkMapping(ctx context.Context, entity *NetworkMapping) error
	GetNetworkMapping(ctx context.Context, id uuid.UUID) (*NetworkMapping, error)
	DeleteNetworkMapping(ctx context.Context, id uuid.UUID) error
	PageNetworkMappingByUserID(ctx context.Context, computerId uuid.UUID, page int32, size int32) ([]*NetworkMapping, int32, error)
	UpdateNetworkMapping(ctx context.Context, entity *NetworkMapping) error
	QueryGatewayIdByAgentId(ctx context.Context, agentId uuid.UUID) (uuid.UUID, error)
	QueryGatewayIdByComputeIds(ctx context.Context, computeInstanceIds []uuid.UUID) (uuid.UUID, error)
	GetNetworkMappingByPublicIpdAndPort(ctx context.Context, ip string, port int32) (*NetworkMapping, error)
}

type Gateway struct {
	ID   uuid.UUID
	Name string
	IP   string
	Port int32
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
	Port        int32
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
	GetGatewayPortByGatewayIdAndPort(ctx context.Context, id uuid.UUID, port int32) (*GatewayPort, error)
}

type NetworkMappingUseCase struct {
	repo              NetworkMappingRepo
	gatewayRepo       GatewayRepo
	gatewayPortRepo   GatewayPortRepo
	taskRepo          TaskRepo
	ciu               *ComputeInstanceUsercase
	domainBindingRepo DomainBindingRepository
	log               *log.Helper
}

func NewNetworkMappingUseCase(repo NetworkMappingRepo,
	gatewayRepo GatewayRepo,
	gatewayPortRepo GatewayPortRepo,
	taskRepo TaskRepo,
	domainBindingRepo DomainBindingRepository,
	ciu *ComputeInstanceUsercase,
	logger log.Logger) *NetworkMappingUseCase {
	return &NetworkMappingUseCase{
		repo:              repo,
		gatewayRepo:       gatewayRepo,
		gatewayPortRepo:   gatewayPortRepo,
		ciu:               ciu,
		taskRepo:          taskRepo,
		domainBindingRepo: domainBindingRepo,
		log:               log.NewHelper(logger),
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
	if maxItem.Count <= 0 {
		return nil, fmt.Errorf("无可用端口")
	}
	// 查询当前 gateway 的空余端口并进行分配
	gp, err := m.gatewayPortRepo.GetGatewayPortFirstByNotUsed(ctx, maxItem.FkGatewayID)
	if err != nil {
		return nil, err
	}
	gatewayPort := gp.Port

	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorize")
	}

	g, err := m.gatewayRepo.GetGateway(ctx, gp.FkGatewayID)
	if err != nil {
		return nil, err
	}

	protocol := nmc.Protocol
	if protocol == "" {
		protocol = "TCP"
	}

	// 保存数据库
	// 进行网络映射转换
	nwp := NetworkMapping{
		ID:       uuid.UUID{},
		Name:     nmc.Name,
		Protocol: protocol,
		// gateway id
		FkGatewayID: gp.FkGatewayID,
		// computer_id
		FkComputerID: nmc.ComputerId,
		// 映射到网关的端口号
		GatewayPort: gatewayPort,
		// 需要映射的虚拟机端口号
		ComputerPort: nmc.ComputerPort,
		//  0 待开始 1 进行中 2 已完成，3 失败
		Status:    0,
		UserId:    claim.GetUserId(),
		GatewayIP: g.IP,
	}
	err = m.repo.CreateNetworkMapping(ctx, &nwp)
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
		Id:           nwp.ID.String(),
		Name:         fmt.Sprintf("NetworkMapping_%s", nwp.ID.String()),
		InstanceId:   ci.ID.String(),
		InstancePort: nwp.ComputerPort,
		RemotePort:   nwp.GatewayPort,
		GatewayId:    nwp.FkGatewayID.String(),
		GatewayIp:    g.IP,
		GatewayPort:  g.Port,
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
	return &nwp, err
}

func (m *NetworkMappingUseCase) PageNetworkMapping(ctx context.Context, userId uuid.UUID, page int32, size int32) ([]*NetworkMapping, int32, error) {
	m.log.WithContext(ctx).Infof("PageNetorkMapping %s %d %d", userId, page, size)
	return m.repo.PageNetworkMappingByUserID(ctx, userId, page, size)
}

func (m *NetworkMappingUseCase) GetNetworkMapping(ctx context.Context, id uuid.UUID) (*NetworkMapping, error) {
	m.log.WithContext(ctx).Infof("GetNetorkMapping %s", id)
	return m.repo.GetNetworkMapping(ctx, id)
}

func (m *NetworkMappingUseCase) DeleteNetworkMapping(ctx context.Context, id uuid.UUID) error {
	m.log.WithContext(ctx).Infof("DeleteNetworkMapping %s", id)

	nwp, err := m.GetNetworkMapping(ctx, id)
	if err != nil {
		return err
	}

	// 判断有无域名绑定
	domains, err := m.domainBindingRepo.ListByNetworkMappingId(ctx, nwp.ID)
	if err != nil {
		return err
	}

	if len(domains) > 0 {
		return errors.New("请先解绑域名")
	}

	gp, err := m.gatewayPortRepo.GetGatewayPortByGatewayIdAndPort(ctx, nwp.FkGatewayID, nwp.GatewayPort)
	if err != nil {
		return err
	}

	instance, err := m.ciu.Get(ctx, nwp.FkComputerID)
	if err != nil {
		return err
	}
	fkGatewayId := gp.FkGatewayID
	gateway, err := m.gatewayRepo.GetGateway(ctx, fkGatewayId)
	if err != nil {
		return err
	}

	nptp := &queue.NatNetworkMappingTaskParamVO{
		Id:           nwp.ID.String(),
		Name:         fmt.Sprintf("NetworkMapping_%s", nwp.ID.String()),
		InstanceId:   instance.ID.String(),
		InstancePort: nwp.ComputerPort,
		RemotePort:   nwp.GatewayPort,
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

	return m.gatewayPortRepo.Update(ctx, gp)

}

func (m *NetworkMappingUseCase) UpdateNetworkMapping(ctx context.Context, id uuid.UUID, status int) error {
	m.log.WithContext(ctx).Infof("UpdateNetworkMapping %s", id, status)
	nm, err := m.repo.GetNetworkMapping(ctx, id)
	if err != nil {
		return err
	}
	nm.Status = status
	return m.repo.UpdateNetworkMapping(ctx, nm)
}

func (m *NetworkMappingUseCase) NextNetworkMapping(ctx context.Context, computeInstanceId uuid.UUID) (*NextNetworkMappingInfo, error) {
	// 1. 判断这个agent 是否有networkmapping 记录
	gatewayId, err := m.repo.QueryGatewayIdByComputeIds(ctx, []uuid.UUID{computeInstanceId})
	if err != nil {
		return nil, err
	}

	g, err := m.gatewayRepo.GetGateway(ctx, gatewayId)
	if err != nil {
		return nil, err
	}

	//2. 查询gateway 的最小可用端口
	gatewayPort, err := m.gatewayPortRepo.GetGatewayPortFirstByNotUsed(ctx, gatewayId)

	return &NextNetworkMappingInfo{
		PublicIP:   g.IP,
		PublicPort: gatewayPort.Port,
	}, err
}

func (m *NetworkMappingUseCase) GetNetworkMappingIP(ctx context.Context, networkMappingId uuid.UUID) (string, error) {
	mapping, err := m.repo.GetNetworkMapping(ctx, networkMappingId)
	if err != nil {
		return "", err
	}
	gateway, err := m.gatewayRepo.GetGateway(ctx, mapping.FkGatewayID)
	if err != nil {
		return "", err
	}

	return gateway.IP, err
}

func (m *NetworkMappingUseCase) GetNextGatewayPort(ctx context.Context, agentId uuid.UUID) (*GatewayPort, error) {
	// 1. 判断这个agent 是否有networkmapping 记录
	gatewayId, err := m.repo.QueryGatewayIdByAgentId(ctx, agentId)
	if err != nil {
		return nil, err
	}

	//2. 查询gateway 的最小可用端口
	return m.gatewayPortRepo.GetGatewayPortFirstByNotUsed(ctx, gatewayId)
}
