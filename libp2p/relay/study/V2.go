package study

import (
	"bufio"
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peerstore"
	"github.com/libp2p/go-libp2p-core/pnet"
	"github.com/libp2p/go-libp2p/p2p/transport/tcp"
	"log"
	"os"
)


// 两个节点交流


func V2()  {

	ctx , cancel := context.WithCancel(context.Background())
	defer cancel()

	file, err := os.Open("E:\\swarm.key")
	if err != nil {
		log.Printf("加载 swarm key失败 %s", err)
	}
	psk , err := pnet.DecodeV1PSK(file)

	opts := []libp2p.Option{
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"),
		//libp2p.DisableRelay(),

		//quic.NewTransport failed: QUIC doesn't support private networks yet
		libp2p.Transport(tcp.NewTCPTransport),
		libp2p.PrivateNetwork(psk),
		libp2p.EnableRelay(),
	}

	h1, err := libp2p.New(opts...)
	if err != nil {
		log.Printf("Failed to create h1: %v", err)
		return
	}

	h2, err := libp2p.New(opts...)
	if err != nil {
		log.Printf("Failed to create h2: %v", err)
		return
	}

	h1.Peerstore().SetAddrs(h2.ID(), h2.Addrs(), peerstore.PermanentAddrTTL)
	h2.Peerstore().SetAddrs(h1.ID(), h1.Addrs(), peerstore.PermanentAddrTTL)


	log.Printf("h1 peers : %s", h1.Peerstore().Peers().String())
	log.Printf("h2 peers : %s", h2.Peerstore().Peers().String())

	h2.SetStreamHandler("kevin", func(s network.Stream) {
		rw := bufio.NewReadWriter(bufio.NewReader(s),bufio.NewWriter(s))
		for{
			str, _ := rw.ReadString('\n')

			log.Printf("h2 读取到数据  --》 %s", str)
		}
	})

	h1stream, err := h1.NewStream(ctx, h2.ID(), "kevin")
	for i := 0; i< 5 ; i++  {
		str := fmt.Sprintf("h1 send maessge : %d \n", i)
		log.Println(str)
		h1stream.Write([]byte(str))
	}

	select {}
}
