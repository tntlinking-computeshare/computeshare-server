package service

import (
	"context"
	"encoding/json"

	pb "github.com/mohaijiang/computeshare-server/api/queue/v1"
)

type QueueTaskService struct {
	pb.UnimplementedQueueTaskServer
	task *pb.QueueTaskVo
}

func NewQueueTaskService() *QueueTaskService {
	vo := pb.NatProxyCreateVO{
		Id:           "uuid",
		Name:         "my-nat-1",
		InstanceName: "name1",
		InstancePort: 10088,
		RemotePort:   6000,
	}
	marshal, _ := json.Marshal(vo)

	task := &pb.QueueTaskVo{
		Id:      1,
		AgentId: "1",
		Cmd:     pb.QueueCmd_NAT_PROXY_CREATE,
		Params:  string(marshal),
		Status:  pb.TaskStatus_CREATED,
	}

	return &QueueTaskService{
		task: task,
	}
}

func (s *QueueTaskService) GetAgentTask(ctx context.Context, req *pb.QueueTaskGetRequest) (*pb.QueueTaskGetResponse, error) {

	if s.task.Status == pb.TaskStatus_CREATED {
		return &pb.QueueTaskGetResponse{
			Code:    200,
			Message: SUCCESS,
			Data:    s.task,
		}, nil
	} else {
		return &pb.QueueTaskGetResponse{
			Code:    200,
			Message: SUCCESS,
		}, nil
	}

}
func (s *QueueTaskService) UpdateAgentTask(ctx context.Context, req *pb.QueueTaskUpdateRequest) (*pb.QueueTaskUpdateResponse, error) {
	s.task.Status = req.Status
	return &pb.QueueTaskUpdateResponse{
		Code:    200,
		Message: SUCCESS,
	}, nil
}
