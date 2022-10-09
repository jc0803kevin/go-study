package main

import (
	"context"
	"encoding/json"

	"github.com/libp2p/go-libp2p-core/peer"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

// ChatRoomBufSize is the number of incoming messages to buffer for each topic.
const ChatRoomBufSize = 128

type ChatRoom struct {
	Messages chan *ChatMessage

	ctx   context.Context
	ps    *pubsub.PubSub
	topic *pubsub.Topic
	sub   *pubsub.Subscription

	roomName string
	self     peer.ID
	nick     string
}

func (cr *ChatRoom) readLoop() {
	for {
		msg, err := cr.sub.Next(cr.ctx)
		if err != nil {
			close(cr.Messages)
			return
		}

		// only forward messages delivered by others
		if msg.ReceivedFrom == cr.self {
			continue
		}

		cm := new(ChatMessage)

		err = json.Unmarshal(msg.Data, cm)
		if err != nil {
			continue
		}

		// send valid messages onto the Messages channel
		cr.Messages <- cm
	}
}

type ChatMessage struct {
	Message    string
	SenderID   string
	SenderNick string
}

func JoinChatRoom(ctx context.Context, ps *pubsub.PubSub, selfID peer.ID, nickname string, roomName string) (*ChatRoom, error) {

	// 利用房间名称 作为主题名称
	topic, err := ps.Join(topicName(roomName))
	if err != nil {
		return nil, err
	}

	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}

	cr := &ChatRoom{
		Messages: make(chan *ChatMessage, ChatRoomBufSize),
		ctx:      ctx,
		ps:       ps,
		topic:    topic,
		sub:      sub,
		roomName: roomName,
		self:     selfID,
		nick:     nickname,
	}

	// start reading messages from the subscription in a loop
	go cr.readLoop()

	return cr, nil
}

func (cr *ChatRoom) Publish(message string) error {
	m := ChatMessage{
		Message:    message,
		SenderID:   cr.self.Pretty(),
		SenderNick: cr.nick,
	}

	msgBytes, err := json.Marshal(m)
	if err != nil {
		return err
	}

	return cr.topic.Publish(cr.ctx, msgBytes)

}

func (cr *ChatRoom) ListPeers() []peer.ID {
	return cr.ps.ListPeers(topicName(cr.roomName))
}

func topicName(roomName string) string {
	return "chat-room:" + roomName
}
