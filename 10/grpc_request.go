package main

import (
	"fmt"

	"golang.org/x/net/context"
	pb "hello.pb"

	"google.golang.org/grpc"
)

func main() {
	address := "localhost:55555"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewHelloClient(conn)

	name := "Nacho Herrera"
	hr := &pb.HelloRequest{Name: name}
	r, err := c.Say(context.Background(), hr)
	if err != nil {
		panic(err)
	}

	fmt.Println(r.Message)
}
