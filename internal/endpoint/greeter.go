package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/fengjx/go-kit-start/internal/current"
	"github.com/fengjx/go-kit-start/internal/logic"
	"github.com/fengjx/go-kit-start/pb"
)

type greeterEndpoints struct {
}

func newGreeterEndpoints() *greeterEndpoints {
	return &greeterEndpoints{}
}

func (e *greeterEndpoints) MakeSayHelloEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logger := current.Logger(ctx)
		logger.Info("greeter say hello")
		helloReq := request.(*pb.HelloReq)
		msg, err := logic.GetInst().HelloLogic.SayHello(ctx, helloReq.Name)
		if err != nil {
			return nil, err
		}
		resp := &pb.HelloResp{
			Message: msg,
		}
		return resp, nil
	}
}
