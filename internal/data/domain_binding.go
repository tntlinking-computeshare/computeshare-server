package data

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/api/global"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"github.com/mohaijiang/computeshare-server/internal/data/ent/domainbinding"
	"github.com/samber/lo"
)

type domainRepositoryImpl struct {
	data *Data
	log  *log.Helper
}

func NewDomainBindingRepository(data *Data, logger log.Logger) biz.DomainBindingRepository {
	return &domainRepositoryImpl{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *domainRepositoryImpl) List(ctx context.Context, userId uuid.UUID) ([]*biz.DomainBinding, error) {
	list, err := r.data.getDomainBinding(ctx).Query().
		Where(domainbinding.UserIDEQ(userId)).
		Order(domainbinding.ByCreateTime(sql.OrderDesc())).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(list, r.toBiz), err
}

func (r *domainRepositoryImpl) PageQuery(ctx context.Context, userId, networkMappingId uuid.UUID, page, size int32) (*global.Page[*biz.DomainBinding], error) {
	list, err := r.data.getDomainBinding(ctx).Query().
		Where(domainbinding.UserIDEQ(userId), domainbinding.FkNetworkMappingIDEQ(networkMappingId)).
		Order(domainbinding.ByCreateTime(sql.OrderDesc())).
		Offset(int(page - 1)).Limit(int(size)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	total, err := r.data.getDomainBinding(ctx).Query().
		Where(domainbinding.UserIDEQ(userId)).
		Order(domainbinding.ByCreateTime(sql.OrderDesc())).Count(ctx)
	if err != nil {
		return nil, err
	}

	return &global.Page[*biz.DomainBinding]{
		Size:  size,
		Total: int64(total),
		Page:  page,
		Data:  lo.Map(list, r.toBiz),
	}, nil
}

func (r *domainRepositoryImpl) Save(ctx context.Context, entity *biz.DomainBinding) error {
	data, err := r.data.getDomainBinding(ctx).Create().
		SetCreateTime(entity.CreateTime).
		SetName(entity.Name).
		SetDomain(entity.Domain).
		SetFkComputeInstanceID(entity.FkComputeInstanceID).
		SetFkNetworkMappingID(entity.FkNetworkMappingID).
		SetGatewayPort(entity.GatewayPort).
		SetUserID(entity.UserID).
		Save(ctx)
	if err == nil {
		entity.ID = data.ID
	}
	return err
}
func (r *domainRepositoryImpl) Get(ctx context.Context, id uuid.UUID) (*biz.DomainBinding, error) {
	data, err := r.data.getDomainBinding(ctx).Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return r.toBiz(data, 0), nil
}
func (r *domainRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {

	return r.data.getDomainBinding(ctx).DeleteOneID(id).Exec(ctx)
}

func (r *domainRepositoryImpl) ListByNetworkMappingId(ctx context.Context, networkMappingId uuid.UUID) ([]*biz.DomainBinding, error) {
	list, err := r.data.getDomainBinding(ctx).Query().Where(domainbinding.FkNetworkMappingIDEQ(networkMappingId)).All(ctx)
	if err != nil {
		return nil, err
	}
	return lo.Map(list, r.toBiz), err
}

func (r *domainRepositoryImpl) toBiz(item *ent.DomainBinding, _ int) *biz.DomainBinding {
	if item == nil {
		return nil
	}
	return &biz.DomainBinding{
		ID: item.ID,
		// 用户ID
		UserID: item.UserID,
		// 实例ID
		FkComputeInstanceID: item.FkComputeInstanceID,
		// 网络映射id
		FkNetworkMappingID: item.FkNetworkMappingID,
		// 映射名
		Name: item.Name,
		// 域名
		Domain: item.Domain,
		// 映射到gateway的端口
		GatewayPort: item.GatewayPort,
		// 创建时间
		CreateTime: item.CreateTime,
	}
}
