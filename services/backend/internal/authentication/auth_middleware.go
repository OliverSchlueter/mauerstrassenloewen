package authentication

import (
	"errors"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/usermanagement"
	"github.com/OliverSchlueter/mauerstrassenloewen/common/sloki"
	"log/slog"
	"net/http"
	"strings"
)

const (
	authTokenHeader    = "X-Auth-Token"
	authUserHeader     = "X-Auth-Username"
	authPasswordHeader = "X-Auth-Password"
)

var (
	ErrMissingAuthToken    = errors.New("missing auth token")
	ErrMissingAuthUsername = errors.New("missing auth username")
	ErrMissingAuthPassword = errors.New("missing auth password")
)

func (s *Store) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if !strings.HasPrefix(path, "/api") || strings.HasPrefix(path, "/api/v1/user/register") || strings.HasPrefix(path, "/api/v1/user/login") {
			next.ServeHTTP(w, r)
			return
		}

		// Try username and password from the header
		username, password, err := getUsernameAndPasswordFromHeader(&r.Header)
		if err != nil {
			slog.Warn("Could not get auth user and password from header", sloki.WrapError(err))
		}

		var u *usermanagement.User
		if len(username) > 0 && len(password) > 0 {
			u, err = s.IsAuthUserValid(r.Context(), username, password)
			if err != nil {
				slog.Warn("Could not check auth user", sloki.WrapError(err))
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		}

		// Try to get the user from the auth token
		if u == nil {
			token, err := getAuthTokenFromHeader(&r.Header)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			u, err = s.IsAuthTokenValid(r.Context(), token)
			if err != nil {
				slog.Warn("Could not check auth token", sloki.WrapError(err))
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		}

		if u == nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		r.WithContext(writeUser(r.Context(), u))
		next.ServeHTTP(w, r)
	})
}

func getUsernameAndPasswordFromHeader(h *http.Header) (username, password string, err error) {
	username, err = getAuthUsernameFromHeader(h)
	if err != nil {
		if errors.Is(err, ErrMissingAuthUsername) {
			return "", "", nil
		} else {
			slog.Warn("Could not get auth username from header", sloki.WrapError(err))
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

	return username, password, nil
}

func getAuthTokenFromHeader(h *http.Header) (string, error) {
	token := h.Get(authTokenHeader)
	if token == "" {
		return "", ErrMissingAuthToken
	}

	return token, nil
}

func getAuthUsernameFromHeader(h *http.Header) (string, error) {
	username := h.Get(authUserHeader)
	if username == "" {
		return "", ErrMissingAuthUsername
	}

	return username, nil
}

func getAuthPasswordFromHeader(h *http.Header) (string, error) {
	password := h.Get(authPasswordHeader)
	if password == "" {
		return "", ErrMissingAuthPassword
	}

	return password, nil
}
