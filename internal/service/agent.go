package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/global/consts"
	"github.com/samber/lo"
	"strconv"

	//"github.com/ipfs/go-ipfs/core"
	"github.com/mohaijiang/computeshare-server/internal/biz"

	pb "github.com/mohaijiang/computeshare-server/api/agent/v1"

	computepb "github.com/mohaijiang/computeshare-server/api/compute/v1"
)

type AgentService struct {
	pb.UnimplementedAgentServer

	log *log.Helper

	uc *biz.AgentUsecase
}

func NewAgentService(uc *biz.AgentUsecase, logger log.Logger) *AgentService {
	return &AgentService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *AgentService) CreateAgent(ctx context.Context, req *pb.CreateAgentRequest) (*pb.CreateAgentReply, error) {
	s.log.Infof("input data %v", req)

	agent := biz.Agent{
		MAC:            req.Mac,
		Hostname:       req.Hostname,
		TotalCPU:       req.TotalCpu,
		TotalMemory:    req.TotalMemory,
		OccupiedCPU:    req.OccupiedCpu,
		OccupiedMemory: req.OccupiedMemory,
		IP:             req.Ip,
		Active:         true,
	}
	id, err := s.uc.Create(ctx, &agent)
	return &pb.CreateAgentReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.CreateAgentReply_Data{
			Id: id.String(),
		},
	}, err
}
func (s *AgentService) UpdateAgent(_ context.Context, req *pb.UpdateAgentRequest) (*pb.UpdateAgentReply, error) {
	s.log.Infof("input data %v", req)
	return &pb.UpdateAgentReply{
		Code:    200,
		Message: SUCCESS,
	}, nil
}
func (s *AgentService) DeleteAgent(_ context.Context, req *pb.DeleteAgentRequest) (*pb.DeleteAgentReply, error) {
	s.log.Infof("input data %v", req)
	return &pb.DeleteAgentReply{
		Code:    200,
		Message: SUCCESS,
	}, nil
}
func (s *AgentService) GetAgent(ctx context.Context, req *pb.GetAgentRequest) (*pb.GetAgentReply, error) {
	s.log.Infof("input data %v", req)
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	agent, err := s.uc.Get(ctx, id)
	return &pb.GetAgentReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.AgentReply{
			Id:  agent.ID.String(),
			MAC: agent.MAC,
		},
	}, err
}
func (s *AgentService) ListAgent(_ context.Context, _ *pb.ListAgentRequest) (*pb.ListAgentReply, error) {
	return &pb.ListAgentReply{
		Code:    200,
		Message: SUCCESS,
	}, nil
}

func (s *AgentService) ListAgentInstance(ctx context.Context, req *pb.ListAgentInstanceReq) (*computepb.ListInstanceReply, error) {
	result, err := s.uc.ListAgentInstance(ctx, req.Mac)
	return &computepb.ListInstanceReply{
		Code:    200,
		Message: SUCCESS,
		Data: lo.Map(result, func(item *biz.ComputeInstance, _ int) *computepb.Instance {
			return &computepb.Instance{
				Id:             item.ID.String(),
				Name:           item.Name,
				Status:         int32(item.Status),
				ExpirationTime: item.ExpirationTime.UnixMilli(),
				ImageName:      item.Image,
				Core:           strconv.Itoa(item.Core),
				Memory:         strconv.Itoa(item.Memory),
				ContainerId:    item.ContainerID,
			}
		}),
	}, err
}

func (s *AgentService) ReportInstanceStatus(ctx context.Context, req *computepb.Instance) (rsp *pb.ReportInstanceStatusReply, err error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, err
	}
	instance := &biz.ComputeInstance{
		ID:          id,
		ContainerID: req.ContainerId,
		AgentId:     req.AgentId,
		Status:      consts.InstanceStatus(req.Status),
	}
	err = s.uc.ReportInstanceStatus(ctx, instance)
	return &pb.ReportInstanceStatusReply{
		Code:    200,
		Message: SUCCESS,
	}, err
}
