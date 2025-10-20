package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/samber/lo"
	"github.com/tntlinking-computeshare/computeshare-server/internal/biz"
	"github.com/tntlinking-computeshare/computeshare-server/internal/data/ent"
	"github.com/tntlinking-computeshare/computeshare-server/internal/data/ent/computeimage"
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
	list, err := csr.data.getComputeImage(ctx).Query().Order(computeimage.BySort()).All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(list, csr.toBiz), err
}

func (csr *computeImageRepo) toBiz(item *ent.ComputeImage, _ int) *biz.ComputeImage {
	return &biz.ComputeImage{
		ID:          item.ID,
		Name:        item.Name,
		Image:       item.Image,
		Tag:         item.Tag,
		OsType:      item.OsType,
		OsVariant:   item.OsVariant,
		Filename:    item.Filename,
		DownloadURL: item.DownloadURL,
		Md5:         item.Md5,
		Arch:        item.Arch,
		BootType:    item.BootType,
	}
}

func (csr *computeImageRepo) Get(ctx context.Context, id int32) (*biz.ComputeImage, error) {
	entity, err := csr.data.getComputeImage(ctx).Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return csr.toBiz(entity, 0), nil
}
