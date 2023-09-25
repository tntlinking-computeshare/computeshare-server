package biz

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/mohaijiang/computeshare-server/third_party/p2p"
	"github.com/tj/assert"
	"os"
	"testing"
	"time"
)

func TestCheckForwardHealth(t *testing.T) {
	ipfsNode, cancel, err := p2p.RunDaemon()
	defer cancel()
	assert.NoError(t, err)
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", "id",
		"service.name", "Name",
		"service.version", "Version",
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	p2pUsecase := NewP2PUsecase(ipfsNode, logger)

	err = p2pUsecase.CheckForwardHealth("/x/ssh", "/p2p/12D3KooWCFGMbTGRTSppFYm5YQBn9awYwwgVDpHgvT1mqYMVSqvi")
	fmt.Println(err)
}

func TestTime(t *testing.T) {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixMilli())
}
