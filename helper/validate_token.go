package helpers

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

// รับ token และ secretKey เพื่อตรวจสอบ token  ถ้าเป็น token ที่สร้างจาก secretKey จะ return true แต่จะ return false ถ้าไม่ใช่
func ValidateToken(stringToken, secretKey string) (bool, error) {
	// Parse the token
	token, err := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
		// Make sure that the token method conforms to "HMAC" algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return false, err
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims) // You can access token claims here e.g. claims["name"]
		return true, nil
	} else {
		return false, errors.New("invalid token")
	}
}
