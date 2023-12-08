package data

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	queue "github.com/mohaijiang/computeshare-server/api/queue/v1"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/task"
	"github.com/samber/lo"
)

type TaskRepo struct {
	data *Data
	log  *log.Helper
}

func NewTaskRepo(data *Data, logger log.Logger) biz.TaskRepo {
	return &TaskRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *TaskRepo) CreateTask(ctx context.Context, entity *biz.Task) error {
	data, err := repo.data.getTaskClient(ctx).Create().
		SetAgentID(entity.AgentID).
		SetCmd(int32(entity.Cmd)).
		SetParams(*entity.Params).
		SetStatus(int(entity.Status)).
		SetCreateTime(entity.CreateTime).
		Save(ctx)

	if err != nil {
		return err
	}
	entity.ID = data.ID
	return err
}

func (repo *TaskRepo) GetTask(ctx context.Context, id uuid.UUID) (*biz.Task, error) {
	instance, err := repo.data.getTaskClient(ctx).Get(ctx, id)
	return repo.toBiz(instance, 0), err
}

func (repo *TaskRepo) ListTaskByAgentID(ctx context.Context, agentID string) ([]*biz.Task, error) {
	list, err := repo.data.getTaskClient(ctx).Query().Where(task.AgentID(agentID)).Order(task.ByCreateTime(sql.OrderAsc())).All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(list, repo.toBiz), nil
}

func (repo *TaskRepo) UpdateTask(ctx context.Context, entity *biz.Task) error {
	return repo.data.getTaskClient(ctx).UpdateOneID(entity.ID).
		SetAgentID(entity.AgentID).
		SetCmd(int32(entity.Cmd)).
		SetParams(*entity.Params).
		SetStatus(int(entity.Status)).
		Exec(ctx)
}

func (repo *TaskRepo) toBiz(item *ent.Task, _ int) *biz.Task {
	if item == nil {
		return nil
	}
	return &biz.Task{
		ID:         item.ID,
		AgentID:    item.AgentID,
		Cmd:        queue.TaskCmd(item.Cmd),
		Params:     item.Params,
		CreateTime: item.CreateTime,
		Status:     queue.TaskStatus(item.Status),
	}
}

func (repo *TaskRepo) GetToDoTaskByAgentId(ctx context.Context, agentId string) (*biz.Task, error) {
	t, err := repo.data.getTaskClient(ctx).Query().
		Where(task.AgentIDEQ(agentId), task.StatusEQ(int(queue.TaskStatus_CREATED))).
		Order(task.ByCreateTime(sql.OrderAsc())).
		First(ctx)
	return repo.toBiz(t, 0), err
}
