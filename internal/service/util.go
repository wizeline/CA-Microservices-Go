package service

import (
	"golang.org/x/crypto/bcrypt"
)

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
