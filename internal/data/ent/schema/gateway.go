package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Gateway holds the schema definition for the Gateway entity.
type Gateway struct {
	ent.Schema
}

// Fields of the Gateway.
func (Gateway) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("name").NotEmpty().MaxLen(50),
		field.String("ip").Comment("网关ip"),
		field.Int32("port").Comment("端口号"),
		field.String("internal_ip").Comment("内网ip"),
	}
}

// Edges of the Gateway.
func (Gateway) Edges() []ent.Edge {
	return nil
}
