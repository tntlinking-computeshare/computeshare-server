package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// NetworkMapping holds the schema definition for the NetworkMapping entity.
type NetworkMapping struct {
	ent.Schema
}

// Fields of the NetworkMapping.
func (NetworkMapping) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("name").NotEmpty().MaxLen(50),
		field.UUID("fk_gateway_id", uuid.UUID{}).Comment("gateway id"),
		field.Int("gateway_port").Comment("映射到网关的端口号"),
		field.Int("computer_port").Comment("需要映射的虚拟机端口号"),
		field.Int("status").Default(0).Comment(" 0 待开始 1 进行中 2 已完成, 3 失败"),
	}
}

// Edges of the NetworkMapping.
func (NetworkMapping) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("fk_computer_id", ComputeInstance.Type).
			Ref("networkMappings").Unique(),
	}
}
