package repository

import "errors"

var (
	ErrUserEmpty   = errors.New("user empty")
	ErrInvalidUser = errors.New("invalid user")
)
