package grpc

import (
	"google.golang.org/grpc"

	"github.com/fengjx/go-kit-start/pb"
)

var server *Server

func Start() {
	server = NewServer()
	server.RegisterServer(func(grpcServer *grpc.Server) {
		pb.RegisterGreeterServer(grpcServer, newGreeterServer())
	}).Start()
}

func Stop() {
	server.Stop()
}
