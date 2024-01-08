package v1

import (
	"encoding/json"
	"errors"
)

func (task *QueueTaskVo) GetTaskParam() (any, error) {
	switch task.Cmd {
	case TaskCmd_VM_CREATE, TaskCmd_VM_DELETE, TaskCmd_VM_START,
		TaskCmd_VM_SHUTDOWN, TaskCmd_VM_RESTART, TaskCmd_VM_RECREATE:
		var vo ComputeInstanceTaskParamVO
		err := json.Unmarshal([]byte(task.GetParams()), &vo)
		return &vo, err

	case TaskCmd_NAT_PROXY_CREATE, TaskCmd_NAT_PROXY_DELETE,
		TaskCmd_NAT_VISITOR_CREATE, TaskCmd_NAT_VISITOR_DELETE:
		var vo NatNetworkMappingTaskParamVO
		err := json.Unmarshal([]byte(task.GetParams()), &vo)
		return &vo, err
	case TaskCmd_STORAGE_CREATE, TaskCmd_STORAGE_DELETE:
		var vo StorageSetupTaskParamVO
		err := json.Unmarshal([]byte(task.GetParams()), &vo)
		return &vo, err
	}
	return nil, errors.New("cannot issue command")
}
