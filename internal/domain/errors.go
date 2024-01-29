package domain

import "errors"

var (
	ErrUserNotFound      = errors.New("UserNotFound")
	ErrUserAlreadyExists = errors.New("UserAlreadyExists")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrInvalidUserName   = errors.New("invalid user_name")
	ErrInvalidFullName   = errors.New("invalid full_name")
	ErrProductNotFound   = errors.New("ProductNotFound")
)
