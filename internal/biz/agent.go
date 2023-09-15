package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"time"
)

type Agent struct {
	ID             uuid.UUID
	PeerId         string
	Active         bool
	LastUpdateTime time.Time
}

type AgentRepo interface {
	//db
	ListAgent(ctx context.Context) ([]*Agent, error)
	GetAgent(ctx context.Context, id uuid.UUID) (*Agent, error)
	CreateAgent(ctx context.Context, agent *Agent) error
	UpdateAgent(ctx context.Context, id uuid.UUID, agent *Agent) error
	DeleteAgent(ctx context.Context, id uuid.UUID) error
	FindByPeerId(ctx context.Context, peerId string) (*Agent, error)
	FindOneActiveAgent(ctx context.Context, cpu string, memory string) (*Agent, error)
}

type AgentUsecase struct {
	repo   AgentRepo
	logger log.Logger
}

func NewAgentUsecase(repo AgentRepo, logger log.Logger) *AgentUsecase {
	return &AgentUsecase{repo: repo, logger: logger}
}

func (uc *AgentUsecase) List(ctx context.Context) (ps []*Agent, err error) {
	ps, err = uc.repo.ListAgent(ctx)
	if err != nil {
		return
	}
	return
}

func (uc *AgentUsecase) Get(ctx context.Context, id uuid.UUID) (p *Agent, err error) {
	p, err = uc.repo.GetAgent(ctx, id)
	if err != nil {
		return
	}
	return
}

func (uc *AgentUsecase) Create(ctx context.Context, agent *Agent) error {
	entity, err := uc.repo.FindByPeerId(ctx, agent.PeerId)
	if err != nil {
		return uc.repo.CreateAgent(ctx, agent)
	} else {
		return uc.repo.UpdateAgent(ctx, entity.ID, agent)
	}
}

func (uc *AgentUsecase) Update(ctx context.Context, id uuid.UUID, agent *Agent) error {
	return uc.repo.UpdateAgent(ctx, id, agent)
}

func (uc *AgentUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.repo.DeleteAgent(ctx, id)
}

func (uc *AgentUsecase) FindOneActiveAgent(ctx context.Context, cpu string, memory string) (*Agent, error) {
	return uc.repo.FindOneActiveAgent(ctx, cpu, memory)
}
