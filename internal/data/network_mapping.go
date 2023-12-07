package data

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/computeinstance"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/networkmapping"
	"github.com/samber/lo"
)

type NetworkMappingRepo struct {
	data *Data
	log  *log.Helper
}

func NewNetworkMappingRepo(data *Data, logger log.Logger) biz.NetworkMappingRepo {
	return &NetworkMappingRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *NetworkMappingRepo) CreateNetworkMapping(ctx context.Context, entity *biz.NetworkMapping) error {
	data, err := repo.data.db.NetworkMapping.Create().
		SetName(entity.Name).
		SetComputerPort(entity.ComputerPort).
		SetFkComputerID(entity.FkComputerID).
		SetFkGatewayID(entity.FkGatewayID).
		SetStatus(entity.Status).
		SetGatewayPort(entity.GatewayPort).
		SetFkUserID(entity.UserId).
		SetGatewayIP(entity.GatewayIP).
		Save(ctx)

	if err != nil {
		return err
	}
	entity.ID = data.ID
	return err
}

func (repo *NetworkMappingRepo) GetNetworkMapping(ctx context.Context, id uuid.UUID) (*biz.NetworkMapping, error) {
	instance, err := repo.data.db.NetworkMapping.Get(ctx, id)
	return repo.toBiz(instance, 0), err
}

func (repo *NetworkMappingRepo) DeleteNetworkMapping(ctx context.Context, id uuid.UUID) error {
	return repo.data.db.NetworkMapping.DeleteOneID(id).Exec(ctx)
}

func (repo *NetworkMappingRepo) PageNetworkMappingByUserID(ctx context.Context, userId uuid.UUID, page int32, size int32) ([]*biz.NetworkMapping, int32, error) {
	count, err := repo.data.db.NetworkMapping.Query().Select(networkmapping.FieldID).Where(networkmapping.FkUserID(userId)).Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	var offset int32
	if page > 0 {
		offset = (page - 1) * size
	} else {
		offset = page * size
	}
	list, err := repo.data.db.NetworkMapping.Query().Where(networkmapping.FkUserID(userId)).Order(networkmapping.ByComputerPort(sql.OrderAsc())).Offset(int(offset)).Limit(int(size)).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	return lo.Map(list, repo.toBiz), int32(count), nil
}

func (repo *NetworkMappingRepo) UpdateNetworkMapping(ctx context.Context, entity *biz.NetworkMapping) error {
	return repo.data.db.NetworkMapping.UpdateOneID(entity.ID).
		SetName(entity.Name).
		SetComputerPort(entity.ComputerPort).
		SetFkComputerID(entity.FkComputerID).
		SetFkGatewayID(entity.FkGatewayID).
		SetStatus(entity.Status).
		SetGatewayPort(entity.GatewayPort).
		SetGatewayIP(entity.GatewayIP).
		Exec(ctx)
}

func (repo *NetworkMappingRepo) toBiz(item *ent.NetworkMapping, _ int) *biz.NetworkMapping {
	if item == nil {
		return nil
	}

	var instanceName string
	instance, err := repo.data.db.ComputeInstance.Get(context.Background(), item.FkComputerID)
	if err == nil {
		instanceName = instance.Name
	}

	return &biz.NetworkMapping{
		ID:                   item.ID,
		Name:                 item.Name,
		FkComputerID:         item.FkComputerID,
		ComputerPort:         item.ComputerPort,
		ComputerInstanceName: instanceName,
		FkGatewayID:          item.FkGatewayID,
		GatewayPort:          item.GatewayPort,
		Status:               item.Status,
		UserId:               item.FkUserID,
		GatewayIP:            item.GatewayIP,
	}
}

func (repo *NetworkMappingRepo) QueryGatewayIdByAgentId(ctx context.Context, agentId uuid.UUID) (uuid.UUID, error) {
	computeInstances, err := repo.data.db.ComputeInstance.Query().Select(computeinstance.FieldID).Where(computeinstance.AgentID(agentId.String())).All(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	if len(computeInstances) == 0 {
		first, err := repo.data.db.Gateway.Query().First(ctx)
		if err != nil {
			return uuid.Nil, err
		}
		return first.ID, nil
	}

	computeInstanceIds := lo.Map(computeInstances, func(item *ent.ComputeInstance, index int) uuid.UUID {
		return item.ID
	})

	type networkMapingGroupByFkGatewayID struct {
		FkGatewayID uuid.UUID `json:"fk_gateway_id,omitempty"`
		Count       int       `json:"count"`
	}

	var v []networkMapingGroupByFkGatewayID

	err = repo.data.db.NetworkMapping.Query().
		Where(networkmapping.FkComputerIDIn(computeInstanceIds...)).
		GroupBy(networkmapping.FieldFkGatewayID).Aggregate(ent.Count()).Scan(ctx, &v)
	if err != nil {
		return uuid.Nil, err
	}

	if len(v) == 0 {
		first, err := repo.data.db.Gateway.Query().First(ctx)
		if err != nil {
			return uuid.Nil, err
		}
		return first.ID, nil
	}

	max := lo.MaxBy(v, func(a networkMapingGroupByFkGatewayID, b networkMapingGroupByFkGatewayID) bool {
		return a.Count > b.Count
	})
	return max.FkGatewayID, nil
}