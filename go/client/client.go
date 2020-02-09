package main

import (
	"context"
	"fmt"
	"os"

	greeter "github.com/thinceller/rails-golang-grpc-example/go/lib"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5300", grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "grpc.Dial: %v\n", err)
		return
	}

	client := greeter.NewGreeterClient(conn)
	req := &greeter.HelloRequest{
		Name: "Kohei",
	}
	res, err := client.SayHello(context.Background(), req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "SayHello: %v\n", err)
		return
	}

	fmt.Fprintf(os.Stdout, "SayHello: %s\n", res.Message)
}
