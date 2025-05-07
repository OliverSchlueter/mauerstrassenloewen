package usermanagement

import "errors"

var (
	ErrUserNotFound      = errors.New("could not find user")
	ErrUserAlreadyExists = errors.New("user already exists")
)
