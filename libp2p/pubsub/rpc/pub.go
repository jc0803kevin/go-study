package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	priv, _ := getPrivateKey(53125)

	host, err := libp2p.New(
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/53125"),
		libp2p.Identity(priv),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(getHostAddress(host))

	ps, err := pubsub.NewGossipSub(ctx, host)

	for {
		ps.Publish("kevin", []byte("hello kevin"))

		time.Sleep(time.Second * 10)
	}

	select {}
}

func getHostAddress(ha host.Host) string {
	// Build host multiaddress
	hostAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/p2p/%s", ha.ID().Pretty()))

	// Now we can build a full multiaddress to reach this host
	// by encapsulating both addresses:
	addr := ha.Addrs()[0]
	return addr.Encapsulate(hostAddr).String()
}

func getPrivateKey(port int) (crypto.PrivKey, error) {

	var generate bool

	fileName := fmt.Sprintf("%d-%s", port, "libp2p-chat.privkey")

	privKeyBytes, err := ioutil.ReadFile(fileName)
	if os.IsNotExist(err) {
		generate = true
	} else if err != nil {
		return nil, err
	}

	if generate {
		privKey, err := generateNewPrivKey()
		if err != nil {
			return nil, err
		}

		privKeyBytes, err := crypto.MarshalPrivateKey(privKey)
		if err != nil {
			return nil, errors.Wrap(err, "marshalling identity private key")
		}

		f, err := os.Create(fileName)
		if err != nil {
			return nil, errors.Wrap(err, "creating identity private key file")
		}
		defer f.Close()

		if _, err := f.Write(privKeyBytes); err != nil {
			return nil, errors.Wrap(err, "writing identity private key to file")
		}

		return privKey, nil
	}

	privKey, err := crypto.UnmarshalPrivateKey(privKeyBytes)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshalling identity private key")
	}

	return privKey, nil
}

func generateNewPrivKey() (crypto.PrivKey, error) {
	privKey, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		return nil, errors.Wrap(err, "generating identity private key")
	}

	return privKey, nil
}
