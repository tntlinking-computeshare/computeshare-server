package consts

type InstanceStatus int8

const (
	InstanceStatusCreating InstanceStatus = iota
	InstanceStatusRunning
	InstanceStatusStarting
	InstanceStatusClosing
	InstanceStatusClosed
	InstanceStatusRestarting
	InstanceStatusDeleting
	InstanceStatusDeleted
	InstanceStatusExpire
)
