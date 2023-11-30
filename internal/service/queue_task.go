package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	pb "github.com/mohaijiang/computeshare-server/api/queue/v1"
	"github.com/mohaijiang/computeshare-server/internal/biz"
)

type QueueTaskService struct {
	pb.UnimplementedQueueTaskServer
	taskUseCase *biz.TaskUseCase
	log         *log.Helper
}

func NewQueueTaskService(taskUseCase *biz.TaskUseCase, logger log.Logger) *QueueTaskService {
	return &QueueTaskService{
		taskUseCase: taskUseCase,
		log:         log.NewHelper(logger),
	}
}

func (s *QueueTaskService) GetAgentTask(ctx context.Context, req *pb.QueueTaskGetRequest) (*pb.QueueTaskGetResponse, error) {
	task, err := s.taskUseCase.GetToDoTaskByAgentId(ctx, req.Id)
	if err != nil {
		return &pb.QueueTaskGetResponse{
			Code:    500,
			Message: "ERROR",
		}, err
	}
	return &pb.QueueTaskGetResponse{
		Code:    200,
		Message: SUCCESS,
		Data:    s.toReply(task, 0),
	}, nil
}
func (s *QueueTaskService) UpdateAgentTask(ctx context.Context, req *pb.QueueTaskUpdateRequest) (*pb.QueueTaskUpdateResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return &pb.QueueTaskUpdateResponse{
			Code:    500,
			Message: "ERROR",
		}, err
	}
	task, err := s.taskUseCase.GetTask(ctx, id)
	if err != nil {
		return &pb.QueueTaskUpdateResponse{
			Code:    500,
			Message: "ERROR",
		}, err
	}
	task.Status = req.Status
	err = s.taskUseCase.UpdateTask(ctx, task)
	if err != nil {
		return &pb.QueueTaskUpdateResponse{
			Code:    500,
			Message: "ERROR",
		}, err
	}
	return &pb.QueueTaskUpdateResponse{
		Code:    200,
		Message: SUCCESS,
	}, nil
}

func (s *QueueTaskService) toReply(p *biz.Task, _ int) *pb.QueueTaskVo {
	if p == nil {
		return nil
	}
	return &pb.QueueTaskVo{
		Id:      p.ID.String(),
		AgentId: p.AgentID,
		Cmd:     p.Cmd,
		Params:  *p.Params,
		Status:  p.Status,
	}
}
