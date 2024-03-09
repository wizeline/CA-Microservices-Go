package service

import (
	"github.com/wizeline/CA-Microservices-Go/internal/domain/entity"
	"github.com/wizeline/CA-Microservices-Go/internal/domain/service/mocks"
)

// We ensure the UserService mock object satisfies the UserService signature.
var _ UserService = &mocks.UserService{}

type UserService interface {
	Add(user entity.User) (int, error)
	Get(id int) (entity.User, error)
	Find(filter, value string) ([]entity.User, error)
	Update(id int, data entity.User) error
	Delete(id int) error

	ValidateLogin(username string, password string) (entity.User, error)
	ChangeEmail(id int, email string) error
	IsActive(id int) (bool, error)
}
