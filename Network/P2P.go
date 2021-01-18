package Network

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/ipfs/go-log"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	discovery "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var logger = log.Logger("NETWORK")

// Start will create a new node and use peer discovery to connect to other nodes in a P2P network.
func Start(port int) {

	ctx := context.Background()
	_ = log.SetLogLevel("NETWORK", os.Getenv("LOG_LEVEL"))

	// Private RSA Key generated in order to identify peer inside the P2P network
	privateK, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, rand.Reader)
	if err != nil {
		panic(err)
	}

	// Constructing a new node in the P2P network
	node, err := libp2p.New(ctx,
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", port)),
		libp2p.Identity(privateK),
		)
	if err != nil {
		panic(err)
	}
	logger.Info("Node ID : ", node.ID()," & Listen addresses : ", node.Addrs())

	node.SetStreamHandler(protocol.ID(os.Getenv("PROTOCOL_ID")), handleStream)

	// Find other nodes inside the P2P network to connect to
	peer2peerDiscovery(node, ctx)

	// wait for a SIGINT or SIGTERM signal to shut down the node
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	if err := node.Close(); err != nil {
		panic(err)
	}
}

// peer2peerDiscovery will use a protocol to discover nodes in the P2P Network so we can connect to them
func peer2peerDiscovery(node host.Host, ctx context.Context) {

	// Create a Distributed Hash Table, used for peer discovery
	networkDHT, err := dht.New(ctx, node)
	if err != nil {
		panic(err)
	}
	if err = networkDHT.Bootstrap(ctx); err != nil {
		panic(err)
	}

	// Checking every addresses in the DHT to find peers to connect to
	var wg sync.WaitGroup
	for _, peerAddress := range dht.DefaultBootstrapPeers {
		peerInfo, _ := peer.AddrInfoFromP2pAddr(peerAddress)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := node.Connect(ctx, *peerInfo); err != nil {
				logger.Debugf("Node Connect failed : %v",err)
			} else {
				logger.Info("Connection established with bootstrap node : ", *peerInfo)
			}
		}()
	}
	wg.Wait()

	logger.Info("Establishing rendez-vous point for other peers to see us")
	routingDiscovery := discovery.NewRoutingDiscovery(networkDHT)
	discovery.Advertise(ctx, routingDiscovery, os.Getenv("RDV"))

	logger.Info("Searching for other peers...")
	peerChan, err := routingDiscovery.FindPeers(ctx, os.Getenv("RDV"))
	if err != nil {
		panic(err)
	}

	for p := range peerChan {
		if p.ID == node.ID() {
			continue
		}
		logger.Debug("Connecting to peer : ", p)
		stream, err := node.NewStream(ctx, p.ID, protocol.ID(os.Getenv("PROTOCOL_ID")))
		if err != nil {
			logger.Debug("Connection failed : ", err)
			continue
		}

		handleStream(stream)
		logger.Info("Connected to : ", p)
	}
}