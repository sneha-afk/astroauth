package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return []byte(""), fmt.Errorf("HashPassword(): %v", err)
	}
	return hashed, nil
}

func VerifyPassword(hashed []byte, attempt string) bool {
	// Returns nil if it matches
	matches := bcrypt.CompareHashAndPassword(hashed, []byte(attempt))
	return matches == nil
}
