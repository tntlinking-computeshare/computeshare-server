package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Cycle holds the schema definition for the Cycle entity.
type CycleOrder struct {
	ent.Schema
}

// Fields of the Cycle.
func (CycleOrder) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("fk_user_id", uuid.UUID{}).Comment("用户id"),
		field.String("order_no").MinLen(16).MaxLen(36).Comment("订单编号"),
		field.String("product_name").MaxLen(50).Comment("产品名字"),
		field.String("product_desc").MaxLen(200).Comment("产品描述"),
		field.String("symbol").MaxLen(1).Comment("symbol"),
		field.Float("cycle").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)",
		}),
		field.String("resource_id").MaxLen(50).Optional().Nillable().Comment("资源id,可为空"),
		field.Time("create_time"),
	}
}

// Edges of the Cycle.
func (CycleOrder) Edges() []ent.Edge {
	//}
	return nil
}
