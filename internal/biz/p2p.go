package biz

import (
	"context"
	"fmt"
	"github.com/fatedier/frp/client"
	"github.com/fatedier/frp/pkg/config"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type P2pClient struct {
}

func NewP2pClient() *P2pClient {

	return &P2pClient{}
}

func (c *P2pClient) ForwardWithRandomPort(id string) (string, string, error) {

	cfgFilePath := "/Users/mohaijiang/app/frp/frpc.toml"
	fmt.Println("hello")
	cfg, pxyCfgs, visitorCfgs, _, err := config.LoadClientConfig(cfgFilePath)

	svr, err := client.NewService(cfg, pxyCfgs, visitorCfgs, cfgFilePath)
	if err != nil {
		panic(err)
	}

	shouldGracefulClose := cfg.Transport.Protocol == "kcp" || cfg.Transport.Protocol == "quic"
	// Capture the exit signal if we use kcp or quic.
	if shouldGracefulClose {
		go handleTermSignal(svr)
	}

	_ = svr.Run(context.Background())

	return "", "", nil
}

func handleTermSignal(svr *client.Service) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	svr.GracefulClose(500 * time.Millisecond)
}
