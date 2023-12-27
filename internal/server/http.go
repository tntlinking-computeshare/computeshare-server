package server

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/gorilla/websocket"
	"github.com/mohaijiang/computeshare-server/internal/data"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	jwt2 "github.com/golang-jwt/jwt/v4"
	agentV1 "github.com/mohaijiang/computeshare-server/api/agent/v1"
	computeV1 "github.com/mohaijiang/computeshare-server/api/compute/v1"
	networkmappingV1 "github.com/mohaijiang/computeshare-server/api/network_mapping/v1"
	queueTaskV1 "github.com/mohaijiang/computeshare-server/api/queue/v1"
	systemv1 "github.com/mohaijiang/computeshare-server/api/system/v1"
	"github.com/mohaijiang/computeshare-server/internal/conf"
	"github.com/mohaijiang/computeshare-server/internal/global"
	"github.com/mohaijiang/computeshare-server/internal/service"
)

func NewWhiteListMatcher() selector.MatchFunc {

	whiteList := make(map[string]struct{})
	whiteList["/api.server.system.v1.User/Login"] = struct{}{}
	whiteList["/api.server.system.v1.User/LoginWithClient"] = struct{}{}
	whiteList["/api.server.system.v1.User/LoginWithValidateCode"] = struct{}{}
	whiteList["/api.server.system.v1.User/SendValidateCode"] = struct{}{}
	whiteList["/api.server.agent.v1.Agent/CreateAgent"] = struct{}{}
	whiteList["/api.server.agent.v1.Agent/ListAgentInstance"] = struct{}{}
	whiteList["/api.server.agent.v1.Agent/ReportInstanceStatus"] = struct{}{}
	whiteList["/api.server.queue.v1.QueueTask/GetAgentTask"] = struct{}{}
	whiteList["/api.server.queue.v1.QueueTask/UpdateAgentTask"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

func TransactionMiddleware(db *ent.Client) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if _, ok := transport.FromServerContext(ctx); ok {
				// Do something on entering
				tx, txErr := db.Tx(ctx)

				if txErr != nil {
					return nil, txErr
				}

				ctx = context.WithValue(ctx, "tx", tx)

				defer func() {
					// Do something on exiting
					if err != nil {
						_ = tx.Rollback()
					} else {
						_ = tx.Commit()
					}
				}()
			}
			reply, err = handler(ctx, req)

			return
		}
	}
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server,
	ac *conf.Auth,
	agenter *service.AgentService,
	queueTaskService *service.QueueTaskService,
	storageService *service.StorageService,
	storageS3Service *service.StorageS3Service,
	userService *service.UserService,
	instanceService *service.ComputeInstanceService,
	powerService *service.ComputePowerService,
	networkMappingService *service.NetworkMappingService,
	domainBindingService *service.DomainBindingService,
	storageProviderService *service.StorageProviderService,
	processService *service.ProcessService,
	job *service.CronJob,
	data *data.Data,
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
			logging.Server(logger),
			jetMiddleware,
			TransactionMiddleware(data.GetDB()),
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
	// vnc websocket
	srv.HandleFunc("/websockify", func(w http.ResponseWriter, r *http.Request) {
		wsHandler(w, r, instanceService, ac)
	})
	agentV1.RegisterAgentHTTPServer(srv, agenter)
	computeV1.RegisterStorageHTTPServer(srv, storageService)
	computeV1.RegisterStorageS3HTTPServer(srv, storageS3Service)
	computeV1.RegisterComputeInstanceHTTPServer(srv, instanceService)
	computeV1.RegisterComputePowerHTTPServer(srv, powerService)
	computeV1.RegisterStorageProviderHTTPServer(srv, storageProviderService)
	computeV1.RegisterProcessHTTPServer(srv, processService)
	systemv1.RegisterUserHTTPServer(srv, userService)
	networkmappingV1.RegisterNetworkMappingHTTPServer(srv, networkMappingService)
	queueTaskV1.RegisterQueueTaskHTTPServer(srv, queueTaskService)
	networkmappingV1.RegisterDomainBindingHTTPServer(srv, domainBindingService)

	srv.Route("/").POST("/v1/storage/upload", computeV1.Storage_UploadFile_Extend_HTTP_Handler(storageService))
	srv.Route("/").POST("/v1/storage/download", computeV1.Storage_DownloadFile_Extend_HTTP_Handler(storageService))
	srv.Route("/").POST("/v1/compute-power/upload/script", computeV1.Compute_Power_UploadSceipt_Extend_HTTP_Handler(powerService))
	srv.Route("/").POST("/v1/compute-power/download", computeV1.Compute_Powere_DownloadResult_Extend_HTTP_Handler(powerService))
	srv.Route("/").POST("/v1/compute-power/download", computeV1.Compute_Powere_DownloadResult_Extend_HTTP_Handler(powerService))
	srv.Route("/").POST("/v1/storage/{bucketName}/objects/upload", computeV1.Storage_S3_UploadFile_Extend_HTTP_Handler(storageS3Service))
	srv.Route("/").GET("/v1/storage/{bucketName}/objects/download", computeV1.Storage_S3_DownloadFile_Extend_HTTP_Handler(storageS3Service))

	job.StartJob()

	return srv
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 允许所有请求，也可以根据需要自定义检查
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request, instanceService *service.ComputeInstanceService, ac *conf.Auth) {
	conn, err := upgrader.Upgrade(w, r, nil)

	fmt.Println("======= 收到websocket请求 ======")
	fmt.Println("")
	fmt.Println("请求地址：", r.RequestURI)
	cookie, _ := r.Cookie("token")
	fmt.Println("请求Header:  cookie: ", cookie.Value)
	tokenString := cookie.Value

	token, err := jwt2.Parse(tokenString, func(token *jwt2.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt2.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(ac.ApiKey), nil
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	var userId string
	if claims, ok := token.Claims.(jwt2.MapClaims); ok {
		userId, _ = claims["UserID"].(string)
	} else {
		fmt.Println("===== websocket失败 ======")
		fmt.Println("       jwt token 验证失败         ")
		fmt.Println("=====              ======")
		return
	}

	//192.168.22.238:5915
	if err != nil {
		fmt.Println("===== websocket失败 ======")
		fmt.Println(err)
		fmt.Println("=====              ======")
		return
	}
	defer conn.Close()

	// 在这里添加 JWT Token 验证逻辑
	// 如果验证失败，可以关闭连接并返回错误

	consoleUrl, err := instanceService.GetInstanceConsole(r.Context(), r.FormValue("instanceId"), userId)
	if err != nil {
		fmt.Println("===== websocket失败 ======")
		fmt.Println(err)
		fmt.Println("=====              ======")
		return
	}

	// 连接到 noVNC 服务
	noVNCConn, _, err := websocket.DefaultDialer.Dial(consoleUrl, nil)
	if err != nil {
		fmt.Println("===== websocket失败 ======")
		fmt.Println(err)
		fmt.Println("=====              ======")
		return
	}
	defer noVNCConn.Close()

	// 开启 goroutine 将 noVNC 发送的消息转发给客户端
	go func() {
		for {
			_, message, err := noVNCConn.ReadMessage()
			if err != nil {
				fmt.Println("===== websocket失败 ======")
				fmt.Println(err)
				fmt.Println("=====              ======")
				return
			}
			err = conn.WriteMessage(websocket.BinaryMessage, message)
			if err != nil {
				fmt.Println("===== websocket失败 ======")
				fmt.Println(err)
				fmt.Println("=====              ======")
				return
			}
		}
	}()

	// 从客户端读取消息并发送到 noVNC 服务
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("===== websocket失败 ======")
			fmt.Println(err)
			fmt.Println("=====              ======")
			return
		}
		err = noVNCConn.WriteMessage(websocket.BinaryMessage, message)
		if err != nil {
			fmt.Println("===== websocket失败 ======")
			fmt.Println(err)
			fmt.Println("=====              ======")
			return
		}
	}
}
