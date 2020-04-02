package main

import (
	"context"
	"go_grpc_helloworld_example/helloworld"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloResponse, error) {
	log.Println("Received name:", in.GetName())
	return &helloworld.HelloResponse{
		Message: "Hello, " + in.GetName(),
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)

	if err != nil {
		log.Println("failed to listen, error:", err)
	}

	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Println("failed to serve, error:", err)
	}
}
