package main

import (
	"context"
	"fmt"
	"github.com/emicklei/go-restful"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var node Node

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cfg := parseFlags()

	node.cfg = cfg
	node.ctx = ctx
	node.RoomMgr.rooms = make(map[string]*ChatRoom)

	self := randomHost(cfg.p2pPort)
	logger.Info("",zap.Any("My peer : ", self.Addrs()[0]))
	node.self = self

	// web api server
	// 创建container，并注册路由
	container := restful.NewContainer()

	// 跨域过滤器
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"X-My-Header"},
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST"},
		CookiesAllowed: false,
		Container:      container}
	container.Filter(cors.Filter)

	container.Filter(container.OPTIONSFilter)

	register(container)

	//web api监听地址
	webApiListenAddr := fmt.Sprintf("0.0.0.0:%d", cfg.webApiListenPort)
	logger.Info("",zap.Any("[*] Listening web api on: ", webApiListenAddr))
	server := newWebserver(webApiListenAddr, container)
	defer server.Close()
	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			logger.Error("Http Server closed under request: %v\n", zap.Error(err))
		} else {
			logger.Error("Could not listen on %s: %v\n", zap.String("WebApiListenAddr", webApiListenAddr), zap.Error(err))
		}
	}


}


// 创建一个webserver
func newWebserver(webApiListenAddr string, container *restful.Container) *http.Server {
	return &http.Server{
		Addr:         webApiListenAddr,
		Handler:      container,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
}

// 注册web路由
func register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/id").To(id))
	ws.Route(ws.GET("/joinRoom").To(JoinRoom))
	ws.Route(ws.GET("/newChat").To(newChat))
	//ws.Route(ws.GET("/participants").To(participants))
	ws.Route(ws.GET("/sendMessage").To(sendMessage))
	//ws.Route(ws.GET("/setNickname").To(setNickname))
	//ws.Route(ws.GET("/messgaeRecord").To(getMessgaeRecord))
	//ws.Route(ws.GET("/rooms").To(getRooms))
	//ws.Route(ws.GET("/getPeers").To(getPeers))

	container.Add(ws)
}
