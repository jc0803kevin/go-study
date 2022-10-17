package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/multiformats/go-multiaddr"
	"log"
)

const CHAT_PID  = "multchat/1.0.0"

func randomHost(port int) (host host.Host) {
	//host, err := libp2p.New(libp2p.ListenAddrs(multiaddr.StringCast("/ip4/127.0.0.1/tcp/0/")))
	host, err := libp2p.New(libp2p.ListenAddrs(multiaddr.StringCast(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", port))))
	if err != nil {
		log.Printf("创建节点失败 err: %s", err)
	}

	InitMdns(host)

	log.Printf("My peer : %s", host.ID().String())
	host.SetStreamHandler(CHAT_PID, handleStream)

	return host
}

func handleStream(s network.Stream) {
	log.Println("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	go readData(rw)

	// stream 's' will stay open until you close it (or the other side closes it).
}


// 接收到新数据
func readData(rw *bufio.ReadWriter) {
	for {
		str, _ := rw.ReadString('\n')

		if str == "" {
			return
		}
		if str != "\n" {
			// Green console colour: 	\x1b[32m
			// Reset console colour: 	\x1b[0m
			fmt.Printf("\x1b[32m%s\x1b[0m> ", str)

			var cm ChatMessage
			if err := json.Unmarshal([]byte(str), &cm); err != nil{
				cr.Messages <- &cm
			}

		}

	}
}

func (cr *ChatRoom)sendMessage(message string, destination string)  {

	m := &ChatMessage{
		Message:    message,
		SenderID:   cr.node.ID().Pretty(),
		SenderNick: cr.nick,
	}
	msgBytes, err := json.Marshal(*m)
	if err != nil {
		panic(err)
	}

	cr.Messages <- m

	// /ip4/127.0.0.1/tcp/3001/p2p/12D3KooWAtyKjJGFABAaccn6SBT61rXwuHtFWKW9mLC2S5gesBpS
	maddr, err := multiaddr.NewMultiaddr(destination)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	// Extract the peer ID from the multiaddr.
	info, err := peer.AddrInfoFromP2pAddr(maddr)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	cr.node.Peerstore().AddAddrs(info.ID, info.Addrs, peerstore.PermanentAddrTTL)

	s, err := cr.node.NewStream(cr.ctx, info.ID, "multchat/1.0.0")
	if s == nil {
		log.Printf("创建流失败")
		panic(err)
	}
	defer s.Close()

	if _, err := s.Write(msgBytes) ; err != nil {
		log.Printf("发送消息失败 %s ",err)
	}
}



//func ()writeData(rw *bufio.ReadWriter) {
//	stdReader := bufio.NewReader(os.Stdin)
//	for {
//		fmt.Print("> ")
//		sendData, err := stdReader.ReadString('\n')
//		if err != nil {
//			log.Println(err)
//			return
//		}
//
//		rw.WriteString(fmt.Sprintf("%s\n", sendData))
//		rw.Flush()
//	}
//}