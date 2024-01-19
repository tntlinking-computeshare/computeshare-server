package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
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
		SetHostname(agent.Hostname).
		SetMAC(agent.MAC).
		SetLastUpdateTime(time.Now()).
		SetIP(agent.IP).
		SetTotalCPU(agent.TotalCPU).
		SetTotalMemory(agent.TotalMemory).
		SetOccupiedCPU(agent.OccupiedCPU).
		SetOccupiedMemory(agent.OccupiedMemory).
		SetActive(agent.Active).
		SetIP(agent.IP).
		Save(ctx)
	if err != nil {
		return err
	}

	agent.ID = result.ID
	return err
}

func (ar *agentRepo) UpdateAgent(ctx context.Context, id uuid.UUID, agent *biz.Agent) error {
	_, err := ar.data.getAgent(ctx).UpdateOneID(id).
		SetHostname(agent.Hostname).
		SetMAC(agent.MAC).
		SetLastUpdateTime(time.Now()).
		SetIP(agent.IP).
		SetTotalCPU(agent.TotalCPU).
		SetTotalMemory(agent.TotalMemory).
		SetOccupiedCPU(agent.OccupiedCPU).
		SetOccupiedMemory(agent.OccupiedMemory).
		SetActive(agent.Active).
		SetIP(agent.IP).
		Save(ctx)
	return err
}

func (ar *agentRepo) UpdateAgentStatus(ctx context.Context, id uuid.UUID, status bool) error {

	_, err := ar.data.getAgent(ctx).UpdateOneID(id).
		SetActive(status).
		SetLastUpdateTime(time.Now()).
		Save(ctx)

	return err
}

func (ar *agentRepo) DeleteAgent(ctx context.Context, id uuid.UUID) error {
	return ar.data.getAgent(ctx).DeleteOneID(id).Exec(ctx)
}

func (ar *agentRepo) FindByMac(ctx context.Context, mac string) (*biz.Agent, error) {

	p, err := ar.data.getAgent(ctx).Query().Where(agent.MACEQ(mac)).First(ctx)

	return ar.toBiz(p, 0), err
}

func (ar *agentRepo) toBiz(p *ent.Agent, _ int) *biz.Agent {
	if p == nil {
		return &biz.Agent{}
	}
	return &biz.Agent{
		ID:       p.ID,
		MAC:      p.MAC,
		Hostname: p.Hostname,
		// 总cpu数
		TotalCPU: p.TotalCPU,
		// 总内存数
		TotalMemory: p.TotalMemory,
		// 占用的cpu
		OccupiedCPU: p.OccupiedCPU,
		// 占用的内存
		OccupiedMemory: p.OccupiedMemory,
		// ip地址
		IP:             p.IP,
		Active:         p.Active,
		LastUpdateTime: p.LastUpdateTime,
	}
}

func (ar *agentRepo) FindOneActiveAgent(ctx context.Context, cpu int, memory int) (*biz.Agent, error) {

	entitys, err := ar.data.getAgent(ctx).Query().
		Where(agent.Active(true)).
		All(ctx)

	if err != nil {
		return nil, err
	}

	for _, entity := range entitys {
		balanceCpu := entity.TotalCPU - entity.OccupiedCPU
		if balanceCpu-int32(cpu) < 0 {
			continue
		}
		balanceMemory := entity.TotalMemory - entity.OccupiedMemory
		if balanceMemory-int32(memory) < 0 {
			continue
		}

		return ar.toBiz(entity, 0), nil
	}

	return nil, errors.New(400, "RESOURCE_INSUFFICIENT", "计算资源不足，请联系管理员")
}
