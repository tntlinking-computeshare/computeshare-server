// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/conf"
	"github.com/mohaijiang/computeshare-server/internal/data"
	"github.com/mohaijiang/computeshare-server/internal/server"
	"github.com/mohaijiang/computeshare-server/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, auth *conf.Auth, logger log.Logger) (*kratos.App, func(), error) {
	grpcServer := server.NewGRPCServer(confServer, logger)
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	agentRepo := data.NewAgentRepo(dataData, logger)
	computeInstanceRepo := data.NewComputeInstanceRepo(dataData, logger)
	agentUsecase := biz.NewAgentUsecase(agentRepo, computeInstanceRepo, logger)
	agentService := service.NewAgentService(agentUsecase, logger)
	taskRepo := data.NewTaskRepo(dataData, logger)
	networkMappingRepo := data.NewNetworkMappingRepo(dataData, logger)
	gatewayPortRepo := data.NewGatewayPortRepo(dataData, logger)
	gatewayRepo := data.NewGatewayRepo(dataData, gatewayPortRepo, logger)
	computeSpecRepo := data.NewComputeSpecRepo(dataData, logger)
	computeImageRepo := data.NewComputeImageRepo(dataData, logger)
	computeInstanceUsercase := biz.NewComputeInstanceUsercase(computeSpecRepo, computeInstanceRepo, computeImageRepo, agentRepo, taskRepo, gatewayRepo, gatewayPortRepo, logger)
	networkMappingUseCase := biz.NewNetworkMappingUseCase(networkMappingRepo, gatewayRepo, gatewayPortRepo, taskRepo, computeInstanceUsercase, logger)
	taskUseCase := biz.NewTaskUseCase(taskRepo, networkMappingUseCase, computeInstanceRepo, logger)
	queueTaskService := service.NewQueueTaskService(taskUseCase, logger)
	storageRepo := data.NewStorageRepo(dataData, logger)
	storagecase := biz.NewStorageUsecase(storageRepo, logger)
	shell := service.NewIpfShell(confData)
	storageService, err := service.NewStorageService(storagecase, shell, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUsercase := biz.NewUserUsecase(auth, userRepo, logger)
	userService := service.NewUserService(userUsercase, logger)
	computeInstanceService := service.NewComputeInstanceService(computeInstanceUsercase, logger)
	scriptRepo := data.NewScriptRepo(dataData, logger)
	scriptExecutionRecordRepo := data.NewScriptExecutionRecordRepo(dataData, logger)
	scriptUseCase := biz.NewScriptUseCase(scriptRepo, scriptExecutionRecordRepo, agentRepo, logger)
	computePowerService, err := service.NewComputePowerService(scriptUseCase, shell, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	domainBindingRepository := data.NewDomainBindingRepository(dataData, logger)
	domainBindingUseCase, err := biz.NewDomainBindingUseCase(domainBindingRepository, networkMappingRepo, gatewayRepo, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	networkMappingService := service.NewNetworkMappingService(networkMappingUseCase, domainBindingUseCase, logger)
	domainBindingService := service.NewDomainBindingService(domainBindingUseCase, networkMappingUseCase)
	storageProviderRepo := data.NewStorageProviderRepo(dataData, logger)
	storageProviderUseCase := biz.NewStorageProviderUseCase(logger, storageProviderRepo, agentRepo, gatewayPortRepo, networkMappingRepo, gatewayRepo, taskRepo, networkMappingUseCase)
	storageProviderService := service.NewStorageProviderService(storageProviderUseCase)
	cronJob := service.NewCronJob(computeInstanceUsercase, agentUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, auth, agentService, queueTaskService, storageService, userService, computeInstanceService, computePowerService, networkMappingService, domainBindingService, storageProviderService, cronJob, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
