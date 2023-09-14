package data

import (
	"computeshare-server/internal/biz"
	"computeshare-server/internal/data/ent"
	"computeshare-server/internal/data/ent/computeinstance"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/samber/lo"
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
	list, err := csr.data.db.ComputeInstance.Query().Where(computeinstance.OwnerEQ(owner)).All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(list, func(item *ent.ComputeInstance, _ int) *biz.ComputeInstance {
		return &biz.ComputeInstance{
			ID:     item.ID,
			Core:   item.Core,
			Memory: item.Memory,
		}
	}), err
}

func (crs *computeInstanceRepo) Create(ctx context.Context, in *biz.ComputeInstance) error {
	entity, err := crs.data.db.ComputeInstance.Create().
		SetOwner(in.Owner).
		SetName(in.Name).
		SetCore(in.Core).
		SetMemory(in.Memory).
		SetImage(in.Image).
		SetExpirationTime(in.ExpirationTime).
		SetStatus(in.Status).
		Save(ctx)

	if err == nil {
		in.ID = entity.ID
	}
	return err
}

func (crs *computeInstanceRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return crs.data.db.ComputeInstance.DeleteOneID(id).Exec(ctx)
}
