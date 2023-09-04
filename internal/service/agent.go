package service

import (
	"computeshare-server/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"

	pb "computeshare-server/api/agent/v1"
)

func NewAgentService(uc *biz.AgentUsecase, logger log.Logger) *AgentService {
	return &AgentService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *AgentService) CreateAgent(ctx context.Context, req *pb.CreateAgentRequest) (*pb.CreateAgentReply, error) {
	s.log.Infof("input data %v", req)
	err := s.uc.Create(ctx, &biz.Agent{
		Name: *req.Id,
	})
	return &pb.CreateAgentReply{}, err
}
func (s *AgentService) UpdateAgent(ctx context.Context, req *pb.UpdateAgentRequest) (*pb.UpdateAgentReply, error) {
	s.log.Infof("input data %v", req)
	return &pb.UpdateAgentReply{}, nil
}
func (s *AgentService) DeleteAgent(ctx context.Context, req *pb.DeleteAgentRequest) (*pb.DeleteAgentReply, error) {
	s.log.Infof("input data %v", req)
	return &pb.DeleteAgentReply{}, nil
}
func (s *AgentService) GetAgent(ctx context.Context, req *pb.GetAgentRequest) (*pb.GetAgentReply, error) {
	return &pb.GetAgentReply{}, nil
}
func (s *AgentService) ListAgent(ctx context.Context, req *pb.ListAgentRequest) (*pb.ListAgentReply, error) {
	return &pb.ListAgentReply{}, nil
}
