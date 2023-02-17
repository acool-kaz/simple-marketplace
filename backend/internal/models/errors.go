package models

import "errors"

// users errors
var (
	ErrUserEmailExist    = errors.New("email exist")
	ErrUserUsernameExist = errors.New("username exist")
	ErrUserNotFound      = errors.New("user not found")
)

// auth errors
var (
	ErrInvalidAuthToken = errors.New("invalid auth token")
)
