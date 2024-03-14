package util

import (
	"regexp"
)

func IsValidEmail(email string) error {
	if email == "" {
		return ErrValueEmpty
	}
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	if !regex.MatchString(email) {
		return ErrInvalidEmail
	}
	return nil
}
