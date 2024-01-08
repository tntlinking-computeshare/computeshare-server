package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Cycle holds the schema definition for the Cycle entity.
type Cycle struct {
	ent.Schema
}

// Fields of the Cycle.
func (Cycle) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("fk_user_id", uuid.UUID{}).Comment("用户id"),
		field.Float("cycle").SchemaType(map[string]string{
			dialect.MySQL: "decimal(10,2)",
		}),
		field.Time("create_time"),
	}
}

// Edges of the Cycle.
func (Cycle) Edges() []ent.Edge {
	//}
	return nil
}
