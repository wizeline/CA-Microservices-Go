package service

import (
	"errors"
	"testing"
	"time"

	"github.com/wizeline/CA-Microservices-Go/internal/entity"
	"github.com/wizeline/CA-Microservices-Go/internal/logger"
	"github.com/wizeline/CA-Microservices-Go/internal/service/mocks"

	"github.com/stretchr/testify/assert"
)

// We ensure the UserRepo mock object satisfies the UserRepo signature.
var _ UserRepo = &mocks.UserRepo{}

func TestCreateUser(t *testing.T) {
	type testcase struct {
		name    string
		wantErr error
		repoErr error
		user    entity.User
	}

	validateTestCases := []testcase{
		{
			name:    "Invalid user (empty username)",
			wantErr: InvalidInputErr{Field: "FirstName"},
			user: entity.User{
				FirstName: "",
				LastName:  "Field",
				Email:     "lisa@field.com",
				Username:  "lisa",
				Passwd:    "pass123",
			},
		},
		{
			name:    "Invalid user (empty lastname)",
			wantErr: InvalidInputErr{Field: "LastName"},
			user: entity.User{
				FirstName: "Lisa",
				LastName:  "",
				Email:     "lisa@field.com",
				Username:  "lisa",
				Passwd:    "pass123",
			},
		},
		{
			name:    "Invalid user (empty lastname)",
			wantErr: InvalidInputErr{Field: "LastName"},
			user: entity.User{
				FirstName: "Lisa",
				LastName:  "",
				Email:     "lisa@field.com",
				Username:  "lisa",
				Passwd:    "pass123",
			},
		},
		{
			name:    "Invalid user (empty email)",
			wantErr: InvalidInputErr{Field: "Email"},
			user: entity.User{
				FirstName: "Lisa",
				LastName:  "Field",
				Email:     "",
				Username:  "lisa",
				Passwd:    "pass123",
			},
		},
		{
			name:    "Invalid user (password must have at least 6 characters)",
			wantErr: InvalidInputErr{Field: "Passwd"},
			user: entity.User{
				FirstName: "Lisa",
				LastName:  "Field",
				Email:     "lisa@field.com",
				Username:  "lisa",
				Passwd:    "12345",
			},
		},
	}

	for _, tt := range validateTestCases {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewUserRepo(t)
			logger := logger.NewZeroLog()
			s := NewUserService(mockRepo, logger)

			gotErr := s.Create(tt.user)

			assert.Error(t, gotErr)
			assert.EqualError(t, gotErr, tt.wantErr.Error())
		})
	}

	createTestCases := []testcase{
		{
			name:    "Valid user",
			wantErr: nil,
			user: entity.User{
				FirstName: "lisa",
				LastName:  "Field",
				Email:     "lisa@field.com",
				Username:  "lisa",
				Passwd:    "pass123",
			},
		},
		{
			name:    "Repository error",
			repoErr: errors.New("Couldn't write user"),
			wantErr: errors.New("Couldn't write user"),
			user: entity.User{
				FirstName: "lisa",
				LastName:  "Field",
				Email:     "lisa@field.com",
				Username:  "lisa",
				Passwd:    "pass123",
			},
		},
	}

	for _, tt := range createTestCases {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewUserRepo(t)
			logger := logger.NewZeroLog()
			s := NewUserService(mockRepo, logger)

			mockRepo.On("Create", tt.user).Return(tt.repoErr)

			gotErr := s.Create(tt.user)

			if tt.wantErr != nil {
				assert.Error(t, gotErr)
				assert.EqualError(t, gotErr, tt.wantErr.Error())
			} else {
				assert.Nil(t, gotErr)
			}
		})
	}

}

