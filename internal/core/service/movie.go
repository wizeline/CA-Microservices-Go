package service

import (
	"github.com/wizeline/CA-Microservices-Go/internal/core/domain/entity"
	repo "github.com/wizeline/CA-Microservices-Go/internal/core/domain/repository"
	svc "github.com/wizeline/CA-Microservices-Go/internal/core/domain/service"
)

var _ svc.MovieService = &MovieService{}

type MovieService struct {
	repo repo.MovieRepository
}

func NewMovie(repo repo.MovieRepository) MovieService {
	return MovieService{
		repo: repo,
	}
}

func (s MovieService) Add(movie entity.Movie) (int, error) {
	// TODO: Implement me!
	return 0, nil
}

func (s MovieService) Get(id int) (entity.Movie, error) {
	// TODO: Implement me!
	return entity.Movie{}, nil
}

func (s MovieService) Find(filter, value string) ([]entity.Movie, error) {
	// TODO: Implement me!
	return nil, nil
}

func (s MovieService) Update(id int, data entity.Movie) error {
	// TODO: Implement me!
	return nil
}

func (s MovieService) Delete(id int) error {
	// TODO: Implement me!
	return nil
}
