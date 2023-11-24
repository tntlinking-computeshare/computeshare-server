package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	pb "github.com/mohaijiang/computeshare-server/api/network_mapping/v1"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/samber/lo"
)

type NetworkMappingService struct {
	pb.UnimplementedNetworkMappingServer
	nm  *biz.NetworkMappingUseCase
	log *log.Helper
}

func NewNetworkMappingService(nm *biz.NetworkMappingUseCase, logger log.Logger) *NetworkMappingService {
	return &NetworkMappingService{
		nm:  nm,
		log: log.NewHelper(logger),
	}
}

func (s *NetworkMappingService) CreateNetworkMapping(ctx context.Context, req *pb.CreateNetworkMappingRequest) (*pb.CreateNetworkMappingReply, error) {
	computerId, err := uuid.Parse(req.ComputerId)
	if err != nil {
		return nil, err
	}
	networkmapping, err := s.nm.CreateNetworkMapping(ctx, &biz.NetworkMappingCreate{
		Name:         req.Name,
		ComputerId:   computerId,
		ComputerPort: req.ComputerPort,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateNetworkMappingReply{
		Code:           200,
		Message:        SUCCESS,
		NetworkMapping: s.toReply(networkmapping, 0),
	}, nil
}
func (s *NetworkMappingService) PageNetworkMapping(ctx context.Context, req *pb.PageNetworkMappingRequest) (*pb.PageNetworkMappingReply, error) {
	computerId, err := uuid.Parse(req.ComputerId)
	if err != nil {
		return nil, err
	}
	list, total, err := s.nm.PageNetworkMapping(ctx, computerId, req.Page, req.Size)
	if err != nil {
		return nil, err
	}
	return &pb.PageNetworkMappingReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.PageNetworkMappingReply_Data{
			List:  lo.Map(list, s.toReply),
			Total: total,
			Page:  req.GetPage(),
			Size:  req.GetSize(),
		},
	}, nil
}
func (s *NetworkMappingService) GetNetworkMapping(ctx context.Context, req *pb.GetNetworkMappingRequest) (*pb.GetNetworkMappingReply, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	networkmapping, err := s.nm.GetNetworkMapping(ctx, id)
	if err != nil {
		return &pb.GetNetworkMappingReply{
			Code:           500,
			Message:        "ERROR",
			NetworkMapping: nil,
		}, err
	}
	return &pb.GetNetworkMappingReply{
		Code:           200,
		Message:        SUCCESS,
		NetworkMapping: s.toReply(networkmapping, 0),
	}, nil
}
func (s *NetworkMappingService) DeleteNetworkMapping(ctx context.Context, req *pb.DeleteNetworkMappingRequest) (*pb.DeleteNetworkMappingReply, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	err = s.nm.DeleteNetworkMapping(ctx, id)
	if err != nil {
		return &pb.DeleteNetworkMappingReply{
			Code:    500,
			Message: "ERROR",
		}, err
	}
	return &pb.DeleteNetworkMappingReply{
		Code:    200,
		Message: SUCCESS,
	}, nil
}

func (s *NetworkMappingService) toReply(p *biz.NetworkMapping, _ int) *pb.NetworkMappingVO {
	if p == nil {
		return nil
	}
	return &pb.NetworkMappingVO{
		Id:           p.ID.String(),
		Name:         p.Name,
		Status:       int32(p.Status),
		GatewayId:    p.FkGatewayID.String(),
		ComputerId:   p.FkComputerID.String(),
		GatewayPort:  int32(p.GatewayPort),
		ComputerPort: int32(p.ComputerPort),
	}
}
