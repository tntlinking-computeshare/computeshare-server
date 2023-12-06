package data

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/storageprovider"
	"github.com/samber/lo"
)

func NewStorageProviderRepo(data *Data, logger log.Logger) biz.StorageProviderRepo {

	return &storageProviderRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type storageProviderRepo struct {
	data *Data
	log  *log.Helper
}

func (r *storageProviderRepo) Create(ctx context.Context, item *biz.StorageProvider) (*biz.StorageProvider, error) {
	save, err := r.data.db.StorageProvider.Create().
		SetAgentID(item.AgentID).
		SetMasterServer(item.MasterServer).
		SetStatus(item.Status).
		SetPublicIP(item.PublicIP).
		SetPublicPort(item.PublicPort).
		SetGrpcPort(item.GrpcPort).
		SetCreatedTime(item.CreatedTime).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.toBiz(save, 0), nil
}
func (r *storageProviderRepo) List(ctx context.Context) ([]*biz.StorageProvider, error) {
	list, err := r.data.db.StorageProvider.Query().Order(storageprovider.ByCreatedTime(sql.OrderDesc())).All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(list, r.toBiz), err
}
func (r *storageProviderRepo) Get(ctx context.Context, id uuid.UUID) (*biz.StorageProvider, error) {
	item, err := r.data.db.StorageProvider.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return r.toBiz(item, 0), err
}
func (r *storageProviderRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.data.db.StorageProvider.DeleteOneID(id).Exec(ctx)
}

func (r *storageProviderRepo) QueryByAgentId(ctx context.Context, id uuid.UUID) (*biz.StorageProvider, error) {
	item, err := r.data.db.StorageProvider.Query().Where(storageprovider.AgentID(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	return r.toBiz(item, 0), err
}

func (r *storageProviderRepo) toBiz(item *ent.StorageProvider, _ int) *biz.StorageProvider {
	if item == nil {
		return nil
	}
	return &biz.StorageProvider{
		ID:           item.ID,
		AgentID:      item.AgentID,
		Status:       item.Status,
		MasterServer: item.MasterServer,
		PublicIP:     item.PublicIP,
		PublicPort:   item.PublicPort,
		GrpcPort:     item.GrpcPort,
		CreatedTime:  item.CreatedTime,
	}
}
