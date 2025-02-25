package helpers

import "golang.org/x/crypto/bcrypt"

// รับค่า hashedPassword กับ password ที่ user ระบุ เพื่อตรวจสอบความถูกต้องของ password
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
