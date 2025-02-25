package helpers

import (
	"errors"
	"fmt"
	"unicode"
)

// ตรวจสอบเงื่อนไขการตั้งค่ารหัสผ่าน
func ValidatePassword(password string, minimun int) error {
	if password == "" {
		return fmt.Errorf("zero value")
	} else if len(password) < minimun {
		return fmt.Errorf("less than min")
	}

	var (
		hasUpper     bool
		hasLower     bool
		hasDigit     bool
		hasSpecial   bool
		allowedChars = "~`!@#$%^&*()_-+={[}]|\\:;\"'<,>.?/"
		allowedRunes = make(map[rune]bool)
	)

	for _, char := range allowedChars {
		allowedRunes[char] = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case allowedRunes[char]:
			hasSpecial = true
		}
	}

	if !hasUpper {
		return errors.New("password must contain at least one uppercase letter")
	} else if !hasLower {
		return errors.New("password must contain at least one lowercase letter")
	} else if !hasDigit {
		return errors.New("password must contain at least one digit letter")
	} else if !hasSpecial {
		return errors.New("password must contain at least one special character")
	}

	return nil
}
