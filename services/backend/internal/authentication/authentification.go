package authentication

import (
	"context"
	"errors"
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/hashing"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/usermanagement"
)

type DB interface {
	CreateAuthToken(ctx context.Context, t Token) error
	GetAuthToken(ctx context.Context, tokenID string) (*Token, error)
	DeleteAuthToken(ctx context.Context, tokenID string) error
	GetAuthTokensByUserID(ctx context.Context, userID string) ([]Token, error)
}

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

func (s *Store) IsAuthTokenValid(ctx context.Context, token string) (*usermanagement.User, error) {
	if token != s.globalToken {
		return nil, nil
	}

	return &usermanagement.User{
		ID:    "global-user",
		Name:  "Global User",
		Email: "globaluser@msl.de",
	}, nil
}

func (s *Store) IsAuthUserValid(ctx context.Context, user, password string) (*usermanagement.User, error) {
	u, err := s.um.GetUser(ctx, user)
	if err != nil {
		if errors.Is(err, usermanagement.ErrUserNotFound) {
			return nil, nil
		}

		return nil, fmt.Errorf("could not get user: %w", err)
	}

	if u.Password != hashing.SHA256(password) {
		return nil, nil
	}

	return u, nil
}
