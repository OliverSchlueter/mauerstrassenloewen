package authentication

import (
	"context"
	"errors"
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/hashing"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/usermanagement"
)

type Store struct {
	um          *usermanagement.Store
	globalToken string
}

type StoreConfiguration struct {
	UM          *usermanagement.Store
	GlobalToken string
}

func NewStore(cfg StoreConfiguration) *Store {
	return &Store{
		um:          cfg.UM,
		globalToken: cfg.GlobalToken,
	}
}

func (s *Store) IsAuthTokenValid(ctx context.Context, token string) (bool, error) {
	return token == s.globalToken, nil
}

func (s *Store) IsAuthUserValid(ctx context.Context, user, password string) (bool, error) {
	u, err := s.um.GetUser(ctx, user)
	if err != nil {
		if errors.Is(err, usermanagement.ErrUserNotFound) {
			return false, nil
		}

		return false, fmt.Errorf("could not get user: %w", err)
	}

	if u.Password != hashing.SHA256(password) {
		return false, nil
	}

	return user == "admin" && password == "admin", nil
}
