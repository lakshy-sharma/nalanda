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
