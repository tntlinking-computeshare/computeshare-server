package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("country_call_coding").NotEmpty().MaxLen(8),
		field.String("telephone_number").NotEmpty().MaxLen(50),
		field.String("password").NotEmpty(),
		field.Time("create_date").Default(time.Now),
		field.Time("last_login_date").Default(time.Now),
		field.String("name").NotEmpty().Comment("用户名"),
		field.String("icon").Comment("头像地址"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		// 唯一约束索引
		index.Fields("id").Unique(),
		index.Fields("country_call_coding", "telephone_number").Unique(),
	}
}
