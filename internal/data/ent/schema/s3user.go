package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// S3User holds the schema definition for the S3User entity.
type S3User struct {
	ent.Schema
}

// Fields of the S3User.
func (S3User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("fk_user_id", uuid.UUID{}).Comment("用户id"),
		field.String("access_key").MaxLen(50).Comment("accessKey"),
		field.String("secret_key").MaxLen(50).Comment("secretKey"),
	}
}

// Edges of the S3User.
func (S3User) Edges() []ent.Edge {
	return []ent.Edge{

		edge.From("buckets", S3Bucket.Type).
			Ref("user").
			// We add the "Required" method to the builder
			// to make this edge required on entity creation.
			// i.e. Card cannot be created without its owner.
			Required(),
	}
}
