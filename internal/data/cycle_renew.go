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
