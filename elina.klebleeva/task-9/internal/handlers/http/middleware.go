package handlers

import (
	"log/slog"
	"net/http"
	"time"
)

// Implement own responseWriter type that captures the status code of a response.
type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		status:         200,
		ResponseWriter: w,
	}
}

func LoggingMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		log := logger.With(slog.String("component", "middleware/logger"))

		fn := func(w http.ResponseWriter, r *http.Request) {
			rw := wrapResponseWriter(w)

			entry := log.With(
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("remote_addr", r.RemoteAddr),
				slog.String("user_agent", r.UserAgent()),
			)

			start := time.Now()
			defer func() {
				entry.Info("request complete",
					slog.Int("status", rw.Status()),
					slog.String("duration", time.Since(start).String()),
				)
			}()

			next.ServeHTTP(rw, r)
		}

		return http.HandlerFunc(fn)
	}
}
