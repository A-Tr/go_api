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
			traceID, error := uuid.NewV4()
			if error != nil {
				logger.WithField("trace-id", traceID).Warn("<could not create uuid")
			}
			w.Header().Set("traceId", traceID.String())
			logger.Info("Soy el middleware", traceID.String())
			h.ServeHTTP(w, r)
		})
	}
}
