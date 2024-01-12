package data

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	global2 "github.com/mohaijiang/computeshare-server/api/global"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/cycletransaction"
)

type cycleTransactionRepo struct {
	data *Data
	log  *log.Helper
}

func NewCycleTransactionRepo(data *Data, logger log.Logger) biz.CycleTransactionRepo {
	return &cycleTransactionRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c *cycleTransactionRepo) PageByUserId(ctx context.Context, userId uuid.UUID, page, size int) (*global2.Page[*biz.CycleTransaction], error) {
	tx := c.data.getCycleTransaction(ctx)

	data, err := tx.Query().Where(cycletransaction.FkUserID(userId)).
		Order(cycletransaction.ByOperationTime(sql.OrderDesc())).
		Limit(size).
		Offset((page - 1) * size).
		All(ctx)

	if err != nil {
		return nil, err
	}

	total, err := tx.Query().Where(cycletransaction.FkUserID(userId)).Count(ctx)
	if err != nil {
		return nil, err
	}

	pageData := &global2.Page[*ent.CycleTransaction]{
		Total: int64(total),
		Size:  int32(size),
		Page:  int32(page),
		Data:  data,
	}
	return global2.Map(pageData, c.toBiz), err
}

func (c *cycleTransactionRepo) toBiz(item *ent.CycleTransaction, _ int) *biz.CycleTransaction {
	if item == nil {
		return nil
	}

	return &biz.CycleTransaction{
		ID: item.ID,
		// cycleId
		FkCycleID: item.FkCycleID,
		// 用户id
		FkUserID: item.FkUserID,
		// fk_cycle_order_id
		FkCycleOrderID: item.FkCycleOrderID,
		// fk_cycle_recharge_id
		FkCycleRechargeID: item.FkCycleRechargeID,
		// 操作
		Operation: item.Operation,
		// symbol
		Symbol: item.Symbol,
		// Cycle holds the value of the "cycle" field.
		Cycle: item.Cycle,
		// 余额
		Balance: item.Balance,
		// 操作时间
		OperationTime: item.OperationTime,
	}
}

func (c *cycleTransactionRepo) Create(ctx context.Context, ct *biz.CycleTransaction) (*biz.CycleTransaction, error) {
	tx := c.data.getCycleTransaction(ctx)

	entity, err := tx.Create().
		SetFkCycleID(ct.FkCycleID).
		SetFkUserID(ct.FkUserID).
		SetFkCycleOrderID(ct.FkCycleOrderID).
		SetFkCycleRechargeID(ct.FkCycleRechargeID).
		SetOperation(ct.Operation).
		SetSymbol(ct.Symbol).
		SetCycle(ct.Cycle).
		SetBalance(ct.Balance).
		SetOperationTime(ct.OperationTime).Save(ctx)
	if err != nil {
		return nil, err
	}

	return c.toBiz(entity, 0), nil
}
