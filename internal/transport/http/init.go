package http

var server *Server

func Start() {
	httpServer := NewServer()
	httpServer.Use(
		traceMiddleware,
	).Add(
		NewHelloHandler(),
	).Start()
}

func Stop() {
	server.Stop()
}
