package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"time"
)

type Agent struct {
	ID uuid.UUID `json:"id,omitempty"`
	// mac 网卡地址
	MAC string `json:"mac,omitempty"`
	// 是否活动
	Active bool `json:"active,omitempty"`
	// 最后更新时间
	LastUpdateTime time.Time `json:"last_update_time,omitempty"`
	// 主机名
	Hostname string `json:"hostname,omitempty"`
	// 总cpu数
	TotalCPU int32 `json:"total_cpu,omitempty"`
	// 总内存数
	TotalMemory int32 `json:"total_memory,omitempty"`
	// 占用的cpu
	OccupiedCPU int32 `json:"occupied_cpu,omitempty"`
	// 占用的内存
	OccupiedMemory int32 `json:"occupied_memory,omitempty"`
	// ip地址
	IP string `json:"ip,omitempty"`
}

type AgentRepo interface {
	ListAgent(ctx context.Context) ([]*Agent, error)
	GetAgent(ctx context.Context, id uuid.UUID) (*Agent, error)
	CreateAgent(ctx context.Context, agent *Agent) error
	UpdateAgent(ctx context.Context, id uuid.UUID, agent *Agent) error
	UpdateAgentStatus(ctx context.Context, id uuid.UUID, status bool) error
	DeleteAgent(ctx context.Context, id uuid.UUID) error
	FindByMac(ctx context.Context, mac string) (*Agent, error)
	FindOneActiveAgent(ctx context.Context, cpu int, memory int) (*Agent, error)
}

type AgentUsecase struct {
	repo         AgentRepo
	log          *log.Helper
	instanceRepo ComputeInstanceRepo
}

func NewAgentUsecase(repo AgentRepo, instanceRepo ComputeInstanceRepo, logger log.Logger) *AgentUsecase {
	return &AgentUsecase{
		repo:         repo,
		instanceRepo: instanceRepo,
		log:          log.NewHelper(logger),
	}
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

func (uc *AgentUsecase) Create(ctx context.Context, agent *Agent) (uuid.UUID, error) {
	entity, err := uc.repo.FindByMac(ctx, agent.MAC)
	if err != nil {
		err := uc.repo.CreateAgent(ctx, agent)
		return agent.ID, err
	} else {
		err = uc.repo.UpdateAgent(ctx, entity.ID, agent)
		return entity.ID, err
	}
}

func (uc *AgentUsecase) Update(ctx context.Context, id uuid.UUID, agent *Agent) error {
	return uc.repo.UpdateAgent(ctx, id, agent)
}

func (uc *AgentUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.repo.DeleteAgent(ctx, id)
}

func (s *AgentUsecase) SyncAgentStatus() {
	ctx := context.Background()
	list, err := s.List(ctx)
	if err != nil {
		s.log.Error(err)
		return
	}

	// TODO ...
	fmt.Println(list)

}

func (uc *AgentUsecase) ListAgentInstance(ctx context.Context, peerId string) ([]*ComputeInstance, error) {
	return uc.instanceRepo.ListByAgentId(ctx, peerId)
}

func (uc *AgentUsecase) ReportInstanceStatus(ctx context.Context, instance *ComputeInstance) error {
	return uc.instanceRepo.Update(ctx, instance.ID, instance)
}
