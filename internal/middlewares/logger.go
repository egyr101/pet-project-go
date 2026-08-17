package middlewares

import (
	"log/slog"
	"net/http"
	"time"
)

type responseWriter struct{
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int){
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func Logger(next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(rw, r)

		slog.Info("http request",
			"method", r.Method,
			"path", r.URL.Path,
			"statusCode", rw.statusCode,
			"duration_ms", time.Since(start).Milliseconds(),
			"request_id", GetIDFromContext(r.Context()),
		)
	})
}