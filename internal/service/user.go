package service

import (
	"errors"
	"slices"

	"github.com/wizeline/CA-Microservices-Go/internal/entity"
	"github.com/wizeline/CA-Microservices-Go/internal/logger"
)

type UserRepo interface {
	Create(user entity.User) error
	Read(id uint64) (entity.User, error)
	ReadAll() ([]entity.User, error)
	Update(user entity.User) error
	Delete(id uint64) error
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
	err := s.validateUser(user)
	if err != nil {
		return err
	}
	return s.repo.Create(user)
}

func (s UserService) Get(id uint64) (entity.User, error) {
	return s.repo.Read(id)
}

func (s UserService) GetAll() ([]entity.User, error) {
	return s.repo.ReadAll()
}

func (s UserService) Find(filter, value string) ([]entity.User, error) {
	err := validateFilter(filter)
	if err != nil {
		return []entity.User{}, err
	}

	users, err := s.repo.ReadAll()
	if err != nil {
		return []entity.User{}, err
	}

	filteredUsers := []entity.User{}
	for _, user := range users {
		switch {
		case filter == "FirstName":
			if user.FirstName == value {
				filteredUsers = append(filteredUsers, user)
			}
		case filter == "LastName":
			if user.LastName == value {
				filteredUsers = append(filteredUsers, user)
			}
		case filter == "Email":
			if user.Email == value {
				filteredUsers = append(filteredUsers, user)
			}
		case filter == "Username":
			if user.Username == value {
				filteredUsers = append(filteredUsers, user)
			}
		}

	}
	return filteredUsers, nil
}

func (s UserService) Update(user entity.User) error {
	err := s.validateUser(user)
	if err != nil {
		return err
	}
	return s.repo.Update(user)
}

func (s UserService) Delete(id uint64) error {
	return s.repo.Delete(id)
}

func (s UserService) Activate(id uint64) error {
	user, err := s.repo.Read(id)
	if err != nil {
		return err
	}
	user.Active = true
	return s.repo.Update(user)
}

func (s UserService) ChangeEmail(id uint64, email string) error {
	user, err := s.repo.Read(id)
	if err != nil {
		return err
	}
	user.Email = email
	return s.repo.Update(user)
}

func (s UserService) ChangePasswd(id uint64, passwd string) error {
	err := validatePassword(passwd)
	if err != nil {
		return err
	}
	user, err := s.repo.Read(id)
	if err != nil {
		return err
	}
	user.Passwd = passwd
	return s.repo.Update(user)
}

func (s UserService) IsActive(id uint64) (bool, error) {
	user, err := s.repo.Read(id)
	if err != nil {
		return false, err
	}
	return user.Active, nil
}

func (s UserService) ValidateLogin(username string, password string) (entity.User, error) {
	users, err := s.Find("Username", username)
	if err != nil {
		return entity.User{}, err
	}

	if len(users) == 0 {
		return entity.User{}, errors.New("user not found")
	}

	if users[0].Passwd == password {
		return users[0], nil
	}

	return entity.User{}, ErrInvalidPassword
}

func (s UserService) validateUser(user entity.User) error {
	switch {
	case user.FirstName == "":
		return InvalidInputErr{Field: "FirstName"}
	case user.LastName == "":
		return InvalidInputErr{Field: "LastName"}
	case user.Email == "":
		return InvalidInputErr{Field: "Email"}
	case user.Username == "":
		return InvalidInputErr{Field: "Username"}
	}
	return validatePassword(user.Passwd)
}

func validatePassword(password string) error {
	if password == "" || len(password) < 6 {
		return InvalidInputErr{Field: "Passwd"}
	}
	return nil
}

func validateFilter(filter string) error {
	validFilters := []string{"FirstName", "LastName", "Email", "Username"}
	if !slices.Contains(validFilters, filter) {
		return InvalidFilter{Filter: filter}
	}
	return nil
}
