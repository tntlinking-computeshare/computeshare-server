package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Employee holds the schema definition for the Employee entity.
type Employee struct {
	ent.Schema
}

// Fields of the Employee.
func (Employee) Fields() []ent.Field {
	return []ent.Field{
		//field.UUID("id",uuid.UUID{}).Default(uuid.New),
		field.String("name"),
		field.Int32("age").Optional(),
	}
}

// Edges of the Employee.
func (Employee) Edges() []ent.Edge {
	return nil
}
