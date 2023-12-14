package http

import "github.com/go-chi/chi/v5/middleware"

var server *Server

func Start() {
	httpServer := NewServer()
	httpServer.Use(
		middleware.RealIP,
		middleware.RequestID,
		middleware.Recoverer,
		traceMiddleware,
		middleware.Logger,
	).Add(
		NewHelloHandler(),
	).Start()
}

func Stop() {
	server.Stop()
}
