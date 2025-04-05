package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func (s *StatusRecorder) WriteHeader(code int) {
	s.Status = code
	s.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		sr := &StatusRecorder{
			ResponseWriter: w,
			Status:         http.StatusOK,
		}

		next.ServeHTTP(sr, r)

		elapsedTime := time.Since(startTime)

		slog.Info(
			"Request received",
			slog.String("method", r.Method),
			slog.String("url", r.URL.String()),
			slog.Int("status", sr.Status),
			slog.Int64("elapsed_time", elapsedTime.Milliseconds()),
		)
	}
}
