package current

import (
	"context"

	"github.com/fengjx/go-halo/logger"
)

var (
	LoggerKey  = "ctx.logger"
	TraceIDKey = "ctx.traceID"
	GoIDKey    = "ctx.goID"
)

func Logger(ctx context.Context) logger.Logger {
	if log, ok := ctx.Value(LoggerKey).(logger.Logger); ok {
		return log
	}
	return nil
}

func WithLogger(ctx context.Context, log logger.Logger) context.Context {
	return context.WithValue(ctx, LoggerKey, log)
}

func TraceID(ctx context.Context) string {
	return ctx.Value(TraceIDKey).(string)
}

func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, TraceIDKey, traceID)
}

func GoID(ctx context.Context) string {
	return ctx.Value(GoIDKey).(string)
}

func WithGoID(ctx context.Context, goID int64) context.Context {
	return context.WithValue(ctx, GoIDKey, goID)
}
