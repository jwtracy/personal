package greeter

import (
	"context"
	"fmt"

	"github.com/golang/glog"
	pb "github.com/johnwtracy/personal/src/apiserver/greeter/pb"
	"github.com/twitchtv/twirp"
)

type Server struct {
	Name     string
	Greeting string
}

func logReceived(ctx context.Context) (context.Context, error) {
	pkg, _ := twirp.PackageName(ctx)
	service, _ := twirp.ServiceName(ctx)
	glog.Infof("received: package: %s, service: %s", pkg, service)
	return ctx, nil
}

func newServerHooks() *twirp.ServerHooks {
	return &twirp.ServerHooks{
		RequestReceived: logReceived,
	}
}

func NewServer(name, greeting string) pb.TwirpServer {
	server := pb.NewGreeterServer(&Server{
		Name:     name,
		Greeting: greeting,
	}, newServerHooks())
	glog.Infof("creating new server: %s", server.PathPrefix())
	return server
}

func (s *Server) Hello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {

	method, _ := twirp.MethodName(ctx)
	glog.Infof("incoming %s request", method)

	return &pb.HelloReply{
		Message: fmt.Sprintf("Hello %s, I am %s. %s", request.Name, s.Name, s.Greeting),
	}, nil
}
