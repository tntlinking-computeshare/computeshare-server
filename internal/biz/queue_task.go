package biz

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	queue "github.com/mohaijiang/computeshare-server/api/queue/v1"
)

type Task struct {
	// ID of the ent.
	ID uuid.UUID
	// AgentID holds the value of the "agent_id" field.
	AgentID string
	//   VM_CREATE = 0; // 创建虚拟机
	//   VM_DELETE = 1;  // 删除虚拟机
	//   VM_START = 2; // 启动虚拟机
	//   VM_SHUTDOWN = 3;  //关闭虚拟机
	//   VM_RESTART = 4;  //关闭虚拟机
	//   VM_VNC_CONNECT = 5;  // vnc 连接
	//   NAT_PROXY_CREATE = 6; // nat 代理创建
	//   NAT_PROXY_DELETE = 7; // nat 代理删除
	//   NAT_VISITOR_CREATE = 8; // nat 访问创建
	//   NAT_VISITOR_DELETE = 9; // nat 访问删除
	Cmd queue.TaskCmd
	// 执行参数，nat 网络类型对应 NatProxyCreateVO, 虚拟机类型对应 ComputeInstanceTaskParamVO
	Params *string
	//   CREATED = 0; //创建
	//   EXECUTING = 1; //执行中
	//   EXECUTED = 2 ; // 执行成功
	//   FAILED = 3 ; // 执行失败
	Status queue.TaskStatus
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time
}

func (task *Task) GetTaskParam() (any, error) {
	switch task.Cmd {
	case queue.TaskCmd_VM_CREATE:
	case queue.TaskCmd_VM_DELETE:
	case queue.TaskCmd_VM_START:
	case queue.TaskCmd_VM_SHUTDOWN:
	case queue.TaskCmd_VM_RESTART:
	case queue.TaskCmd_VM_VNC_CONNECT:
		var vo queue.ComputeInstanceTaskParamVO
		err := json.Unmarshal([]byte(*task.Params), &vo)
		return vo, err
	case queue.TaskCmd_NAT_PROXY_CREATE:
	case queue.TaskCmd_NAT_PROXY_DELETE:
	case queue.TaskCmd_NAT_VISITOR_CREATE:
	case queue.TaskCmd_NAT_VISITOR_DELETE:
		var vo queue.NatNetworkMappingTaskParamVO
		err := json.Unmarshal([]byte(*task.Params), &vo)
		return vo, err
	}
	return nil, errors.New("cannot issue command")
}

type TaskRepo interface {
	CreateTask(ctx context.Context, entity *Task) error
	GetTask(ctx context.Context, id uuid.UUID) (*Task, error)
	ListTaskByAgentID(ctx context.Context, agentID string) ([]*Task, error)
	UpdateTask(ctx context.Context, task *Task) error
}

type TaskUseCase struct {
	repo                  TaskRepo
	networkMappingUseCase *NetworkMappingUseCase
	log                   *log.Helper
}

func NewTaskUseCase(repo TaskRepo, networkMappingUseCase *NetworkMappingUseCase, logger log.Logger) *TaskUseCase {
	return &TaskUseCase{
		repo:                  repo,
		networkMappingUseCase: networkMappingUseCase,
		log:                   log.NewHelper(logger),
	}
}

func (m *TaskUseCase) CreateTask(ctx context.Context, task *Task) (*Task, error) {
	err := m.repo.CreateTask(ctx, task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (m *TaskUseCase) ListTaskByAgentID(ctx context.Context, agentID string) ([]*Task, error) {
	m.log.WithContext(ctx).Infof("ListTaskByAgentID %s %d %d", agentID)
	return m.repo.ListTaskByAgentID(ctx, agentID)
}

func (m *TaskUseCase) GetTask(ctx context.Context, id uuid.UUID) (*Task, error) {
	m.log.WithContext(ctx).Infof("GetTask %s", id)
	return m.repo.GetTask(ctx, id)
}

func (m *TaskUseCase) UpdateTask(ctx context.Context, task *Task) error {
	//进行任务状态后的逻辑处理 TODO
	param, err := task.GetTaskParam()
	if err != nil {
		return err
	}
	switch task.Cmd {
	case queue.TaskCmd_VM_CREATE:
	case queue.TaskCmd_VM_DELETE:
	case queue.TaskCmd_VM_START:
	case queue.TaskCmd_VM_SHUTDOWN:
	case queue.TaskCmd_VM_RESTART:
	case queue.TaskCmd_VM_VNC_CONNECT:
	case queue.TaskCmd_NAT_PROXY_CREATE:
		id, err := uuid.Parse(param.(queue.NatNetworkMappingTaskParamVO).Id)
		if err != nil {
			return err
		}
		m.networkMappingUseCase.UpdateNetorkMapping(ctx, id, int(task.Status))
	case queue.TaskCmd_NAT_PROXY_DELETE:
	case queue.TaskCmd_NAT_VISITOR_CREATE:
	case queue.TaskCmd_NAT_VISITOR_DELETE:
	}
	return m.repo.UpdateTask(ctx, task)
}
