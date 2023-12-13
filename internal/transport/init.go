package transport

import (
	"context"

	"github.com/fengjx/go-kit-start/common/logger"
	"github.com/fengjx/go-kit-start/internal/transport/grpc"
	"github.com/fengjx/go-kit-start/internal/transport/http"
)

func Start(ctx context.Context) {
	logger.Log.Info("transport init")
	http.Start()
	grpc.Start()
}

func Stop() {
	http.Stop()
	grpc.Stop()
}
