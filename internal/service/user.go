package service

import (
	"github.com/wizeline/CA-Microservices-Go/internal/domain/entity"
	repo "github.com/wizeline/CA-Microservices-Go/internal/domain/repository"
	svc "github.com/wizeline/CA-Microservices-Go/internal/domain/service"
)

var _ svc.UserService = &UserService{}

type UserService struct {
	repo repo.UserRepository
}

func NewUser(repo repo.UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func (s UserService) Add(user entity.User) (int, error) {
	// TODO: Implement me!
	return 0, nil
}

func (s UserService) Get(id int) (entity.User, error) {
	// TODO: Implement me!
	return entity.User{}, nil
}

func (s UserService) Find(filter, value string) ([]entity.User, error) {
	// TODO: Implement me!
	return nil, nil
}

func (s UserService) Update(id int, data entity.User) error {
	// TODO: Implement me!
	return nil
}

func (s UserService) Delete(id int) error {
	// TODO: Implement me!
	return nil
}

func (s UserService) ValidateLogin(username string, password string) (entity.User, error) {
	// TODO: Implement me!
	return entity.User{}, nil
}

func (s UserService) ChangeEmail(id int, email string) error {
	// TODO: Implement me!
	return nil
}

func (s UserService) IsActive(id int) (bool, error) {
	// TODO: Implement me!
	return false, nil
}
