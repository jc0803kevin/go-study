package main

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/multiformats/go-multiaddr"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	host, err := libp2p.New()
	if err != nil {
		panic(err)
	}

	relayAddr, err := multiaddr.NewMultiaddr("/ip4/192.168.23.146/tcp/53125/p2p/12D3KooWSjS59wbP3uGKa76KP4u2xQf7AvQJzi56nhmrREpKTZ8L")
	h2, err := peer.AddrInfoFromP2pAddr(relayAddr)

	host.Peerstore().AddAddr(h2.ID, relayAddr, peerstore.PermanentAddrTTL)
	if err := host.Connect(ctx, *h2); err != nil {
		panic(err)
	}

	ps, err := pubsub.NewGossipSub(ctx, host)

	subscription, err := ps.Subscribe("kevin")

	for {
		msg, err := subscription.Next(ctx)
		if err != nil {
			panic(err)
		}

		fmt.Printf("sub rec msg : %s \n", msg.Data)
	}

	select {}
}
