package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/google/uuid"
	global2 "github.com/mohaijiang/computeshare-server/api/global"
	"github.com/mohaijiang/computeshare-server/internal/global"
	"github.com/shopspring/decimal"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// Cycle is a Cycle model.
type Cycle struct {
	ID         uuid.UUID       `json:"id,omitempty"`
	FkUserId   uuid.UUID       `json:"fkUserId,omitempty"`
	Cycle      decimal.Decimal `json:"cycle,omitempty"`
	CreateTime time.Time       `json:"createTime,omitempty"`
	UpdateTime time.Time       `json:"updateTime,omitempty"`
}

// CycleRepo is a Cycle repo.
type CycleRepo interface {
	FindByUserID(context.Context, uuid.UUID) (*Cycle, error)
}

type CycleOrderRepo interface {
	PageByUserId(ctx context.Context, userId uuid.UUID, page, size int) (*global2.Page[*CycleOrder], error)
}

// OrderUseCase is a cycle UseCase.
type OrderUseCase struct {
	cycleRepo CycleRepo
	orderRepo CycleOrderRepo
	log       *log.Helper
}

func (c *OrderUseCase) OrderList(ctx context.Context, page, size int32) (*global2.Page[*CycleOrder], error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New(400, "unauthorized", "unauthorized")
	}

	userId := claim.GetUserId()
	return c.orderRepo.PageByUserId(ctx, userId, int(page), int(size))

}

// NewOrderUseCase new a cycle UseCase.
func NewOrderUseCase(cycleRepo CycleRepo, orderRepo CycleOrderRepo, logger log.Logger) *OrderUseCase {
	return &OrderUseCase{
		cycleRepo: cycleRepo,
		orderRepo: orderRepo,
		log:       log.NewHelper(logger),
	}
}
