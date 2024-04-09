package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wizeline/CA-Microservices-Go/internal/entity"
	"github.com/wizeline/CA-Microservices-Go/internal/util"
)

func Test_validateUser(t *testing.T) {
	tests := []struct {
		name    string
		user    entity.User
		err     error
		wantErr bool
	}{
		{
			name:    "Entity Empty",
			user:    entity.User{},
			err:     &EntityEmptyErr{Name: "User"},
			wantErr: true,
		},
		{
			name: "Invalid Field",
			user: entity.User{
				FirstName: "foo",
				LastName:  "",
				Email:     "foo@example.com",
			},
			err:     &InvalidFieldErr{Name: "LastName", Err: ErrFieldEmpty},
			wantErr: true,
		},
		{
			name: "Invalid Email",
			user: entity.User{
				FirstName: "foo",
				LastName:  "bar",
				Email:     "foo@example",
			},
			err:     &InvalidFieldErr{Name: "Email", Err: util.ErrInvalidEmail},
			wantErr: true,
		},
		{
			name: "Valid",
			user: entity.User{
				FirstName: "foo",
				LastName:  "bar",
				Email:     "foo@example.com",
				Username:  "foouser",
				Passwd:    "foopasswd",
			},
			err:     nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateUser(tt.user)
			if tt.wantErr {
				assert.Equal(t, tt.err, err)
				return
			}
			assert.Nil(t, err)
		})
	}
}
