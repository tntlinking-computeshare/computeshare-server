package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"

	pb "github.com/mohaijiang/computeshare-server/api/compute/v1"
)

type ProcessService struct {
	pb.UnimplementedProcessServer
	cis *ComputeInstanceService
	nms *NetworkMappingService
	log *log.Helper
}

func NewProcessService(cis *ComputeInstanceService, nms *NetworkMappingService, logger log.Logger) *ProcessService {
	return &ProcessService{
		cis: cis,
		nms: nms,
		log: log.NewHelper(logger),
	}
}

func (s *ProcessService) CreateInstanceProcess(ctx context.Context, req *pb.CreateInstanceProcessRequest) (*pb.CreateInstanceProcessReply, error) {

	s.log.Info("通过配置流程创建虚拟机")
	s.log.Info("开始创建虚拟机")
	instanceReply, err := s.cis.Create(ctx, req.Instance)
	if err != nil {
		return nil, err
	}
	instanceId := instanceReply.Data.GetId()
	s.log.Info("创建虚拟机的ID:", instanceId)
	s.log.Info("开始创建网络映射")
	for _, mapping := range req.NetworkMapping {
		mapping.ComputerId = instanceId
		networkMappingReply, err := s.nms.CreateNetworkMapping(ctx, mapping)
		if err != nil {
			return nil, err
		}
		s.log.Info("创建的网络映射id:", networkMappingReply.NetworkMapping.Id)
	}
	return &pb.CreateInstanceProcessReply{
		Code:    200,
		Message: SUCCESS,
	}, nil
}
