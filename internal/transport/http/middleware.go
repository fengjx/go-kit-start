package http

import (
	"net/http"

	"github.com/fengjx/go-halo/halo"
	"go.uber.org/zap"

	"github.com/fengjx/go-kit-start/common/logger"
	"github.com/fengjx/go-kit-start/internal/current"
)

func traceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		goid := halo.GetGoID()
		log := logger.Log.With(zap.Int64("goid", goid))
		ctx := current.WithGoID(r.Context(), goid)
		ctx = current.WithLogger(r.Context(), log)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
