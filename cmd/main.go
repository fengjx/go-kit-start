package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/fengjx/go-kit-start/common/logger"
	"github.com/fengjx/go-kit-start/internal/endpoint"
	"github.com/fengjx/go-kit-start/internal/facade"
	"github.com/fengjx/go-kit-start/internal/integration/dbcli"
	"github.com/fengjx/go-kit-start/internal/integration/rediscli"
	"github.com/fengjx/go-kit-start/internal/service"
	"github.com/fengjx/go-kit-start/internal/transport"
)

func main() {
	logger.Log.Info("app start")
	ctx, cancel := context.WithCancel(context.Background())
	dbcli.Init()
	rediscli.Init()
	facade.Init()
	service.Init()
	endpoint.Init()
	transport.Start(ctx)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	<-quit
	cancel()
	transport.Stop()
}
