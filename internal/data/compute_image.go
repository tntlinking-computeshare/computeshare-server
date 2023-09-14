package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/samber/lo"
)

type computeImageRepo struct {
	data *Data
	log  *log.Helper
}

func NewComputeImageRepo(data *Data, logger log.Logger) biz.ComputeImageRepo {
	return &computeImageRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (csr *computeImageRepo) List(ctx context.Context) ([]*biz.ComputeImage, error) {
	list, err := csr.data.db.ComputeImage.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(list, csr.toBiz), err
}

func (csr *computeImageRepo) toBiz(item *ent.ComputeImage, _ int) *biz.ComputeImage {
	return &biz.ComputeImage{
		ID:    int32(item.ID),
		Name:  item.Name,
		Image: item.Image,
		Tag:   item.Tag,
		Port:  item.Port,
	}
}

func (csr *computeImageRepo) Get(ctx context.Context, id int32) (*biz.ComputeImage, error) {
	entity, err := csr.data.db.ComputeImage.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return csr.toBiz(entity, 0), nil
}
