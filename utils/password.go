package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Use bcrypt to hash password
func HashPassword(password string) (string, error) {
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedpassword), nil
}

// Checks if Input password is correct
func CheckPassword(password string, hashedpassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
}
