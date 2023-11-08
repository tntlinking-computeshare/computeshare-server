package v1

import (
	"encoding/json"
	"errors"
)

func (task *QueueTaskVo) GetTaskParam() (any, error) {
	switch task.Cmd {
	case TaskCmd_VM_CREATE:
	case TaskCmd_VM_DELETE:
	case TaskCmd_VM_START:
	case TaskCmd_VM_SHUTDOWN:
	case TaskCmd_VM_RESTART:
	case TaskCmd_VM_VNC_CONNECT:

		var vo ComputeInstanceTaskParamVO
		err := json.Unmarshal([]byte(task.GetParams()), &vo)
		return vo, err

	case TaskCmd_NAT_PROXY_CREATE:
	case TaskCmd_NAT_PROXY_DELETE:
	case TaskCmd_NAT_VISITOR_CREATE:
	case TaskCmd_NAT_VISITOR_DELETE:
		var vo NatNetworkMappingTaskParamVO
		err := json.Unmarshal([]byte(task.GetParams()), &vo)
		return vo, err
	}
	return nil, errors.New("")
}
