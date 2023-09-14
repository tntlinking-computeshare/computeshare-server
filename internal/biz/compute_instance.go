package biz

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/global"
	"time"
)

type ComputeSpec struct {
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// Core holds the value of the "core" field.
	Core string `json:"core,omitempty"`
	// Memory holds the value of the "memory" field.
	Memory string `json:"memory,omitempty"`
}

type ComputeInstance struct {
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Owner holds the value of the "owner" field.
	Owner string `json:"owner,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Core holds the value of the "core" field.
	Core string `json:"core,omitempty"`
	// Memory holds the value of the "memory" field.
	Memory string `json:"memory,omitempty"`
	// Image holds the value of the "image" field.
	Image string `json:"image,omitempty"`
	// ExpirationTime holds the value of the "expiration_time" field.
	ExpirationTime time.Time `json:"expiration_time,omitempty"`
	// 0: 启动中,1:运行中,2:连接中断, 3:过期
	Status int8 `json:"status,omitempty"`
	// 容器id
	ContainerID string `json:"container_id,omitempty"`
	// p2p agent Id
	PeerID string `json:"peer_id,omitempty"`
}

const (
	InstanceStatusStarting int8 = iota
	InstanceStatusRunning
	InstanceStatusTerminal
	InstanceStatusExpire
)

type ComputeInstanceCreate struct {
	SpecId   int32
	ImageId  int32
	Duration int32
	Name     string
}

type ComputeImage struct {
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// 显示名
	Name string `json:"name,omitempty"`
	// 镜像名
	Image string `json:"image,omitempty"`
	// 版本名
	Tag string `json:"tag,omitempty"`
	// 端口号
	Port int32 `json:"port,omitempty"`
}

type ComputeSpecRepo interface {
	List(ctx context.Context) ([]*ComputeSpec, error)
	Get(ctx context.Context, id int32) (*ComputeSpec, error)
}

type ComputeInstanceRepo interface {
	List(ctx context.Context, owner string) ([]*ComputeInstance, error)
	Create(ctx context.Context, instance *ComputeInstance) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type ComputeImageRepo interface {
	List(ctx context.Context) ([]*ComputeImage, error)
	Get(ctx context.Context, id int32) (*ComputeImage, error)
}

type ComputeInstanceUsercase struct {
	specRepo     ComputeSpecRepo
	instanceRepo ComputeInstanceRepo
	imageRepo    ComputeImageRepo
	logger       log.Logger
}

func NewComputeInstanceUsercase(specRepo ComputeSpecRepo, instanceRepo ComputeInstanceRepo, imageRepo ComputeImageRepo, logger log.Logger) *ComputeInstanceUsercase {
	return &ComputeInstanceUsercase{
		specRepo:     specRepo,
		instanceRepo: instanceRepo,
		imageRepo:    imageRepo,
		logger:       logger,
	}
}

func (uc *ComputeInstanceUsercase) ListComputeSpec(ctx context.Context) ([]*ComputeSpec, error) {
	return uc.specRepo.List(ctx)
}

func (uc *ComputeInstanceUsercase) ListComputeImage(ctx context.Context) ([]*ComputeImage, error) {
	return uc.imageRepo.List(ctx)
}

func (uc *ComputeInstanceUsercase) Create(ctx context.Context, cic *ComputeInstanceCreate) (*ComputeInstance, error) {

	claim, ok := global.FromContext(ctx)
	if !ok {
		return nil, errors.New("cannot get user ID")
	}

	computeSpec, err := uc.specRepo.Get(ctx, cic.SpecId)
	if err != nil {
		return nil, err
	}
	computeImage, err := uc.imageRepo.Get(ctx, cic.ImageId)
	if err != nil {
		return nil, err
	}

	instance := &ComputeInstance{
		Owner:          claim.UserID,
		Name:           cic.Name,
		Core:           computeSpec.Core,
		Memory:         computeSpec.Memory,
		Image:          fmt.Sprintf("%s:%s", computeImage.Image, computeImage.Tag),
		ExpirationTime: time.Now().AddDate(0, int(cic.Duration), 0),
		Status:         InstanceStatusStarting,
	}

	err = uc.instanceRepo.Create(ctx, instance)
	return instance, err
}

func (uc *ComputeInstanceUsercase) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.instanceRepo.Delete(ctx, id)
}
