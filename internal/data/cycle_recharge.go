package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/cyclerecharge"
	"github.com/mohaijiang/computeshare-server/internal/global/consts"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

type cycleRechargeRepo struct {
	data *Data
	log  *log.Helper
}

func NewCycleRechargeRepo(data *Data, logger log.Logger) biz.CycleRechargeRepo {
	return &cycleRechargeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c *cycleRechargeRepo) FindByOutTradeNo(ctx context.Context, outTradeNo string) (*biz.CycleRecharge, error) {
	cycleRecharge, err := c.data.getCycleRecharge(ctx).Query().Where(cyclerecharge.OutTradeNo(outTradeNo)).First(ctx)
	if err != nil {
		return nil, err
	}
	return c.toBiz(cycleRecharge, 0), nil
}

func (c *cycleRechargeRepo) CreateCycleRecharge(ctx context.Context, bizCycleRecharge *biz.CycleRecharge) (*biz.CycleRecharge, error) {
	cycleRecharge, err := c.data.getCycleRecharge(ctx).Query().Where(cyclerecharge.OutTradeNo(bizCycleRecharge.OutTradeNo)).First(ctx)
	if cycleRecharge != nil && err == nil {
		return nil, errors.New("OutTradeNo 已经存在")
	}
	payAmount, _ := bizCycleRecharge.PayAmount.Round(2).Float64()
	buyCycle, _ := bizCycleRecharge.BuyCycle.Round(2).Float64()
	totalAmount, _ := bizCycleRecharge.TotalAmount.Round(2).Float64()
	recharge, err := c.data.getCycleRecharge(ctx).Create().SetFkUserID(bizCycleRecharge.FkUserID).SetOutTradeNo(bizCycleRecharge.OutTradeNo).
		SetAlipayTradeNo(bizCycleRecharge.AlipayTradeNo).SetRechargeChannel(bizCycleRecharge.RechargeChannel).SetRedeemCode(bizCycleRecharge.RedeemCode).
		SetState(string(consts.WaitBuyerPay)).SetPayAmount(payAmount).SetBuyCycle(buyCycle).SetTotalAmount(totalAmount).SetCreateTime(time.Now()).
		SetUpdateTime(time.Now()).Save(ctx)
	if err != nil {
		return nil, err
	}
	return c.toBiz(recharge, 0), nil
}

func (c *cycleRechargeRepo) UpdateCycleRecharge(ctx context.Context, bizCycleRecharge *biz.CycleRecharge) error {
	cycleRecharge, err := c.data.getCycleRecharge(ctx).Query().Where(cyclerecharge.OutTradeNo(bizCycleRecharge.OutTradeNo)).First(ctx)
	if cycleRecharge == nil && err != nil {
		return errors.New("OutTradeNo 不存在")
	}
	totalAmount, _ := bizCycleRecharge.TotalAmount.Round(2).Float64()
	return c.data.getCycleRecharge(ctx).UpdateOne(cycleRecharge).SetAlipayTradeNo(bizCycleRecharge.AlipayTradeNo).
		SetState(bizCycleRecharge.State).SetTotalAmount(totalAmount).SetUpdateTime(time.Now()).Exec(ctx)
}
func (c *cycleRechargeRepo) CountRechargeCycle(ctx context.Context) (decimal.Decimal, error) {
	cycleSum, err := c.data.getCycleRecharge(ctx).Query().Where(cyclerecharge.AlipayTradeNoNEQ(""), cyclerecharge.StateIn(string(consts.TradeSuccess), string(consts.TradeFinished))).
		Aggregate(ent.Sum(cyclerecharge.FieldBuyCycle)).Float64(ctx)
	if err == nil {
		return decimal.NewFromFloat(cycleSum), nil
	}
	if err != nil && strings.Contains(err.Error(), "converting NULL to float64 is unsupported") {
		cycleSum = 0.00
		return decimal.NewFromFloat(cycleSum), nil
	}
	return decimal.Decimal{}, err

}

func (c *cycleRechargeRepo) toBiz(p *ent.CycleRecharge, _ int) *biz.CycleRecharge {
	if p == nil {
		return nil
	}
	return &biz.CycleRecharge{
		ID:              p.ID,
		FkUserID:        p.FkUserID,
		OutTradeNo:      p.OutTradeNo,
		AlipayTradeNo:   p.AlipayTradeNo,
		RechargeChannel: p.RechargeChannel,
		RedeemCode:      p.RedeemCode,
		State:           p.State,
		PayAmount:       decimal.NewFromFloat(p.PayAmount),
		TotalAmount:     decimal.NewFromFloat(p.TotalAmount),
		BuyCycle:        decimal.NewFromFloat(p.BuyCycle),
		CreateTime:      p.CreateTime,
		UpdateTime:      p.UpdateTime,
	}
}
