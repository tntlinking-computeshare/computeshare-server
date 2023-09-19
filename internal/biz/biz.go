package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewGreeterUsecase,
	NewAgentUsecase,
	NewStorageUsecase,
	NewUserUsecase,
	NewComputeInstanceUsercase,
	NewP2PUsecase,
)
