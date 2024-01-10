package data

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	global2 "github.com/mohaijiang/computeshare-server/api/global"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/cyclerenewal"
	"github.com/samber/lo"
	"time"
)

func NewCycleRenewalRepo(data *Data, logger log.Logger) biz.CycleRenewalRepo {

	return &cycleRenewalRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type cycleRenewalRepo struct {
	data *Data
	log  *log.Helper
}

func (r *cycleRenewalRepo) PageByUserId(ctx context.Context, userId uuid.UUID, page, size int) (*global2.Page[*biz.CycleRenewal], error) {

	tx := r.data.getCycleRenewal(ctx)

	data, err := tx.Query().Where(cyclerenewal.FkUserID(userId)).
		Order(cyclerenewal.ByDueTime(sql.OrderDesc())).
		Limit(size).
		Offset((page - 1) * size).
		All(ctx)

	if err != nil {
		return nil, err
	}

	total, err := tx.Query().Where(cyclerenewal.FkUserID(userId)).Count(ctx)
	if err != nil {
		return nil, err
	}

	pageData := &global2.Page[*ent.CycleRenewal]{
		Total: int64(total),
		Size:  int32(size),
		Page:  int32(page),
		Data:  data,
	}
	return global2.Map(pageData, r.toBiz), err
}

func (r *cycleRenewalRepo) toBiz(item *ent.CycleRenewal, _ int) *biz.CycleRenewal {
	if item == nil {
		return nil
	}

	return &biz.CycleRenewal{
		ID:       item.ID,
		FkUserID: item.FkUserID,
		// 资源ID
		ResourceID: item.ResourceID,
		// 资源类型
		ResourceType: item.ResourceType,
		// 产品名字
		ProductName: item.ProductName,
		// 产品描述
		ProductDesc: item.ProductDesc,
		// 状态
		State: item.State,
		// 延长时间
		ExtendDay: item.ExtendDay,
		// 额外的价格
		ExtendPrice: item.ExtendPrice,
		// 到期时间
		DueTime: item.DueTime,
		// 续费时间
		RenewalTime: item.RenewalTime,
		// 自动续费
		AutoRenewal: item.AutoRenewal,
	}
}

func (r *cycleRenewalRepo) GetById(ctx context.Context, renewalId uuid.UUID) (*biz.CycleRenewal, error) {
	tx := r.data.getCycleRenewal(ctx)
	data, err := tx.Get(ctx, renewalId)
	return r.toBiz(data, 0), err

}

func (r *cycleRenewalRepo) Create(ctx context.Context, renewal *biz.CycleRenewal) (*biz.CycleRenewal, error) {
	tx := r.data.getCycleRenewal(ctx)
	create := tx.Create().
		SetFkUserID(renewal.FkUserID).
		SetResourceID(renewal.ResourceID).
		SetResourceType(renewal.ResourceType).
		SetProductName(renewal.ProductName).
		SetProductDesc(renewal.ProductDesc).
		SetState(renewal.State).
		SetExtendDay(renewal.ExtendDay).
		SetExtendPrice(renewal.ExtendPrice).
		SetAutoRenewal(renewal.AutoRenewal)

	if renewal.DueTime != nil {
		create.SetDueTime(*renewal.DueTime)
	}
	if renewal.RenewalTime != nil {
		create.SetRenewalTime(*renewal.RenewalTime)
	}

	entity, err := create.Save(ctx)
	return r.toBiz(entity, 0), err
}

func (r *cycleRenewalRepo) Update(ctx context.Context, id uuid.UUID, renewal *biz.CycleRenewal) error {
	tx := r.data.getCycleRenewal(ctx)
	return tx.UpdateOneID(id).
		SetFkUserID(renewal.FkUserID).
		SetResourceID(renewal.ResourceID).
		SetResourceType(renewal.ResourceType).
		SetProductName(renewal.ProductName).
		SetProductDesc(renewal.ProductDesc).
		SetState(renewal.State).
		SetExtendDay(renewal.ExtendDay).
		SetExtendPrice(renewal.ExtendPrice).
		SetAutoRenewal(renewal.AutoRenewal).
		SetNillableRenewalTime(renewal.RenewalTime).
		SetNillableDueTime(renewal.DueTime).Exec(ctx)
}

func (r *cycleRenewalRepo) QueryDailyRenew(ctx context.Context) ([]*biz.CycleRenewal, error) {
	tx := r.data.getCycleRenewal(ctx)

	currentTime := time.Now()

	// 将当前时间设置为当天的 0 点整
	zeroTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 0, 0, 0, currentTime.Location())

	list, err := tx.Query().Where(cyclerenewal.AutoRenewalEQ(true), cyclerenewal.RenewalTime(zeroTime)).All(ctx)

	return lo.Map(list, r.toBiz), err
}
