package main

import (
	"context"
	"fmt"
	"net"
	"time"

	test1pb "github.com/martin-lin-cw/test1/gen/proto/test1/v1"
	"google.golang.org/grpc"
)

type Test1Service struct {
	test1pb.UnsafeTest1ServiceServer
}

// Hello1 implements test1pb.Test1ServiceServer.
func (t *Test1Service) Hello1(context.Context, *test1pb.Hello1Request) (*test1pb.Hello1Response, error) {
	return &test1pb.Hello1Response{Result: "Hello1"}, nil
}

func NewTest1Service() test1pb.Test1ServiceServer {
	return &Test1Service{}
}

func main() {
	fmt.Println("Hello, test1")

	server := grpc.NewServer()

	test1Server := NewTest1Service()

	test1pb.RegisterTest1ServiceServer(server, test1Server)

	ln, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	go server.Serve(ln)

	time.Sleep(time.Second * 3)
}
