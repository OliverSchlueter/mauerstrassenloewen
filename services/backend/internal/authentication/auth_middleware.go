package authentication

import (
	"errors"
	"net/http"
	"strings"
)

const (
	authTokenHeader = "X-Auth-Token"
)

var (
	ErrMissingAuthToken = errors.New("missing auth token")
)

func (s *Store) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if !strings.HasPrefix(path, "/api") {
			next.ServeHTTP(w, r)
			return
		}

		token, err := getAuthTokenFromHeader(&r.Header)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		if !s.IsAuthTokenValid(token) {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}

		next.ServeHTTP(w, r)
	})
}

func getAuthTokenFromHeader(h *http.Header) (string, error) {
	token := h.Get(authTokenHeader)
	if token == "" {
		return "", ErrMissingAuthToken
	}

	return token, nil
}
