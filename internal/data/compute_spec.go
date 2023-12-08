package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/samber/lo"
)

type computeSpecRepo struct {
	data *Data
	log  *log.Helper
}

func NewComputeSpecRepo(data *Data, logger log.Logger) biz.ComputeSpecRepo {
	return &computeSpecRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (csr *computeSpecRepo) List(ctx context.Context) ([]*biz.ComputeSpec, error) {
	list, err := csr.data.getComputeSpec(ctx).Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(list, csr.toBiz), err
}

func (csr *computeSpecRepo) toBiz(item *ent.ComputeSpec, _ int) *biz.ComputeSpec {
	return &biz.ComputeSpec{
		ID:     item.ID,
		Core:   item.Core,
		Memory: item.Memory,
	}
}

func (csr *computeSpecRepo) Get(ctx context.Context, id int32) (*biz.ComputeSpec, error) {
	entity, err := csr.data.getComputeSpec(ctx).Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return csr.toBiz(entity, 0), nil
}
