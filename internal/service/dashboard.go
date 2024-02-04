package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/mohaijiang/computeshare-server/api/dashboard/v1"
	"github.com/mohaijiang/computeshare-server/internal/biz"
)

type DashboardService struct {
	pb.UnimplementedDashboardServer
	dashboardUseCase *biz.DashboardUseCase
	log              *log.Helper
}

func NewDashboardService(dashboardUseCase *biz.DashboardUseCase, logger log.Logger) *DashboardService {
	return &DashboardService{
		dashboardUseCase: dashboardUseCase,
		log:              log.NewHelper(logger),
	}
}

func (d *DashboardService) ProvidersCount(ctx context.Context, req *pb.ProvidersCountRequest) (*pb.ProvidersCountReply, error) {
	count, err := d.dashboardUseCase.ProvidersCount(ctx)
	return &pb.ProvidersCountReply{
		Code:    200,
		Message: SUCCESS,
		Data:    int32(count),
	}, err
}
func (d *DashboardService) GatewaysCount(ctx context.Context, req *pb.GatewaysCountRequest) (*pb.GatewaysCountReply, error) {
	count, err := d.dashboardUseCase.GatewaysCount(ctx)
	return &pb.GatewaysCountReply{
		Code:    200,
		Message: SUCCESS,
		Data:    int32(count),
	}, err
}
func (d *DashboardService) StoragesCount(ctx context.Context, req *pb.StoragesCountRequest) (*pb.StoragesCountReply, error) {
	return &pb.StoragesCountReply{}, nil
}
func (d *DashboardService) ProvidersList(ctx context.Context, req *pb.ProvidersListRequest) (*pb.ProvidersListReply, error) {
	list, err := d.dashboardUseCase.ProvidersList(ctx)
	return &pb.ProvidersListReply{
		Code:    200,
		Message: SUCCESS,
		Data:    list,
	}, err
}
func (d *DashboardService) GatewaysList(ctx context.Context, req *pb.GatewaysListRequest) (*pb.GatewaysListReply, error) {
	list, err := d.dashboardUseCase.GatewaysList(ctx)
	return &pb.GatewaysListReply{
		Code:    200,
		Message: SUCCESS,
		Data:    list,
	}, err
}
func (d *DashboardService) CyclesCount(ctx context.Context, req *pb.CyclesCountRequest) (*pb.CyclesCountReply, error) {
	count, err := d.dashboardUseCase.CyclesCount(ctx)
	return &pb.CyclesCountReply{
		Code:    200,
		Message: SUCCESS,
		Data:    count,
	}, err
}
func (d *DashboardService) SandboxCount(ctx context.Context, req *pb.SandboxCountRequest) (*pb.SandboxCountReply, error) {
	return &pb.SandboxCountReply{}, nil
}
func (d *DashboardService) LastComputeInstancesCount(ctx context.Context, req *pb.LastComputeInstancesCountRequest) (*pb.LastComputeInstancesCountReply, error) {
	list, err := d.dashboardUseCase.LastComputeInstancesCount(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.LastComputeInstancesCountReply{
		Code:    200,
		Message: SUCCESS,
		Data:    list,
	}, nil
}