func TestGetUser(t *testing.T) {
	testsCases := []struct {
		name     string
		userID   uint64
		repoErr  error
		wantErr  error
		wantUser entity.User
	}{
		{
			name:   "Valid user ID",
			userID: 1,
			wantUser: entity.User{
				ID:        1,
				FirstName: "lisa",
				LastName:  "Field",
				Email:     "lisa@field.com",
				BirthDay:  time.Date(1998, 12, 20, 0, 0, 0, 0, time.UTC),
				Username:  "lisa",
				Passwd:    "pass123",
				Active:    true,
			},
		},
		{
			name:     "Repository fails to get user",
			userID:   1,
			repoErr:  errors.New("error from repo"),
			wantErr:  errors.New("error from repo"),
			wantUser: entity.User{},
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewUserRepo(t)
			logger := logger.NewZeroLog()
			s := NewUserService(mockRepo, logger)

			mockRepo.On("Read", tt.userID).Return(tt.wantUser, tt.repoErr)

			gotUser, gotErr := s.Get(tt.userID)

			if tt.wantErr != nil {
				assert.Error(t, gotErr)
				assert.EqualError(t, gotErr, tt.wantErr.Error())
			} else {
				assert.Nil(t, gotErr)
			}

			assert.Equal(t, tt.wantUser, gotUser)
		})
	}

}

