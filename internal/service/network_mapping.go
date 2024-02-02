package service

import (
	"context"
	"errors"
	"github.com/mohaijiang/computeshare-server/internal/global"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	pb "github.com/mohaijiang/computeshare-server/api/network_mapping/v1"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/samber/lo"
)

type NetworkMappingService struct {
	pb.UnimplementedNetworkMappingServer
	nm  *biz.NetworkMappingUseCase
	dm  *biz.DomainBindingUseCase
	log *log.Helper
}

func NewNetworkMappingService(nm *biz.NetworkMappingUseCase, dm *biz.DomainBindingUseCase, logger log.Logger) *NetworkMappingService {
	return &NetworkMappingService{
		nm:  nm,
		dm:  dm,
		log: log.NewHelper(logger),
	}
}

func (s *NetworkMappingService) CheckPermission(ctx context.Context, networkMappingId string) error {
	id, err := uuid.Parse(networkMappingId)
	if err != nil {
		return err
	}
	networkMapping, err := s.nm.GetNetworkMapping(ctx, id)
	if err != nil {
		return err
	}
	claim, ok := global.FromContext(ctx)
	if !ok {
		return errors.New("unauthorized")
	}

	userId := claim.GetUserId()
	if networkMapping.UserId == userId {
		return nil
	}
	return errors.New("no permission")
}

func (s *NetworkMappingService) CreateNetworkMapping(ctx context.Context, req *pb.CreateNetworkMappingRequest) (*pb.CreateNetworkMappingReply, error) {
	computerId, err := uuid.Parse(req.ComputerId)
	if err != nil {
		return nil, err
	}
	networkmapping, err := s.nm.CreateNetworkMapping(ctx, &biz.NetworkMappingCreate{
		Name:         req.Name,
		Protocol:     req.Protocol,
		ComputerId:   computerId,
		ComputerPort: req.ComputerPort,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateNetworkMappingReply{
		Code:           200,
		Message:        SUCCESS,
		NetworkMapping: s.toReply(ctx, networkmapping, 0),
	}, nil
}
func (s *NetworkMappingService) PageNetworkMapping(ctx context.Context, req *pb.PageNetworkMappingRequest) (*pb.PageNetworkMappingReply, error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("unauthorize")
	}
	list, total, err := s.nm.PageNetworkMapping(ctx, claim.GetUserId(), req.Page, req.Size)
	if err != nil {
		return nil, err
	}
	return &pb.PageNetworkMappingReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.PageNetworkMappingReply_Data{
			List: lo.Map(list, func(item *biz.NetworkMapping, index int) *pb.NetworkMappingVO {
				return s.toReply(ctx, item, index)
			}),
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

	err = s.CheckPermission(ctx, req.Id)
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
		NetworkMapping: s.toReply(ctx, networkmapping, 0),
	}, nil
}
func (s *NetworkMappingService) DeleteNetworkMapping(ctx context.Context, req *pb.DeleteNetworkMappingRequest) (*pb.DeleteNetworkMappingReply, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	err = s.CheckPermission(ctx, req.Id)
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

func (s *NetworkMappingService) toReply(ctx context.Context, p *biz.NetworkMapping, _ int) *pb.NetworkMappingVO {
	if p == nil {
		return nil
	}
	list, _ := s.dm.ListByNetworkMappingId(ctx, p.ID)

	return &pb.NetworkMappingVO{
		Id:           p.ID.String(),
		Name:         p.Name,
		Status:       int32(p.Status),
		GatewayId:    p.FkGatewayID.String(),
		InstanceId:   p.FkComputerID.String(),
		InstanceName: p.ComputerInstanceName,
		GatewayPort:  p.GatewayPort,
		InstancePort: p.ComputerPort,
		Domains:      list,
		Protocol:     p.Protocol,
		GatewayIp:    p.GatewayIP,
	}
}

func (s *NetworkMappingService) NextNetworkMapping(ctx context.Context, req *pb.NextNetworkMappingRequest) (*pb.NextNetworkMappingReply, error) {
	id, err := uuid.Parse(req.GetComputerId())
	if err != nil {
		return nil, err
	}
	next, err := s.nm.NextNetworkMapping(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.NextNetworkMappingReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.NextNetworkMappingReply_Data{
			PublicPort: next.PublicPort,
			PublicIp:   next.PublicIP,
		},
	}, err
}

func (s *NetworkMappingService) UpdateNetworkMapping(ctx context.Context, req *pb.UpdateNetworkMappingRequest) (*pb.UpdateNetworkMappingReply, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	err = s.CheckPermission(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	computerId, err := uuid.Parse(req.ComputerId)
	if err != nil {
		return nil, err
	}

	networkmapping, err := s.nm.UpdateNetworkMapping(ctx, id, &biz.NetworkMappingCreate{
		Name:         req.Name,
		Protocol:     req.Protocol,
		ComputerId:   computerId,
		ComputerPort: req.ComputerPort,
	})

	return &pb.UpdateNetworkMappingReply{
		Code:           200,
		Message:        SUCCESS,
		NetworkMapping: s.toReply(ctx, networkmapping, 0),
	}, nil
}
