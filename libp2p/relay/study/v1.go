package study

import (
	"context"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"
	"log"
)

/**

h1 h2 都开启中继模式

 */


func V1()  {

	h1, err := libp2p.New(libp2p.EnableRelay())
	if err != nil {
		log.Printf("Failed to create h1: %v", err)
		return
	}

	h2, err := libp2p.New(libp2p.EnableRelay())
	if err != nil {
		log.Printf("Failed to create h2: %v", err)
		return
	}

	h2info := peer.AddrInfo{
		ID:    h2.ID(),
		Addrs: h2.Addrs(),
	}

	// h1 h2 建立连接
	if err := h1.Connect(context.Background(), h2info); err != nil {
		log.Printf("Failed to connect h1 and h2: %v", err)
		return
	}else {
		log.Println(" h1  h2  connect success..")
	}

}


