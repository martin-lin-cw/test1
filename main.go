package main

import (
	"context"
	"fmt"
	"net"
	"time"

	test1pb "github.com/martin-lin-cw/test1/gen/proto/test1/v1"
	test2pb "github.com/martin-lin-cw/test1/gen/proto/test2/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(fmt.Errorf("new client: %w", err))
	}

	client := test2pb.NewTest2ServiceClient(conn)
	fmt.Println("client created")

	for i := 0; i < 10; i++ {
		resp, err := client.Hello2(context.Background(), &test2pb.Hello2Request{})
		if err != nil {
			fmt.Println(fmt.Errorf("call test2 Hello2: %w", err))
		}

		fmt.Println(resp.GetResult())
	}

	time.Sleep(time.Second * 10)
	conn.Close()
}
