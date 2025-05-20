package mongo

import (
	"context"
	"errors"
	"fmt"
	"github.com/OliverSchlueter/mauerstrassenloewen/backend/internal/lessons"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DB struct {
	coll *mongo.Collection
}

type Configuration struct {
	Mongo *mongo.Database
}

func NewDB(config *Configuration) *DB {
	coll := config.Mongo.Collection("lessons")

	return &DB{
		coll: coll,
	}
}

func (db *DB) GetLesson(ctx context.Context, id string) (*lessons.Lesson, error) {
	res := db.coll.FindOne(ctx, bson.D{
		{"user_id", id},
	})

	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return nil, lessons.ErrLessonNotFound
		}
		return nil, fmt.Errorf("could not find lesson: %w", res.Err())
	}

	var lesson lessons.Lesson
	if err := res.Decode(&lesson); err != nil {
		return nil, fmt.Errorf("could not decode lesson: %w", err)
	}

	return &lesson, nil
}

func (db *DB) UpsertLesson(ctx context.Context, lesson *lessons.Lesson) error {
	filter := bson.D{{"user_id", lesson.UserID}}

	_, err := db.coll.UpdateOne(ctx, filter, bson.M{"$set": lesson}, options.UpdateOne().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("could not update lesson: %w", err)
	}

	return nil
}
