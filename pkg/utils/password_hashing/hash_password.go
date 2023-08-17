package password_hashing

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string, saltSize int) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), saltSize)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func IsPasswordsMatch(password string, hashedPassword []byte) bool {
	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)); err != nil {
		return false
	}
	return true
}
