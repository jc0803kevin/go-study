package main

import (
	"bufio"
	"context"
	"crypto/rand"
	"flag"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"
	"io"
	"log"
	mrand "math/rand"

	golog "github.com/ipfs/go-log/v2"
	ma "github.com/multiformats/go-multiaddr"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// LibP2P code uses golog to log messages. They log with different
	// string IDs (i.e. "swarm"). We can control the verbosity level for
	// all loggers with:
	golog.SetAllLoggers(golog.LevelInfo) // Change to INFO for extra info

	// Parse options from the command line
	listenF := flag.Int("l", 0, "wait for incoming connections")
	targetF := flag.String("d", "", "target peer to dial")
	insecureF := flag.Bool("insecure", false, "use an unencrypted connection")
	seedF := flag.Int64("seed", 0, "set random seed for id generation")
	flag.Parse()

	if *listenF == 0 {
		log.Fatal("Please provide a port to bind on with -l")
	}

	// 建立链接的对等 需要保持加密 是否一致。要么同时加密，要么同时不加密
	ha, err := makeBasicHost(*listenF, *insecureF, *seedF)
	if err != nil {
		log.Fatal(err)
	}

	if *targetF == "" {
		startListener(ctx, ha, *listenF, *insecureF)
		// Run until canceled.
		<-ctx.Done()
	} else {
		runSender(ctx, ha, *targetF)
	}

}

func makeBasicHost(listenPort int, insecure bool, randseed int64) (host.Host, error) {

	var r io.Reader
	if randseed == 0 {
		r = rand.Reader
	} else {
		r = mrand.New(mrand.NewSource(randseed))
	}

	priv, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
	if err != nil {
		return nil, err
	}

	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", listenPort)),
		libp2p.Identity(priv),
		libp2p.DisableRelay(),
	}

	if insecure {
		opts = append(opts, libp2p.NoSecurity)
	}

	return libp2p.New(opts...)

}

func startListener(ctx context.Context, ha host.Host, listenPort int, insecure bool) {

	fullAddr := getHostAddress(ha)
	log.Printf("I am %s\n", fullAddr)

	// Set a stream handler on host A. /echo/1.0.0 is
	// a user-defined protocol name.
	ha.SetStreamHandler("/echo/1.0.0", func(s network.Stream) {
		log.Println("listener received new stream")
		if err := doEcho(s); err != nil {
			log.Println(err)
			s.Reset()
		} else {
			s.Close()
		}
	})

	log.Println("listening for connections")

	if insecure {
		log.Printf("Now run \"./echo -l %d -d %s -insecure\" on a different terminal\n", listenPort+1, fullAddr)
	} else {
		log.Printf("Now run \"./echo -l %d -d %s\" on a different terminal\n", listenPort+1, fullAddr)
	}

}

func runSender(ctx context.Context, ha host.Host, targetPeer string) {
	fullAddr := getHostAddress(ha)
	log.Printf("I am %s\n", fullAddr)

	ha.SetStreamHandler("/echo/1.0.0", func(s network.Stream) {
		log.Println("sender received new stream")
		if err := doEcho(s); err != nil {
			log.Println(err)
			s.Reset()
		} else {
			s.Close()
		}
	})

	// 目标对等节点
	maddr, err := ma.NewMultiaddr(targetPeer)
	if err != nil {
		log.Println(err)
		return
	}

	info, err := peer.AddrInfoFromP2pAddr(maddr)
	if err != nil {
		log.Println(err)
		return
	}

	// 将对等节点 添加到自身对等库
	ha.Peerstore().AddAddrs(info.ID, info.Addrs, peerstore.PermanentAddrTTL)

	log.Println("sender opening stream")

	s, err := ha.NewStream(context.Background(), info.ID, "/echo/1.0.0")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("sender saying hello")

	_, err = s.Write([]byte("Hello world ! \n"))
	if err != nil {
		log.Println(err)
		return
	}

	out, err := io.ReadAll(s)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("read reply: %q\n", out)

}

func doEcho(s network.Stream) error {
	buf := bufio.NewReader(s)
	str, err := buf.ReadString('\n')
	if err != nil {
		return err
	}

	log.Printf("read: %s", str)

	_, err = s.Write([]byte(str))

	return err
}

func getHostAddress(ha host.Host) string {
	// Build host multiaddress
	hostAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/p2p/%s", ha.ID().Pretty()))

	// Now we can build a full multiaddress to reach this host
	// by encapsulating both addresses:
	addr := ha.Addrs()[0]
	return addr.Encapsulate(hostAddr).String()
}