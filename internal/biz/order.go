package biz

import (
	"context"
	"github.com/google/uuid"
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

// OrderUseCase is a cycle UseCase.
type OrderUseCase struct {
	cycleRepo CycleRepo
	log       *log.Helper
}

// NewOrderUseCase new a cycle UseCase.
func NewOrderUseCase(cycleRepo CycleRepo, logger log.Logger) *OrderUseCase {
	return &OrderUseCase{cycleRepo: cycleRepo, log: log.NewHelper(logger)}
}
