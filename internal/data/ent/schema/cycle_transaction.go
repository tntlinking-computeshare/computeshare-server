package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CycleTransaction holds the schema definition for the CycleTransaction entity.
type CycleTransaction struct {
	ent.Schema
}

// Fields of the CycleTransaction.
func (CycleTransaction) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("fk_cycle_id", uuid.UUID{}).Comment("cycleId"),
		field.UUID("fk_user_id", uuid.UUID{}).Comment("用户id"),
		field.UUID("fk_cycle_order_id", uuid.UUID{}).Comment("fk_cycle_order_id"),
		field.UUID("fk_cycle_recharge_id", uuid.UUID{}).Comment("fk_cycle_recharge_id"),
		field.String("operation").MaxLen(40).Comment("操作"),
		field.String("symbol").MaxLen(1).Comment("symbol"),
		field.Float("cycle").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)",
		}),
		field.Float("balance").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)",
		}).Comment("余额"),
		field.Time("operation_time").Comment("操作时间"),
	}
}

// Edges of the CycleTransaction.
func (CycleTransaction) Edges() []ent.Edge {
	//}
	return nil
}
