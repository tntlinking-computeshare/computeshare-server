package server

import (
	agentV1 "computeshare-server/api/agent/v1"
	computeV1 "computeshare-server/api/compute/v1"
	v1 "computeshare-server/api/helloworld/v1"
	"computeshare-server/internal/conf"
	"computeshare-server/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server,
	greeter *service.GreeterService,
	agenter *service.AgentService,
	storageService *service.StorageService,
	logger log.Logger) *http.Server {

	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
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
	srv.Route("/").POST("/v1/storage/upload", computeV1.Storage_UploadFile_Extend_HTTP_Handler(storageService))
	return srv
}
