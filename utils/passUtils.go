package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(plainPass string) (string, error) {
	pass := []byte(plainPass)
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	return string(hash), err
}

func ComparePasswords(hashedPass string, plainPass string) (bool, error) {
	pass := []byte(plainPass)
	byteHash := []byte(hashedPass)
	err := bcrypt.CompareHashAndPassword(byteHash, pass)
	if err != nil {
		return false, err
	}
	return true, nil
}
