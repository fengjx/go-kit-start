package grpc

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"

	"github.com/fengjx/go-kit-start/common/logger"
	"github.com/fengjx/go-kit-start/internal/current"
	"github.com/fengjx/go-kit-start/internal/endpoint"
	"github.com/fengjx/go-kit-start/pb"
)

type GreeterServer struct {
	pb.UnimplementedGreeterServer
	sayHello grpctransport.Handler
}

func (s *GreeterServer) SayHello(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	_, resp, err := s.sayHello.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.HelloResp), nil
}

func decodeSayHello(_ context.Context, req interface{}) (interface{}, error) {
	helloReq := req.(*pb.HelloReq)
	return &pb.HelloReq{
		Name: helloReq.Name,
	}, nil
}

func encodeSayHello(_ context.Context, resp interface{}) (interface{}, error) {
	helloResp := resp.(*pb.HelloResp)
	return &pb.HelloResp{
		Message: helloResp.Message,
	}, nil
}

func newGreeterServer() pb.GreeterServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerBefore(func(ctx context.Context, md metadata.MD) context.Context {
			traceID := uuid.NewString()
			if len(md.Get(current.TraceIDKey)) > 0 {
				traceID = md.Get(current.TraceIDKey)[0]
			}
			ctx = current.WithTraceID(ctx, traceID)
			log := logger.Log.With(zap.String("traceID", traceID))
			ctx = current.WithLogger(ctx, log)
			return ctx
		}),
		grpctransport.ServerErrorHandler(NewLogErrorHandler()),
	}
	return &GreeterServer{
		sayHello: grpctransport.NewServer(
			endpoint.GetInst().GreeterEndpoints.MakeSayHelloEndpoint(),
			decodeSayHello,
			encodeSayHello,
			options...,
		),
	}
}
