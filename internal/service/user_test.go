package service

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wizeline/CA-Microservices-Go/internal/entity"
	"github.com/wizeline/CA-Microservices-Go/internal/service/mocks"
	"golang.org/x/crypto/bcrypt"
)

var (
	// We ensure the UserRepo mock object satisfies the UserRepo signature.
	_ UserRepo = &mocks.UserRepo{}

	errRepoTest = errors.New("some repo error")
)

func TestUserService_Create(t *testing.T) {
	type repo struct {
		args entity.User
		err  error
	}
	tests := []struct {
		name string
		repo repo
		args UserCreateArgs
		err  error
	}{
		{
			name: "Empty firstName",
			args: UserCreateArgs{
				FirstName: "",
				LastName:  "Field",
				Email:     "lisa@field.com",
				Username:  "lisa",
				Passwd:    "pass123",
			},
			err: &InvalidInputErr{Field: "FirstName", Err: ErrEmptyValue},
		},
		{
			name: "Empty lastname)",
			repo: repo{},
			args: UserCreateArgs{
				FirstName: "Lisa",
				LastName:  "",
				Email:     "lisa@field.com",
				Username:  "lisa",
				Passwd:    "pass123",
			},
			err: &InvalidInputErr{Field: "LastName", Err: ErrEmptyValue},
		},
		{
			name: "Empty email",
			repo: repo{},
			args: UserCreateArgs{
				FirstName: "Lisa",
				LastName:  "Field",
				Email:     "",
				Username:  "lisa",
				Passwd:    "pass123",
			},
			err: &InvalidInputErr{Field: "Email", Err: ErrEmptyValue},
		},
		{
			name: "Invalid password",
			repo: repo{},
			args: UserCreateArgs{
				FirstName: "Lisa",
				LastName:  "Field",
				Email:     "lisa@field.com",
				Username:  "lisa",
				Passwd:    "12345",
			},
			err: &InvalidInputErr{Field: "Passwd", Err: ErrInvalidPasswd},
		},
		{
			name: "Repository error",
			repo: repo{
				args: entity.User{
					ID:        uint64(0),
					FirstName: "lisa",
					LastName:  "Field",
					Email:     "lisa@field.com",
					Username:  "lisa",
					Active:    false,
				},
				err: errRepoTest,
			},
			args: UserCreateArgs{
				FirstName: "lisa",
				LastName:  "Field",
				Email:     "lisa@field.com",
				Username:  "lisa",
				Passwd:    "pass123",
			},
			err: errRepoTest,
		},
		{
			name: "Valid user",
			repo: repo{
				args: entity.User{
					ID:        uint64(0),
					FirstName: "lisa",
					LastName:  "Field",
					Email:     "lisa@field.com",
					Username:  "lisa",
					Active:    false,
				},
				err: nil,
			},
			args: UserCreateArgs{
				FirstName: "lisa",
				LastName:  "Field",
				Email:     "lisa@field.com",
				Username:  "lisa",
				Passwd:    "pass123",
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.UserRepo{}
			// TODO: migrate the Create mocked function to validate the expected arguments to the repository. Currently, there are some issues due to the hashed password.
			// mockRepo.On("Create", tt.repo.args).Return(tt.repo.err)
			mockRepo.On("Create", mock.AnythingOfType("entity.User")).Return(tt.repo.err)
			svc := NewUserService(mockRepo)

			err := svc.Create(tt.args)

			if tt.err != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err.Error())
				return
			}

			assert.Nil(t, err)
		})
	}
}

