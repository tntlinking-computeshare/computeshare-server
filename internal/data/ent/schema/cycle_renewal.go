package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CycleRenewal holds the schema definition for the CycleRenewal entity.
type CycleRenewal struct {
	ent.Schema
}

// Fields of the CycleRenewal.
func (CycleRenewal) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("fk_user_id", uuid.UUID{}).Comment("用户id"),
		field.UUID("resource_id", uuid.UUID{}).Comment("资源ID"),
		field.Int("resource_type").Max(1).Comment("资源类型"),
		field.String("product_name").MaxLen(50).Comment("产品名字"),
		field.String("product_desc").MaxLen(200).Comment("产品描述"),
		field.Int8("state").Comment("状态"),
		field.Int8("extend_day").Comment("延长时间"),
		field.Float("extend_price").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)",
		}).Comment("额外的价格"),
		field.Time("due_time").Comment("到期时间"),
		field.Time("renewal_time").Comment("续费时间"),
		field.Bool("auto_renewal").Comment("自动续费"),
	}
}

// Edges of the CycleRenewal.
func (CycleRenewal) Edges() []ent.Edge {
	//}
	return nil
}
