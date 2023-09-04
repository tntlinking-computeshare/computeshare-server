package service

import (
	pb "computeshare-server/api/agent/v1"
	"computeshare-server/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGreeterService, NewAgentService)

type AgentService struct {
	pb.UnimplementedAgentServer

	log *log.Helper

	uc *biz.AgentUsecase
}
