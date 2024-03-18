package service

import (
	"errors"
	"fmt"
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
}

func (e InvalidInputErr) Error() string {
	return fmt.Sprintf("Invalid value for field: %s", e.Field)
}

type InvalidFilter struct {
	Filter string
}

func (e InvalidFilter) Error() string {
	return fmt.Sprintf("Invalid filter for search: %s", e.Filter)
}

var InvalidPassword = errors.New("Can't login, passwords doesn't match")
