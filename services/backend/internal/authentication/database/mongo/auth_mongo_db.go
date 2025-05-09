package mongo

import (
	"context"
	"errors"
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/authentication"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type DB struct {
	coll *mongo.Collection
}

type Configuration struct {
	DB *mongo.Database
}

func NewDB(cfg Configuration) *DB {
	coll := cfg.DB.Collection("auth_tokens")

	return &DB{
		coll: coll,
	}
}

func (db *DB) CreateAuthToken(ctx context.Context, t authentication.Token) error {
	_, err := db.coll.InsertOne(ctx, t)
	if err != nil {
		return fmt.Errorf("could not create auth token: %w", err)
	}

	return nil
}

func (db *DB) GetAuthToken(ctx context.Context, tokenID string) (*authentication.Token, error) {
	filter := bson.D{
		{"tokenId", tokenID},
	}

	res := db.coll.FindOne(ctx, filter)
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return nil, authentication.ErrTokenNotFound
		}
		return nil, fmt.Errorf("could not find auth token: %w", res.Err())
	}

	var token authentication.Token
	if err := res.Decode(&token); err != nil {
		return nil, fmt.Errorf("could not decode auth token: %w", err)
	}

	return &token, nil
}

func (db *DB) DeleteAuthToken(ctx context.Context, tokenID string) error {
	filter := bson.D{
		{"tokenId", tokenID},
	}

	_, err := db.coll.DeleteOne(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil
		}
		return fmt.Errorf("could not delete auth token: %w", err)
	}

	return nil
}

func (db *DB) GetAuthTokensByUserID(ctx context.Context, userID string) ([]authentication.Token, error) {
	filter := bson.D{
		{"userId", userID},
	}

	cursor, err := db.coll.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("could not find auth tokens: %w", err)
	}
	defer cursor.Close(ctx)

	var tokens []authentication.Token
	if err := cursor.All(ctx, &tokens); err != nil {
		return nil, fmt.Errorf("could not decode auth tokens: %w", err)
	}
	return tokens, nil
}
