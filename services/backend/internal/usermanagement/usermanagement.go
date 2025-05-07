package usermanagement

import (
	"context"
	"github.com/google/uuid"
)

type Database interface {
	GetUserByID(ctx context.Context, id string) (*User, error)
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

func (s *Store) GetUser(ctx context.Context, id string) (*User, error) {
	return s.db.GetUserByID(ctx, id)
}

func (s *Store) CreateUser(ctx context.Context, user *User) error {
	user.ID = uuid.NewString()
	//TODO: Mailabfrage implementieren
	//TODO: Pflichtfelderprüfung implementierern

	if err := validateUser(user); err != nil {
		return err
	}

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
	//TODO: Passwortanforderung implementieren
	return nil
}
