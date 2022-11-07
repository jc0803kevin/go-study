package main

import (
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
	"log"
)

//输出节点的完整信息

func V4()  {

	basicHost , _ := libp2p.New()

	// 12D3KooWAkRX6EZH1ueDf22dNVm33z8KftU4HMY1oH8A6t3JKU8z
	// 建立主机的Multiaddr，libp2p使用一种独特的Mutliaddr，而非传统的IP+端口，用于节点之间互相发现
	hostAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/p2p/%s", basicHost.ID().Pretty()))//服务器的ipfs地址，用于被其它节点发现
	// 现在，我们可以通过封装两个地址来构建一个可抵达主机的完整的Multiaddr
	addr := basicHost.Addrs()[0]

	//  /ip4/192.168.23.146/tcp/54621
	log.Printf("addr %s\n", addr)

	// 全地址
	// /ip4/192.168.23.146/tcp/54621/p2p/12D3KooWAkRX6EZH1ueDf22dNVm33z8KftU4HMY1oH8A6t3JKU8z
	fullAddr := addr.Encapsulate(hostAddr)
	log.Printf("I am %s\n", fullAddr)


	dncapsulate()
}

func dncapsulate()  {

	info , _ := peer.AddrInfoFromString("/ip4/192.168.23.146/tcp/54621/p2p/12D3KooWAkRX6EZH1ueDf22dNVm33z8KftU4HMY1oH8A6t3JKU8z")

	// 12D3KooWAkRX6EZH1ueDf22dNVm33z8KftU4HMY1oH8A6t3JKU8z
	log.Printf("ID   : %s ", info.ID)

	// /ip4/192.168.23.146/tcp/54621
	log.Printf("Addr : %s ", info.Addrs[0])

}