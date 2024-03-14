package repository

import "github.com/wizeline/CA-Microservices-Go/internal/domain/entity"

func validateUser(user entity.User) error {
	if user == (entity.User{}) {
		return ErrUserEmpty
	}
	// TODO: add validations
	return nil
}
