package middleware

import (
	"github.com/OliverSchlueter/mauerstrassenloewen/common/sloki"
	"log/slog"
	"net/http"
	"runtime/debug"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Log the panic and stack trace
				slog.Error("Panic recovered", sloki.WrapError(err.(error)))
				debug.PrintStack()

				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
