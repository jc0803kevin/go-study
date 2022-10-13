package main
//
//import (
//	"bufio"
//	"context"
//	"fmt"
//	"github.com/libp2p/go-libp2p"
//	"github.com/libp2p/go-libp2p-core/network"
//	"github.com/libp2p/go-libp2p-core/peer"
//	"github.com/libp2p/go-libp2p-core/pnet"
//	"github.com/libp2p/go-libp2p/p2p/transport/tcp"
//	"github.com/multiformats/go-multiaddr"
//	"log"
//	"os"
//)
//
//func V2_bak()  {
//
//	ctx , cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	file, err := os.Open("E:\\swarm.key")
//	if err != nil {
//		log.Printf("加载 swarm key失败 %s", err)
//	}
//	psk , err := pnet.DecodeV1PSK(file)
//
//	opts := []libp2p.Option{
//		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"),
//		//libp2p.DisableRelay(),
//
//		//quic.NewTransport failed: QUIC doesn't support private networks yet
//		libp2p.Transport(tcp.NewTCPTransport),
//		libp2p.PrivateNetwork(psk),
//		libp2p.EnableRelay(),
//	}
//
//	h1, err := libp2p.New(opts...)
//	if err != nil {
//		log.Printf("Failed to create h1: %v", err)
//		return
//	}
//
//	h2, err := libp2p.New(opts...)
//	if err != nil {
//		log.Printf("Failed to create h2: %v", err)
//		return
//	}
//
//
//	relayAddr , err := multiaddr.NewMultiaddr("/ip4/112.124.44.121/tcp/4001/p2p/12D3KooWKff4bsrFYwKmn6HEDi4jaBxN63oW1iMinqhN36sFKzyQ")
//	relayInfo , err := peer.AddrInfoFromP2pAddr(relayAddr)
//	if err != nil {
//		log.Printf(" AddrInfoFromP2pAddr 失败 %s", err)
//	}
//
//
//	if err := h1.Connect(ctx, *relayInfo); err != nil {
//		log.Printf("Failed h1 connect relayInfo: %v", err)
//		return
//	}
//
//	if err := h2.Connect(ctx, *relayInfo); err != nil {
//		log.Printf("Failed h1 connect relayInfo: %v", err)
//		return
//	}
//
//	//h2info := peer.AddrInfo{
//	//	ID:    h2.ID(),
//	//	Addrs: h2.Addrs(),
//	//}
//
//	log.Printf("h1 peers : %s", h1.Peerstore().Peers().String())
//	log.Printf("h2 peers : %s", h2.Peerstore().Peers().String())
//
//	//h1.Peerstore().SetAddrs(h2.ID(), h2.Addrs(), peerstore.PermanentAddrTTL)
//	//h2.Peerstore().SetAddrs(h1.ID(), h1.Addrs(), peerstore.PermanentAddrTTL)
//	//
//	h2.SetStreamHandler("kevin", func(s network.Stream) {
//		rw := bufio.NewReadWriter(bufio.NewReader(s),bufio.NewWriter(s))
//
//		log.Println(rw.ReadString('\n'))
//	})
//	//
//	//h1.SetStreamHandler("kevin", func(s network.Stream) {
//	//	rw := bufio.NewReadWriter(bufio.NewReader(s),bufio.NewWriter(s))
//	//
//	//	log.Println(rw.ReadString('\n'))
//	//
//	//})
//
//	h1stream, err := h1.NewStream(ctx, h2.ID(), "kevin")
//	for i := 0; i< 5 ; i++  {
//		str := fmt.Sprintf("send maessge : %d ", i)
//		log.Println(str)
//		h1stream.Write([]byte(str))
//	}
//
//	//h1.SetStreamHandler("kevin", handleStream)
//	//h2.SetStreamHandler("kevin", handleStream)
//
//	//if err := h1.Connect(ctx, h2info); err != nil {
//	//	panic(err)
//	//}
//	//
//	//log.Printf("h1 peers : %s", h1.Peerstore().Peers().String())
//	//log.Printf("h2 peers : %s", h2.Peerstore().Peers().String())
//
//
//	select {}
//}
//
//
//
//func handleStream(s network.Stream) {
//	log.Println("Got a new stream!")
//
//	// Create a buffer stream for non blocking read and write.
//	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))
//
//	go readData(rw)
//	go writeData(rw)
//
//	// stream 's' will stay open until you close it (or the other side closes it).
//}
//
//func readData(rw *bufio.ReadWriter) {
//	for {
//		str, _ := rw.ReadString('\n')
//
//		if str == "" {
//			return
//		}
//		if str != "\n" {
//			// Green console colour: 	\x1b[32m
//			// Reset console colour: 	\x1b[0m
//			fmt.Printf("\x1b[32m%s\x1b[0m> ", str)
//		}
//
//	}
//}
//
//func writeData(rw *bufio.ReadWriter) {
//	for i := 0; i< 5 ; i++  {
//		str := fmt.Sprintf("send maessge : %d \n", i)
//		log.Println(str)
//		rw.WriteString(str)
//		rw.Flush()
//	}
//
///*	for {
//		fmt.Print("> ")
//		sendData, err := stdReader.ReadString('\n')
//		if err != nil {
//			log.Println(err)
//			return
//		}
//
//		rw.WriteString(fmt.Sprintf("%s\n", sendData))
//		rw.Flush()
//	}*/
//}