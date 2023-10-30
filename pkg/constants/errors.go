package constants

import "errors"

// error conventions

var (
	ErrBunNotNotFound   = errors.New("sql: no rows in result set")
	ErrMenteeNotFound   = errors.New("mentee not found")
	ErrUserNotFound     = errors.New("user not found")
	ErrSessionNotFound  = errors.New("session not found")
	ErrDeviceNotFound   = errors.New("device not found")
	ErrInterestNotFound = errors.New("interest not found")
	ErrKeyNotFound      = errors.New("key not found")

	ErrUsernameAlreadyTaken = errors.New("username already taken")
	ErrEmailAlreadyTaken    = errors.New("email already taken")

	ErrPasswordIncorrect = errors.New("password incorrect")
	ErrCredentialInvalid = errors.New("invalid username or password")
)
