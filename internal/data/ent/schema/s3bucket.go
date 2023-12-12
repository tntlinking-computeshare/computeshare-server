package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// S3Bucket holds the schema definition for the S3Bucket entity.
type S3Bucket struct {
	ent.Schema
}

// Fields of the S3Bucket.
func (S3Bucket) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("bucket").MaxLen(50).Comment("bucketName").Unique(),
		field.Time("createdTime"),
	}
}

// Edges of the S3Bucket.
func (S3Bucket) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("s3_user", S3User.Type).
			Unique(),
	}
}
