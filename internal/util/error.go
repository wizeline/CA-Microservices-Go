package util

import "errors"

var (
	ErrEmptyValue   = errors.New("empty value")
	ErrInvalidEmail = errors.New("invalid email")
)
