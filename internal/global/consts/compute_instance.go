package consts

type InstanceStatus int8

const (
	// InstanceStatusCreating 创建中
	InstanceStatusCreating InstanceStatus = iota
	// InstanceStatusRunning  实例运行中
	InstanceStatusRunning
	// InstanceStatusStarting 实例启动中
	InstanceStatusStarting
	// InstanceStatusClosing 实例关闭中
	InstanceStatusClosing
	// InstanceStatusClosed 实例关闭
	InstanceStatusClosed
	// InstanceStatusRestarting 实例重启中
	InstanceStatusRestarting
	// InstanceStatusDeleting 实例删除中
	InstanceStatusDeleting
	// InstanceStatusDeleted 实例已删除
	InstanceStatusDeleted
	// InstanceStatusExpire 实例已过期
	InstanceStatusExpire
)
