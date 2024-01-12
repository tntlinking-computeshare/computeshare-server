package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	global2 "github.com/mohaijiang/computeshare-server/api/global"
	"github.com/mohaijiang/computeshare-server/internal/global"
)

type CycleTransactionRepo interface {
	PageByUserId(ctx context.Context, userId uuid.UUID, page, size int) (*global2.Page[*CycleTransaction], error)
	Create(ctx context.Context, transaction *CycleTransaction) (*CycleTransaction, error)
}

type CycleTransactionUseCase struct {
	log  *log.Helper
	repo CycleTransactionRepo
}

func NewCycleTransactionUseCase(logger log.Logger, repo CycleTransactionRepo) *CycleTransactionUseCase {
	return &CycleTransactionUseCase{
		log:  log.NewHelper(logger),
		repo: repo,
	}
}

func (c *CycleTransactionUseCase) PageByUser(ctx context.Context, page int32, size int32) (*global2.Page[*CycleTransaction], error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New(400, "unauthorized", "unauthorized")
	}

	userId := claim.GetUserId()
	return c.repo.PageByUserId(ctx, userId, int(page), int(size))
}
