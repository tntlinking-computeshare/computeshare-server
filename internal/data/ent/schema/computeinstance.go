package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/internal/global/consts"
)

// ComputeInstance holds the schema definition for the ComputeInstance entity.
type ComputeInstance struct {
	ent.Schema
}

// Fields of the ComputeInstance.
func (ComputeInstance) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("owner").NotEmpty(),
		field.String("name").NotEmpty(),
		field.String("core").NotEmpty(),
		field.String("memory").NotEmpty(),
		field.String("image").NotEmpty(),
		field.Int32("image_id").Comment("镜像id"),
		field.String("port").Optional().Comment("容器端口"),
		field.Time("expiration_time"),
		field.Int8("status").GoType(consts.InstanceStatus(0)).Comment("0: 启动中,1:运行中,2:连接中断, 3:过期"),
		field.String("container_id").Optional().Comment("容器id"),
		field.String("agent_id").Optional().Comment("p2p agent Id"),
		field.String("vnc_ip").Comment("vnc 内网链接ip"),
		field.Int32("vnc_port").Comment("vnc 内网链接端口号"),
		field.Text("docker_compose").Comment("初始化的docker容器"),
	}
}

// Edges of the ComputeInstance.
func (ComputeInstance) Edges() []ent.Edge {
	return nil
}

func (ComputeInstance) Indexes() []ent.Index {
	return []ent.Index{
		// 索引
		index.Fields("id").Unique(),
		index.Fields("owner"),
	}
}
