package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	pb "github.com/mohaijiang/computeshare-server/api/dashboard/v1"
	"github.com/mohaijiang/computeshare-server/internal/global/consts"
)

type DashboardUseCase struct {
	agentRepo           AgentRepo
	gatewayRepo         GatewayRepo
	gatewayPortRepo     GatewayPortRepo
	cycleRedeemCodeRepo CycleRedeemCodeRepo
	cycleRechargeRepo   CycleRechargeRepo
	logger              log.Logger
}

func NewDashboardUseCase(
	agentRepo AgentRepo,
	gatewayRepo GatewayRepo,
	gatewayPortRepo GatewayPortRepo,
	cycleRedeemCodeRepo CycleRedeemCodeRepo,
	cycleRechargeRepo CycleRechargeRepo,
	logger log.Logger,
) *DashboardUseCase {
	return &DashboardUseCase{
		agentRepo:           agentRepo,
		gatewayRepo:         gatewayRepo,
		gatewayPortRepo:     gatewayPortRepo,
		cycleRedeemCodeRepo: cycleRedeemCodeRepo,
		cycleRechargeRepo:   cycleRechargeRepo,
		logger:              logger,
	}
}

func (d *DashboardUseCase) ProvidersCount(ctx context.Context) (count int, err error) {
	countAgent, err := d.agentRepo.CountAgent(ctx)
	if err != nil {
		return 0, err
	}
	return countAgent, nil
}

func (d *DashboardUseCase) ProvidersList(ctx context.Context) (list []*pb.ProvidersListReply_ProvidersList, err error) {
	agents, err := d.agentRepo.ListAgent(ctx)
	if err != nil {
		return nil, err
	}
	for _, agent := range agents {
		var providersList pb.ProvidersListReply_ProvidersList
		providersList.Type = consts.Agent
		providersList.Active = agent.Active
		providersList.Mac = agent.MAC
		providersList.Ip = agent.IP
		providersList.TotalCpu = agent.TotalCPU
		providersList.TotalMemory = agent.TotalMemory
		list = append(list, &providersList)
	}
	return list, nil
}

func (d *DashboardUseCase) GatewaysCount(ctx context.Context) (count int, err error) {
	countGateway, err := d.gatewayRepo.CountGateway(ctx)
	if err != nil {
		return 0, err
	}
	return countGateway, nil
}

func (d *DashboardUseCase) GatewaysList(ctx context.Context) (list []*pb.GatewaysListReply_GatewaysList, err error) {
	gatewayList, err := d.gatewayRepo.ListGateway(ctx)
	if err != nil {
		return nil, err
	}
	gatewayPorts, err := d.gatewayPortRepo.CountGatewayPort(ctx)
	if err != nil {
		return nil, err
	}
	gatewayPortByIsUsed, err := d.gatewayPortRepo.CountGatewayPortByIsUsed(ctx, true)
	if err != nil {
		return nil, err
	}
	publicGatewayPortByIsUsed, err := d.gatewayPortRepo.CountPublicGatewayPortByIsUsed(ctx, true)
	if err != nil {
		return nil, err
	}
	gatewayPortMap := make(map[uuid.UUID]int32)
	gatewayPortByIsUsedMap := make(map[uuid.UUID]int32)
	publicGatewayPortByIsUsedMap := make(map[uuid.UUID]int32)
	for _, gatewayPort := range gatewayPorts {
		gatewayPortMap[gatewayPort.FkGatewayID] = gatewayPort.Count
	}
	for _, gatewayPortCount := range gatewayPortByIsUsed {
		gatewayPortByIsUsedMap[gatewayPortCount.FkGatewayID] = gatewayPortCount.Count
	}
	for _, publicGatewayPortCount := range publicGatewayPortByIsUsed {
		publicGatewayPortByIsUsedMap[publicGatewayPortCount.FkGatewayID] = publicGatewayPortCount.Count
	}
	for _, gateway := range gatewayList {
		var gatewaysList pb.GatewaysListReply_GatewaysList
		gatewaysList.Ip = gateway.IP
		gatewaysList.Name = gateway.Name
		gatewaysList.TotalPort = gatewayPortMap[gateway.ID]
		gatewaysList.UseIntranetPort = gatewayPortByIsUsedMap[gateway.ID]
		gatewaysList.UsePublicPort = publicGatewayPortByIsUsedMap[gateway.ID]
		list = append(list, &gatewaysList)
	}
	return list, nil
}

func (d *DashboardUseCase) CyclesCount(ctx context.Context) (count *pb.CyclesCountReply_CyclesCount, err error) {
	var cyclesCount pb.CyclesCountReply_CyclesCount
	countCycleRecoveryTotal, err := d.cycleRedeemCodeRepo.CountCycleRecoveryTotal(ctx)
	if err != nil {
		return nil, err
	}
	cyclesCount.RecoveryTotal = countCycleRecoveryTotal.StringFixed(2)
	countCycleUseTotal, err := d.cycleRedeemCodeRepo.CountCycleUseTotal(ctx)
	if err != nil {
		return nil, err
	}
	cyclesCount.GrantVouchersTotal = countCycleUseTotal.StringFixed(2)
	countRechargeCycle, err := d.cycleRechargeRepo.CountRechargeCycle(ctx)
	if err != nil {
		return nil, err
	}
	cyclesCount.RechargedTotal = countRechargeCycle.StringFixed(2)
	return &cyclesCount, nil
}
