package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserResourceLimit holds the schema definition for the UserResourceLimit entity.
type UserResourceLimit struct {
	ent.Schema
}

// Fields of the UserResourceLimit.
func (UserResourceLimit) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("fk_user_id", uuid.UUID{}).Default(uuid.New).Comment("用户id"),
		field.Int32("max_cpu").Default(20),
		field.Int32("max_memory").Default(40),
		field.Int32("max_network_mapping").Default(10),
	}
}

// Edges of the UserResourceLimit.
func (UserResourceLimit) Edges() []ent.Edge {
	return nil
}
