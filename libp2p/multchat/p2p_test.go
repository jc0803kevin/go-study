package main

import (
	"encoding/json"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/peer"
	"log"
	"testing"
)

func TestNode_Invited(t *testing.T) {

	imsg := InvitedMessage{
		Invite:"11111111111111",
		RoomID:"roomID",
		RoomName:"roomName",
	}

	// todo kevin 只要是可导出成员（变量首字母大写），都可以转成json
	imsgByte , err := json.Marshal(imsg)
	if err != nil {
		log.Printf("Invited编码消息失败")
		panic(err)
	}

	log.Printf("发送邀请消息 imsgByte：%s", string(imsgByte))

	m := &ChatMessage{
		Type:		Invited,
		Message:    string(imsgByte),
		SenderID:   "node.self.ID().Pretty()",
		SenderNick: "node.cfg.nick",
	}

	msgBytes, err := json.Marshal(*m)
	if err != nil {
		panic(err)
	}

	var cm ChatMessage
	if err := json.Unmarshal([]byte(string(msgBytes)), &cm); err != nil{
		//cr.Messages <- &cm
		return
	}

	switch cm.Type {
	case SEND:
		break
	case Invited:
		// InvitedMessage

		var imsg InvitedMessage
		if err := json.Unmarshal([]byte(cm.Message), &imsg); err != nil {
			log.Printf("解析邀请消息异常 err : %s", err)
			return
		}

		log.Printf("接收到邀请消息 imsg ： %s", &imsg)



	}

}


func TestToString(t *testing.T) {

	host , _ := libp2p.New()

	info := peer.AddrInfo{
		ID :   host.ID(),
		Addrs : host.Addrs(),
	}
	fmt.Println(ToString(info))
}
