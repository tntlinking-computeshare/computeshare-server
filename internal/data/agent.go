package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/agent"
	"github.com/samber/lo"
	"time"
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
	ps, err := ar.data.getAgent(ctx).Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(ps, ar.toBiz), err
}

func (ar *agentRepo) GetAgent(ctx context.Context, id uuid.UUID) (*biz.Agent, error) {
	p, err := ar.data.getAgent(ctx).Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return ar.toBiz(p, 0), nil
}

func (ar *agentRepo) CreateAgent(ctx context.Context, agent *biz.Agent) error {
	result, err := ar.data.getAgent(ctx).
		Create().
		SetPeerID(agent.PeerId).
		SetActive(agent.Active).
		Save(ctx)
	if err != nil {
		return err
	}

	agent.ID = result.ID
	return err
}

func (ar *agentRepo) UpdateAgent(ctx context.Context, id uuid.UUID, agent *biz.Agent) error {
	p, err := ar.data.getAgent(ctx).Get(ctx, id)
	if err != nil {
		return err
	}
	_, err = p.Update().
		SetActive(agent.Active).
		SetLastUpdateTime(time.Now()).
		Save(ctx)
	return err
}

func (ar *agentRepo) DeleteAgent(ctx context.Context, id uuid.UUID) error {
	return ar.data.getAgent(ctx).DeleteOneID(id).Exec(ctx)
}

func (ar *agentRepo) FindByPeerId(ctx context.Context, peerId string) (*biz.Agent, error) {

	p, err := ar.data.getAgent(ctx).Query().Where(agent.PeerIDEQ(peerId)).First(ctx)

	return ar.toBiz(p, 0), err
}

func (ar *agentRepo) toBiz(p *ent.Agent, _ int) *biz.Agent {
	if p == nil {
		return &biz.Agent{}
	}
	return &biz.Agent{
		ID:             p.ID,
		PeerId:         p.PeerID,
		Active:         p.Active,
		LastUpdateTime: p.LastUpdateTime,
	}
}

func (ar *agentRepo) FindOneActiveAgent(ctx context.Context, cpu string, memory string) (*biz.Agent, error) {

	entity, err := ar.data.getAgent(ctx).Query().Where(agent.Active(true)).First(ctx)
	return ar.toBiz(entity, 0), err
}
