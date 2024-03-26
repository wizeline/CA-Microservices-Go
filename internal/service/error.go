package service

import "fmt"

type Err struct {
	Err error
}

func (e Err) Error() string {
	return fmt.Sprintf("service: %s", e.Err)
}

func (e Err) Unwrap() error {
	return e.Err
}
