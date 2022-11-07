package main


import (
	"context"

	"github.com/ipfs/go-log/v2"

	"github.com/libp2p/go-libp2p"

	"github.com/multiformats/go-multiaddr"

	//ds "github.com/ipfs/go-datastore"
	//dsync "github.com/ipfs/go-datastore/sync"
	dht "github.com/libp2p/go-libp2p-kad-dht"
)

var logger = log.Logger("rendezvous")

func main() {
	log.SetAllLoggers(log.LevelWarn)
	log.SetLogLevel("rendezvous", "debug")
	logger.Info("crate host")

	host, err := libp2p.New(libp2p.ListenAddrs([]multiaddr.Multiaddr(nil)...))
	if err != nil {
		panic(err)
	}
	logger.Info("Host created. We are:", host.ID())
	logger.Info(host.Addrs())

	ctx := context.Background()
	//dstore := dsync.MutexWrap(ds.NewMapDatastore())
	// kademliaDHT := dht.New(ctx, host)
	//kademliaDHT := dht.NewDHT(ctx, host, dstore)
	kademliaDHT, err := dht.New(ctx, host, dht.Mode(dht.ModeAutoServer))
	if err != nil {
		panic(err)
	}

	// Bootstrap the DHT. In the default configuration, this spawns a Background
	// thread that will refresh the peer table every five minutes.
	logger.Info("Bootstrapping the DHT")
	if err = kademliaDHT.Bootstrap(ctx); err != nil {
		panic(err)
	}

	// logger.Infof("%v%v", host.ID().Pretty(), host.Addrs())

	select {}

}