func TestGetAllUsers(t *testing.T) {
	testsCases := []struct {
		name      string
		users     []entity.User
		repoErr   error
		wantErr   error
		wantUsers []entity.User
	}{
		{
			name: "Sucessful retrieval of users",
			users: []entity.User{
				{ID: 1, Username: "lisafield@mail.com"},
				{ID: 2, Username: "mat123@mail.com"},
				{ID: 3, Username: "juan@mail.com"},
			},
			wantUsers: []entity.User{
				{ID: 1, Username: "lisafield@mail.com"},
				{ID: 2, Username: "mat123@mail.com"},
				{ID: 3, Username: "juan@mail.com"},
			},
		},
		{
			name:      "Repository fails to get users",
			users:     []entity.User{},
			repoErr:   errors.New("Repo failed to get users"),
			wantErr:   errors.New("Repo failed to get users"),
			wantUsers: []entity.User{},
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewUserRepo(t)
			logger := logger.NewZeroLog()
			s := NewUserService(mockRepo, logger)

			mockRepo.On("ReadAll").Return(tt.users, tt.repoErr)

			gotUsers, gotErr := s.GetAll()

			if tt.wantErr != nil {
				assert.Error(t, gotErr)
				assert.EqualError(t, gotErr, tt.wantErr.Error())
			} else {
				assert.Nil(t, gotErr)
			}

			assert.Equal(t, tt.wantUsers, gotUsers)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	type testcase struct {
		name    string
		repoErr error
		wantErr error
		user    entity.User
	}

	validateTestCases := []testcase{
		{
			name:    "Invalid user (empty firstname)",
			wantErr: InvalidInputErr{Field: "FirstName"},
			user: entity.User{
				ID:        2,
				FirstName: "",
				LastName:  "Field",
				Email:     "lisa@field.com",
				Username:  "lisa",
				Passwd:    "pass123",
			},
		},
		{
			name:    "Invalid user (empty lastname)",
			wantErr: InvalidInputErr{Field: "LastName"},
			user: entity.User{
				ID:        3,
				FirstName: "Lisa",
				LastName:  "",
				Email:     "lisa@field.com",
				Username:  "lisa",
				Passwd:    "pass123",
			},
		},
		{
			name:    "Invalid user (empty email)",
			wantErr: InvalidInputErr{Field: "Email"},
			user: entity.User{
				ID:        3,
				FirstName: "Lisa",
				LastName:  "Field",
				Email:     "",
				Username:  "lisa",
				Passwd:    "pass123",
			},
		},
		{
			name:    "Invalid user (empty username)",
			wantErr: InvalidInputErr{Field: "Username"},
			user: entity.User{
				ID:        3,
				FirstName: "Lisa",
				LastName:  "Field",
				Email:     "lisa@field.com",
				Username:  "",
				Passwd:    "pass123",
			},
		},
		{
			name:    "Invalid user (passwd less than 6 characters)",
			wantErr: InvalidInputErr{Field: "Passwd"},
			user: entity.User{
				ID:        3,
				FirstName: "Lisa",
				LastName:  "Field",
				Email:     "lisa@field.com",
				Username:  "lisa",
				Passwd:    "pass1",
			},
		},
	}

	for _, tt := range validateTestCases {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewUserRepo(t)
			logger := logger.NewZeroLog()

			s := NewUserService(mockRepo, logger)

			gotErr := s.Update(tt.user)

			assert.Error(t, gotErr)
			assert.EqualError(t, gotErr, tt.wantErr.Error())
		})
	}

	updateTestCases := []testcase{
		{
			name: "Valid user update",
			user: entity.User{
				ID:        1,
				FirstName: "lisa",
				LastName:  "Field",
				Email:     "lisa@field.com",
				Username:  "lisa",
				Passwd:    "pass123",
			},
		},
		{
			name:    "Repo error",
			repoErr: errors.New("repo: Couldn't write user"),
			wantErr: errors.New("repo: Couldn't write user"),
			user: entity.User{
				ID:        1,
				FirstName: "lisa",
				LastName:  "Field",
				Email:     "lisa@field.com",
				Username:  "lisa",
				Passwd:    "pass123",
			},
		},
	}

	for _, tt := range updateTestCases {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewUserRepo(t)
			logger := logger.NewZeroLog()

			s := NewUserService(mockRepo, logger)

			mockRepo.On("Update", tt.user).Return(tt.repoErr)

			gotErr := s.Update(tt.user)

			if tt.wantErr != nil {
				assert.Error(t, gotErr)
				assert.EqualError(t, gotErr, tt.wantErr.Error())
			} else {
				assert.Nil(t, gotErr)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	testsCases := []struct {
		name    string
		userID  uint64
		repoErr error
		wantErr error
	}{
		{
			name:   "Valid user deletion",
			userID: 1,
		},
		{
			name:    "Repository fails to delete user",
			userID:  1,
			repoErr: errors.New("Some error happened while deleting user"),
			wantErr: errors.New("Some error happened while deleting user"),
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewUserRepo(t)
			logger := logger.NewZeroLog()
			s := NewUserService(mockRepo, logger)

			mockRepo.On("Delete", tt.userID).Return(tt.repoErr)

			gotErr := s.Delete(tt.userID)

			if tt.wantErr != nil {
				assert.Error(t, gotErr)
				assert.EqualError(t, gotErr, tt.wantErr.Error())
			} else {
				assert.Nil(t, gotErr)
			}
		})
	}
}

func TestUserIsActive(t *testing.T) {
	testsCases := []struct {
		name         string
		user         entity.User
		wantIsActive bool
		repoErr      error
		wantErr      error
	}{
		{
			name: "Active user",
			user: entity.User{
				ID:     1,
				Active: true,
			},
			wantIsActive: true,
		},
		{
			name: "Inactive user",
			user: entity.User{
				ID:     1,
				Active: false,
			},
			wantIsActive: false,
		},
		{
			name:         "Repository fails to get user",
			repoErr:      errors.New("Some error happened in repo"),
			wantErr:      errors.New("Some error happened in repo"),
			wantIsActive: false,
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewUserRepo(t)
			logger := logger.NewZeroLog()
			s := NewUserService(mockRepo, logger)

			mockRepo.On("Read", tt.user.ID).Return(tt.user, tt.repoErr)

			gotIsActive, gotErr := s.IsActive(tt.user.ID)

			if tt.wantErr != nil {
				assert.Error(t, gotErr)
				assert.EqualError(t, gotErr, tt.wantErr.Error())
			} else {
				assert.Nil(t, gotErr)
			}
			assert.Equal(t, tt.wantIsActive, gotIsActive)
		})
	}
}

