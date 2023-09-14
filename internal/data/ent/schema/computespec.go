package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ComputeSpec holds the schema definition for the ComputeSpec entity.
type ComputeSpec struct {
	ent.Schema
}

// Fields of the ComputeSpec.
func (ComputeSpec) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id"),
		field.String("core").NotEmpty(),
		field.String("memory").NotEmpty(),
	}
}

// Edges of the ComputeSpec.
func (ComputeSpec) Edges() []ent.Edge {
	return nil
}

func (ComputeSpec) Indexes() []ent.Index {
	return []ent.Index{
		// 索引
		index.Fields("id").Unique(),
	}
}
