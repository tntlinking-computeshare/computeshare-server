package data

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/computeinstance"
	"github.com/mohaijiang/computeshare-server/internal/global/consts"
	"github.com/samber/lo"
	"time"
)

type computeInstanceRepo struct {
	data *Data
	log  *log.Helper
}

func NewComputeInstanceRepo(data *Data, logger log.Logger) biz.ComputeInstanceRepo {
	return &computeInstanceRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (csr *computeInstanceRepo) List(ctx context.Context, owner string) ([]*biz.ComputeInstance, error) {
	list, err := csr.data.getComputeInstance(ctx).Query().
		Where(computeinstance.OwnerEQ(owner)).
		Order(computeinstance.ByExpirationTime(sql.OrderDesc())).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(list, csr.toBiz), err
}

func (csr *computeInstanceRepo) ListByPeerId(ctx context.Context, agentId string) ([]*biz.ComputeInstance, error) {
	list, err := csr.data.getComputeInstance(ctx).Query().
		Where(computeinstance.AgentIDEQ(agentId)).
		Order(computeinstance.ByExpirationTime(sql.OrderDesc())).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(list, csr.toBiz), err
}

func (crs *computeInstanceRepo) Create(ctx context.Context, in *biz.ComputeInstance) error {
	entity, err := crs.data.getComputeInstance(ctx).Create().
		SetOwner(in.Owner).
		SetName(in.Name).
		SetCore(in.Core).
		SetMemory(in.Memory).
		SetImage(in.Image).
		SetExpirationTime(in.ExpirationTime).
		SetStatus(in.Status).
		SetPort(in.Port).
		SetAgentID(in.AgentId).
		Save(ctx)

	if err != nil {
		return err
	}

	in.ID = entity.ID
	return err
}

func (crs *computeInstanceRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return crs.data.getComputeInstance(ctx).DeleteOneID(id).Exec(ctx)
}

func (crs *computeInstanceRepo) Update(ctx context.Context, id uuid.UUID, instance *biz.ComputeInstance) error {
	return crs.data.getComputeInstance(ctx).UpdateOneID(id).
		SetStatus(instance.Status).
		SetAgentID(instance.AgentId).
		SetContainerID(instance.ContainerID).
		Exec(ctx)
}

func (crs *computeInstanceRepo) toBiz(item *ent.ComputeInstance, _ int) *biz.ComputeInstance {
	if item == nil {
		return nil
	}
	return &biz.ComputeInstance{
		ID:             item.ID,
		Owner:          item.Owner,
		Name:           item.Name,
		Core:           item.Core,
		Memory:         item.Memory,
		Image:          item.Image,
		Port:           item.Port,
		ExpirationTime: item.ExpirationTime,
		Status:         item.Status,
		ContainerID:    item.ContainerID,
		AgentId:        item.AgentID,
	}
}

func (crs *computeInstanceRepo) Get(ctx context.Context, id uuid.UUID) (*biz.ComputeInstance, error) {
	instance, err := crs.data.getComputeInstance(ctx).Get(ctx, id)
	return crs.toBiz(instance, 0), err
}

func (crs *computeInstanceRepo) ListAll(ctx context.Context) ([]*biz.ComputeInstance, error) {
	result, err := crs.data.getComputeInstance(ctx).Query().Where(computeinstance.StatusEQ(consts.InstanceStatusRunning)).All(ctx)
	if err != nil {
		return []*biz.ComputeInstance{}, err
	}

	return lo.Map(result, crs.toBiz), err
}

func stateKey(id uuid.UUID) string {
	return fmt.Sprintf("compute_instance:stats:%s", id.String())
}

func (crs *computeInstanceRepo) SaveInstanceStats(ctx context.Context, id uuid.UUID, rdbInstance *biz.ComputeInstanceRds) error {
	key := stateKey(id)
	err := crs.data.rdb.RPush(ctx, key, rdbInstance).Err()
	if err != nil {
		return err
	}
	length, err := crs.data.rdb.LLen(ctx, key).Result()
	if err != nil {
		return err
	}

	extraSize := length - 10

	if extraSize > 0 {
		crs.data.rdb.LPop(ctx, key)
	}

	return nil

}
func (crs *computeInstanceRepo) GetInstanceStats(ctx context.Context, id uuid.UUID) ([]*biz.ComputeInstanceRds, error) {
	var result []*biz.ComputeInstanceRds
	err := crs.data.rdb.LRange(ctx, stateKey(id), 0, 10).ScanSlice(&result)
	return result, err
}

func (crs *computeInstanceRepo) SetInstanceExpiration(ctx context.Context) error {
	return crs.data.getComputeInstance(ctx).Update().
		SetStatus(consts.InstanceStatusExpire).
		Where(
			computeinstance.ExpirationTimeLT(time.Now()),
			computeinstance.StatusNEQ(consts.InstanceStatusExpire),
		).
		Exec(ctx)
}

func (crs *computeInstanceRepo) UpdateStatus(ctx context.Context, id uuid.UUID, status consts.InstanceStatus) error {
	return crs.data.getComputeInstance(ctx).UpdateOneID(id).SetStatus(status).Exec(ctx)
}
