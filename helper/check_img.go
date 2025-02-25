package helpers

import (
	"mime/multipart"
	"net/http"
	"strings"
)

// ตรวจสอบว่า file ที่ส่งมาเป็นไฟล์ภาพหรือไม่
func CheckIMG(fileHeader *multipart.FileHeader) (bool, error) {
	// Open the file
	file, err := fileHeader.Open()
	if err != nil {
		return false, err
	}
	defer file.Close()

	// Read the first 512 bytes to determine the content type
	buf := make([]byte, 512)
	if _, err := file.Read(buf); err != nil {
		return false, err
	}

	// Use http.DetectContentType to determine the content type
	contentType := http.DetectContentType(buf)

	// Check if the content type starts with "image/"
	if contentType != "" && strings.HasPrefix(contentType, "image/") {
		return true, nil
	}

	return false, nil
}
