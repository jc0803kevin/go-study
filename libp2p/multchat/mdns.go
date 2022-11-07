package main

import (
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
)

const (
	ServiceName = mdns.ServiceName
)

type discoveryNotifee struct {
	PeerChan chan peer.AddrInfo
}

//interface to be called when new  peer is found
//发现新的对等节点回调
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	n.PeerChan <- pi

	//cr.JoinRoom(pi)
}

func InitMdns(peerhost host.Host) chan peer.AddrInfo {
	// register with service so that we get notified about peer discovery
	n := &discoveryNotifee{}
	n.PeerChan = make(chan peer.AddrInfo)

	// An hour might be a long long period in practical applications. But this is fine for us
	ser := mdns.NewMdnsService(peerhost, ServiceName, n)
	if err := ser.Start(); err != nil {
		panic(err)
	}
	return n.PeerChan
}
