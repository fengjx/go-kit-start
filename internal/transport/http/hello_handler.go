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
)

func decodeSayHelloRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	logger := current.Logger(ctx)
	request := &endpoint.SayHelloReq{}
	err := binding.ShouldBind(r, request)
	if err != nil {
		logger.Info("decode say hello err", zap.Error(err))
		errn := &errno.Errno{
			Code:     4,
			HttpCode: http.StatusBadRequest,
			Msg:      err.Error(),
		}
		return nil, errn
	}
	logger.Info("decode say hello", zap.Any("req", request))
	return request, nil
}

type HelloHandler struct {
}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

func (c HelloHandler) Bind(router *chi.Mux) {
	router.Route(openAPI+"/hello", func(r chi.Router) {
		r.Handle("/hi", c.sayHello())
	})
}

func (c HelloHandler) sayHello() *httptransport.Server {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
	}
	return httptransport.NewServer(
		endpoint.MakeSayHelloEndpoint(),
		decodeSayHelloRequest,
		encodeResponse,
		options...,
	)
}
