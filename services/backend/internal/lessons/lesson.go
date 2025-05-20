package lessons

import (
	"context"
)

type Database interface {
	GetLesson(ctx context.Context, id string) (*Lesson, error)
	UpsertLesson(ctx context.Context, lesson *Lesson) error
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

func (s *Store) GetLesson(ctx context.Context, id string) (*Lesson, error) {
	return s.db.GetLesson(ctx, id)
}

func (s *Store) UpsertLesson(ctx context.Context, userID string, lessonID string) error {
	return s.db.UpsertLesson(ctx, lesson)
}

// TODO: LessonID zusammenbauen
