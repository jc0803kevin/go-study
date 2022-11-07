package main

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
)

func ConnPeers(addr string) {

	host, err := libp2p.New()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	peerInfo, err := getPeerInfoByDest(addr)

	err = host.Connect(ctx, *peerInfo)
	if err != nil {
		panic(err)
	}
}

func getPeerInfoByDest(relay string) (*peer.AddrInfo, error) {
	relayAddr, err := multiaddr.NewMultiaddr(relay)
	if err != nil {
		return nil, err
	}
	pid, err := relayAddr.ValueForProtocol(multiaddr.P_P2P)
	if err != nil {
		return nil, err
	}
	relayPeerID, err := peer.Decode(pid)
	if err != nil {
		return nil, err
	}

	relayPeerAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/p2p/%s", pid))
	relayAddress := relayAddr.Decapsulate(relayPeerAddr)
	peerInfo := &peer.AddrInfo{
		ID:    relayPeerID,
		Addrs: []multiaddr.Multiaddr{relayAddress},
	}
	return peerInfo, err
}
