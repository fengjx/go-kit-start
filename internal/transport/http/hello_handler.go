package http

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
	"go.uber.org/zap"

	"github.com/fengjx/go-kit-start/common/errno"
	"github.com/fengjx/go-kit-start/internal/current"
	"github.com/fengjx/go-kit-start/internal/endpoint"
	"github.com/fengjx/go-kit-start/internal/transport/http/binding"
	"github.com/fengjx/go-kit-start/pb"
)

func (h *HelloHandler) decodeSayHelloRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	logger := current.Logger(ctx)
	helloReq := &pb.HelloReq{}
	err := binding.ShouldBind(r, helloReq)
	if err != nil {
		logger.Info("decode say hello err", zap.Error(err))
		errn := &errno.Errno{
			Code:     4,
			HttpCode: http.StatusBadRequest,
			Msg:      err.Error(),
		}
		return nil, errn
	}
	logger.Info("decode say hello", zap.Any("req", helloReq))
	return helloReq, nil
}

type HelloHandler struct {
}

func newHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

func (h *HelloHandler) Bind(router *chi.Mux) {
	router.Route(openAPI+"/greeter", func(r chi.Router) {
		r.Handle("/hi", h.sayHello())
	})
}

func (h *HelloHandler) sayHello() *httptransport.Server {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
	}
	return httptransport.NewServer(
		endpoint.GetInst().GreeterEndpoints.MakeSayHelloEndpoint(),
		h.decodeSayHelloRequest,
		encodeResponse,
		options...,
	)
}
