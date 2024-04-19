package util

import (
	"golang.org/x/crypto/bcrypt"
)

// hashPassword generates a hashed version of the password using bcrypt
func HashPassword(passwd string) (string, error) {
	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPasswd), nil
}

func CompareHashAndPassword(hashedPassword, passwd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwd))
	if err != nil {
		return err
	}
	return nil
}
