package authentication

import (
	"errors"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/sloki"
	"log/slog"
	"net/http"
	"strings"
)

const (
	authTokenHeader    = "X-Auth-Token"
	authUserHeader     = "X-Auth-User"
	authPasswordHeader = "X-Auth-Password"
)

var (
	ErrMissingAuthToken    = errors.New("missing auth token")
	ErrMissingAuthUser     = errors.New("missing auth user")
	ErrMissingAuthPassword = errors.New("missing auth password")
)

func (s *Store) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if !strings.HasPrefix(path, "/api") {
			next.ServeHTTP(w, r)
			return
		}

		user, password, err := getUserAndPassword(&r.Header)
		if err != nil {
			slog.Warn("Could not get auth user and password from header", sloki.WrapError(err))
		}

		if len(user) == 0 && len(password) == 0 {
			if !s.IsAuthUserValid(user, password) {
				slog.Warn("Auth user is not valid")
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
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

func getUserAndPassword(h *http.Header) (user, password string, err error) {
	user, err = getAuthUserFromHeader(h)
	if err != nil {
		if errors.Is(err, ErrMissingAuthUser) {
			return "", "", nil
		} else {
			slog.Warn("Could not get auth user from header", sloki.WrapError(err))
		}
	}

	password, err = getAuthPasswordFromHeader(h)
	if err != nil {
		if errors.Is(err, ErrMissingAuthPassword) {
			return "", "", nil
		} else {
			slog.Warn("Could not get auth password from header", sloki.WrapError(err))
		}
	}

	return user, password, nil
}

func getAuthTokenFromHeader(h *http.Header) (string, error) {
	token := h.Get(authTokenHeader)
	if token == "" {
		return "", ErrMissingAuthToken
	}

	return token, nil
}

func getAuthUserFromHeader(h *http.Header) (string, error) {
	user := h.Get(authUserHeader)
	if user == "" {
		return "", ErrMissingAuthUser
	}

	return user, nil
}

func getAuthPasswordFromHeader(h *http.Header) (string, error) {
	password := h.Get(authPasswordHeader)
	if password == "" {
		return "", ErrMissingAuthPassword
	}

	return password, nil
}
