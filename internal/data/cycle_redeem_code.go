package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/cycleredeemcode"
	"github.com/shopspring/decimal"
	"time"
)

func NewCycleRedeemCodeRepo(data *Data, logger log.Logger) biz.CycleRedeemCodeRepo {

	return &cycleRedeemCodeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type cycleRedeemCodeRepo struct {
	data *Data
	log  *log.Helper
}

func (c *cycleRedeemCodeRepo) FindByRedeemCode(ctx context.Context, redeemCode string) (*biz.CycleRedeemCode, error) {
	cycleRedeemCode, err := c.data.getCycleRedeemCode(ctx).Query().Where(cycleredeemcode.RedeemCode(redeemCode)).First(ctx)
	if err != nil {
		return nil, err
	}
	return c.toBiz(cycleRedeemCode, 0), nil
}

func (c *cycleRedeemCodeRepo) Update(ctx context.Context, cycleRedeemCode *biz.CycleRedeemCode) error {
	tx := c.data.getCycleRedeemCode(ctx)
	first, err := tx.Query().Where(cycleredeemcode.RedeemCode(cycleRedeemCode.RedeemCode)).First(ctx)
	if err != nil {
		return err
	}
	return tx.UpdateOneID(first.ID).SetFkUserID(cycleRedeemCode.FkUserID).
		SetState(cycleRedeemCode.State).SetUseTime(time.Now()).Exec(ctx)
}

func (r *cycleRedeemCodeRepo) toBiz(p *ent.CycleRedeemCode, _ int) *biz.CycleRedeemCode {
	if p == nil {
		return nil
	}
	return &biz.CycleRedeemCode{
		ID:         p.ID,
		FkUserID:   p.FkUserID,
		RedeemCode: p.RedeemCode,
		Cycle:      decimal.NewFromFloat(p.Cycle),
		State:      p.State,
		CreateTime: p.CreateTime,
		UseTime:    p.UseTime,
	}
}
