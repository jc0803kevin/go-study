package main

import (
	"flag"
)

type Config struct {
	nick string
	roomName string
	p2pPort int
	webApiListenPort int
}


func parseFlags() (config Config){

	nickFlag := flag.String("nick", "", "nickname to use in chat. will be generated if empty")
	//roomFlag := flag.String("room", "", "name of chat room to join")
	p2pPortFlag := flag.Int("p2p-port", 0, "p2p listen port")
	webApiListenPortFlag := flag.Int("web-port", 0, "web api listen port")
	flag.Parse()

	if *nickFlag == "" {
		panic("请设置节点名称")
	}

	//if *roomFlag == "" {
	//	panic("请设置房间名称")
	//}

	config = Config{
		nick:*nickFlag,
		//roomName:*roomFlag,
		p2pPort:*p2pPortFlag,
		webApiListenPort:*webApiListenPortFlag,
	}

	return config
}
