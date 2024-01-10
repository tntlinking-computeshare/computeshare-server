package biz

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/mohaijiang/computeshare-server/internal/global/consts"
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
	case queue.TaskCmd_VM_CREATE, queue.TaskCmd_VM_DELETE, queue.TaskCmd_VM_START,
		queue.TaskCmd_VM_SHUTDOWN, queue.TaskCmd_VM_RESTART, queue.TaskCmd_VM_RECREATE:
		var vo queue.ComputeInstanceTaskParamVO
		err := json.Unmarshal([]byte(*task.Params), &vo)
		return &vo, err
	case queue.TaskCmd_NAT_PROXY_CREATE, queue.TaskCmd_NAT_PROXY_DELETE,
		queue.TaskCmd_NAT_VISITOR_CREATE, queue.TaskCmd_NAT_VISITOR_DELETE:
		var vo queue.NatNetworkMappingTaskParamVO
		err := json.Unmarshal([]byte(*task.Params), &vo)
		return &vo, err
	case queue.TaskCmd_STORAGE_CREATE, queue.TaskCmd_STORAGE_DELETE:
		var vo queue.StorageSetupTaskParamVO
		err := json.Unmarshal([]byte(*task.Params), &vo)
		return &vo, err
	}

	return nil, errors.New("cannot issue command")
}

type TaskRepo interface {
	CreateTask(ctx context.Context, entity *Task) error
	GetTask(ctx context.Context, id uuid.UUID) (*Task, error)
	ListTaskByAgentID(ctx context.Context, agentID string) ([]*Task, error)
	UpdateTask(ctx context.Context, task *Task) error
	GetToDoTaskByAgentId(ctx context.Context, id string) (*Task, error)
}

type TaskUseCase struct {
	repo                   TaskRepo
	networkMappingRepo     NetworkMappingRepo
	computeInstanceRepo    ComputeInstanceRepo
	storageProviderUseCase *StorageProviderUseCase
	cycleRenewalRepo       CycleRenewalRepo
	log                    *log.Helper
}