func TestActivateUser(t *testing.T) {
	testsCases := []struct {
		name            string
		userID          uint64
		repoReadError   error
		repoUpdateError error
		wantErr         error
		user            entity.User
		userToStore     entity.User
	}{
		{
			name:   "Activate valid user",
			userID: 1,
			user: entity.User{
				ID:     1,
				Active: false,
			},
			userToStore: entity.User{
				ID:     1,
				Active: true,
			},
		},
		{
			name:          "Activate inexistent user",
			userID:        1,
			repoReadError: errors.New("repo: user doesn't exists"),
			wantErr:       errors.New("repo: user doesn't exists"),
		},
		{
			name:            "Can't activate user",
			userID:          1,
			repoUpdateError: errors.New("repo: user can't be updated"),
			wantErr:         errors.New("repo: user can't be updated"),
			user: entity.User{
				ID:     1,
				Active: false,
			},
			userToStore: entity.User{
				ID:     1,
				Active: true,
			},
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewUserRepo(t)
			logger := logger.NewZeroLog()
			s := NewUserService(mockRepo, logger)

			mockRepo.On("Read", tt.userID).Return(tt.user, tt.repoReadError)

			if tt.repoReadError == nil {
				mockRepo.On("Update", tt.userToStore).Return(tt.repoUpdateError)
			}

			gotErr := s.Activate(tt.userID)

			if tt.wantErr != nil {
				assert.Error(t, gotErr)
				assert.EqualError(t, gotErr, tt.wantErr.Error())
			} else {
				assert.Nil(t, gotErr)
			}
		})
	}
}

func TestChangeUserEmail(t *testing.T) {
	testsCases := []struct {
		name            string
		userID          uint64
		repoReadError   error
		repoUpdateError error
		wantErr         error
		user            entity.User
		userToStore     entity.User
		newEmail        string
	}{
		{
			name:     "Change email of valid user",
			userID:   1,
			newEmail: "testemail@email.com",
			user: entity.User{
				ID:    1,
				Email: "lisa@field.com",
			},
			userToStore: entity.User{
				ID:    1,
				Email: "testemail@email.com",
			},
		},
		{
			name:          "Change email of inexistent user",
			userID:        1,
			repoReadError: errors.New("repo: user doesn't exists"),
			wantErr:       errors.New("repo: user doesn't exists"),
			user: entity.User{
				ID:    1,
				Email: "lisa@field.com",
			},
		},
		{
			name:            "Can't change email of user",
			userID:          1,
			newEmail:        "testemail@email.com",
			repoUpdateError: errors.New("repo: user can't be updated"),
			wantErr:         errors.New("repo: user can't be updated"),
			user: entity.User{
				ID:    1,
				Email: "lisa@field.com",
			},
			userToStore: entity.User{
				ID:    1,
				Email: "testemail@email.com",
			},
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewUserRepo(t)
			logger := logger.NewZeroLog()
			s := NewUserService(mockRepo, logger)

			mockRepo.On("Read", tt.userID).Return(tt.user, tt.repoReadError)

			if tt.repoReadError == nil {
				mockRepo.On("Update", tt.userToStore).Return(tt.repoUpdateError)
			}

			gotErr := s.ChangeEmail(tt.userID, tt.newEmail)

			if tt.wantErr != nil {
				assert.Error(t, gotErr)
				assert.EqualError(t, gotErr, tt.wantErr.Error())
			} else {
				assert.Nil(t, gotErr)
			}
		})
	}
}

