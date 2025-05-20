package usermanagement

import (
	"context"
	"errors"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/hashing"
	"github.com/google/uuid"
	"strings"
	"unicode"
)

type Database interface {
	GetUserByID(ctx context.Context, id string) (*User, error)
	GetUserByName(ctx context.Context, name string) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, id string) error
}

type Store struct {
	db Database
}

type Configuration struct {
	DB Database
}

func NewStore(config *Configuration) *Store {
	return &Store{
		db: config.DB,
	}
}

func (s *Store) GetUserByID(ctx context.Context, id string) (*User, error) {
	return s.db.GetUserByID(ctx, id)
}

func (s *Store) GetUserByName(ctx context.Context, name string) (*User, error) {
	return s.db.GetUserByName(ctx, name)
}

func (s *Store) CreateUser(ctx context.Context, user *User) error {
	user.ID = uuid.NewString()
	//TODO: Mailabfrage implementieren

	if err := validateUser(user); err != nil {
		return err
	}

	user.Password = hashing.SHA256(user.Password)

	return s.db.CreateUser(ctx, user)
}

func (s *Store) UpdateUser(ctx context.Context, user *User) error {
	if err := validateUser(user); err != nil {
		return err
	}

	return s.db.UpdateUser(ctx, user)
}

func (s *Store) DeleteUser(ctx context.Context, id string) error {
	return s.db.DeleteUser(ctx, id)
}

func validateUser(user *User) error {

	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if err := validatePassword(user.Password, user.Name); err != nil {
		return err
	}

	return nil
}

func validatePassword(password string, firstName string) error {

	if strings.TrimSpace(password) == "" {
		return errors.New("password can't be empty")
	}

	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	var hasNumber bool
	var hasSpecial bool

	for _, ch := range password {
		switch {
		case unicode.IsNumber(ch):
			hasNumber = true
		case unicode.IsSymbol(ch) || unicode.IsPunct(ch):
			hasSpecial = true
		}
	}

	if !hasNumber {
		return errors.New("password must contain at least one number")
	}
	if !hasSpecial {
		return errors.New("password must include at least one special character")
	}

	if strings.EqualFold(password, firstName) {
		return errors.New("password must not match the name")
	}

	return nil
}
