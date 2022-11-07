package main

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/emicklei/go-restful"
	"io"
	"log"
)

//获取当前用户id
func id(request *restful.Request, response *restful.Response) {
	io.WriteString(response.ResponseWriter, node.self.ID().String())
}

//func getPeers(request *restful.Request, response *restful.Response)  {
//	peersJson, _ := json.Marshal(cr.peers)
//
//	io.WriteString(response.ResponseWriter, string(peersJson))
//}

func newChat(request *restful.Request, response *restful.Response)  {
	roomName := request.QueryParameter("roomName")
	log.Printf("开始创建一个聊天房间 roomName： %s", roomName)
	if roomName == "" {
		io.WriteString(response.ResponseWriter, "roomName 不能为空")
		return
	}

	h := md5.New()
	h.Write([]byte(roomName))
	roomID := hex.EncodeToString(h.Sum(nil))

	_ , isPersent := node.RoomMgr.rooms[roomID]
	if isPersent {
		io.WriteString(response.ResponseWriter, "房间名称以存在 : " + roomID)
		return
	}
	log.Printf("开始创建一个聊天房间 roomID： %s", roomID)
	cr, _ := NewChatRoom(roomID, node.ctx, node.self, node.cfg.nick, roomName)

	node.RoomMgr.rooms[roomID] = &cr

	io.WriteString(response.ResponseWriter, "创建房间成功 房间名称："+ roomID)
}

// 群里邀请新成员，
//
// 群里历史成员 都需要添加新成员

func JoinRoom(request *restful.Request, response *restful.Response)  {
	destination := request.QueryParameter("destination")
	roomID := request.QueryParameter("roomID")

	cr , isPersent := node.RoomMgr.rooms[roomID]
	if !isPersent {
		io.WriteString(response.ResponseWriter, "该房间不存，请正确选择房间或创建一个新房间")
		return
	}

	// 新成员
	destInfo := GetInfoByMultiaddr(destination)
	//node.RoomMgr.rooms[roomID] = cr

	// 向被邀请人 发送消息 以相互加入群
	node.Invited(cr, destInfo)

	cr.JoinRoom(destInfo)
	io.WriteString(response.ResponseWriter, "ok")
}

func sendMessage(request *restful.Request, response *restful.Response)  {
 	message := request.QueryParameter("message")
	roomID := request.QueryParameter("roomID")

	log.Printf("sendMessage roomID : %s", roomID)
	log.Printf("sendMessage message : %s", message)

	if message == "" {
		io.WriteString(response.ResponseWriter, "message is nil")
		return
	}
	if roomID == "" {
		io.WriteString(response.ResponseWriter, "roomID is nil")
		return
	}


	cr , isPersent := node.RoomMgr.rooms[roomID]
	if !isPersent {
		io.WriteString(response.ResponseWriter, "该房间不存，请正确选择房间或创建一个新房间")
		return
	}

	// 将消息广播 到房间中所有的成员
	node.sendMessage(cr, message)
	io.WriteString(response.ResponseWriter, "ok")
}
