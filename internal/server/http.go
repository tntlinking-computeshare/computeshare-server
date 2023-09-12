package server

import (
	agentV1 "computeshare-server/api/agent/v1"
	computeV1 "computeshare-server/api/compute/v1"
	v1 "computeshare-server/api/helloworld/v1"
	systemv1 "computeshare-server/api/system/v1"
	"computeshare-server/internal/conf"
	"computeshare-server/internal/global"
	"computeshare-server/internal/service"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	jwt2 "github.com/golang-jwt/jwt/v4"
)

func NewWhiteListMatcher() selector.MatchFunc {

	whiteList := make(map[string]struct{})
	whiteList["/api.system.v1.User/CreateUser"] = struct{}{}
	whiteList["/api.system.v1.User/Login"] = struct{}{}
	whiteList["/api.system.v1.User/LoginWithValidateCode"] = struct{}{}
	whiteList["/api.system.v1.User/SendValidateCode"] = struct{}{}
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
	greeter *service.GreeterService,
	agenter *service.AgentService,
	storageService *service.StorageService,
	userService *service.UserService,
	logger log.Logger) *http.Server {

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			selector.Server(
				jwt.Server(func(token *jwt2.Token) (interface{}, error) {
					return []byte(ac.ApiKey), nil
				}, jwt.WithSigningMethod(jwt2.SigningMethodHS256), jwt.WithClaims(func() jwt2.Claims {
					return &global.ComputeServerClaim{}
				})),
			).
				Match(NewWhiteListMatcher()).
				Build(),
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
	//opts = append(opts, http.Middleware(metadata.Server()))
	//opts = append(opts, http.Middleware(jwt.Server(func(token *jwtv4.Token) (interface{}, error) {
	//	return []byte(testKey), nil
	//})))
	srv := http.NewServer(opts...)
	openAPIhandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIhandler)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	agentV1.RegisterAgentHTTPServer(srv, agenter)
	computeV1.RegisterStorageHTTPServer(srv, storageService)
	systemv1.RegisterUserHTTPServer(srv, userService)

	srv.Route("/").POST("/v1/storage/upload", computeV1.Storage_UploadFile_Extend_HTTP_Handler(storageService))
	return srv
}
