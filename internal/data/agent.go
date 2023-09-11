package data

import (
	"computeshare-server/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type agentRepo struct {
	data *Data
	log  *log.Helper
}

func NewAgentRepo(data *Data, logger log.Logger) biz.AgentRepo {
	return &agentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ar *agentRepo) ListAgent(ctx context.Context) ([]*biz.Agent, error) {
	ps, err := ar.data.db.Agent.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Agent, 0)
	for _, p := range ps {
		rv = append(rv, &biz.Agent{
			ID:   p.ID,
			Name: p.Name,
		})
	}
	return rv, nil
}

func (ar *agentRepo) GetAgent(ctx context.Context, id uuid.UUID) (*biz.Agent, error) {
	p, err := ar.data.db.Agent.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.Agent{
		ID:   p.ID,
		Name: p.Name,
	}, nil
}

func (ar *agentRepo) CreateAgent(ctx context.Context, agent *biz.Agent) error {
	result, err := ar.data.db.Agent.
		Create().
		SetName(agent.Name).
		Save(ctx)

	agent.ID = result.ID
	return err
}

func (ar *agentRepo) UpdateAgent(ctx context.Context, id uuid.UUID, agent *biz.Agent) error {
	p, err := ar.data.db.Agent.Get(ctx, id)
	if err != nil {
		return err
	}
	_, err = p.Update().
		SetName(agent.Name).
		Save(ctx)
	return err
}

func (ar *agentRepo) DeleteAgent(ctx context.Context, id uuid.UUID) error {
	return ar.data.db.Agent.DeleteOneID(id).Exec(ctx)
}