func TestChangeUserPasswd(t *testing.T) {
	type testcase struct {
		name            string
		userID          uint64
		newPasswd       string
		repoReadError   error
		user            entity.User
		repoUpdateError error
		userToStore     entity.User
		wantErr         error
	}
	invalidPasswordTestCases := []testcase{
		{
			name:      "Can't change passwd of user to an invalid passwd",
			userID:    1,
			newPasswd: "p",
			wantErr:   InvalidInputErr{Field: "Passwd"},
			user: entity.User{
				ID:     1,
				Passwd: "pass123",
			},
		},
	}

	for _, tt := range invalidPasswordTestCases {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewUserRepo(t)
			logger := logger.NewZeroLog()
			s := NewUserService(mockRepo, logger)

			gotErr := s.ChangePasswd(tt.userID, tt.newPasswd)

			assert.Error(t, gotErr)
			assert.EqualError(t, gotErr, tt.wantErr.Error())
		})
	}

	updatePasswordTestCases := []testcase{
		{
			name:      "Change passwd of valid user",
			userID:    1,
			newPasswd: "newpass",
			user: entity.User{
				ID:     1,
				Passwd: "pass123",
			},
			userToStore: entity.User{
				ID:     1,
				Passwd: "newpass",
			},
		},
		{
			name:          "Change passwd of inexistent user",
			userID:        1,
			newPasswd:     "newpassword",
			repoReadError: errors.New("mockRepo: user doesn't exists"),
			wantErr:       errors.New("mockRepo: user doesn't exists"),
		},
		{
			name:            "Can't change password of user",
			userID:          1,
			newPasswd:       "newpassword",
			repoUpdateError: errors.New("mockRepo: user can't be updated"),
			wantErr:         errors.New("mockRepo: user can't be updated"),
			user: entity.User{
				ID:     1,
				Passwd: "pass123",
			},
			userToStore: entity.User{
				ID:     1,
				Passwd: "newpassword",
			},
		},
	}

	for _, tt := range updatePasswordTestCases {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewUserRepo(t)
			logger := logger.NewZeroLog()
			s := NewUserService(mockRepo, logger)

			mockRepo.On("Read", tt.userID).Return(tt.user, tt.repoReadError)

			if tt.repoReadError == nil {
				mockRepo.On("Update", tt.userToStore).Return(tt.repoUpdateError)
			}

			gotErr := s.ChangePasswd(tt.userID, tt.newPasswd)

			if tt.wantErr != nil {
				assert.Error(t, gotErr)
				assert.EqualError(t, gotErr, tt.wantErr.Error())
			} else {
				assert.Nil(t, gotErr)
			}
		})
	}
}

