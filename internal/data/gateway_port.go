package data

import (
	"context"
	"github.com/google/uuid"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/gatewayport"
)

type GatewayPortRepo struct {
	data *Data
	log  *log.Helper
}

func NewGatewayPortRepo(data *Data, logger log.Logger) biz.GatewayPortRepo {
	return &GatewayPortRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *GatewayPortRepo) CountGatewayPortByIsUsed(ctx context.Context, isUsed bool) ([]*biz.GatewayPortCount, error) {
	var counts []*biz.GatewayPortCount
	err := repo.data.db.GatewayPort.Query().
		Select(gatewayport.FieldFkGatewayID).
		Where(gatewayport.IsUse(isUsed)).
		GroupBy(gatewayport.FieldFkGatewayID).
		Aggregate(ent.Count()).
		Scan(ctx, &counts)
	return counts, err
}

func (repo *GatewayPortRepo) GetGatewayPortFirstByNotUsed(ctx context.Context, gatewayID string) (*biz.GatewayPort, error) {
	instance, err := repo.data.db.GatewayPort.Query().
		Where(gatewayport.FkGatewayID(gatewayID), gatewayport.IsUse(false)).
		Order(gatewayport.ByPort(sql.OrderAsc())).First(ctx)
	return repo.toBiz(instance, 0), err
}

func (repo *GatewayPortRepo) toBiz(item *ent.GatewayPort, _ int) *biz.GatewayPort {
	if item == nil {
		return nil
	}
	return &biz.GatewayPort{
		ID:          item.ID,
		FkGatewayID: item.FkGatewayID,
		Port:        item.Port,
		IsUse:       item.IsUse,
	}
}

func (repo *GatewayPortRepo) Update(ctx context.Context, gp *biz.GatewayPort) error {
	if gp == nil {
		return nil
	}

	return repo.data.db.GatewayPort.UpdateOneID(gp.ID).
		SetIsUse(gp.IsUse).
		Exec(ctx)
}

func (repo *GatewayPortRepo) GetGatewayPortByGatewayIdAndPort(ctx context.Context, gatewayId uuid.UUID, port int) (*biz.GatewayPort, error) {
	data, err := repo.data.db.GatewayPort.Query().
		Where(gatewayport.FkGatewayIDEQ(gatewayId), gatewayport.PortEQ(int64(port))).
		First(ctx)

	return repo.toBiz(data, 0), err
}
