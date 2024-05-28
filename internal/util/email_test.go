package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		err     error
		wantErr bool
	}{
		{
			name:    "Empty",
			email:   "",
			err:     ErrEmptyValue,
			wantErr: true,
		},
		{
			name:    "Invalid",
			email:   "foo@example",
			err:     ErrInvalidEmail,
			wantErr: true,
		},
		{
			name:    "Valid",
			email:   "foo@example.com",
			err:     nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEmail(tt.email)
			if tt.wantErr {
				assert.ErrorIs(t, err, tt.err)
				return
			}
			assert.Nil(t, err)
		})
	}
}
