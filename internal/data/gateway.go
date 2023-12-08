package data

import (
	"context"
	"errors"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/networkmapping"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/samber/lo"
)

type GatewayRepo struct {
	data            *Data
	gatewayPortRepo biz.GatewayPortRepo
	log             *log.Helper
}

func NewGatewayRepo(data *Data, gatewayPortRepo biz.GatewayPortRepo, logger log.Logger) biz.GatewayRepo {
	return &GatewayRepo{
		data:            data,
		gatewayPortRepo: gatewayPortRepo,
		log:             log.NewHelper(logger),
	}
}

func (repo *GatewayRepo) GetGateway(ctx context.Context, id uuid.UUID) (*biz.Gateway, error) {
	instance, err := repo.data.getGateway(ctx).Get(ctx, id)
	return repo.toBiz(instance, 0), err
}

func (repo *GatewayRepo) ListGateway(ctx context.Context) ([]*biz.Gateway, error) {
	list, err := repo.data.getGateway(ctx).Query().All(ctx)
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

func (repo *GatewayRepo) FindInstanceSuitableGateway(ctx context.Context, instanceId uuid.UUID) (*biz.Gateway, error) {

	networkMapping, err := repo.data.getNetworkMapping(ctx).Query().Where(networkmapping.FkComputerIDEQ(instanceId)).First(ctx)
	if err == nil {
		return repo.GetGateway(ctx, networkMapping.FkGatewayID)
	}

	counts, err := repo.gatewayPortRepo.CountGatewayPortByIsUsed(ctx, false)
	if err != nil {
		return nil, err
	}
	if len(counts) == 0 {
		return nil, errors.New("no available gateway")
	}
	maxCount := lo.MaxBy(counts, func(item *biz.GatewayPortCount, max *biz.GatewayPortCount) bool {
		return item.Count > max.Count
	})

	if maxCount.Count == 0 {
		return nil, errors.New("no available gateway")
	}

	return repo.GetGateway(ctx, maxCount.FkGatewayID)

}
