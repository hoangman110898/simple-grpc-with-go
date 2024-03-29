package main

import (
	"context"
	"fmt"
	"github.com/toransahu/grpc-eg-go/hello"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {

}

func (*server) Hello(ctx context.Context, request *hello.HelloRequest) (*hello.HelloResponse, error) {
	name := request.Name
	response := &hello.HelloResponse{
		Greeting: "Hello " + name,
	}
	return response, nil
}



func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	hello.RegisterHelloServiceServer(s, &server{})

	s.Serve(lis)
}