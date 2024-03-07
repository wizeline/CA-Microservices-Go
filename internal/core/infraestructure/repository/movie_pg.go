package repository

import (
	"github.com/wizeline/CA-Microservices-Go/internal/core/domain/entity"
	repo "github.com/wizeline/CA-Microservices-Go/internal/core/domain/repository"
	"github.com/wizeline/CA-Microservices-Go/internal/core/infraestructure/db"
)

var _ repo.MovieRepository = &MovieRepositoryPg{}

type MovieRepositoryPg struct {
	conn *db.DBPgConn
}

func NewMovieRepositoryPg(conn *db.DBPgConn) MovieRepositoryPg {
	return MovieRepositoryPg{
		conn: conn,
	}
}

func (r MovieRepositoryPg) Create(movie entity.Movie) (int, error) {
	// TODO: Implement me!
	return 0, nil
}

func (r MovieRepositoryPg) Read(id int) (entity.Movie, error) {
	// TODO: Implement me!
	return entity.Movie{}, nil
}

func (r MovieRepositoryPg) ReadAll() ([]entity.Movie, error) {
	// TODO: Implement me!
	return nil, nil
}

func (r MovieRepositoryPg) Update(e entity.Movie) error {
	// TODO: Implement me!
	return nil
}

func (r MovieRepositoryPg) Delete(id int) error {
	// TODO: Implement me!
	return nil
}
