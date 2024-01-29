package password

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword is a function to hash password
func Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hash), err
}
