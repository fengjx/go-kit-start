package http

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	httptransport "github.com/go-kit/kit/transport/http"
	"go.uber.org/zap"

	"github.com/fengjx/go-kit-start/internal/current"
	"github.com/fengjx/go-kit-start/internal/endpoint"
)

func decodeSayHelloRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	name := r.FormValue("name")
	request := &endpoint.SayHelloReq{Name: name}
	logger := current.Logger(ctx)
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