func TestUserService_Get(t *testing.T) {
	type repoResp struct {
		user entity.User
		err  error
	}
	type repo struct {
		args uint64
		resp repoResp
	}
	tests := []struct {
		name   string
		repo   repo
		userID uint64
		exp    entity.User
		err    error
	}{
		{
			name:   "ID zero value",
			userID: 0,
			exp:    entity.User{},
			err:    ErrZeroValue,
		},
		{
			name:   "Repository error",
			userID: 1,
			repo: repo{
				args: 1,
				resp: repoResp{
					user: entity.User{},
					err:  errRepoTest,
				},
			},
			err: errRepoTest,
			exp: entity.User{},
		},
		{
			name: "Valid user ID",
			repo: repo{
				args: 1,
				resp: repoResp{
					user: entity.User{
						ID:        1,
						FirstName: "lisa",
						LastName:  "Field",
						Email:     "lisa@field.com",
						BirthDay:  time.Date(1998, 12, 20, 0, 0, 0, 0, time.UTC),
						Username:  "lisa",
						Passwd:    "pass123",
						Active:    true,
					},
					err: nil,
				},
			},
			userID: 1,
			exp: entity.User{
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.UserRepo{}
			mockRepo.On("Read", tt.repo.args).Return(tt.repo.resp.user, tt.repo.resp.err)
			svc := NewUserService(mockRepo)
			out, err := svc.Get(tt.userID)

			if tt.err != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err.Error())
				return
			}

			assert.Nil(t, err)
			assert.Equal(t, tt.exp, out)
		})
	}

}

func TestUserService_GetAll(t *testing.T) {
	type repo struct {
		users []entity.User
		err   error
	}
	tests := []struct {
		name string
		repo repo
		exp  []entity.User
		err  error
	}{
		{
			name: "Repository error",
			repo: repo{
				users: []entity.User{},
				err:   errRepoTest,
			},
			exp: []entity.User{},
			err: errRepoTest,
		},
		{

			name: "Sucessful",
			repo: repo{
				users: []entity.User{
					{ID: 1, Username: "lisafield@mail.com"},
					{ID: 2, Username: "mat123@mail.com"},
					{ID: 3, Username: "juan@mail.com"},
				},
				err: nil,
			},
			exp: []entity.User{
				{ID: 1, Username: "lisafield@mail.com"},
				{ID: 2, Username: "mat123@mail.com"},
				{ID: 3, Username: "juan@mail.com"},
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewUserRepo(t)
			mockRepo.On("ReadAll").Return(tt.repo.users, tt.repo.err)
			svc := NewUserService(mockRepo)

			out, err := svc.GetAll()

			if tt.err != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err.Error())
				return
			}

			assert.Nil(t, err)
			assert.Equal(t, tt.exp, out)
		})
	}
}

