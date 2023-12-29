package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"

	pb "github.com/mohaijiang/computeshare-server/api/compute/v1"
)

type SandboxService struct {
	pb.UnimplementedSandboxServer
	cis *ComputeInstanceService
	nms *NetworkMappingService
	log *log.Helper
}

func NewSandboxService(cis *ComputeInstanceService, nms *NetworkMappingService, logger log.Logger) *SandboxService {
	return &SandboxService{
		cis: cis,
		nms: nms,
		log: log.NewHelper(logger),
	}
}

func (s *SandboxService) CreateInstanceProcess(ctx context.Context, req *pb.CreateSandboxRequest) (*pb.CreateSandboxReply, error) {

	s.log.Info("通过配置流程创建虚拟机")
	s.log.Info("开始创建虚拟机")
	instanceReply, err := s.cis.Create(ctx, req.Instance)
	if err != nil {
		return nil, err
	}
	instanceId := instanceReply.Data.GetId()
	s.log.Info("创建虚拟机的ID:", instanceId)
	s.log.Info("开始创建网络映射")
	var mappings []*pb.CreateSandboxReply_CreateSandbox_NetworkMapping
	for _, mapping := range req.NetworkMapping {
		mapping.ComputerId = instanceId
		networkMappingReply, err := s.nms.CreateNetworkMapping(ctx, mapping)
		if err != nil {
			s.log.Errorf("创建网络映射失败：映射名:%s, 资源id: %s, 资源端口号： %d", mapping.Name, mapping.ComputerId, mapping.ComputerPort)
			continue
		}
		s.log.Info("创建的网络映射id:", networkMappingReply.NetworkMapping.Id)
		mappings = append(mappings, &pb.CreateSandboxReply_CreateSandbox_NetworkMapping{
			Id:           networkMappingReply.NetworkMapping.Id,
			Name:         networkMappingReply.NetworkMapping.Name,
			ComputerPort: networkMappingReply.NetworkMapping.InstancePort,
			ServerIp:     networkMappingReply.NetworkMapping.GatewayIp,
			ServerPort:   networkMappingReply.NetworkMapping.GatewayPort,
		})
	}

	return &pb.CreateSandboxReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.CreateSandboxReply_CreateSandboxReply_Data{
			InstanceId:      instanceId,
			NetworkMappings: mappings,
		},
	}, nil
}
