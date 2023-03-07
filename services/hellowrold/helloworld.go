package hellowrold

import (
	"context"
	"fmt"

	pb "main.go/proto/hellowrold"
)

type Server struct {
	pb.UnimplementedHellowroldServer
}
func (s Server)SayHello(ctx context.Context,r *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: fmt.Sprintf("Hello %v",r.GetName()),
	},nil
}