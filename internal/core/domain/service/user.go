package service

import (
	"github.com/wizeline/CA-Microservices-Go/internal/core/domain/entity"
)

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
