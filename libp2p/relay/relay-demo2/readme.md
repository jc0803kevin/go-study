

通过指定的中继，让在nat后面的两个节点进行通信

参考
https://github.com/zfunnily/hole-punching-go-libp2p
https://zfunnily.github.io/2021/10/gop2pthree/

/mnt/e/workspace/gowork/src/github.com/jc0803kevin/go-study/libp2p/relay/relay-demo2

先实现中继服务
监听一条协议RelayProtocolRequest，当有一个节点上来，需要记录，当有两个节点时，返回两个节点的信息给对方节点。
有一个公网IP


节点需要的组件:

创建host
监听RelayMsg协议 和RelayProtocolResponse协议
连接中继服务 n.ConnectRelay()
节点发送RelayProtocolRequest协议给中继服务,服务会返回RelayProtocolResponse协议给节点 其他节点的peerID
通过中继服务连接其他节点n.ConnectNode() ,
并且往其他节点发送RelayMsg协议, n.CmdRelay


第一步，双方节点先连接中继节点 
第二步，其中一节点h1通过中继服务h2获取到另一方的节点h3地址(这个例子直接就能通过内存获取到对方的地址) 
第三步，h1连接h3，并通过h3监听的协议"/cats"新建数据流，用于收发数据。


## 运行

### 编译

```shell script
kevin@DESKTOP-SMS02PU:/mnt/e/workspace/gowork/src/github.com/jc0803kevin/go-study/libp2p/relay/relay-demo2$ go build -o chat chat.go
go: downloading github.com/libp2p/go-libp2p v0.22.0
go: downloading github.com/ipfs/go-cid v0.2.0
go: downloading github.com/multiformats/go-multicodec v0.5.0
go: downloading github.com/libp2p/go-yamux/v3 v3.1.2
go: downloading github.com/klauspost/compress v1.15.1
go: downloading github.com/lucas-clemente/quic-go v0.28.1
go: downloading github.com/prometheus/client_golang v1.12.1
go: downloading go.uber.org/zap v1.22.0
go: downloading github.com/multiformats/go-base32 v0.0.4
go: downloading github.com/klauspost/cpuid/v2 v2.1.0
go: downloading golang.org/x/net v0.0.0-20220812174116-3211cb980234
go: downloading github.com/marten-seemann/qtls-go1-17 v0.1.2
kevin@DESKTOP-SMS02PU:/mnt/e/workspace/gowork/src/github.com/jc0803kevin/go-study/libp2p/relay/relay-demo2$ cd server/
kevin@DESKTOP-SMS02PU:/mnt/e/workspace/gowork/src/github.com/jc0803kevin/go-study/libp2p/relay/relay-demo2/server$
kevin@DESKTOP-SMS02PU:/mnt/e/workspace/gowork/src/github.com/jc0803kevin/go-study/libp2p/relay/relay-demo2/server$ go build -o relay_server relay-server.go
kevin@DESKTOP-SMS02PU:/mnt/e/workspace/gowork/src/github.com/jc0803kevin/go-study/libp2p/relay/relay-demo2/server$ ll
total 29704
drwxrwxrwx 1 kevin kevin      512 Oct 13 11:09 ./
drwxrwxrwx 1 kevin kevin      512 Oct 13 11:09 ../
-rwxrwxrwx 1 kevin kevin     3900 Oct 13 10:46 relay-server.go*
-rwxrwxrwx 1 kevin kevin 30411642 Oct 13 11:09 relay_server*
kevin@DESKTOP-SMS02PU:/mnt/e/workspace/gowork/src/github.com/jc0803kevin/go-study/libp2p/relay/relay-demo2/server$
```

### 启动中继节点
```shell script
ubuntu@VM-16-15-ubuntu:~/libp2p$ ./relay_server -sp 5001
2022/10/13 11:15:45 Run './chat -relay /ip4/10.0.16.15/tcp/5001/p2p/12D3KooWRgBS7MboAKshjGAXu3kBdtBS3xHQEM16ZFrLNPUf2WDe' on another console.
2022/10/13 11:15:45 You can replace 192.168.0.100 with public IP as well.
2022/10/13 11:15:45 Waiting for incoming connection
2022/10/13 11:15:45 
2022/10/13 11:17:31 a new Stream relay req, remotePeerID: 12D3KooWDef74MKwHZci27aWUHJfE1BnEYrMYyWoBFGqM4dqszXp; 
remoteAddr: /ip4/10.0.16.15/tcp/44048, localAddr: /ip4/10.0.16.15/tcp/5001
2022/10/13 11:18:00 a new Stream relay req, remotePeerID: 12D3KooWNJpkYPmiYvUrGCS4Tdj3gyuYFBDtrjPuezqGi6tmpfSW; 
remoteAddr: /ip4/116.228.34.130/tcp/57980, localAddr: /ip4/10.0.16.15/tcp/5001
2022/10/13 11:18:00 start send peer:  12D3KooWNJpkYPmiYvUrGCS4Tdj3gyuYFBDtrjPuezqGi6tmpfSW
2022/10/13 11:18:00 start send peer:  12D3KooWNJpkYPmiYvUrGCS4Tdj3gyuYFBDtrjPuezqGi6tmpfSW

```

