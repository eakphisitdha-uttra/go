package helpers

// ตรวจสอบสิทธิ์ รับ permissions และ accessKeys จะ return true ถ้ามีสิทธิ์
func CheckPermission(permissions, accessKeys []string) bool {
	hasPermission := false
	for _, item := range accessKeys {
		if contains(permissions, item) {
			hasPermission = true
			break
		}
	}
	return hasPermission
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
