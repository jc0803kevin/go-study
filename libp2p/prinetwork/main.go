package main

import (
	"context"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/pnet"
	"github.com/libp2p/go-libp2p/p2p/transport/tcp"
	"github.com/multiformats/go-multiaddr"
	"log"
	"os"
)

// 利用libp2p 加载swarm key 加入私有网络
// 私有网络中的任意两个节点 之间 可以正常通信


func main() {

	ctx, cancel := context.WithCancel(context.Background())
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
	}

	privhost , err := libp2p.New(opts...)
	if err != nil {
		log.Printf(" 启动节点失败 %s", err)
	}
	log.Printf("privhost 节点 peer id %s ", privhost.ID())

	privhost2 , err := libp2p.New(opts...)
	if err != nil {
		log.Printf(" 启动节点失败 %s", err)
	}
	log.Printf("privhost2 节点 peer id %s ", privhost2.ID())
	h2info := peer.AddrInfo{
		ID:    privhost2.ID(),
		Addrs: privhost2.Addrs(),
	}

	// 公有网络
	normalhost , err := libp2p.New()
	if err != nil {
		log.Printf(" 启动节点失败 %s", err)
	}
	log.Printf("normalhost 节点 peer id %s ", normalhost.ID())

	// 私有网络 引导节点
	muladdr , err := multiaddr.NewMultiaddr("/ip4/112.124.44.121/tcp/4001/p2p/12D3KooWKff4bsrFYwKmn6HEDi4jaBxN63oW1iMinqhN36sFKzyQ")

	info , err := peer.AddrInfoFromP2pAddr(muladdr)
	if err != nil {
		log.Printf(" AddrInfoFromP2pAddr 失败 %s", err)
	}

	if err := privhost.Connect(ctx, *info); err != nil {
		log.Printf("privhost 节点 连接失败 %s ", info.ID)
	}
	if err := privhost2.Connect(ctx, *info); err != nil {
		log.Printf("privhost2 节点 连接失败 %s ", info.ID)
	}

	// 加入私有网络 中的两个随机几点 可以正常通信
	if err := privhost.Connect(ctx, h2info); err != nil {
		log.Printf("privhost connect privhost2  节点 连接失败 %s ", info.ID)
	}else {
		log.Printf("privhost [%s] connect privhost2 [%s] success . ",privhost.ID(), privhost2.ID())
	}

	// 不是私有网络 连接会失败
	if err := normalhost.Connect(ctx, *info); err != nil {
		log.Printf("normalhost 节点 连接失败 %s ", info.ID)
	}


	log.Println(privhost.Peerstore().PeersWithAddrs().String())
	log.Println(privhost2.Peerstore().PeersWithAddrs().String())
	log.Println(normalhost.Peerstore().PeersWithAddrs().String())

	select {}
}
