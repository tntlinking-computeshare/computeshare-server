package data

import (
	"computeshare-server/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ur *userRepo) ListUser(ctx context.Context) ([]*biz.User, error) {
	ps, err := ur.data.db.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.User, 0)
	for _, p := range ps {
		rv = append(rv, &biz.User{
			ID:            p.ID,
			Name:          p.Name,
			CreateDate:    p.CreateDate,
			LastLoginDate: p.LastLoginDate,
		})
	}
	return rv, nil
}
func (ur *userRepo) GetUser(ctx context.Context, id uuid.UUID) (*biz.User, error) {
	p, err := ur.data.db.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &biz.User{
		ID:            p.ID,
		Name:          p.Name,
		CreateDate:    p.CreateDate,
		LastLoginDate: p.LastLoginDate,
	}, nil
}

func (ur *userRepo) GetUserPassword(ctx context.Context, id uuid.UUID) (*biz.User, error) {
	p, err := ur.data.db.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &biz.User{
		ID:            p.ID,
		Name:          p.Name,
		Password:      p.Password,
		CreateDate:    p.CreateDate,
		LastLoginDate: p.LastLoginDate,
	}, nil
}

func (ur *userRepo) CreateUser(ctx context.Context, user *biz.User) error {
	result, err := ur.data.db.User.
		Create().
		SetName(user.Name).
		SetPassword(user.Password).
		Save(ctx)

	user.ID = result.ID
	return err
}
func (ur *userRepo) UpdateUser(ctx context.Context, id uuid.UUID, user *biz.User) error {
	p, err := ur.data.db.User.Get(ctx, id)
	if err != nil {
		return err
	}
	_, err = p.Update().
		SetName(user.Name).
		SetLastLoginDate(*p.LastLoginDate).
		Save(ctx)
	return err
}
func (ur *userRepo) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return ur.data.db.User.DeleteOneID(id).Exec(ctx)
}
