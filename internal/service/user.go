package service

import (
	"fmt"

	"github.com/wizeline/CA-Microservices-Go/internal/domain/entity"
	repo "github.com/wizeline/CA-Microservices-Go/internal/domain/repository"
	svc "github.com/wizeline/CA-Microservices-Go/internal/domain/service"
)

// We ensure the UserService implementation satisfies the UserService signature in the domain
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
	err := s.repo.Create(user)
	if err != nil {
		return -1, fmt.Errorf("Couldn't create user: %w", err)
	}
	return user.ID, nil
}

func (s UserService) Get(id int) (entity.User, error) {
	user, err := s.repo.Read(id)
	if err != nil {
		return entity.User{}, fmt.Errorf("Couldn't get user: %w", err)
	}
	return user, nil
}

func (s UserService) Find(filter, value string) ([]entity.User, error) {
	// TODO: Implement me!
	return nil, nil
}

func (s UserService) Update(id int, data entity.User) error {
	err := s.repo.Update(data)
	if err != nil {
		return fmt.Errorf("Couldn't update user: %w", err)
	}
	return nil
}

func (s UserService) Delete(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("Couldn't delete user: %w", err)
	}
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
	user, err := s.repo.Read(id)
	if err != nil {
		return false, fmt.Errorf("Couldn't get user: %w", err)
	}
	return user.Active, nil
}
