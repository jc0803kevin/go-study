package main

import (
	"encoding/hex"
	"fmt"
	"github.com/libp2p/go-libp2p-core/peer"
	"testing"

	ic "github.com/libp2p/go-libp2p-core/crypto"
)

// https://github.com/libp2p/specs/blob/master/peer-ids/peer-ids.md
// PeerID生成的过程：创建私钥 => 导出公钥 => hash函数 => 编码
// bafzbeie5745rpv2m6tjyuugywy4d5ewrqgqqhfnf445he3omzpjbx5xqxe-- Peer ID (sha256) 编码为 CID ( inspect )。
// QmYyQSo1c1Ym7orWxLYvCrM2EmxFTANf8wXmmE7DWjhx5N-- Peer ID (sha256) 编码为原始 base58btc 多哈希。
// 12D3KooWD3eckifWpRn9wQpMG9R9hX3sD158z7EqHWmweQAJU5SA-- 对等 ID（ed25519，使用“身份”多重哈希）编码为原始 base58btc 多重哈希。
func TestPeerIdEncode(t *testing.T) {

	//fmt.Println(PeerIdEncode("/ip4/127.0.0.1/tcp/3001"))
	//fmt.Println(PeerIdEncode("/ip4/127.0.0.1/tcp/3001/p2p/QmdXGaeGiVA745XorV1jr11RHxB9z4fqykm6xCUPX1aTJo"))

	//_, pubKey, _ := ic.GenerateKeyPair(ic.Secp256k1, 2048)
	//pubByte, _ := pubKey.Raw()
	//pubKeyHex := hex.EncodeToString(pubByte)
	//fmt.Println("pubKeyHex -->" ,pubKeyHex )

	// peeId 是根据 启动节点时候的私钥 -》 公钥 ，利用公告的来的
	pubByte, _ := hex.DecodeString("02c97dec1771eb564be5ab7f1a203d11c81f14dbae27fbdd2a4837f9e4843c3dab")
	pubKey, _ := ic.UnmarshalSecp256k1PublicKey(pubByte)

	//16Uiu2HAm8zDHydQLQoy2CHdTcAt4vVDNpZBHyzXjqpgNVSh2hZtE
	//bafzaajiiaijccawjpxwbo4plkzf6lk37diqd2eoid4knxlrh7posusbx7hsiipb5vm
	peerId, _ := peer.IDFromPublicKey(pubKey)
	fmt.Println("peerId   :", peerId)
	fmt.Println("ToCid    :", peer.ToCid(peerId))

}
