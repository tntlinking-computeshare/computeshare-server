package biz

import (
	"fmt"
	"github.com/mohaijiang/computeshare-server/third_party/p2p"
	"github.com/tj/assert"
	"testing"
)

func TestCheckForwardHealth(t *testing.T) {
	ipfsNode, cancel, err := p2p.RunDaemon()
	defer cancel()
	assert.NoError(t, err)
	p2pUsecase := NewP2PUsecase(ipfsNode)

	err = p2pUsecase.CheckForwardHealth("/x/ssh", "/p2p/12D3KooWCFGMbTGRTSppFYm5YQBn9awYwwgVDpHgvT1mqYMVSqvi")
	fmt.Println(err)
}
