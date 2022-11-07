package main

import (
	"context"
	"github.com/libp2p/go-libp2p-core/host"
)

// 当前节点
type Node struct {

	RoomMgr RoomMgr

	self host.Host

	cfg Config

	ctx context.Context
}

// 房间管理
type RoomMgr struct {

	rooms map[string]*ChatRoom

}
