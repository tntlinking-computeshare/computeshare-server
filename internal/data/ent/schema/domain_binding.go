package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// DomainBindding holds the schema definition for the DomainBindding entity.
type DomainBinding struct {
	ent.Schema
}

// Fields of the DomainBindding.
func (DomainBinding) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}).Comment("用户ID"),
		field.UUID("fk_compute_instance_id", uuid.UUID{}).Comment("实例ID"),
		field.UUID("fk_network_mapping_id", uuid.UUID{}).Comment("网络映射id"),
		field.String("name").NotEmpty().MaxLen(50).Comment("映射名"),
		field.String("domain").MaxLen(255).Comment("域名"),
		field.Int("gateway_port").Comment("映射到gateway的端口"),
		field.Time("create_time").Comment("创建时间"),
	}
}