func TestFindUsers(t *testing.T) {
	type testcase struct {
		name      string
		filter    string
		value     string
		users     []entity.User
		repoErr   error
		wantErr   error
		wantUsers []entity.User
	}
	users := []entity.User{
		{ID: 1, FirstName: "Lisa", LastName: "LastnameA", Username: "lisa1", Email: "lisa1@email.com"},
		{ID: 2, FirstName: "Joan", LastName: "LastnameB", Username: "joan1", Email: "joan1@email.com"},
		{ID: 3, FirstName: "Juan", LastName: "LastnameC", Username: "juan1", Email: "juan1@email.com"},
	}
	findUsersByFiltersTestCases := []testcase{
		{
			name:   "Find by FirstName",
			users:  users,
			filter: "FirstName",
			value:  "Lisa",
			wantUsers: []entity.User{
				{
					ID:        1,
					FirstName: "Lisa",
					LastName:  "LastnameA",
					Username:  "lisa1",
					Email:     "lisa1@email.com",
				},
			},
		},
		{
			name:   "Find by LastName",
			users:  users,
			filter: "LastName",
			value:  "LastnameB",
			wantUsers: []entity.User{
				{
					ID:        2,
					FirstName: "Joan",
					LastName:  "LastnameB",
					Username:  "joan1",
					Email:     "joan1@email.com",
				},
			},
		},
		{
			name:   "Find by Email",
			users:  users,
			filter: "Email",
			value:  "juan1@email.com",
			wantUsers: []entity.User{
				{
					ID:        3,
					FirstName: "Juan",
					LastName:  "LastnameC",
					Username:  "juan1",
					Email:     "juan1@email.com",
				},
			},
		},
		{
			name:   "Find by Username",
			users:  users,
			filter: "Username",
			value:  "joan1",
			wantUsers: []entity.User{
				{
					ID:        2,
					FirstName: "Joan",
					LastName:  "LastnameB",
					Username:  "joan1",
					Email:     "joan1@email.com",
				},
			},
		},
		{
			name:      "Repository error",
			users:     []entity.User{},
			repoErr:   errors.New("repo failed to fetch users"),
			filter:    "LastName",
			value:     "lisa1",
			wantUsers: []entity.User{},
			wantErr:   errors.New("repo failed to fetch users"),
		},
	}

	for _, tt := range findUsersByFiltersTestCases {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewUserRepo(t)
			logger := logger.NewZeroLog()
			s := NewUserService(mockRepo, logger)

			mockRepo.On("ReadAll").Return(tt.users, tt.repoErr)

			gotUsers, gotErr := s.Find(tt.filter, tt.value)

			if tt.wantErr != nil {
				assert.Error(t, gotErr)
				assert.EqualError(t, gotErr, tt.wantErr.Error())
			} else {
				assert.Nil(t, gotErr)
			}

			assert.Equal(t, tt.wantUsers, gotUsers)
		})
	}

	validateFiltersTestCases := []testcase{
		{
			name: "Unsopported filter",
			users: []entity.User{
				{ID: 1, Username: "lisafield@mail.com"},
			},
			filter:    "TestFilter",
			value:     "lisafield@mail.com",
			wantUsers: []entity.User{},
			wantErr:   InvalidFilter{Filter: "TestFilter"},
		},
	}

	for _, tt := range validateFiltersTestCases {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewUserRepo(t)
			logger := logger.NewZeroLog()
			s := NewUserService(mockRepo, logger)

			gotUsers, gotErr := s.Find(tt.filter, tt.value)

			assert.Error(t, gotErr)
			assert.EqualError(t, gotErr, tt.wantErr.Error())

			assert.Equal(t, tt.wantUsers, gotUsers)
		})
	}
}

func TestValidateLogin(t *testing.T) {
	users := []entity.User{
		{ID: 1, Username: "user1", Passwd: "pass123"},
		{ID: 2, Username: "user2", Passwd: "pass456"},
		{ID: 3, Username: "user3", Passwd: "pass789"},
	}

	testsCases := []struct {
		name     string
		username string
		password string
		repoErr  error
		users    []entity.User
		wantErr  error
		wantUser entity.User
	}{
		{
			name:     "Valid password",
			username: "user1",
			password: "pass123",
			users:    users,
			wantUser: entity.User{
				ID:       1,
				Username: "user1",
				Passwd:   "pass123",
			},
		},
		{
			name:     "Invalid password",
			username: "user1",
			password: "pass567",
			users:    users,
			wantUser: entity.User{},
			wantErr:  ErrInvalidPassword,
		},
		{
			name:     "Repository error",
			username: "user1",
			password: "pass567",
			repoErr:  errors.New("Repository failed to fetch users"),
			users:    []entity.User{},
			wantUser: entity.User{},
			wantErr:  errors.New("Repository failed to fetch users"),
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewUserRepo(t)
			logger := logger.NewZeroLog()
			s := NewUserService(mockRepo, logger)

			mockRepo.On("ReadAll").Return(tt.users, tt.repoErr)

			gotUser, gotErr := s.ValidateLogin(tt.username, tt.password)

			if tt.wantErr != nil {
				assert.Error(t, gotErr)
				assert.EqualError(t, gotErr, tt.wantErr.Error())
			} else {
				assert.Nil(t, gotErr)
			}

			assert.Equal(t, tt.wantUser, gotUser)
		})
	}
}
