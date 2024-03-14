package repository

import (
	"errors"
	"fmt"
)

var (
	ErrFieldEmpty = errors.New("field value empty")
)

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
