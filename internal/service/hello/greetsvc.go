package hello

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/fengjx/go-kit-start/internal/current"
)

type greetService struct {
}

func newGreetService() *greetService {
	return &greetService{}
}

func (svc *greetService) sayHello(ctx context.Context, name string) string {
	logger := current.Logger(ctx)
	logger.Info("say hello", zap.Any("name", name))
	return fmt.Sprintf("hello: %s", name)
}
