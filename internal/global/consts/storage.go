package consts

type StorageProviderStatus int

const (
	StorageProviderStatus_NOT_RUN    StorageProviderStatus = iota // 未运行
	StorageProviderStatus_SETUPING                                // 启动中
	StorageProviderStatus_SETUP_FAIL                              //启动失败
	StorageProviderStatus_RUNNING                                 //运行中
	StorageProviderStatus_RUN_FAIL                                //运行失败
)
