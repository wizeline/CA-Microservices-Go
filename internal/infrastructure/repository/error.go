package repository

import (
	"errors"
	"fmt"
)

var (
	ErrFieldEmpty = errors.New("field value empty")
)

type Err struct {
	Err error
}

func (e Err) Error() string {
	return fmt.Sprintf("repository: %s", e.Err)
}

func (e Err) Unwrap() error {
	return e.Err
}

type EntityEmptyErr struct {
	Name string
}

func (e EntityEmptyErr) Error() string {
	return fmt.Sprintf("entity %v empty", e.Name)
}

type InvalidFieldErr struct {
	Name string
	Err  error
}

func (e InvalidFieldErr) Error() string {
	return fmt.Sprintf("invalid field %v: %s", e.Name, e.Err)
}

func (e InvalidFieldErr) Unwrap() error {
	return e.Err
}
