package main

import (
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
)

type discoveryNotifee struct {
	PeerChan chan peer.AddrInfo
}

// 发现一个新的对等节点
// interface to be called when new  peer is found
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	n.PeerChan <- pi
}

// Initialize the MDNS service
func initMDNS(peerhost host.Host, rendezvous string) chan peer.AddrInfo {

	notifee := &discoveryNotifee{}
	notifee.PeerChan = make(chan peer.AddrInfo)
	mdnsService := mdns.NewMdnsService(peerhost, rendezvous, notifee)

	if err := mdnsService.Start(); err != nil {
		panic(err)
	}

	return notifee.PeerChan
}
