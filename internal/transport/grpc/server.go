package grpc

import (
	userpb "github.com/TouchGlass/project-protos/proto/user"
	"github.com/TouchGlass/users-service/internal/user"
	"google.golang.org/grpc"
	"net"
)

func RunGRPC(svc user.UserService) error {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	grpcSrv := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcSrv, NewHandler(svc))

	return grpcSrv.Serve(listener)

}
