package service

import (
	"github.com/wizeline/CA-Microservices-Go/internal/entity"
	"github.com/wizeline/CA-Microservices-Go/internal/logger"
)

type UserRepo interface {
	Create(user entity.User) error
	Read(id int) (entity.User, error)
	ReadAll() ([]entity.User, error)
	Update(user entity.User) error
	Delete(id int) error
}

type UserService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo, l logger.ZeroLog) UserService {
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
