
https://developers.google.com/protocol-buffers/docs/proto3


### 安装 protoc-gen-go


// 在goland 使用proto 需要安装对应的protoc，protoc-gen-go
// 在终端运行命令 go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
// 会在gopath 目录下面生成一个 protoc-gen-go.exe
// 如果提示不是内部命令 就将这个复制到 C:\Users\Administrator\go\bin 即可


### 引入第三方的 proto文件
```
import "github.com/gogo/protobuf/gogoproto/gogo.proto";
import "github.com/mwitkow/go-proto-validators/validator.proto";
```

对应这种引入第三方的 proto文件，在编译的时候出现找不到 则需要将其下载到

protoc 安装目录下的include 中，

按照对应的目录将文件放入
```shell script
kevin@DESKTOP-SMS02PU:/mnt/e/softInstall/Java/protoc-21.6-win64/include$ tree
.
├── github.com
│   ├── gogo
│   │   └── protobuf
│   │       └── gogoproto
│   │           └── gogo.proto
│   └── mwitkow
│       └── go-proto-validators
│           └── validator.proto
├── google
│   ├── api
│   │   ├── annotations.proto
│   │   └── http.proto
│   └── protobuf                        #在下载安装的时候有
│       ├── any.proto
│       ├── api.proto
│       ├── compiler
│       │   └── plugin.proto
│       ├── descriptor.proto
│       ├── duration.proto
│       ├── empty.proto
│       ├── field_mask.proto
│       ├── source_context.proto
│       ├── struct.proto
│       ├── timestamp.proto
│       ├── type.proto
│       └── wrappers.proto
└── protoc-gen-swagger
    └── options
        ├── annotations.proto
        └── openapiv2.proto

12 directories, 18 files
kevin@DESKTOP-SMS02PU:/mnt/e/softInstall/Java/protoc-21.6-win64/include$
```