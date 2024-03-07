package repository

import (
	"github.com/wizeline/CA-Microservices-Go/internal/core/domain/entity"
)

type MovieRepository interface {
	Create(movie entity.Movie) (int, error)
	Read(id int) (entity.Movie, error)
	ReadAll() ([]entity.Movie, error)
	Update(e entity.Movie) error
	Delete(id int) error
}
