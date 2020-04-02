package main

import (
	"context"
	"go_grpc_helloworld_example/helloworld"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "HelloWorld"
)

func main() {
	// 建立连接
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("can not connect, error: ", err)
	}

	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)

	// 构造request需要的参数
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// 构建context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// grpc调用
	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: name})
	if err != nil {
		log.Println("could not SayHello, error:", err)
	}
	log.Println("Get response:", r.GetMessage())
}
