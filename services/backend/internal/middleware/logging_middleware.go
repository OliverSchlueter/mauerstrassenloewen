package middleware

import (
	"github.com/prometheus/client_golang/prometheus"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

var (
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "backend_requests_total",
			Help: "Total number of requests processed by the backend http server.",
		},
		[]string{"method", "path", "status", "elapsed_time"},
	)
)

type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func (s *StatusRecorder) WriteHeader(code int) {
	s.Status = code
	s.ResponseWriter.WriteHeader(code)
}

func RegisterPrometheusHttpLogging() {
	prometheus.MustRegister(RequestCount)
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

		RequestCount.WithLabelValues(
			r.Method,
			r.URL.Path,
			strconv.Itoa(sr.Status),
			strconv.FormatInt(elapsedTime.Milliseconds(), 10),
		).Inc()
	}
}
