package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CompareHashPassword(hashPassword string, password string) error {
	// return nil when password matches hash
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