func TestUserService_Update(t *testing.T) {
	type repoReadResp struct {
		user entity.User
		err  error
	}
	type repoRead struct {
		id   uint64
		resp repoReadResp
	}
	type repoUpdate struct {
		args entity.User
		err  error
	}
	type repo struct {
		read   repoRead
		update repoUpdate
	}
	tests := []struct {
		name string
		repo repo
		args UserUpdateArgs
		err  error
	}{
		{
			name: "No arguments",
			repo: repo{},
			args: UserUpdateArgs{},
			err:  ErrEmptyArgs,
		},
		{
			name: "ID zero value",
			repo: repo{},
			args: UserUpdateArgs{
				ID:        0,
				FirstName: "foo",
			},
			err: &InvalidInputErr{Field: "ID", Err: ErrZeroValue},
		},
		{
			name: "Update FirstName",
			args: UserUpdateArgs{
				ID:        1,
				FirstName: "Lisa",
			},
			err: nil,
			repo: repo{
				read: repoRead{
					id: 1,
					resp: repoReadResp{
						user: entity.User{ID: 1, FirstName: "Laura"},
						err:  nil,
					},
				},
				update: repoUpdate{
					args: entity.User{ID: 1, FirstName: "Lisa"},
				},
			},
		},
		{
			name: "Update LastName",
			args: UserUpdateArgs{
				ID:       1,
				LastName: "Lawrence",
			},
			err: nil,
			repo: repo{
				read: repoRead{
					id: 1,
					resp: repoReadResp{
						user: entity.User{ID: 1, LastName: "Miller"},
						err:  nil,
					},
				},
				update: repoUpdate{
					args: entity.User{ID: 1, LastName: "Lawrence"},
				},
			},
		},
		{
			name: "Update Birthday",
			args: UserUpdateArgs{
				ID:       1,
				BirthDay: time.Date(2000, 12, 20, 0, 0, 0, 0, time.UTC),
			},
			err: nil,
			repo: repo{
				read: repoRead{
					id: 1,
					resp: repoReadResp{
						user: entity.User{
							ID:       1,
							BirthDay: time.Date(1998, 12, 20, 0, 0, 0, 0, time.UTC),
						},
						err: nil,
					},
				},
				update: repoUpdate{
					args: entity.User{
						ID:       1,
						BirthDay: time.Date(2000, 12, 20, 0, 0, 0, 0, time.UTC),
					},
					err: nil,
				},
			},
		},
		{
			name: "Can update all fields at once",
			args: UserUpdateArgs{
				ID:        1,
				FirstName: "lisa",
				LastName:  "Lauwrence",
				BirthDay:  time.Date(2000, 12, 20, 0, 0, 0, 0, time.UTC),
			},
			err: nil,
			repo: repo{
				read: repoRead{
					id: 1,
					resp: repoReadResp{
						user: entity.User{
							ID:        1,
							FirstName: "laura",
							LastName:  "Field",
							BirthDay:  time.Date(1998, 12, 20, 0, 0, 0, 0, time.UTC),
						},
						err: nil,
					},
				},
				update: repoUpdate{
					args: entity.User{
						ID:        1,
						FirstName: "lisa",
						LastName:  "Lauwrence",
						BirthDay:  time.Date(2000, 12, 20, 0, 0, 0, 0, time.UTC),
					},
					err: nil,
				},
			},
		},
		{
			name: "Repository fails to get users",
			args: UserUpdateArgs{
				ID:        1,
				FirstName: "lisa",
			},
			err: errRepoTest,
			repo: repo{
				read: repoRead{
					id: 1,
					resp: repoReadResp{
						user: entity.User{},
						err:  errRepoTest,
					},
				},
				update: repoUpdate{},
			},
		},
		{
			name: "Repository fails to update users",
			args: UserUpdateArgs{
				ID:        1,
				FirstName: "lisa",
			},
			err: errRepoTest,
			repo: repo{
				read: repoRead{
					id: 1,
					resp: repoReadResp{
						user: entity.User{
							ID:        1,
							FirstName: "laura",
							LastName:  "Field",
						},
						err: nil,
					},
				},
				update: repoUpdate{
					args: entity.User{
						ID:        1,
						FirstName: "lisa",
						LastName:  "Field",
					},
					err: errRepoTest,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.UserRepo{}
			mockRepo.On("Read", tt.repo.read.id).Return(tt.repo.read.resp.user, tt.repo.read.resp.err)
			mockRepo.On("Update", tt.repo.update.args).Return(tt.repo.update.err)
			svc := NewUserService(mockRepo)

			err := svc.Update(tt.args)

			if tt.err != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err.Error())
				return
			}

			assert.Nil(t, err)
		})
	}
}

func TestUserService_Delete(t *testing.T) {
	type repo struct {
		id  uint64
		err error
	}
	tests := []struct {
		name string
		repo repo
		id   uint64
		err  error
	}{
		{
			name: "ID zero value",
			repo: repo{id: 0},
			id:   0,
			err:  &InvalidInputErr{Field: "id", Err: ErrZeroValue},
		},
		{
			name: "Repository error",
			repo: repo{
				id:  1,
				err: errRepoTest,
			},
			id:  1,
			err: errRepoTest,
		},

		{
			name: "User deleted",
			repo: repo{id: 1, err: nil},
			id:   1,
			err:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &mocks.UserRepo{}
			mockRepo.On("Delete", tt.repo.id).Return(tt.repo.err)
			svc := NewUserService(mockRepo)

			err := svc.Delete(tt.id)

			if tt.err != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err.Error())
				return
			}
			assert.Nil(t, err)
		})
	}
}

