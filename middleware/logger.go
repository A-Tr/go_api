package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/satori/go.uuid"
)

func LoggerMiddleware(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		traceID := uuid.Must(uuid.NewV4())

		inner.ServeHTTP(w, r)
		log.Print("HOLI")
		log.Printf(
			"%s\t%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
			traceID.String(),
		)
	})
}
