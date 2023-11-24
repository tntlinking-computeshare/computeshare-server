package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// GatewayPort holds the schema definition for the GatewayPort entity.
type GatewayPort struct {
	ent.Schema
}

// Fields of the GatewayPort.
func (GatewayPort) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("fk_gateway_id", uuid.UUID{}).Comment("gateway id"),
		field.Int64("port").Comment("端口号"),
		field.Bool("is_use").Comment("是否使用").Default(false),
	}
}

// Edges of the GatewayPort.
func (GatewayPort) Edges() []ent.Edge {
	return nil
}
