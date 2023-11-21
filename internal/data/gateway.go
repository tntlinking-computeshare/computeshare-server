package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/samber/lo"
)

type GatewayRepo struct {
	data *Data
	log  *log.Helper
}

func NewGatewayRepo(data *Data, logger log.Logger) biz.GatewayRepo {
	return &GatewayRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *GatewayRepo) GetGateway(ctx context.Context, id uuid.UUID) (*biz.Gateway, error) {
	instance, err := repo.data.db.Gateway.Get(ctx, id)
	return repo.toBiz(instance, 0), err
}

func (repo *GatewayRepo) ListGateway(ctx context.Context) ([]*biz.Gateway, error) {
	list, err := repo.data.db.Gateway.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(list, repo.toBiz), nil
}

func (repo *GatewayRepo) toBiz(item *ent.Gateway, _ int) *biz.Gateway {
	if item == nil {
		return nil
	}
	return &biz.Gateway{
		ID:   item.ID,
		Name: item.Name,
		IP:   item.IP,
		Port: item.Port,
	}
}
