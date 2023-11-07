package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	jwt2 "github.com/golang-jwt/jwt/v4"
	agentV1 "github.com/mohaijiang/computeshare-server/api/agent/v1"
	computeV1 "github.com/mohaijiang/computeshare-server/api/compute/v1"
	queueTaskV1 "github.com/mohaijiang/computeshare-server/api/queue/v1"
	systemv1 "github.com/mohaijiang/computeshare-server/api/system/v1"
	"github.com/mohaijiang/computeshare-server/internal/conf"
	"github.com/mohaijiang/computeshare-server/internal/global"
	"github.com/mohaijiang/computeshare-server/internal/service"
)

func NewWhiteListMatcher() selector.MatchFunc {

	whiteList := make(map[string]struct{})
	whiteList["/api.system.v1.User/CreateUser"] = struct{}{}
	whiteList["/api.system.v1.User/Login"] = struct{}{}
	whiteList["/api.system.v1.User/LoginWithValidateCode"] = struct{}{}
	whiteList["/api.system.v1.User/SendValidateCode"] = struct{}{}
	whiteList["/api.agent.v1.Agent/CreateAgent"] = struct{}{}
	whiteList["/api.agent.v1.Agent/ListAgentInstance"] = struct{}{}
	whiteList["/api.agent.v1.Agent/ReportInstanceStatus"] = struct{}{}
	whiteList["/github.com.mohaijiang.api.queue.v1.QueueTask/GetAgentTask"] = struct{}{}
	whiteList["/github.com.mohaijiang.api.queue.v1.QueueTask/UpdateAgentTask"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server,
	ac *conf.Auth,
	agenter *service.AgentService,
	queueTaskService *service.QueueTaskService,
	storageService *service.StorageService,
	userService *service.UserService,
	instanceService *service.ComputeInstanceService,
	powerService *service.ComputePowerService,
	job *service.CronJob,
	logger log.Logger) *http.Server {

	jetMiddleware := selector.Server(
		jwt.Server(func(token *jwt2.Token) (interface{}, error) {
			return []byte(ac.ApiKey), nil
		}, jwt.WithSigningMethod(jwt2.SigningMethodHS256), jwt.WithClaims(func() jwt2.Claims {
			return &global.ComputeServerClaim{}
		})),
	).Match(NewWhiteListMatcher()).Build()

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			jetMiddleware,
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	openAPIhandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIhandler)
	agentV1.RegisterAgentHTTPServer(srv, agenter)
	computeV1.RegisterStorageHTTPServer(srv, storageService)
	computeV1.RegisterComputeInstanceHTTPServer(srv, instanceService)
	computeV1.RegisterComputePowerHTTPServer(srv, powerService)
	systemv1.RegisterUserHTTPServer(srv, userService)
	queueTaskV1.RegisterQueueTaskHTTPServer(srv, queueTaskService)

	srv.Route("/").POST("/v1/storage/upload", computeV1.Storage_UploadFile_Extend_HTTP_Handler(storageService))
	srv.Route("/").POST("/v1/storage/download", computeV1.Storage_DownloadFile_Extend_HTTP_Handler(storageService))
	srv.Route("/").POST("/v1/compute-power/upload/script", computeV1.Compute_Power_UploadSceipt_Extend_HTTP_Handler(powerService))
	srv.Route("/").POST("/v1/compute-power/download", computeV1.Compute_Powere_DownloadResult_Extend_HTTP_Handler(powerService))
	srv.HandleFunc("/v1/vm/terminal", instanceService.Terminal)

	job.StartJob()

	return srv
}
