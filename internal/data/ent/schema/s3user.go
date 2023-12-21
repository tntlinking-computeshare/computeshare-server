package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
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
		field.Int8("type").Comment("类型"),
		field.String("access_key").MaxLen(50).Comment("accessKey").Unique(),
		field.String("secret_key").MaxLen(50).Comment("secretKey"),
		field.Time("create_time").Default(time.Now).Comment("创建时间"),
		field.Time("update_time").Default(time.Now).Comment("修改时间"),
	}
}

// Edges of the S3User.
func (S3User) Edges() []ent.Edge {
	//return []ent.Edge{
	//	edge.From("buckets", S3Bucket.Type).
	//		Ref("s3_user"),
	//}
	return nil
}
