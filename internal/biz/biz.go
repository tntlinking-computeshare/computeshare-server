package biz

import (
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewAgentUsecase,
	NewStorageUsecase,
	NewUserUsecase,
	NewComputeInstanceUsercase,
	NewScriptUseCase,
	NewNetworkMappingUseCase,
	NewTaskUseCase,
	NewDomainBindingUseCase,
	NewStorageS3UseCase,
	NewStorageProviderUseCase,
	NewOrderUseCase,
	NewCycleTransactionUseCase,
	NewCycleRenewalUseCase,
	NewDashboardUseCase,
)
