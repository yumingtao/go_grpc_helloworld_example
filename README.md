# 安装grpc
- git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc
- git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
- git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text
- git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto
- go get -u github.com/golang/protobuf/proto
- go get -u github.com/golang/protobuf/protoc-gen-go
- cd $GOPATH/src/
- go install google.golang.org/grpc

# 开发helloworld
- 创建helloworld目录
- 编写helloworld.proto
- 在项目根目录运行
protoc helloworld/helloworld.proto --go_out=plugins=grpc:.
- 编写服务端
- 编写client端
- 运行服务端, go run hello_server.go
- 运行client端, go run hello_client.go

# 参考
- https://www.qfgolang.com/?special=grpcweifuwukuangjia&pid=2668
- https://blog.didiyun.com/index.php/2018/12/12/grpc-golang-1/
- https://github.com/grpc/grpc-go/tree/master/examples/helloworld
