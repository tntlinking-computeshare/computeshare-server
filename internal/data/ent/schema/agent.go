package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"time"
)

// Agent holds the schema definition for the Agent entity.
type Agent struct {
	ent.Schema
}

// Fields of the Agent.
func (Agent) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("peer_id").Unique().NotEmpty(),
		field.Bool("active").Default(true).Comment("是否活动"),
		field.Time("last_update_time").Default(time.Now).UpdateDefault(time.Now).Comment("最后更新时间"),
	}
}

// Edges of the Agent.
func (Agent) Edges() []ent.Edge {
	return nil
}

func (Agent) Indexes() []ent.Index {
	return []ent.Index{
		// 唯一约束索引
		index.Fields("id").
			Unique(),
	}
}
