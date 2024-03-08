package repository

import (
	"github.com/wizeline/CA-Microservices-Go/internal/domain/entity"
	repo "github.com/wizeline/CA-Microservices-Go/internal/domain/repository"
	"github.com/wizeline/CA-Microservices-Go/internal/infrastructure/db"
)

var _ repo.UserRepository = &UserRepositoryPg{}

type UserRepositoryPg struct {
	conn *db.DBPgConn
}

func NewUserRepositoryPg(conn *db.DBPgConn) UserRepositoryPg {
	return UserRepositoryPg{
		conn: conn,
	}
}

func (r UserRepositoryPg) Create(user entity.User) (int, error) {
	// TODO: Implement me!
	return 0, nil
}

func (r UserRepositoryPg) Read(id int) (entity.User, error) {
	// TODO: Implement me!
	return entity.User{}, nil
}

func (r UserRepositoryPg) ReadAll() ([]entity.User, error) {
	// TODO: Implement me!
	return nil, nil
}

func (r UserRepositoryPg) Update(id int, data entity.User) error {
	// TODO: Implement me!
	return nil
}

func (r UserRepositoryPg) Delete(id int) error {
	// TODO: Implement me!
	return nil
}
