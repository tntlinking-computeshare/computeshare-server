package data

import (
	"context"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/cycle"
	"github.com/shopspring/decimal"

	"github.com/go-kratos/kratos/v2/log"
)

type cycleRepo struct {
	data *Data
	log  *log.Helper
}

// NewCycleRepo .
func NewCycleRepo(data *Data, logger log.Logger) biz.CycleRepo {
	return &cycleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c *cycleRepo) FindByUserID(ctx context.Context, userId uuid.UUID) (*biz.Cycle, error) {
	entity, err := c.data.getCycle(ctx).Query().Where(cycle.FkUserID(userId)).First(ctx)
	if err != nil {
		return nil, err
	}
	return c.toBiz(entity, 0), nil
}

func (c *cycleRepo) Update(ctx context.Context, entity *biz.Cycle) error {
	tx := c.data.getCycle(ctx)
	balance, _ := entity.Cycle.Float64()
	return tx.UpdateOneID(entity.ID).
		SetCycle(balance).
		SetUpdateTime(entity.UpdateTime).Exec(ctx)
}

func (c *cycleRepo) toBiz(p *ent.Cycle, _ int) *biz.Cycle {
	if p == nil {
		return nil
	}
	return &biz.Cycle{
		ID:         p.ID,
		FkUserId:   p.FkUserID,
		Cycle:      decimal.NewFromFloat(p.Cycle),
		CreateTime: p.CreateTime,
		UpdateTime: p.UpdateTime,
	}
}
