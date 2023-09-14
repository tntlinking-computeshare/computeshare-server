package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ComputeImage holds the schema definition for the ComputeImage entity.
type ComputeImage struct {
	ent.Schema
}

// Fields of the ComputeImage.
func (ComputeImage) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id"),
		field.String("name").NotEmpty().Comment("显示名"),
		field.String("image").NotEmpty().Comment("镜像名"),
		field.String("tag").NotEmpty().Comment("版本名"),
		field.Int32("port").Comment("端口号"),
	}
}

// Edges of the ComputeImage.
func (ComputeImage) Edges() []ent.Edge {
	return nil
}

func (ComputeImage) Indexes() []ent.Index {
	return []ent.Index{
		// 索引
		index.Fields("id").Unique(),
	}
}
