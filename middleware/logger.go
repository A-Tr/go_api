package middleware

import (
	"net/http"

	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

type Adapter func(http.Handler) http.Handler

func LoggerMiddleware(logger *log.Logger) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			traceID := uuid.Must(uuid.NewV4()).String()
			w.Header().Set("traceId", traceID)
			logger.Info("Soy el middleware", traceID)
			h.ServeHTTP(w, r)
		})
	}
}
