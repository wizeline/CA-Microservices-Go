package repository

import (
	"github.com/wizeline/CA-Microservices-Go/internal/domain/entity"
	"github.com/wizeline/CA-Microservices-Go/internal/domain/repository/mocks"
)

// We ensure the UserRepository mock object satisfies the UserRepository signature.
var _ UserRepository = &mocks.UserRepository{}

type UserRepository interface {
	Create(user entity.User) error
	Read(id int) (entity.User, error)
	ReadAll() ([]entity.User, error)
	Update(user entity.User) error
	Delete(id int) error
}
