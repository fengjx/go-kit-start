package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/fengjx/go-kit-start/internal/current"
	"github.com/fengjx/go-kit-start/internal/service/hello"
)

type SayHelloReq struct {
	Name string `json:"name"`
}

type SayHelloResp struct {
	Msg string `json:"msg"`
}

func MakeSayHelloEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logger := current.Logger(ctx)
		logger.Info("endpoint say hello")
		req := request.(*SayHelloReq)
		msg := hello.SayHello(ctx, req.Name)
		resp := &SayHelloResp{Msg: msg}
		return resp, nil
	}
}
