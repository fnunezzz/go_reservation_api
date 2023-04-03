package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}


func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

// TODO
// Maybe add a more complex strength checker with https://github.com/wagslane/go-password-validator
func CheckPasswordStrength(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("pasword must have at least 8 characters")
	}
	return nil
}