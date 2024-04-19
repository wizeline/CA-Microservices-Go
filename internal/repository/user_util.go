package repository

import (
	"github.com/wizeline/CA-Microservices-Go/internal/entity"
	"github.com/wizeline/CA-Microservices-Go/internal/util"
)

func validateUser(user entity.User) error {
	if user == (entity.User{}) {
		return &EntityEmptyErr{Name: "User"}
	}
	if user.FirstName == "" {
		return &InvalidFieldErr{Name: "FirstName", Err: ErrFieldEmpty}
	}
	if user.LastName == "" {
		return &InvalidFieldErr{Name: "LastName", Err: ErrFieldEmpty}
	}
	if err := util.IsValidEmail(user.Email); err != nil {
		return &InvalidFieldErr{Name: "Email", Err: err}
	}
	if user.Username == "" {
		return &InvalidFieldErr{Name: "Username", Err: ErrFieldEmpty}
	}
	if user.Passwd == "" {
		return &InvalidFieldErr{Name: "Passwd", Err: ErrFieldEmpty}
	}

	return nil
}
