package service

import (
	"github.com/wizeline/CA-Microservices-Go/internal/core/domain/entity"
)

type MovieService interface {
	Add(movie entity.Movie) (int, error)
	Get(id int) (entity.Movie, error)
	Find(filter, value string) ([]entity.Movie, error)
	Update(id int, data entity.Movie) error
	Delete(id int) error
}
