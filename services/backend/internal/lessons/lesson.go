package lessons

import (
	"context"
	"errors"
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
	l, err := s.GetLesson(ctx, userID)
	if err != nil {
		if errors.Is(err, ErrLessonNotFound) {
			l = &Lesson{
				UserID: userID,
				Done: map[string]bool{
					lessonID: true,
				},
			}
		}
		return err
	}

	l.Done[lessonID] = true

	return s.db.UpsertLesson(ctx, l)
}
