package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// ScriptExecutionRecord holds the schema definition for the ScriptExecutionRecord entity.
type ScriptExecutionRecord struct {
	ent.Schema
}

// Fields of the ScriptExecutionRecord.
func (ScriptExecutionRecord) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id"),
		field.String("user_id"),
		field.Int32("fk_script_id").Positive(),
		field.Text("script_content").NotEmpty(),
		field.String("file_address"),
		field.Int32("execute_state"),
		field.Text("execute_result").NotEmpty(),
		field.Time("create_time").Default(time.Now()),
		field.Time("update_time").Default(time.Now()),
	}
}

// Edges of the ScriptExecutionRecord.
func (ScriptExecutionRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("script", Script.Type).
			Ref("scriptExecutionRecords").
			Unique(),
	}
}
