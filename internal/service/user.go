package service

import (
	"fmt"
	"time"

	"github.com/wizeline/CA-Microservices-Go/internal/entity"
)

type UserRepo interface {
	Create(user entity.User) error
	Read(id uint64) (entity.User, error)
	ReadAll() ([]entity.User, error)
	Update(user entity.User) error
	Delete(id uint64) error
}

type UserCreateArgs struct {
	FirstName string
	LastName  string
	Email     string
	BirthDay  time.Time
	Username  string
	Passwd    string
}

type UserUpdateArgs struct {
	ID        uint64
	FirstName string
	LastName  string
	BirthDay  time.Time
}

type UserService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) UserService {
	return UserService{
		repo: repo,
	}
}

func (s UserService) Create(args UserCreateArgs) error {
	if err := validateUserCreate(args); err != nil {
		return err
	}
	hashedPwd, err := hashPasswd(args.Passwd)
	if err != nil {
		return err
	}
	return s.repo.Create(entity.User{
		FirstName: args.FirstName,
		LastName:  args.LastName,
		Email:     args.Email,
		BirthDay:  args.BirthDay,
		Username:  args.Username,
		Passwd:    hashedPwd,
	})
}

func (s UserService) Get(id uint64) (entity.User, error) {
	if id == 0 {
		return entity.User{}, ErrZeroValue
	}
	return s.repo.Read(id)
}

func (s UserService) GetAll() ([]entity.User, error) {
	return s.repo.ReadAll()
}

func (s UserService) Find(filter, value string) ([]entity.User, error) {
	if err := validateUserFilter(filter); err != nil {
		return nil, err
	}
	users, err := s.repo.ReadAll()
	if err != nil {
		return nil, err
	}

	filteredUsers := make([]entity.User, 0)
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

func (s UserService) Update(args UserUpdateArgs) error {
	if err := validateUserUpdate(args); err != nil {
		return err
	}
	user, err := s.repo.Read(args.ID)
	if err != nil {
		return err
	}

	if args.FirstName != "" {
		user.FirstName = args.FirstName
	}
	if args.LastName != "" {
		user.LastName = args.LastName
	}
	if !args.BirthDay.IsZero() {
		user.BirthDay = args.BirthDay
	}
	return s.repo.Update(user)
}

func (s UserService) Delete(id uint64) error {
	if id == 0 {
		return &InvalidInputErr{Field: "id", Err: ErrZeroValue}
	}
	return s.repo.Delete(id)
}

func (s UserService) Activate(id uint64) error {
	if id == 0 {
		return &InvalidInputErr{Field: "id", Err: ErrZeroValue}
	}
	user, err := s.repo.Read(id)
	if err != nil {
		return err
	}
	user.Active = true
	return s.repo.Update(user)
}

func (s UserService) ChangeEmail(id uint64, email string) error {
	if id == 0 {
		return &InvalidInputErr{Field: "id", Err: ErrZeroValue}
	}
	if err := validateEmail(email); err != nil {
		return &InvalidInputErr{Field: "email", Err: err}
	}

	user, err := s.repo.Read(id)
	if err != nil {
		return err
	}
	user.Email = email
	return s.repo.Update(user)
}

func (s UserService) ChangePasswd(id uint64, passwd string) error {
	if err := validateUserPasswd(passwd); err != nil {
		return err
	}
	user, err := s.repo.Read(id)
	if err != nil {
		return err
	}

	hashedPasswd, err := hashPasswd(passwd)
	if err != nil {
		return err
	}
	user.Passwd = hashedPasswd
	return s.repo.Update(user)
}

func (s UserService) IsActive(id uint64) (bool, error) {
	user, err := s.repo.Read(id)
	if err != nil {
		return false, err
	}
	return user.Active, nil
}

func (s UserService) ValidateLogin(username string, passwd string) (entity.User, error) {
	if username == "" {
		return entity.User{}, &InvalidInputErr{Field: "username", Err: ErrEmptyValue}
	}
	if err := validateUserPasswd(passwd); err != nil {
		return entity.User{}, err
	}
	users, err := s.Find("Username", username)
	if err != nil {
		return entity.User{}, err
	}
	if total := len(users); total != 1 {
		return entity.User{}, fmt.Errorf("expected one user got %d", total)
	}
	if err := compareHashAndPassword(users[0].Passwd, passwd); err != nil {
		return entity.User{}, ErrPasswdDoNotMatch
	}
	return users[0], nil
}
