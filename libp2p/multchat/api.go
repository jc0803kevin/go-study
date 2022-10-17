package main

import (
	"encoding/json"
	"github.com/emicklei/go-restful"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
	"io"
	"log"
)

//获取当前用户id
func id(request *restful.Request, response *restful.Response) {
	io.WriteString(response.ResponseWriter, self.ID().String())
}

func getPeers(request *restful.Request, response *restful.Response)  {
	peersJson, _ := json.Marshal(cr.peers)

	io.WriteString(response.ResponseWriter, string(peersJson))
}

func newChat(request *restful.Request, response *restful.Response)  {

	roomName := request.QueryParameter("roomName")

	if roomName == "" {
		io.WriteString(response.ResponseWriter, "roomName 不能为空")
		return
	}

	//ctx , cancel := context.WithCancel(context.Background())
	//defer cancel()

	cr, _ = NewChatRoom(ctx, self, cfg.nick, roomName)
	//if err != nil {
	//	panic("创建房间失败。。。")
	//}

	//crBytes, err := json.Marshal(&cr)
	//if err != nil {
	//	panic("创建房间失败。。。")
	//}

	io.WriteString(response.ResponseWriter, "创建房间成功 房间名称："+ roomName)
}

func JoinRoom(request *restful.Request, response *restful.Response)  {
	destination := request.QueryParameter("destination")

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

	cr.JoinRoom(*info)

	io.WriteString(response.ResponseWriter, "ok")
}

func sendMessage(request *restful.Request, response *restful.Response)  {
	message := request.QueryParameter("message")
	destination := request.QueryParameter("destination")

	log.Printf("sendMessage destination : %s", destination)
	log.Printf("sendMessage message : %s", message)

	//target , _ := peer.Decode(peerID)
	cr.sendMessage(message, destination)
	io.WriteString(response.ResponseWriter, "ok")
}
