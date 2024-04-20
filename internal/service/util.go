package service

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func validateEmail(email string) error {
	if email == "" {
		return ErrEmptyValue
	}
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	if !regex.MatchString(email) {
		return ErrInvalidEmail
	}
	return nil
}

// hashPasswd generates a hashed version of the password using bcrypt
func hashPasswd(passwd string) (string, error) {
	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPasswd), nil
}

func compareHashAndPassword(hashedPassword, passwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwd))
}
