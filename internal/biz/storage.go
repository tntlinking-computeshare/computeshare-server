package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"time"
)

type Storage struct {
	ID uuid.UUID `json:"id,omitempty"`
	// Owner holds the value of the "owner" field.
	Owner string `json:"owner,omitempty"`
	// 0: DIR, 1:file
	Type int32 `json:"type,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Cid holds the value of the "cid" field.
	Cid string `json:"cid,omitempty"`
	// size
	Size int32 `json:"size,omitempty"`
	// LastModify holds the value of the "last_modify" field.
	LastModify time.Time `json:"last_modify,omitempty"`

	// ParentID holds the value of the "parent_id" field.
	ParentID string `json:"parent_id,omitempty"`
}

type StorageRepo interface {
	ListStorage(ctx context.Context, owner string, parentId string) ([]*Storage, error)
	GetStorage(ctx context.Context, id uuid.UUID) (*Storage, error)
	CreateStorage(ctx context.Context, storage *Storage) error
	UpdateStorage(ctx context.Context, id uuid.UUID, storage *Storage) error
	DeleteStorage(ctx context.Context, id uuid.UUID) error
}

type Storagecase struct {
	repo   StorageRepo
	logger log.Logger
}

func NewStorageUsecase(repo StorageRepo, logger log.Logger) *Storagecase {
	return &Storagecase{
		repo:   repo,
		logger: logger,
	}
}

func (uc *Storagecase) List(ctx context.Context, owner string, parentId string) (ps []*Storage, err error) {
	ps, err = uc.repo.ListStorage(ctx, owner, parentId)
	if err != nil {
		return
	}
	return
}

func (uc *Storagecase) Get(ctx context.Context, id uuid.UUID) (p *Storage, err error) {
	p, err = uc.repo.GetStorage(ctx, id)
	if err != nil {
		return
	}
	return
}

func (uc *Storagecase) Create(ctx context.Context, storage *Storage) error {
	return uc.repo.CreateStorage(ctx, storage)
}

func (uc *Storagecase) Update(ctx context.Context, id uuid.UUID, storage *Storage) error {
	return uc.repo.UpdateStorage(ctx, id, storage)
}

func (uc *Storagecase) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.repo.DeleteStorage(ctx, id)
}
