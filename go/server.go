package main

import (
	"context"
	"log"
	"net"

	greeter "github.com/thinceller/rails-golang-grpc-example/go/lib"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":5300")
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
		return
	}

	grpcSrv := grpc.NewServer()
	greeter.RegisterGreeterServer(grpcSrv, &server{})
	log.Printf("Greeter service is runnning!")
	grpcSrv.Serve(listener)
}

type server struct{}

func (s *server) SayHello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloReply, error) {
	requestName := req.GetName()
	out := &greeter.HelloReply{
		Message: "Hello, " + requestName,
	}
	return out, nil
}
