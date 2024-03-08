package repository

import (
	"database/sql"

	"github.com/wizeline/CA-Microservices-Go/internal/domain/entity"
	repo "github.com/wizeline/CA-Microservices-Go/internal/domain/repository"
)

// We ensure the UserRepositoryPg implementation satisfies the UserRepository signature in the domain
var _ repo.UserRepository = &UserRepositoryPg{}

type PgConn interface {
	DB() *sql.DB
	Close()
}

type UserRepositoryPg struct {
	conn PgConn
}

func NewUserRepositoryPg(conn PgConn) UserRepositoryPg {
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
