package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	//"github.com/ipfs/go-ipfs/core"
	"github.com/mohaijiang/computeshare-server/internal/biz"

	pb "github.com/mohaijiang/computeshare-server/api/agent/v1"
)

type AgentService struct {
	pb.UnimplementedAgentServer

	log *log.Helper

	uc *biz.AgentUsecase

	//node *core.IpfsNode
}

func NewAgentService(uc *biz.AgentUsecase, logger log.Logger) *AgentService {
	return &AgentService{
		uc:  uc,
		log: log.NewHelper(logger),
		//node: node,
	}
}

func (s *AgentService) CreateAgent(ctx context.Context, req *pb.CreateAgentRequest) (*pb.CreateAgentReply, error) {
	s.log.Infof("input data %v", req)

	agent := &biz.Agent{
		PeerId: req.GetName(),
		Active: true,
	}
	err := s.uc.Create(ctx, agent)
	return &pb.CreateAgentReply{
		Code:    200,
		Message: SUCCESS,
		Data: &pb.CreateAgentReply_Data{
			Id: agent.ID.String(),
		},
	}, err
}
func (s *AgentService) UpdateAgent(ctx context.Context, req *pb.UpdateAgentRequest) (*pb.UpdateAgentReply, error) {
	s.log.Infof("input data %v", req)
	return &pb.UpdateAgentReply{
		Code:    200,
		Message: SUCCESS,
	}, nil
}
func (s *AgentService) DeleteAgent(ctx context.Context, req *pb.DeleteAgentRequest) (*pb.DeleteAgentReply, error) {
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
			Id:   agent.ID.String(),
			Name: agent.PeerId,
		},
	}, err
}
func (s *AgentService) ListAgent(ctx context.Context, req *pb.ListAgentRequest) (*pb.ListAgentReply, error) {
	return &pb.ListAgentReply{
		Code:    200,
		Message: SUCCESS,
	}, nil
}