func TestUserService_IsActive(t *testing.T) {
	type repoResp struct {
		user entity.User
		err  error
	}
	type repo struct {
		id   uint64
		resp repoResp
	}
	tests := []struct {
		name string
		repo repo
		id   uint64
		exp  bool
		err  error
	}{
		{
			name: "Repository error",
			repo: repo{
				id:   1,
				resp: repoResp{user: entity.User{}, err: errRepoTest}},
			id:  1,
			exp: false,
			err: errRepoTest,
		},
		{
			name: "Inactive",
			repo: repo{
				id: 1,
				resp: repoResp{
					user: entity.User{ID: 1, Active: false},
					err:  nil,
				},
			},
			id:  1,
			exp: false,
			err: nil,
		},
		{
			name: "Active",
			repo: repo{
				id: 1,
				resp: repoResp{
					user: entity.User{
						ID:     1,
						Active: true,
					},
					err: nil,
				},
			},
			id:  1,
			exp: true,
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewUserRepo(t)
			mockRepo.On("Read", tt.repo.id).Return(tt.repo.resp.user, tt.repo.resp.err)
			svc := NewUserService(mockRepo)

			out, err := svc.IsActive(tt.id)

			if tt.err != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.err.Error())
				return
			}

			assert.Nil(t, err)
			assert.Equal(t, tt.exp, out)
		})
	}
}

func TestUserService_Activate(t *testing.T) {
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
			repoReadError: errors.New("mockRepo: user doesn't exists"),
			wantErr:       errors.New("mockRepo: user doesn't exists"),
		},
		{
			name:            "Can't activate user",
			userID:          1,
			repoUpdateError: errors.New("mockRepo: user can't be updated"),
			wantErr:         errors.New("mockRepo: user can't be updated"),
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
			svc := NewUserService(mockRepo)

			mockRepo.On("Read", tt.userID).Return(tt.user, tt.repoReadError)

			if tt.repoReadError == nil {
				mockRepo.On("Update", tt.userToStore).Return(tt.repoUpdateError)
			}

			gotErr := svc.Activate(tt.userID)

			if tt.wantErr != nil {
				assert.Error(t, gotErr)
				assert.EqualError(t, gotErr, tt.wantErr.Error())
			} else {
				assert.Nil(t, gotErr)
			}
		})
	}
}

func TestUserService_ChangeEmail(t *testing.T) {
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
			newEmail:      "testemail@email.com",
			repoReadError: errors.New("mockRepo: user doesn't exists"),
			wantErr:       errors.New("mockRepo: user doesn't exists"),
			user: entity.User{
				ID:    1,
				Email: "lisa@field.com",
			},
		},
		{
			name:            "Can't change email of user",
			userID:          1,
			newEmail:        "testemail@email.com",
			repoUpdateError: errors.New("mockRepo: user can't be updated"),
			wantErr:         errors.New("mockRepo: user can't be updated"),
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
			mockRepo := &mocks.UserRepo{}
			mockRepo.On("Read", tt.userID).Return(tt.user, tt.repoReadError)
			if tt.repoReadError == nil {
				mockRepo.On("Update", tt.userToStore).Return(tt.repoUpdateError)
			}
			svc := NewUserService(mockRepo)

			gotErr := svc.ChangeEmail(tt.userID, tt.newEmail)

			if tt.wantErr != nil {
				assert.Error(t, gotErr)
				assert.EqualError(t, gotErr, tt.wantErr.Error())
			} else {
				assert.Nil(t, gotErr)
			}
		})
	}
}

