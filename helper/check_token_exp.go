package helpers

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// รับ token ที่เป็น string และ return true ถ้่า token หมดอายุแล้ว
func CheckTokenExp(stringToken string) (bool, error) {
	token, err := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
		// Normally you will provide the key for verification here
		// return []byte("your-256-bit-secret"), nil
		return nil, nil
	})

	if err != nil {
		return false, err // Return false and the error
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expTime := int64(claims["exp"].(float64))

		if time.Now().Unix() > expTime {
			return true, nil // Token is expired
		}
	}
	return false, errors.New("token is not expired") // Token is not expired
}
