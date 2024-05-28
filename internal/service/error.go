package service

import (
	"errors"
	"fmt"
)

var (
	ErrNotSupported = errors.New("not supported")
	ErrZeroValue    = errors.New("zero value")

	ErrEmptyArgs        = errors.New("empty arguments")
	ErrInvalidPasswd    = errors.New("invalid password")
	ErrPasswdDoNotMatch = errors.New("passwords do not match")
)

type Err struct {
	Err error
}

func (e Err) Error() string {
	return fmt.Sprintf("service: %s", e.Err)
}

func (e Err) Unwrap() error {
	return e.Err
}

type InvalidInputErr struct {
	Field string
	Err   error
}

func (e InvalidInputErr) Error() string {
	return fmt.Sprintf("invalid input for field %q : %s", e.Field, e.Err)
}

func (e InvalidInputErr) Unwrap() error {
	return e.Err
}

type InvalidFilterErr struct {
	Filter string
	Err    error
}

func (e InvalidFilterErr) Error() string {
	return fmt.Sprintf("invalid filter %q : %s", e.Filter, e.Err)
}

func (e InvalidFilterErr) Unwrap() error {
	return e.Err
}
