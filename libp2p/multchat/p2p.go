package main

import (
	"bufio"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/peerstore"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/multiformats/go-multiaddr"
	"io/ioutil"
	"log"
	"os"
)

const CHAT_SEND_PID  = "multchat/send/1.0.0"
const CHAT_INVITED_PID  = "multchat/Invited/1.0.0"

func randomHost(port int) (host host.Host) {
	//host, err := libp2p.New(libp2p.ListenAddrs(multiaddr.StringCast("/ip4/127.0.0.1/tcp/0/")))

	var priv crypto.PrivKey
	if node.cfg.configPeerIdentity == ""{
		priv, _, _ = crypto.GenerateKeyPair(crypto.RSA, 2048)
		writeIdentityJson(priv, port)
	}else {
		log.Printf("load local Identity Json : %s", node.cfg.configPeerIdentity)
		priv, _, _ = readIdentityJson(node.cfg.configPeerIdentity)
	}

	host, err := libp2p.New(
		libp2p.ListenAddrs(multiaddr.StringCast(fmt.Sprintf("/ip4/127.0.0.1/tcp/%d", port))),
		libp2p.Identity(priv),
		)
	if err != nil {
		log.Printf("创建节点失败 err: %s", err)
	}

	InitMdns(host)

	log.Printf("My peer : %s", host.ID().String())
	host.SetStreamHandler(CHAT_SEND_PID, handleStream)
	host.SetStreamHandler(CHAT_INVITED_PID, handleStream)

	return host
}

func handleStream(s network.Stream) {
	log.Println("Got a new stream!")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	go readData(rw)

	// stream 's' will stay open until you close it (or the other side closes it).
}


// 接收到新数据
func readData(rw *bufio.ReadWriter) {
	for {
		str, _ := rw.ReadString('\n')

		if str == "" {
			return
		}
		if str != "\n" {
			// Green console colour: 	\x1b[32m
			// Reset console colour: 	\x1b[0m
			fmt.Printf("\x1b[32m%s\x1b[0m> \n", str)

			var cm ChatMessage
			if err := json.Unmarshal([]byte(str), &cm); err != nil{
				//cr.Messages <- &cm
				return
			}

			switch cm.Type {
			case SEND:
				break
			case Invited:
				// InvitedMessage

				fmt.Printf("接收到邀请消息 cm.Message ： %s\n", cm.Message)

				var imsg InvitedMessage
				if err := json.Unmarshal([]byte(cm.Message), &imsg); err != nil {
					fmt.Printf("解析邀请消息异常 err : %s\n", err)
					return
				}

				cr , isPersent := node.RoomMgr.rooms[imsg.RoomID]
				if isPersent {
					fmt.Printf("已经创建该房间 roomID：%s\n", imsg.RoomID)
					return
				}else {
					newcr, _ := NewChatRoom(imsg.RoomID, node.ctx, node.self, node.cfg.nick, imsg.RoomName)
					node.RoomMgr.rooms[imsg.RoomID] = &newcr
					cr = &newcr
				}

				info := GetInfoByMultiaddr(imsg.Invite)
				cr.JoinRoom(info)

			}

		}

	}
}


func (node *Node)Invited(cr *ChatRoom, destInfo peer.AddrInfo)  {

	//hostAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/p2p/%s", node.self.ID().Pretty()))
	//addr := node.self.Addrs()[0]
	//fullAddr := addr.Encapsulate(hostAddr)
	//node.RoomMgr.rooms.Messages <- m

	imsg := &InvitedMessage{
		Invite:ToString(destInfo),
		RoomID:cr.id,
		RoomName:cr.roomName,
	}

	imsgByte , err := json.Marshal(imsg)
	if err != nil {
		log.Printf("Invited编码消息失败")
		panic(err)
	}

	log.Printf("发送邀请消息 imsgByte：%s", string(imsgByte))

	m := &ChatMessage{
		Type:		Invited,
		Message:    string(imsgByte),
		SenderID:   node.self.ID().Pretty(),
		SenderNick: node.cfg.nick,
	}

	msgBytes, err := json.Marshal(*m)
	if err != nil {
		panic(err)
	}

	for e := range cr.peers {
		info := cr.peers[e]
		node.self.Peerstore().AddAddrs(info.ID, info.Addrs, peerstore.PermanentAddrTTL)

		s, err := node.self.NewStream(node.ctx, info.ID, "multchat/Invited/1.0.0")
		if s == nil {
			log.Printf("Invited创建流失败")
			panic(err)
		}
		defer s.Close()

		if _, err := s.Write(msgBytes) ; err != nil {
			log.Printf("Invited发送消息失败 %s ",err)
		}

	}

}


