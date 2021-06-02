
package main

import (
	"context"
	"fmt"
	"github.com/toransahu/grpc-eg-go/hello"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello client ...")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := hello.NewHelloServiceClient(cc)
	request := &hello.HelloRequest{Name: "Jeremy"}

	resp, _ := client.Hello(context.Background(), request)
	fmt.Printf("Receive response => [%v]", resp.Greeting)
}