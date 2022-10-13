

## libp2p options 配置项解析


### `ListenAddrStrings`
    配置 libp2p 监听给定的（未解析的）地址

### `ListenAddrs` 
    配置 libp2p 以监听给定的地址

### `NoSecurity`
    是一个完全禁用所有传输安全性的选项，它与所有其他传输安全协议不兼容

### `Muxer`
    配置为使用给定的流复用器

### `Transport`
	配置为使用给定的传输

### `Peerstore`
	配置为使用给定的 peerstore

### `PrivateNetwork`
	 配置为使用给定的专用网络
### `BandwidthReporter`
    配置为使用给定的带宽报告器
### `Identity`
    配置为使用给定的私钥来标识自己
### `ConnectionManager`
    配置为使用给定的连接管理器
### `AddrsFactory`
    配置 libp2p 以使用给定的地址工厂
### `EnableRelay`
    启用中继传输 （默认：启用）
### `DisableRelay`
    禁用中继传输
### `EnableRelayService`
	选择中继实现方式  relayv2
### `EnableAutoRelay`
	启用 AutoRelay 子系统
### `StaticRelays`
	配置已知中继
### `DefaultStaticRelays`
	默认中继
### `ForceReachabilityPublic`
	覆盖 AutoNAT 子系统中的自动可达性检测，强制本地节点相信它是外部可达的。
### `ForceReachabilityPrivate`
	覆盖 AutoNAT 子系统中的自动可达性检测，强制本地节点相信它在 NAT 后面并且外部不可访问。
### `EnableNATService`
	开启NAT服务
### `AutoNATServiceRateLimit`
	配置的默认速率限制
### `Ping`
	支持 ping 服务；默认启用
### `Routing`
	配置路由方式
### `NoListenAddrs`
	配置为默认不监听
### `NoTransports`
	配置为不启用任何传输
### `ProtocolVersion`
	识别协议
### `UserAgent`
	设置与识别协议一起发送的 libp2p 用户代理
### `MultiaddrResolver`
	设置 libp2p dns 解析器
### `EnableHolePunching`
    开启打孔
    通过启用 NATT 的对等方来启动和响应打孔尝试来启用 NAT 穿越，创建与其他对等方的直接/NAT 遍历连接。（默认：禁用）
### `WithDialTimeout`
    超时时间





