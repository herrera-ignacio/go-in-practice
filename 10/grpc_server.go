package main

import (
	"context"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "hello.proto"
)

type server struct{}

func (s *server) Say(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	msg := "Hello " + in.Name + "!"
	return &pb.HelloResponse{Message: msg}, nil
}

func main() {
	l, err := net.Listen("tcp", ":55555")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})
	s.Serve(l)
}
