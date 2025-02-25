package helpers

import "golang.org/x/crypto/bcrypt"

// รับ password และ return hashed password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
