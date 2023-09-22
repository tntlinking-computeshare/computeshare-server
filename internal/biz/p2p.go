package biz

import (
	"context"
	"errors"
	"fmt"
	"github.com/ipfs/kubo/core/commands"
	"github.com/ipfs/kubo/p2p"
	"github.com/libp2p/go-libp2p/core/peer"
	pstore "github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/p2p/protocol/ping"
	pb "github.com/mohaijiang/computeshare-client/api/network/v1"
	ma "github.com/multiformats/go-multiaddr"
	madns "github.com/multiformats/go-multiaddr-dns"
	"strconv"
	"strings"
	"time"

	"github.com/ipfs/kubo/core"
)

const P2PProtoPrefix = "/x/"

var resolveTimeout = 10 * time.Second

type P2PUsecase struct {
	node *core.IpfsNode
}

func NewP2PUsecase(ipfsNode *core.IpfsNode) *P2PUsecase {
	return &P2PUsecase{
		node: ipfsNode,
	}
}

func (s *P2PUsecase) Ping(ctx context.Context, peerId string) bool {
	addr, pid, err := commands.ParsePeerParam(peerId)
	if err != nil {
		return false
	}
	if addr != nil {
		s.node.Peerstore.AddAddr(pid, addr, time.Minute*2) // temporary
	}
	pings := ping.Ping(ctx, s.node.PeerHost, pid)
	r, ok := <-pings
	if !ok {
		fmt.Println("peerId ", pid.String(), " cannot reach")
	}
	fmt.Println(r.RTT)
	return ok
}

func (s *P2PUsecase) CreateForward(ctx context.Context, protoOpt string, listenOpt string, targetOpt string) error {

	proto := protocol.ID(protoOpt)

	listen, err := ma.NewMultiaddr(listenOpt)
	if err != nil {
		return err
	}

	targets, err := parseIpfsAddr(targetOpt)
	if err != nil {
		return err
	}

	allowCustom := false

	if !allowCustom && !strings.HasPrefix(string(proto), P2PProtoPrefix) {
		return errors.New("protocol name must be within '" + P2PProtoPrefix + "' namespace")
	}

	err = forwardLocal(s.node.Context(), s.node.P2P, s.node.Peerstore, proto, listen, targets)
	return err
}
func (s *P2PUsecase) CloseListen(ctx context.Context, protoOpt string, listenOpt string, targetOpt string) error {

	var proto protocol.ID
	proto = protocol.ID(protoOpt)

	var target, listen ma.Multiaddr
	var err error

	if listenOpt != "" {
		listen, err = ma.NewMultiaddr(listenOpt)
		if err != nil {
			return err
		}
	}

	if targetOpt != "" {
		target, err = ma.NewMultiaddr(targetOpt)
		if err != nil {
			return err
		}
	}

	match := func(listener p2p.Listener) bool {
		if proto != "" && proto != listener.Protocol() {
			return false
		}
		if listen != nil && !listen.Equal(listener.ListenAddress()) {
			return false
		}
		if target != nil && !target.Equal(listener.TargetAddress()) {
			return false
		}
		return true
	}

	done := s.node.P2P.ListenersLocal.Close(match)
	done += s.node.P2P.ListenersP2P.Close(match)
	fmt.Println("close connection : ", done)
	return nil
}

// checkPort checks whether target multiaddr contains tcp or udp protocol
// and whether the port is equal to 0
func (s *P2PUsecase) CheckPort(target ma.Multiaddr) error {
	// get tcp or udp port from multiaddr
	getPort := func() (string, error) {
		sport, _ := target.ValueForProtocol(ma.P_TCP)
		if sport != "" {
			return sport, nil
		}

		sport, _ = target.ValueForProtocol(ma.P_UDP)
		if sport != "" {
			return sport, nil
		}
		return "", fmt.Errorf("address does not contain tcp or udp protocol")
	}

	sport, err := getPort()
	if err != nil {
		return err
	}

	port, err := strconv.Atoi(sport)
	if err != nil {
		return err
	}

	if port == 0 {
		return fmt.Errorf("port can not be 0")
	}

	return nil
}

// parseIpfsAddr is a function that takes in addr string and return ipfsAddrs
func parseIpfsAddr(addr string) (*peer.AddrInfo, error) {
	multiaddr, err := ma.NewMultiaddr(addr)
	if err != nil {
		return nil, err
	}

	pi, err := peer.AddrInfoFromP2pAddr(multiaddr)
	if err == nil {
		return pi, nil
	}

	// resolve multiaddr whose protocol is not ma.P_IPFS
	ctx, cancel := context.WithTimeout(context.Background(), resolveTimeout)
	defer cancel()
	addrs, err := madns.Resolve(ctx, multiaddr)
	if err != nil {
		return nil, err
	}
	if len(addrs) == 0 {
		return nil, errors.New("fail to resolve the multiaddr:" + multiaddr.String())
	}
	var info peer.AddrInfo
	for _, addr := range addrs {
		taddr, id := peer.SplitAddr(addr)
		if id == "" {
			// not an ipfs addr, skipping.
			continue
		}
		switch info.ID {
		case "":
			info.ID = id
		case id:
		default:
			return nil, fmt.Errorf(
				"ambiguous multiaddr %s could refer to %s or %s",
				multiaddr,
				info.ID,
				id,
			)
		}
		info.Addrs = append(info.Addrs, taddr)
	}
	return &info, nil
}

// forwardLocal forwards local connections to a libp2p service
func forwardLocal(ctx context.Context, p *p2p.P2P, ps pstore.Peerstore, proto protocol.ID, bindAddr ma.Multiaddr, addr *peer.AddrInfo) error {
	ps.AddAddrs(addr.ID, addr.Addrs, pstore.TempAddrTTL)
	// TODO: return some info
	_, err := p.ForwardLocal(ctx, addr.ID, proto, bindAddr)
	return err
}

func (s *P2PUsecase) ListListen(ctx context.Context, req *pb.ListListenRequest) (*pb.ListListenReply, error) {
	output := &pb.ListListenReply{}

	s.node.P2P.ListenersLocal.Lock()
	for _, listener := range s.node.P2P.ListenersLocal.Listeners {
		output.Result = append(output.Result, &pb.ListenReply{
			Protocol:      string(listener.Protocol()),
			ListenAddress: listener.ListenAddress().String(),
			TargetAddress: listener.TargetAddress().String(),
		})
	}
	s.node.P2P.ListenersLocal.Unlock()

	s.node.P2P.ListenersP2P.Lock()
	for _, listener := range s.node.P2P.ListenersP2P.Listeners {
		output.Result = append(output.Result, &pb.ListenReply{
			Protocol:      string(listener.Protocol()),
			ListenAddress: listener.ListenAddress().String(),
			TargetAddress: listener.TargetAddress().String(),
		})
	}
	s.node.P2P.ListenersP2P.Unlock()
	return output, nil
}

// CheckForwardHealth check if the remote node is connected
func (c *P2PUsecase) CheckForwardHealth(protoOpt, target string) error {
	targets, err := parseIpfsAddr(target)
	proto := protocol.ID(protoOpt)
	if err != nil {
		return err
	}
	cctx, cancel := context.WithTimeout(context.Background(), time.Second*3) //TODO: configurable?
	defer cancel()
	stream, err := c.node.PeerHost.NewStream(cctx, targets.ID, proto)
	if err != nil {
		return err
	} else {
		_ = stream.Close()
		return nil
	}
}
