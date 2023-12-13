package http

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/fengjx/go-kit-start/common/config"
	"github.com/fengjx/go-kit-start/common/errno"
	"github.com/fengjx/go-kit-start/common/logger"
	"github.com/fengjx/go-kit-start/internal/current"
)

const (
	openAPI  = "/open/api"
	innerAPI = "/inner/api"
	adminAPI = "/admin/api"
)

type result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

// 统一返回值处理
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	res := &result{
		Msg:  "ok",
		Data: response,
	}
	traceLog := current.Logger(ctx)
	traceLog.Info("http response", zap.Any("data", res))
	return json.NewEncoder(w).Encode(res)
}

// 统一异常处理
func errorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	traceLog := current.Logger(ctx)
	traceLog.Error("handler error", zap.Error(err))
	httpCode := 500
	msg := errno.SystemErr.Msg
	var errn *errno.Errno
	ok := errors.As(err, &errn)
	if ok && errn.HttpCode > 0 {
		httpCode = errn.HttpCode
		msg = errn.Msg
	}
	w.WriteHeader(httpCode)
	res := &result{
		Code: httpCode,
		Msg:  msg,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		traceLog.Error("write error msg fail", zap.Error(err))
	}
}

type Server struct {
	httpServer *http.Server
	router     *chi.Mux
}

func NewServer() *Server {
	serverConfig := config.GetConfig().Server
	router := chi.NewRouter()
	httpServer := &http.Server{
		Addr:    serverConfig.Listen,
		Handler: router,
	}
	svr := &Server{
		httpServer: httpServer,
		router:     router,
	}
	return svr
}

func (s *Server) Start() {
	logger.Log.Infof("server listening on %s", s.httpServer.Addr)
	if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Panicf("listen: %s\n", err)
	}
}

func (s *Server) Stop() {
	// Graceful HTTP server shutdown.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.httpServer.Shutdown(ctx); err != nil {
		logger.Log.Error("http server stop err: %v", zap.Error(err))
	} else {
		logger.Log.Infof("http server was shutdown gracefully")
	}
}

type Middleware func(http.Handler) http.Handler

func (s *Server) Use(middlewares ...Middleware) *Server {
	for _, middleware := range middlewares {
		s.router.Use(middleware)
	}
	return s
}

func (s *Server) Add(handlers ...Handler) *Server {
	for _, handler := range handlers {
		handler.Bind(s.router)
	}
	return s
}

type Handler interface {
	Bind(router *chi.Mux)
}
