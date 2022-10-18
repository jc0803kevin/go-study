package main

import (
	"context"
	"encoding/json"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"log"
	"sync"
)

const ChatRoomBufSize = 128

type ChatMessage struct {
	Type 	   int
	Message    string
	SenderID   string
	SenderNick string
}

// 邀请消息
type InvitedMessage struct {
	// 邀请人 /ip4/127.0.0.1/tcp/3001/p2p/12D3KooWAtyKjJGFABAaccn6SBT61rXwuHtFWKW9mLC2S5gesBpS
	Invite string

	RoomID string
	RoomName string
}


const (
	// 发送消息
	SEND = iota

	// 邀请入群
	Invited
)


type ChatRoom struct {
	// 群ID
	id string

	Messages chan *ChatMessage
	ctx   context.Context

	node host.Host
	self peer.ID
	roomName string
	nick     string

	// 对等列表
	peers []peer.AddrInfo

	lock         sync.RWMutex

}

func (cr *ChatRoom)  peersToString() (peersJsonStr string) {
	bytes , err := json.Marshal(cr.peers)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}


func (cr *ChatRoom)JoinRoom(info peer.AddrInfo)  {
	cr.lock.Lock()
	defer cr.lock.Unlock()

	for e := range cr.peers {
		if cr.peers[e].ID.String() == info.ID.String() {
			log.Printf("该对等节点已存在 peerID:%s", info.ID.String())
			return
		}
	}

	log.Printf("自动添加新的节点 %s", info.String())
	cr.peers = append(cr.peers, info)


}

func NewChatRoom(roomID string, ctx context.Context,  node host.Host, nickname string, roomName string) (ChatRoom, error) {

	cr := ChatRoom{
		id:      roomID,
		ctx:      ctx,
		node:     node,
		self:	  node.ID(),
		nick:     nickname,
		roomName: roomName,
		Messages: make(chan *ChatMessage, ChatRoomBufSize),
	}

	// 房间对等加上自己
	//hostAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/p2p/%s", node.ID().Pretty()))
	//addr := node.Addrs()[0]
	//fullAddr := addr.Encapsulate(hostAddr)
	//myselfInfo := GetInfoByMultiaddr(fullAddr)
	//cr.JoinRoom(myselfInfo)

	//// start reading messages from the subscription in a loop
	//go cr.readLoop()
	return cr, nil
}