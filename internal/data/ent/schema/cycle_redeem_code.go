package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CycleRedeemCode holds the schema definition for the Cycle entity.
type CycleRedeemCode struct {
	ent.Schema
}

// Fields of the Cycle.
func (CycleRedeemCode) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("fk_user_id", uuid.UUID{}).Comment("用户id"),
		field.String("redeem_code").Unique().Comment("兑换码"),
		field.Float("cycle").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)",
		}).Comment("兑换码对应的周期"),
		field.Bool("state").Comment("状态"),
		field.Time("create_time").Comment("创建时间"),
		field.Time("use_time").Comment("使用"),
	}
}

// Edges of the Cycle.
func (CycleRedeemCode) Edges() []ent.Edge {
	//}
	return nil
}
