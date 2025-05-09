package authentication

import (
	"context"
	"errors"
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/hashing"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/usermanagement"
	"github.com/google/uuid"
	"time"
)

type DB interface {
	CreateAuthToken(ctx context.Context, t Token) error
	GetAuthToken(ctx context.Context, tokenID string) (*Token, error)
	DeleteAuthToken(ctx context.Context, tokenID string) error
	GetAuthTokensByUserID(ctx context.Context, userID string) ([]Token, error)
}

type Store struct {
	db          DB
	um          *usermanagement.Store
	globalToken string
}

type StoreConfiguration struct {
	DB          DB
	UM          *usermanagement.Store
	GlobalToken string
}

func NewStore(cfg StoreConfiguration) *Store {
	return &Store{
		db:          cfg.DB,
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

func (s *Store) CreateAuthToken(ctx context.Context, userID string) (string, error) {
	rawToken := "msl_" + uuid.NewString()

	t := Token{
		ID:        uuid.NewString(),
		UserID:    userID,
		Hash:      hashing.SHA256(rawToken),
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
	}

	if err := s.db.CreateAuthToken(ctx, t); err != nil {
		return "", fmt.Errorf("could not create auth token: %w", err)
	}

	return rawToken, nil
}

func (s *Store) GetAuthToken(ctx context.Context, tokenID string) (*Token, error) {
	t, err := s.db.GetAuthToken(ctx, tokenID)
	if err != nil {
		return nil, fmt.Errorf("could not get auth token: %w", err)
	}

	if t.ExpiresAt.Before(time.Now()) {
		if err := s.db.DeleteAuthToken(ctx, t.ID); err != nil {
			return nil, fmt.Errorf("could not delete expired auth token: %w", err)
		}
		return nil, ErrTokenExpired
	}

	return t, nil
}

func (s *Store) GetAuthTokensByUserID(ctx context.Context, userID string) ([]Token, error) {
	tokens, err := s.db.GetAuthTokensByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("could not get auth tokens: %w", err)
	}

	for i := range tokens {
		if tokens[i].ExpiresAt.Before(time.Now()) {
			if err := s.db.DeleteAuthToken(ctx, tokens[i].ID); err != nil {
				return nil, fmt.Errorf("could not delete expired auth token: %w", err)
			}
			tokens = append(tokens[:i], tokens[i+1:]...)
			i--
		}
	}

	return tokens, nil
}

func (s *Store) DeleteAuthToken(ctx context.Context, tokenID string) error {
	if err := s.db.DeleteAuthToken(ctx, tokenID); err != nil {
		return fmt.Errorf("could not delete auth token: %w", err)
	}

	return nil
}