func NewTaskUseCase(repo TaskRepo,
	networkMappingRepo NetworkMappingRepo,
	computeInstanceRepo ComputeInstanceRepo,
	storageProviderUseCase *StorageProviderUseCase,
	cycleRenewalRepo CycleRenewalRepo,
	logger log.Logger) *TaskUseCase {
	return &TaskUseCase{
		repo:                   repo,
		networkMappingRepo:     networkMappingRepo,
		computeInstanceRepo:    computeInstanceRepo,
		storageProviderUseCase: storageProviderUseCase,
		cycleRenewalRepo:       cycleRenewalRepo,
		log:                    log.NewHelper(logger),
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

func (m *TaskUseCase) GetToDoTaskByAgentId(ctx context.Context, agentID string) (*Task, error) {
	m.log.WithContext(ctx).Infof("GetToDoTaskByAgentId %s %d %d", agentID)
	return m.repo.GetToDoTaskByAgentId(ctx, agentID)
}

func (m *TaskUseCase) GetTask(ctx context.Context, id uuid.UUID) (*Task, error) {
	m.log.WithContext(ctx).Infof("GetTask %s", id)
	return m.repo.GetTask(ctx, id)
}

func (m *TaskUseCase) UpdateTask(ctx context.Context, task *Task) error {
	//TODO ...进行任务状态后的逻辑处理
	param, err := task.GetTaskParam()
	if err != nil {
		return err
	}

	getInstanceId := func(param any) (uuid.UUID, error) {
		vo, ok := param.(*queue.ComputeInstanceTaskParamVO)
		if !ok {
			return uuid.Nil, errors.New("get task param error")
		}
		instanceId, err := uuid.Parse(vo.InstanceId)
		if err != nil {
			return uuid.Nil, err
		}
		return instanceId, nil
	}

	switch task.Status {

	case queue.TaskStatus_CREATED:

	case queue.TaskStatus_EXECUTING:

		switch task.Cmd {
		case queue.TaskCmd_VM_CREATE:
			instanceId, err := getInstanceId(param)
			if err != nil {
				return err
			}
			_ = m.computeInstanceRepo.UpdateStatus(ctx, instanceId, consts.InstanceStatusCreating)
		case queue.TaskCmd_VM_DELETE:
			instanceId, err := getInstanceId(param)
			if err != nil {
				return err
			}
			_ = m.computeInstanceRepo.UpdateStatus(ctx, instanceId, consts.InstanceStatusDeleting)
		case queue.TaskCmd_VM_START:
			instanceId, err := getInstanceId(param)
			if err != nil {
				return err
			}
			_ = m.computeInstanceRepo.UpdateStatus(ctx, instanceId, consts.InstanceStatusStarting)
		case queue.TaskCmd_VM_SHUTDOWN:
			instanceId, err := getInstanceId(param)
			if err != nil {
				return err
			}
			instance, err := m.computeInstanceRepo.Get(ctx, instanceId)
			if err != nil {
				return err
			}
			if instance.Status != consts.InstanceStatusExpire {
				_ = m.computeInstanceRepo.UpdateStatus(ctx, instanceId, consts.InstanceStatusClosing)
			}
		case queue.TaskCmd_VM_RESTART:
			instanceId, err := getInstanceId(param)
			if err != nil {
				return err
			}
			_ = m.computeInstanceRepo.UpdateStatus(ctx, instanceId, consts.InstanceStatusRestarting)
		case queue.TaskCmd_VM_RECREATE:
			{
				instanceId, err := getInstanceId(param)
				if err != nil {
					return err
				}
				_ = m.computeInstanceRepo.UpdateStatus(ctx, instanceId, consts.InstanceStatusReCreating)
			}
		case queue.TaskCmd_NAT_PROXY_CREATE,
			queue.TaskCmd_NAT_PROXY_DELETE,
			queue.TaskCmd_NAT_VISITOR_CREATE,
			queue.TaskCmd_NAT_VISITOR_DELETE:
			id, err := uuid.Parse(param.(*queue.NatNetworkMappingTaskParamVO).Id)
			if err != nil {
				return err
			}
			mapping, err := m.networkMappingRepo.GetNetworkMapping(ctx, id)
			if err != nil {
				return err
			}
			mapping.Status = int(task.Status)
			_ = m.networkMappingRepo.UpdateNetworkMapping(ctx, mapping)

		case queue.TaskCmd_STORAGE_CREATE:
			id, err := uuid.Parse(param.(*queue.StorageSetupTaskParamVO).Id)
			if err != nil {
				return err
			}
			_ = m.storageProviderUseCase.UpdateStorageProviderStatus(ctx, id, consts.StorageProviderStatus_SETUPING)

		case queue.TaskCmd_STORAGE_DELETE:
			id, err := uuid.Parse(param.(*queue.StorageSetupTaskParamVO).Id)
			if err != nil {
				return err
			}
			_ = m.storageProviderUseCase.UpdateStorageProviderStatus(ctx, id, consts.StorageProviderStatus_NOT_RUN)
		}

	case queue.TaskStatus_EXECUTED:

		switch task.Cmd {
		case queue.TaskCmd_VM_CREATE:
			instanceId, err := getInstanceId(param)
			if err != nil {
				return err
			}
			_ = m.computeInstanceRepo.UpdateStatus(ctx, instanceId, consts.InstanceStatusRunning)
		case queue.TaskCmd_VM_DELETE:
			instanceId, err := getInstanceId(param)
			if err != nil {
				return err
			}
			_ = m.computeInstanceRepo.UpdateStatus(ctx, instanceId, consts.InstanceStatusDeleted)
			_ = m.computeInstanceRepo.Delete(ctx, instanceId)
			renewal, err := m.cycleRenewalRepo.QueryByResourceId(ctx, instanceId)
			if err != nil {
				renewal.State = int8(consts.RenewalState_STOP)
				renewal.RenewalTime = nil
				renewal.DueTime = nil
				_ = m.cycleRenewalRepo.Update(ctx, renewal.ID, renewal)
			}
		case queue.TaskCmd_VM_START:
			instanceId, err := getInstanceId(param)
			if err != nil {
				return err
			}
			_ = m.computeInstanceRepo.UpdateStatus(ctx, instanceId, consts.InstanceStatusRunning)
		case queue.TaskCmd_VM_SHUTDOWN:
			instanceId, err := getInstanceId(param)
			if err != nil {
				return err
			}
			instance, err := m.computeInstanceRepo.Get(ctx, instanceId)
			if err != nil {
				return err
			}
			if instance.Status != consts.InstanceStatusExpire {
				_ = m.computeInstanceRepo.UpdateStatus(ctx, instanceId, consts.InstanceStatusClosed)
			}
		case queue.TaskCmd_VM_RESTART:
			instanceId, err := getInstanceId(param)
			if err != nil {
				return err
			}
			_ = m.computeInstanceRepo.UpdateStatus(ctx, instanceId, consts.InstanceStatusRunning)
		case queue.TaskCmd_VM_RECREATE:
			instanceId, err := getInstanceId(param)
			if err != nil {
				return err
			}
			_ = m.computeInstanceRepo.UpdateStatus(ctx, instanceId, consts.InstanceStatusRunning)
		case queue.TaskCmd_NAT_PROXY_CREATE,
			queue.TaskCmd_NAT_VISITOR_CREATE:

			id, err := uuid.Parse(param.(*queue.NatNetworkMappingTaskParamVO).Id)
			if err != nil {
				return err
			}
			mapping, err := m.networkMappingRepo.GetNetworkMapping(ctx, id)
			if err != nil {
				return err
			}
			mapping.Status = int(task.Status)
			_ = m.networkMappingRepo.UpdateNetworkMapping(ctx, mapping)

		case queue.TaskCmd_NAT_PROXY_DELETE, queue.TaskCmd_NAT_VISITOR_DELETE:
			id, err := uuid.Parse(param.(*queue.NatNetworkMappingTaskParamVO).Id)
			if err != nil {
				return err
			}
			err = m.networkMappingRepo.DeleteNetworkMapping(ctx, id)
		case queue.TaskCmd_STORAGE_CREATE:
			id, err := uuid.Parse(param.(*queue.StorageSetupTaskParamVO).Id)
			if err != nil {
				return err
			}
			_ = m.storageProviderUseCase.UpdateStorageProviderStatus(ctx, id, consts.StorageProviderStatus_RUNNING)

		case queue.TaskCmd_STORAGE_DELETE:
			id, err := uuid.Parse(param.(*queue.StorageSetupTaskParamVO).Id)
			if err != nil {
				return err
			}
			_ = m.storageProviderUseCase.UpdateStorageProviderStatus(ctx, id, consts.StorageProviderStatus_NOT_RUN)
		}

	case queue.TaskStatus_FAILED:

		switch task.Cmd {

		case queue.TaskCmd_STORAGE_CREATE:
			id, err := uuid.Parse(param.(*queue.StorageSetupTaskParamVO).Id)
			if err != nil {
				return err
			}
			_ = m.storageProviderUseCase.UpdateStorageProviderStatus(ctx, id, consts.StorageProviderStatus_SETUP_FAIL)

		case queue.TaskCmd_STORAGE_DELETE:
			id, err := uuid.Parse(param.(*queue.StorageSetupTaskParamVO).Id)
			if err != nil {
				return err
			}
			_ = m.storageProviderUseCase.UpdateStorageProviderStatus(ctx, id, consts.StorageProviderStatus_NOT_RUN)
		}

	}

	return m.repo.UpdateTask(ctx, task)
}
