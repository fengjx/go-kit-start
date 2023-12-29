package transport

import (
	"context"

	"github.com/fengjx/go-kit-start/common/logger"
	"github.com/fengjx/go-kit-start/internal/transport/grpc"
	"github.com/fengjx/go-kit-start/internal/transport/http"
)

func Start(_ context.Context) {
	logger.Log.Info("transport init")
	go http.Start()
	go grpc.Start()
}

func Stop() {
	http.Stop()
	grpc.Stop()
}
