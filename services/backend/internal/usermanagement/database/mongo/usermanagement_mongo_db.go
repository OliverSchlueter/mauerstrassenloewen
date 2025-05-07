package mongo

import (
	"context"
	"errors"
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/usermanagement"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type DB struct {
	coll *mongo.Collection
}

type Configuration struct {
	Mongo *mongo.Database
}

func NewDB(config *Configuration) *DB {
	coll := config.Mongo.Collection("users")

	return &DB{
		coll: coll,
	}
}

func (db *DB) GetUserByID(ctx context.Context, id string) (*usermanagement.User, error) {
	res := db.coll.FindOne(ctx, usermanagement.User{
		ID: id,
	})

	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return nil, usermanagement.ErrUserNotFound
		}
		return nil, fmt.Errorf("could not find user: %w", res.Err())
	}

	var user usermanagement.User
	if err := res.Decode(&user); err != nil {
		return nil, fmt.Errorf("could not decode user: %w", err)
	}

	return &user, nil
}

func (db *DB) CreateUser(ctx context.Context, user *usermanagement.User) error {
	_, err := db.coll.InsertOne(ctx, user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return usermanagement.ErrUserAlreadyExists
		}

		return fmt.Errorf("could not create user: %w", err)
	}

	return nil
}

func (db *DB) UpdateUser(ctx context.Context, user *usermanagement.User) error {
	_, err := db.coll.UpdateOne(ctx, usermanagement.User{ID: user.ID}, bson.M{"$set": user})
	if err != nil {
		return fmt.Errorf("could not update user: %w", err)
	}

	return nil
}

func (db *DB) DeleteUser(ctx context.Context, id string) error {
	_, err := db.coll.DeleteOne(ctx, usermanagement.User{ID: id})
	if err != nil {
		return fmt.Errorf("could not delete user: %w", err)
	}

	return nil
}
