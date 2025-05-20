package lessons

import (
	"errors"
)

var (
	ErrLessonNotFound = errors.New("lesson does not exists")
)
