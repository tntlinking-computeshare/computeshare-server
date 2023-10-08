package p2p

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	goipfsp2p "github.com/mohaijiang/go-ipfs-p2p"
	"testing"
	"time"
)

func TestCheckHealth(t *testing.T) {

	priv, _, err := crypto.GenerateKeyPair(crypto.RSA, 2048)
	skbytes, err := crypto.MarshalPrivateKey(priv)
	if err != nil {
		panic(err)
	}
	privateKey := base64.StdEncoding.EncodeToString(skbytes)

	swarmkey := "/key/swarm/psk/1.0.0/\n/base16/\n2108249f85354ed11ecf999a4500e9b616f71516b6c222ce630d14e434ef5562"
	bootstrap := "/ip4/61.172.179.6/tcp/32002/p2p/12D3KooWN89csfSa1Pa2u3HmanFe4Cx12qvuPkPd2uLcX7whMyZm"
	cli, err := goipfsp2p.NewP2pClient(14001, privateKey, swarmkey, []string{bootstrap})
	if err != nil {
		panic(err)
	}

	peerId := "QmVAJsf8zyWm7Siv5QvMXLGZ2MXRiMaK3Eb3RUgaj2ZSHf"
	_, _, err = cli.ForwardWithRandomPort(peerId)
	fmt.Println("forward: ", err)

	id, _ := peer.Decode(peerId)

	proto := protocol.ID("/x/ssh")

	ctx := context.Background()
	for {
		_, err = cli.RoutedHost.NewStream(ctx, id, proto)
		fmt.Println("err:", err)

		time.Sleep(time.Second)
	}
}
