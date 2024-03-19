package service

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/wizeline/CA-Microservices-Go/internal/domain/entity"
	"github.com/wizeline/CA-Microservices-Go/internal/domain/repository/mocks"
)

func TestAddUser(t *testing.T) {
	user := entity.User{
		ID:        1,
		FirstName: "lisa",
		LastName:  "Field",
		Email:     "lisa@field.com",
		BirthDay:  time.Date(1998, 12, 20, 0, 0, 0, 0, time.UTC),
		Username:  "lisa",
		Passwd:    "pass123",
		Active:    true,
	}
	testsCases := []struct {
		name           string
		wantsErr       bool
		wantID         int
		wantErrMessage string
		repoError      error
	}{
		{
			name:           "success",
			wantsErr:       false,
			wantErrMessage: "",
			repoError:      nil,
		},
		{
			name:           "failure",
			wantsErr:       true,
			wantErrMessage: "Couldn't create user",
			repoError:      errors.New("Couldn't write to db"),
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewUserRepository(t)
			s := NewUser(mockRepo)
			mockRepo.On("Create", user).Return(tt.repoError)

			_, gotErr := s.Add(user)

			if tt.wantsErr {
				assert.Error(t, gotErr)
				assert.Contains(t, gotErr.Error(), tt.wantErrMessage)
			} else {
				assert.Nil(t, gotErr)
			}
		})
	}

}

func TestGetUser(t *testing.T) {

	user := entity.User{
		ID:        1,
		FirstName: "lisa",
		LastName:  "Field",
		Email:     "lisa@field.com",
		BirthDay:  time.Date(1998, 12, 20, 0, 0, 0, 0, time.UTC),
		Username:  "lisa",
		Passwd:    "pass123",
		Active:    true,
	}
	testsCases := []struct {
		name           string
		userID         int
		repoErr        error
		wantsErr       bool
		wantErrMessage string
		wantUser       entity.User
	}{
		{
			name:           "success",
			userID:         user.ID,
			repoErr:        nil,
			wantsErr:       false,
			wantErrMessage: "",
			wantUser:       user,
		},
		{
			name:           "failure",
			userID:         user.ID,
			repoErr:        errors.New("some error happened"),
			wantsErr:       true,
			wantErrMessage: "Couldn't get user",
			wantUser:       entity.User{},
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewUserRepository(t)
			s := NewUser(repo)
			repo.On("Read", tt.userID).Return(tt.wantUser, tt.repoErr)

			gotUser, gotErr := s.Get(tt.userID)

			if tt.wantsErr {
				assert.Error(t, gotErr)
				assert.Contains(t, gotErr.Error(), tt.wantErrMessage)
			} else {
				assert.Nil(t, gotErr)
			}

			assert.Equal(t, tt.wantUser, gotUser)
		})
	}

}

func TestUpdateUser(t *testing.T) {

	userData := entity.User{
		ID:        1,
		FirstName: "lisa",
		LastName:  "Field",
		Email:     "lisa@field.com",
		BirthDay:  time.Date(1998, 12, 20, 0, 0, 0, 0, time.UTC),
		Username:  "lisa",
		Passwd:    "pass123",
		Active:    true,
	}
	testsCases := []struct {
		name           string
		userID         int
		repoErr        error
		wantsErr       bool
		wantErrMessage string
		userData       entity.User
	}{
		{
			name:           "success",
			userID:         userData.ID,
			repoErr:        nil,
			wantsErr:       false,
			wantErrMessage: "",
			userData:       userData,
		},
		{
			name:           "failure",
			userID:         userData.ID,
			repoErr:        errors.New("some error happened"),
			wantsErr:       true,
			wantErrMessage: "Couldn't update user",
			userData:       userData,
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewUserRepository(t)
			s := NewUser(repo)
			repo.On("Update", tt.userData).Return(tt.repoErr)

			gotErr := s.Update(tt.userID, tt.userData)

			if tt.wantsErr {
				assert.Error(t, gotErr)
				assert.Contains(t, gotErr.Error(), tt.wantErrMessage)
			} else {
				assert.Nil(t, gotErr)
			}
		})
	}

}

func TestDeleteUser(t *testing.T) {
	testsCases := []struct {
		name           string
		userID         int
		repoErr        error
		wantsErr       bool
		wantErrMessage string
	}{
		{
			name:           "success",
			userID:         1,
			repoErr:        nil,
			wantsErr:       false,
			wantErrMessage: "",
		},
		{
			name:           "failure",
			userID:         1,
			repoErr:        errors.New("some error happened"),
			wantsErr:       true,
			wantErrMessage: "Couldn't delete user",
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewUserRepository(t)
			s := NewUser(repo)
			repo.On("Delete", tt.userID).Return(tt.repoErr)

			gotErr := s.Delete(tt.userID)

			if tt.wantsErr {
				assert.Error(t, gotErr)
				assert.Contains(t, gotErr.Error(), tt.wantErrMessage)
			} else {
				assert.Nil(t, gotErr)
			}
		})
	}

}

func TestUserIsActive(t *testing.T) {
	user := entity.User{
		ID:        1,
		FirstName: "lisa",
		LastName:  "Field",
		Email:     "lisa@field.com",
		BirthDay:  time.Date(1998, 12, 20, 0, 0, 0, 0, time.UTC),
		Username:  "lisa",
		Passwd:    "pass123",
		Active:    true,
	}
	testsCases := []struct {
		name           string
		user           entity.User
		userID         int
		wantIsActive   bool
		repoErr        error
		wantsErr       bool
		wantErrMessage string
	}{
		{
			name:           "success",
			user:           user,
			userID:         user.ID,
			wantIsActive:   true,
			repoErr:        nil,
			wantsErr:       false,
			wantErrMessage: "",
		},
		{
			name:           "failure",
			user:           entity.User{},
			userID:         1,
			wantIsActive:   false,
			repoErr:        errors.New("some error happened"),
			wantsErr:       true,
			wantErrMessage: "Couldn't get user",
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			repo := mocks.NewUserRepository(t)
			s := NewUser(repo)
			repo.On("Read", tt.userID).Return(tt.user, tt.repoErr)

			gotIsActive, gotErr := s.IsActive(tt.userID)

			if tt.wantsErr {
				assert.Error(t, gotErr)
				assert.Contains(t, gotErr.Error(), tt.wantErrMessage)
			} else {
				assert.Nil(t, gotErr)
			}
			assert.Equal(t, tt.wantIsActive, gotIsActive)
		})
	}

}
