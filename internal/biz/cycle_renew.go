package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	global2 "github.com/mohaijiang/computeshare-server/api/global"
	"github.com/mohaijiang/computeshare-server/internal/global"
)

type CycleRenewalRepo interface {
	PageByUserId(ctx context.Context, id uuid.UUID, page, size int) (*global2.Page[*CycleRenewal], error)
}

type CycleRenewalUseCase struct {
	log  *log.Helper
	repo CycleRenewalRepo
}

func NewCycleRenewalUseCase(logger log.Logger, repo CycleRenewalRepo) *CycleRenewalUseCase {
	return &CycleRenewalUseCase{
		log:  log.NewHelper(logger),
		repo: repo,
	}
}

func (c *CycleRenewalUseCase) PageByUser(ctx context.Context, page, size int32) (*global2.Page[*CycleRenewal], error) {
	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New(400, "unauthorized", "unauthorized")
	}

	userId := claim.GetUserId()
	return c.repo.PageByUserId(ctx, userId, int(page), int(size))
}
