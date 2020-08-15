package greeter

import (
	"context"
	"fmt"

	pb "github.com/johnwtracy/personal/src/apiserver/greeter/pb"
)

type Server struct {
	Name     string
	Greeting string
	Port     int
}

func NewServer(name, greeting string, port int) *Server {
	return &Server{
		Name:     name,
		Greeting: greeting,
		Port:     port,
	}
}

func (s *Server) Hello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: fmt.Sprintf("Hello %s, I am %s. %s", request.Name, s.Name, s.Greeting),
	}, nil
}
