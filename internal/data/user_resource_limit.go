package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/userresourcelimit"
)

type userResourceLimitRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserResourceLimitRepo(data *Data, logger log.Logger) biz.UserResourceLimitRepo {
	return &userResourceLimitRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *userResourceLimitRepo) Create(ctx context.Context, userRL *biz.UserResourceLimit) (*biz.UserResourceLimit, error) {
	save, err := repo.data.GetUserResourceLimit(ctx).Create().
		SetFkUserID(userRL.FkUserID).Save(ctx)
	if err != nil {
		return nil, err
	}
	return repo.toBiz(save, 0), nil
}
func (repo *userResourceLimitRepo) ExistsByUserId(ctx context.Context, userId uuid.UUID) bool {
	exists, err := repo.data.GetUserResourceLimit(ctx).Query().Where(userresourcelimit.FkUserID(userId)).Exist(ctx)
	if err != nil {
		return false
	}
	return exists
}
func (repo *userResourceLimitRepo) GetByUserId(ctx context.Context, userId uuid.UUID) (*biz.UserResourceLimit, error) {
	exists := repo.ExistsByUserId(ctx, userId)
	var item *ent.UserResourceLimit
	var err error
	if exists {
		item, err = repo.data.GetUserResourceLimit(ctx).Query().Where(userresourcelimit.FkUserID(userId)).First(ctx)
	} else {
		item, err = repo.data.GetUserResourceLimit(ctx).Create().
			SetFkUserID(userId).Save(ctx)
	}
	if err != nil {
		return nil, err
	}
	return repo.toBiz(item, 0), nil
}
func (repo *userResourceLimitRepo) Update(ctx context.Context, id uuid.UUID, limit *biz.UserResourceLimit) error {
	return repo.data.GetUserResourceLimit(ctx).UpdateOneID(id).SetMaxCPU(limit.MaxCPU).SetMaxMemory(limit.MaxMemory).SetMaxNetworkMapping(limit.MaxNetworkMapping).Exec(ctx)
}

func (repo *userResourceLimitRepo) toBiz(item *ent.UserResourceLimit, _ int) *biz.UserResourceLimit {
	if item == nil {
		return nil
	}

	return &biz.UserResourceLimit{
		ID:                item.ID,
		FkUserID:          item.FkUserID,
		MaxCPU:            item.MaxCPU,
		MaxMemory:         item.MaxMemory,
		MaxNetworkMapping: item.MaxNetworkMapping,
	}
}
