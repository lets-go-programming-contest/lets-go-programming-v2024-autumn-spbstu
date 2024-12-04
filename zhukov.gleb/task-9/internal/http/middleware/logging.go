package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrappedWriter := &responseWriterStatus{ResponseWriter: w}

		defer func() {
			logger.Printf("%d %s %s %s %s %dms\n",
				wrappedWriter.statusCode,
				r.Method,
				r.URL.Path,
				r.RemoteAddr,
				r.UserAgent(),
				time.Since(start).Milliseconds())
		}()

		next.ServeHTTP(wrappedWriter, r)
	})
}

type responseWriterStatus struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriterStatus) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
