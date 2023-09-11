package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID            uuid.UUID
	Name          string
	Password      string
	CreateDate    time.Time
	LastLoginDate *time.Time
}

type UserRepo interface {
	ListUser(ctx context.Context) ([]*User, error)
	GetUser(ctx context.Context, id uuid.UUID) (*User, error)
	GetUserPassword(ctx context.Context, id uuid.UUID) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, id uuid.UUID, user *User) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type UserUsercase struct {
	repo   UserRepo
	logger log.Logger
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsercase {
	return &UserUsercase{
		repo:   repo,
		logger: logger,
	}
}

func (uc *UserUsercase) List(ctx context.Context) (ps []*User, err error) {
	ps, err = uc.repo.ListUser(ctx)
	if err != nil {
		return
	}
	return
}

func (uc *UserUsercase) Get(ctx context.Context, id uuid.UUID) (p *User, err error) {
	p, err = uc.repo.GetUser(ctx, id)
	if err != nil {
		return
	}
	return
}

func (uc *UserUsercase) Create(ctx context.Context, user *User) error {
	return uc.repo.CreateUser(ctx, user)
}

func (uc *UserUsercase) Update(ctx context.Context, id uuid.UUID, user *User) error {
	return uc.repo.UpdateUser(ctx, id, user)
}

func (uc *UserUsercase) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.repo.DeleteUser(ctx, id)
}