func TestUserService_ChangePasswd(t *testing.T) {
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
			wantErr:   ErrInvalidPasswd,
			user: entity.User{
				ID:     1,
				Passwd: "pass123",
			},
		},
	}

	for _, tt := range invalidPasswordTestCases {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewUserRepo(t)
			svc := NewUserService(mockRepo)

			gotErr := svc.ChangePasswd(tt.userID, tt.newPasswd)

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
			name:   "Repository fails to get user",
			userID: 1, newPasswd: "newpass",
			repoReadError: errors.New("mockRepo: user doesn't exists"),
			wantErr:       errors.New("mockRepo: user doesn't exists"),
		},
		{
			name:            "Repository fails to update user",
			userID:          1,
			newPasswd:       "newpass",
			repoUpdateError: errors.New("mockRepo: user can't be updated"),
			wantErr:         errors.New("mockRepo: user can't be updated"),
			user: entity.User{
				ID:     1,
				Passwd: "pass123",
			},
			userToStore: entity.User{
				ID:     1,
				Passwd: "newpass",
			},
		},
	}

	for _, tt := range updatePasswordTestCases {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewUserRepo(t)
			mockRepo.On("Read", tt.userID).Return(tt.user, tt.repoReadError)
			if tt.repoReadError == nil {
				mockRepo.On("Update", mock.AnythingOfType("entity.User")).Return(tt.repoUpdateError).Once().Run(func(args mock.Arguments) {
					userArg := args.Get(0).(entity.User)

					assert.Equal(t, tt.userID, userArg.ID)
					assert.NoError(t, bcrypt.CompareHashAndPassword([]byte(userArg.Passwd), []byte(tt.newPasswd)))
				})
			}
			svc := NewUserService(mockRepo)

			gotErr := svc.ChangePasswd(tt.userID, tt.newPasswd)

			if tt.wantErr != nil {
				assert.Error(t, gotErr)
				assert.EqualError(t, gotErr, tt.wantErr.Error())
			} else {
				assert.Nil(t, gotErr)
			}
		})
	}
}

func TestUserService_Find(t *testing.T) {
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
			mockRepo.On("ReadAll").Return(tt.users, tt.repoErr)
			svc := NewUserService(mockRepo)

			out, err := svc.Find(tt.filter, tt.value)

			if tt.wantErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.wantErr.Error())
				return
			}

			assert.Nil(t, err)
			assert.Equal(t, tt.wantUsers, out)
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
			wantUsers: nil,
			wantErr:   &InvalidFilterErr{Filter: "TestFilter", Err: ErrNotSupported},
		},
	}

	for _, tt := range validateFiltersTestCases {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewUserRepo(t)
			svc := NewUserService(mockRepo)

			gotUsers, gotErr := svc.Find(tt.filter, tt.value)

			assert.Error(t, gotErr)
			assert.EqualError(t, gotErr, tt.wantErr.Error())
			assert.Equal(t, tt.wantUsers, gotUsers)
		})
	}
}

func TestUserService_ValidateLogin(t *testing.T) {
	userPasswd := "mypass"
	hashedPasswd, _ := hashPasswd(userPasswd)

	users := []entity.User{
		{
			ID:       1,
			Username: "user1",
			Passwd:   hashedPasswd,
		},
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
			password: userPasswd,
			users:    users,
			wantUser: entity.User{
				ID:       1,
				Username: "user1",
				Passwd:   hashedPasswd,
			},
		},
		{
			name:     "User doesn't exists",
			username: "user2",
			password: userPasswd,
			users:    users,
			wantErr:  errors.New("expected one user got 0"),
		},
		{
			name:     "Invalid password",
			username: "user1",
			password: "pass567",
			users:    users,
			wantUser: entity.User{},
			wantErr:  ErrPasswdDoNotMatch,
		},
		{
			name:     "Repository fails to get user",
			username: "user4",
			password: userPasswd,
			repoErr:  errors.New("mockRepo: user doesn't exists"),
			wantErr:  errors.New("mockRepo: user doesn't exists"),
			users:    []entity.User{},
			wantUser: entity.User{},
		},
	}

	for _, tt := range testsCases {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewUserRepo(t)
			svc := NewUserService(mockRepo)

			mockRepo.On("ReadAll").Return(tt.users, tt.repoErr)

			gotUser, gotErr := svc.ValidateLogin(tt.username, tt.password)

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
