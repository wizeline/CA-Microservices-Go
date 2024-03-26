package service

import (
	"github.com/wizeline/CA-Microservices-Go/internal/domain/entity"
	repo "github.com/wizeline/CA-Microservices-Go/internal/domain/repository"
	svc "github.com/wizeline/CA-Microservices-Go/internal/domain/service"

	"github.com/wizeline/CA-Microservices-Go/internal/infrastructure/logger"
)

// We ensure the UserService implementation satisfies the UserService signature in the domain
var _ svc.UserService = &UserService{}

type UserService struct {
	repo repo.UserRepository
}

func NewUserService(repo repo.UserRepository, l logger.ZeroLog) UserService {
	return UserService{
		repo: repo,
	}
}

func (s UserService) Create(user entity.User) error {
	// TODO: Implement me!
	return nil
}

func (s UserService) Get(id uint64) (entity.User, error) {
	// TODO: Implement me!
	return entity.User{}, nil
}

func (s UserService) GetAll() ([]entity.User, error) {
	// TODO: Implement me!
	return nil, nil
}

func (s UserService) Find(filter, value string) ([]entity.User, error) {
	// TODO: Implement me!
	return nil, nil
}

func (s UserService) Update(user entity.User) error {
	// TODO: Implement me!
	return nil
}

func (s UserService) Delete(id uint64) error {
	// TODO: Implement me!
	return nil
}

func (s UserService) Activate(id uint64) error {
	// TODO: Implement me!
	return nil
}

func (s UserService) ChangeEmail(id uint64, email string) error {
	// TODO: Implement me!
	return nil
}

func (s UserService) ChangePasswd(id uint64, passwd string) error {
	// TODO: Implement me!
	return nil
}

func (s UserService) IsActive(id uint64) (bool, error) {
	// TODO: Implement me!
	return false, nil
}

func (s UserService) ValidateLogin(username string, password string) (entity.User, error) {
	// TODO: Implement me!
	return entity.User{}, nil
}