func (node *Node)sendMessage(cr *ChatRoom, message string)  {

	m := &ChatMessage{
		Type:		SEND,
		Message:    message,
		SenderID:   node.self.ID().Pretty(),
		SenderNick: node.cfg.nick,
	}
	msgBytes, err := json.Marshal(*m)
	if err != nil {
		panic(err)
	}

	cr.Messages <- m

	for e := range cr.peers {
		info := cr.peers[e]
		log.Printf("发送消息 to  %s ",ToString(info))

		if node.self.ID().String() == info.ID.String(){
			continue
		}

		node.self.Peerstore().AddAddrs(info.ID, info.Addrs, peerstore.PermanentAddrTTL)

		s, _ := node.self.NewStream(node.ctx, info.ID,  "multchat/send/1.0.0")
		defer s.Close()

		if s == nil {
			log.Printf("创建流失败")
		}

		if _, err := s.Write(msgBytes) ; err != nil {
			log.Printf("发送消息失败 %s ",err)
		}

	}

}

///ip4/127.0.0.1/tcp/3001/p2p/12D3KooWAtyKjJGFABAaccn6SBT61rXwuHtFWKW9mLC2S5gesBpS
func GetInfoByMultiaddr(destination string) (p peer.AddrInfo) {
	maddr, err := multiaddr.NewMultiaddr(destination)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	// Extract the peer ID from the multiaddr.
	info, err := peer.AddrInfoFromP2pAddr(maddr)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return *info
}

//func ()writeData(rw *bufio.ReadWriter) {
//	stdReader := bufio.NewReader(os.Stdin)
//	for {
//		fmt.Print("> ")
//		sendData, err := stdReader.ReadString('\n')
//		if err != nil {
//			log.Println(err)
//			return
//		}
//
//		rw.WriteString(fmt.Sprintf("%s\n", sendData))
//		rw.Flush()
//	}
//}

type PeerIdentityJson struct {
	Name string
	Id string
	PrivKey string
	PubKey string
}

func readIdentityJson(filePath string) (crypto.PrivKey, crypto.PubKey, error) {
	// Open our jsonFile
	jsonFile, err := os.Open(filePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, nil, err
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var peerIdentity PeerIdentityJson

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &peerIdentity)

	privKeyByte, err := b64.StdEncoding.DecodeString(peerIdentity.PrivKey)
	if err != nil {
		return nil, nil, err
	}

	pvtKey, err := crypto.UnmarshalRsaPrivateKey(privKeyByte)
	if err != nil {
		return nil, nil, err
	}

	return pvtKey, pvtKey.GetPublic(), nil
}


// Create json file with identity information
func writeIdentityJson(privKey crypto.PrivKey, port int) {
	ID, err := peer.IDFromPrivateKey(privKey)
	if err != nil {
		panic(err)
	}
	_ = ID

	pvtBytes, err := privKey.Raw()
	if err != nil {
		panic(err)
	}
	_ = pvtBytes
	pubBytes, err := privKey.GetPublic().Raw()
	if err != nil {
		panic(err)
	}
	_ = pubBytes

	identityJson := &PeerIdentityJson {
		Name: "Relay Lich 1.0",
		Id: ID.Pretty(),
		PrivKey: b64.StdEncoding.EncodeToString(pvtBytes),
		PubKey: b64.StdEncoding.EncodeToString(pubBytes),
	}

	file, err := json.MarshalIndent(identityJson, "", " ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))

	_ = ioutil.WriteFile(fmt.Sprintf("relay-peer-key-%d.json", port), file, 0644)
}


func ToString(pi peer.AddrInfo) string {
	return fmt.Sprintf("%v/p2p/%v", pi.Addrs[0], pi.ID)
}