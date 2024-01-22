package data

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	global2 "github.com/mohaijiang/computeshare-server/api/global"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/cycleorder"
)

func NewCycleOrderRepo(data *Data, logger log.Logger) biz.CycleOrderRepo {

	return &cycleOrderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type cycleOrderRepo struct {
	data *Data
	log  *log.Helper
}

func (r *cycleOrderRepo) PageByUserId(ctx context.Context, userId uuid.UUID, page, size int) (*global2.Page[*biz.CycleOrder], error) {

	tx := r.data.getCycleOrder(ctx)

	data, err := tx.Query().Where(cycleorder.FkUserID(userId)).
		Order(cycleorder.ByCreateTime(sql.OrderDesc())).
		Limit(size).
		Offset((page - 1) * size).
		All(ctx)

	if err != nil {
		return nil, err
	}

	total, err := tx.Query().Where(cycleorder.FkUserID(userId)).Count(ctx)
	if err != nil {
		return nil, err
	}

	pageData := &global2.Page[*ent.CycleOrder]{
		Total: int64(total),
		Size:  int32(size),
		Page:  int32(page),
		Data:  data,
	}
	return global2.Map(pageData, r.toBiz), err
}

func (r *cycleOrderRepo) toBiz(item *ent.CycleOrder, _ int) *biz.CycleOrder {
	if item == nil {
		return nil
	}

	return &biz.CycleOrder{
		ID:          item.ID,
		FkUserID:    item.FkUserID,
		OrderNo:     item.OrderNo,
		ProductName: item.ProductName,
		ProductDesc: item.ProductDesc,
		Symbol:      item.Symbol,
		Cycle:       item.Cycle,
		CreateTime:  item.CreateTime,
	}
}

func (r *cycleOrderRepo) CheckOrderNoExists(ctx context.Context, orderNo string) bool {
	tx := r.data.getCycleOrder(ctx)
	exist, err := tx.Query().Where(cycleorder.OrderNo(orderNo)).Exist(ctx)
	if err != nil {
		return false
	}

	return exist
}

func (r *cycleOrderRepo) Create(ctx context.Context, order *biz.CycleOrder) (*biz.CycleOrder, error) {
	tx := r.data.getCycleOrder(ctx)
	entity, err := tx.Create().
		SetFkUserID(order.FkUserID).
		SetOrderNo(order.OrderNo).
		SetProductName(order.ProductName).
		SetProductDesc(order.ProductDesc).
		SetSymbol(order.Symbol).
		SetCycle(order.Cycle).
		SetCreateTime(order.CreateTime).
		SetNillableResourceID(order.ResourceId).
		Save(ctx)

	return r.toBiz(entity, 0), err
}
