package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"time"
)

// Task holds the schema definition for the Storage entity.
type Task struct {
	ent.Schema
}

// Fields of the Storage.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("agent_id").NotEmpty().MaxLen(50),
		field.Int32("cmd").Comment("  VM_CREATE = 0; // 创建虚拟机\n  VM_DELETE = 1;  // 删除虚拟机\n  VM_START = 2; // 启动虚拟机\n  VM_SHUTDOWN = 3;  //关闭虚拟机\n  VM_RESTART = 4;  //关闭虚拟机\n  VM_VNC_CONNECT = 5;  // vnc 连接\n  NAT_PROXY_CREATE = 6; // nat 代理创建\n  NAT_PROXY_DELETE = 7; // nat 代理删除\n  NAT_VISITOR_CREATE = 8; // nat 访问创建\n  NAT_VISITOR_DELETE = 9; // nat 访问删除").Default(0),
		field.String("params").Nillable().MaxLen(1024).Comment("执行参数，nat 网络类型对应 NatProxyCreateVO, 虚拟机类型对应 ComputeInstanceTaskParamVO"),
		field.Int("status").Comment("  CREATED = 0; //创建\n  EXECUTING = 1; //执行中\n  EXECUTED = 2 ; // 执行成功\n  FAILED = 3 ; // 执行失败"),
		field.Time("create_time").Default(time.Now),
	}
}

// Edges of the Storage.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Task) Indexes() []ent.Index {
	return []ent.Index{
		// 索引
		index.Fields("agent_id"),
		index.Fields("create_time"),
	}
}
