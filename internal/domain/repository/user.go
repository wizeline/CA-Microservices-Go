package repository

import (
	"github.com/wizeline/CA-Microservices-Go/internal/domain/entity"
)

type UserRepository interface {
	Create(user entity.User) (int, error)
	Read(id int) (entity.User, error)
	ReadAll() ([]entity.User, error)
	Update(id int, data entity.User) error
	Delete(id int) error
}
