package service

import (
	"context"

	pb "github.com/mohaijiang/computeshare-server/api/queue/v1"
)

type QueueTaskService struct {
	pb.UnimplementedQueueTaskServer
}

func NewQueueTaskService() *QueueTaskService {
	return &QueueTaskService{}
}

func (s *QueueTaskService) GetAgentTask(ctx context.Context, req *pb.QueueTaskGetRequest) (*pb.QueueTaskGetResponse, error) {
	return &pb.QueueTaskGetResponse{}, nil
}
func (s *QueueTaskService) UpdateAgentTask(ctx context.Context, req *pb.QueueTaskUpdateRequest) (*pb.QueueTaskUpdateResponse, error) {
	return &pb.QueueTaskUpdateResponse{}, nil
}
