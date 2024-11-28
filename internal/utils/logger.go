package utils

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// LoggingMiddleware logs details of each HTTP request and response
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		reqID := uuid.New().String()

		// Log incoming request
		log.Info().
			Str("reqId", reqID).
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Str("host", r.Host).
			Str("remoteAddress", r.RemoteAddr).
			Msg("incoming request")

		// Capture the status code
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: 200}
		next.ServeHTTP(lrw, r)

		// Log response
		duration := time.Since(start).Seconds() * 1000 // in milliseconds
		log.Info().
			Str("reqId", reqID).
			Int("statusCode", lrw.statusCode).
			Float64("responseTime", duration).
			Msg("request completed")
	})
}

// loggingResponseWriter is a wrapper to capture HTTP status codes
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
