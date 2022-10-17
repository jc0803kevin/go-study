package main

import (
	"context"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"log"
	"sync"
)

const ChatRoomBufSize = 128

type ChatMessage struct {
	Message    string
	SenderID   string
	SenderNick string
}

type ChatRoom struct {
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

func NewChatRoom(ctx context.Context,  node host.Host, nickname string, roomName string) (ChatRoom, error) {

	cr := ChatRoom{
		ctx:      ctx,
		node:     node,
		self:	  node.ID(),
		nick:     nickname,
		roomName: roomName,
		Messages: make(chan *ChatMessage, ChatRoomBufSize),
	}

	//// start reading messages from the subscription in a loop
	//go cr.readLoop()
	return cr, nil
}