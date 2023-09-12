package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"time"
)

// Storage holds the schema definition for the Storage entity.
type Storage struct {
	ent.Schema
}

// Fields of the Storage.
func (Storage) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("owner").NotEmpty().MaxLen(50),
		field.Int32("type").Comment("0: DIR, 1:file").Default(0),
		field.String("name").NotEmpty().MaxLen(50),
		field.String("cid").Nillable().MaxLen(80),
		field.Int32("size").Comment("size"),
		field.Time("last_modify").Default(time.Now),
		field.String("parent_id").MaxLen(80),
	}
}

// Edges of the Storage.
func (Storage) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Storage) Indexes() []ent.Index {
	return []ent.Index{
		// 索引
		index.Fields("owner"),
	}
}