### 在中继所在服务器内网 启动一个节点

```shell script
ubuntu@VM-16-15-ubuntu:~/libp2p$ ./chat -relay /ip4/10.0.16.15/tcp/5001/p2p/12D3KooWRgBS7MboAKshjGAXu3kBdtBS3xHQEM16ZFrLNPUf2WDe
2022/10/13 11:17:31 Connect relay success Next Login...
2022/10/13 11:17:31 12D3KooWDef74MKwHZci27aWUHJfE1BnEYrMYyWoBFGqM4dqszXp login to relay
2022/10/13 11:17:31 Im  12D3KooWDef74MKwHZci27aWUHJfE1BnEYrMYyWoBFGqM4dqszXp


2022/10/13 11:18:00 a new Stream Relay rsp,addrs.len 2, peerInfo : 12D3KooWNJpkYPmiYvUrGCS4Tdj3gyuYFBDtrjPuezqGi6tmpfSW=/ip4/116.228.34.130/tcp/57980
12D3KooWDef74MKwHZci27aWUHJfE1BnEYrMYyWoBFGqM4dqszXp=/ip4/10.0.16.15/tcp/44048

2022/10/13 11:18:00 start ConnectNodeByRelay.... 12D3KooWNJpkYPmiYvUrGCS4Tdj3gyuYFBDtrjPuezqGi6tmpfSW /ip4/116.228.34.130/tcp/57980
2022/10/13 11:18:00 receive onFirstP2PMsg:  Hello Im 12D3KooWDef74MKwHZci27aWUHJfE1BnEYrMYyWoBFGqM4dqszXp
2022/10/13 11:18:00 NewStream err
2022/10/13 11:18:00 protocol not supported
2022/10/13 11:18:05 Cmdp2p start ....
2022/10/13 11:18:05 opening p2p chat stream
2022/10/13 11:18:05 [INFO] p2p chat connected!
> > > > Hello this is  msgkkk
> I am romote
> 
```

### 在本机启动一个 节点
```shell script
kevin@DESKTOP-SMS02PU:/mnt/e/workspace/gowork/src/github.com/jc0803kevin/go-study/libp2p/relay/relay-demo2$ ./chat -relay /ip4/110.42.187.187/tcp/5001/p2p/12D3KooWRgBS7MboAKshjGAXu3kBdtBS3xHQEM16ZFrLNPUf2WDe
2022/10/13 11:18:01 Connect relay success Next Login...
2022/10/13 11:18:01 12D3KooWNJpkYPmiYvUrGCS4Tdj3gyuYFBDtrjPuezqGi6tmpfSW login to relay
2022/10/13 11:18:01 Im  12D3KooWNJpkYPmiYvUrGCS4Tdj3gyuYFBDtrjPuezqGi6tmpfSW
2022/10/13 11:18:01 a new Stream Relay rsp,addrs.len 2, peerInfo : 12D3KooWNJpkYPmiYvUrGCS4Tdj3gyuYFBDtrjPuezqGi6tmpfSW=/ip4/116.228.34.130/tcp/57980
12D3KooWDef74MKwHZci27aWUHJfE1BnEYrMYyWoBFGqM4dqszXp=/ip4/10.0.16.15/tcp/44048

2022/10/13 11:18:01 start ConnectNodeByRelay.... 12D3KooWDef74MKwHZci27aWUHJfE1BnEYrMYyWoBFGqM4dqszXp /ip4/10.0.16.15/tcp/44048
2022/10/13 11:18:01 receive onFirstP2PMsg:  Hello Im 12D3KooWNJpkYPmiYvUrGCS4Tdj3gyuYFBDtrjPuezqGi6tmpfSW
2022/10/13 11:18:01 NewStream err
2022/10/13 11:18:01 protocol not supported
2022/10/13 11:18:06 Cmdp2p start ....
2022/10/13 11:18:06 opening p2p chat stream
2022/10/13 11:18:06 [INFO] p2p chat connected!
> Hello this is  msg
> > kkk
> I am romote
>

```
