package p2p

import (
	"context"
	"errors"
	"fmt"
	"github.com/ipfs/boxo/coreiface/options"
	"github.com/ipfs/kubo/config"
	"github.com/ipfs/kubo/core/commands"
	"github.com/ipfs/kubo/p2p"
	"github.com/ipfs/kubo/plugin/loader"
	"github.com/ipfs/kubo/repo/fsrepo"
	peer "github.com/libp2p/go-libp2p/core/peer"
	pstore "github.com/libp2p/go-libp2p/core/peerstore"
	protocol "github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/p2p/protocol/ping"
	pb "github.com/mohaijiang/computeshare-client/api/network/v1"
	ma "github.com/multiformats/go-multiaddr"
	madns "github.com/multiformats/go-multiaddr-dns"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	libp2p2 "github.com/ipfs/kubo/core/node/libp2p"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/ipfs/kubo/core"
)

const SwarmKey = "/key/swarm/psk/1.0.0/\n/base16/\n2108249f85354ed11ecf999a4500e9b616f71516b6c222ce630d14e434ef5562"
const P2PProtoPrefix = "/x/"

var resolveTimeout = 10 * time.Second

func init() {
	ipfsPath, err := fsrepo.BestKnownPath()
	plugins, err := loader.NewPluginLoader(ipfsPath)
	if err != nil {
		log.Errorf("error loading plugins: %s", err)
	}

	if err := plugins.Initialize(); err != nil {
		log.Errorf("error initializing plugins: %s", err)
	}

	if err := plugins.Inject(); err != nil {
		log.Errorf("error initializing plugins: %s", err)
	}
}

func RunDaemon() (*core.IpfsNode, func(), error) {

	ctx := context.Background()
	ipfsPath, err := fsrepo.BestKnownPath()

	if !fsrepo.IsInitialized(ipfsPath) {
		identity, err := config.CreateIdentity(os.Stdout, []options.KeyGenerateOption{
			options.Key.Type(options.Ed25519Key),
		})
		if err != nil {
			log.Error("create identity error : ", err)
			return nil, nil, err
		}
		conf, err := config.InitWithIdentity(identity)
		if err != nil {
			log.Error("InitWithIdentity error: ", err)
			return nil, nil, err
		}

		conf.Bootstrap = []string{}
		conf.Swarm.RelayClient.Enabled = config.True
		conf.Swarm.RelayService.Enabled = config.True
		err = fsrepo.Init(
			ipfsPath,
			conf,
		)
		if err != nil {
			log.Error("fsrepo  init fail : ", err)
			return nil, nil, err
		}
	}
	swarmKeyFile := filepath.Join(ipfsPath, "swarm.key")

	_, err = os.Lstat(swarmKeyFile)
	if err != nil {
		err = os.WriteFile(swarmKeyFile, []byte(SwarmKey), 0644)
		if err != nil {
			log.Error("init swarm.key fail", err)
			return nil, nil, err
		}
	}

	repo, err := fsrepo.Open(ipfsPath)
	if err != nil {
		log.Error("fsrepo is not initialization: ", err)
		return nil, nil, err
	}
	ncfg := &core.BuildCfg{
		Repo:                        repo,
		Permanent:                   true,
		Online:                      true,
		DisableEncryptedConnections: false,
		ExtraOpts: map[string]bool{
			"pubsub": false,
			"ipnsps": false,
		},
		Routing: libp2p2.DHTOption,
	}

	node, err := core.NewNode(ctx, ncfg)
	if err != nil {
		log.Error("error from node construction: ", err)
		return nil, nil, err
	}
	node.IsDaemon = true

	printSwarmAddrs(node)
	cleanup := func() {
		_ = node.Close()
	}
	return node, cleanup, nil
}

// printSwarmAddrs prints the addresses of the host
func printSwarmAddrs(node *core.IpfsNode) {
	if !node.IsOnline {
		fmt.Println("Swarm not listening, running in offline mode.")
		return
	}

	var lisAddrs []string
	ifaceAddrs, err := node.PeerHost.Network().InterfaceListenAddresses()
	if err != nil {
		log.Errorf("failed to read listening addresses: %s", err)
	}
	for _, addr := range ifaceAddrs {
		lisAddrs = append(lisAddrs, addr.String())
	}
	sort.Strings(lisAddrs)
	for _, addr := range lisAddrs {
		fmt.Printf("Swarm listening on %s\n", addr)
	}

	var addrs []string
	for _, addr := range node.PeerHost.Addrs() {
		addrs = append(addrs, addr.String())
	}
	sort.Strings(addrs)
	for _, addr := range addrs {
		fmt.Printf("Swarm announcing %s\n", addr)
	}
}

type P2pService struct {
	node *core.IpfsNode
}

func NewP2pService(node *core.IpfsNode) *P2pService {
	return &P2pService{
		node: node,
	}
}

func (s *P2pService) Ping(ctx context.Context, peerId string) bool {
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

func (s *P2pService) CreateForward(ctx context.Context, protoOpt string, listenOpt string, targetOpt string) error {

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
func (s *P2pService) CloseListen(ctx context.Context, protoOpt string, listenOpt string, targetOpt string) error {

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
func (s *P2pService) CheckPort(target ma.Multiaddr) error {
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

func (s *P2pService) ListListen(ctx context.Context, req *pb.ListListenRequest) (*pb.ListListenReply, error) {
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
