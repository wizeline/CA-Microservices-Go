package service

import (
	"github.com/wizeline/CA-Microservices-Go/internal/domain/entity"
)

type UserService interface {
	Create(user entity.User) error
	Get(id uint64) (entity.User, error)
	GetAll() ([]entity.User, error)
	Find(filter, value string) ([]entity.User, error)
	Update(user entity.User) error
	Delete(id uint64) error

	Activate(id uint64) error
	ChangeEmail(id uint64, email string) error
	ChangePasswd(id uint64, passwd string) error
	IsActive(id uint64) (bool, error)
	ValidateLogin(username string, password string) (entity.User, error)
}
