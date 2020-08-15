package main

import (
	"fmt"
	"log"
	"net"

	"github.com/johnwtracy/personal/src/apiserver/greeter"
	pb "github.com/johnwtracy/personal/src/apiserver/greeter/pb"
	"google.golang.org/grpc"
)

func main() {
	server := greeter.NewServer("John Tracy", "See you, space cowboy!", 8080)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", server.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, server)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
