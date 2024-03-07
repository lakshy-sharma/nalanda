package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// This function is used for hashing the password.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte{}, bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("couldnt hash password %w", err)
	}
	return string(hashedPassword), nil
}

// This function is used for verifying the password.
func VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}
