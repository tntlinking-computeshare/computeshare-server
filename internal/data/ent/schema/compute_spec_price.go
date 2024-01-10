package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ComputeSpecPrice holds the schema definition for the ComputeSpecPrice entity.
type ComputeSpecPrice struct {
	ent.Schema
}

// Fields of the ComputeSpecPrice.
func (ComputeSpecPrice) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id"),
		field.Int32("fk_compute_spec_id").Unique().Comment("资源规格id"),
		field.Int32("day").Default(30).Comment("天数"),
		field.Float32("price").Default(50000).Comment("此天数的价格"),
	}
}

// Edges of the ComputeSpecPrice.
func (ComputeSpecPrice) Edges() []ent.Edge {
	return nil
}

func (ComputeSpecPrice) Indexes() []ent.Index {
	return []ent.Index{
		// 索引
		index.Fields("id").Unique(),
	}
}
