package consts

type ExecuteState int32

const (
	UnExecuted = iota + 1
	Executing
	Completed
	ExecutionFailed
	Canceled
)
