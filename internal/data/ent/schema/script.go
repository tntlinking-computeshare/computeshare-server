package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Script holds the schema definition for the Script entity.
type Script struct {
	ent.Schema
}

// Fields of the Script.
func (Script) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id"),
		field.String("user_id"),
		field.Int32("task_number").Positive(),
		field.String("script_name").NotEmpty(),
		field.String("file_address"),
		field.Text("script_content").NotEmpty(),
		field.Int32("execute_state").Default(0).Comment("Latest execution status"),
		field.Text("execute_result").Comment("Latest execution results"),
		field.Time("create_time").Default(time.Now()),
		field.Time("update_time").Default(time.Now()),
	}
}

// Edges of the Script.
func (Script) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("scriptExecutionRecords", ScriptExecutionRecord.Type),
	}
}
