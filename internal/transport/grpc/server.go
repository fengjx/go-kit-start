package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/fengjx/go-halo/addr"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/fengjx/go-kit-start/common/logger"
	"github.com/fengjx/go-kit-start/internal/current"
)

type Server struct {
	grpcServer *grpc.Server
	address    string
}

func NewServer() *Server {
	server := grpc.NewServer(grpc.UnaryInterceptor(grpctransport.Interceptor))
	return &Server{
		grpcServer: server,
		address:    ":8090",
	}
}

func (s *Server) Start() {
	ln, err := net.Listen("tcp", s.address)
	if err != nil {
		panic(err)
	}
	address := ln.Addr().String()
	host, port, err := addr.ExtractHostPort(address)
	if err != nil {
		panic(err)
	}
	s.address = fmt.Sprintf("%s:%s", host, port)
	logger.Log.Infof("grpc server listening on %s", s.address)
	if err = s.grpcServer.Serve(ln); err != nil {
		panic(err)
	}
}

func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
}

type RegisterHandler func(grpcServer *grpc.Server)

func (s *Server) RegisterServer(rh RegisterHandler) *Server {
	rh(s.grpcServer)
	return s
}

type LogErrorHandler struct {
}

func NewLogErrorHandler() *LogErrorHandler {
	return &LogErrorHandler{}
}

func (h *LogErrorHandler) Handle(ctx context.Context, err error) {
	log := current.Logger(ctx)
	log.Warn("handle grpc err", zap.String("traceId", current.TraceID(ctx)), zap.Error(err))
}
