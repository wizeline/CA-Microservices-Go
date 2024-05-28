package util

import (
	"regexp"
)

func ValidateEmail(v string) error {
	if v == "" {
		return ErrEmptyValue
	}
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	if !regex.MatchString(v) {
		return ErrInvalidEmail
	}
	return nil
}
