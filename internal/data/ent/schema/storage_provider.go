package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/global/consts"
)

// StorageProvider holds the schema definition for the StorageProvider entity.
type StorageProvider struct {
	ent.Schema
}

// Fields of the StorageProvider.
func (StorageProvider) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("agent_id", uuid.UUID{}).Unique().Comment("agent 节点ID"),
		field.Int("status").GoType(consts.StorageProviderStatus(0)).Default(0).Comment("提供状态： 0：未运行，1：启动中，2： 启动失败，3： 运行中，4： 运行失败"),
		field.String("master_server").NotEmpty().Comment("存储节点master http地址"),
		field.String("public_ip").NotEmpty().Comment("存储volume的nat 映射IP"),
		field.Int32("public_port").Comment("存储节点volume的http nat映射端口"),
		field.Int32("grpc_port").Comment("存储节点volume的grpc nat映射端口"),
		field.Time("created_time").Comment("创建时间"),
	}
}

// Edges of the StorageProvider.
func (StorageProvider) Edges() []ent.Edge {
	return nil
}
