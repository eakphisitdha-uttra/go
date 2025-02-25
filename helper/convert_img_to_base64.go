package helpers

import (
	"encoding/base64"
	"io/ioutil"
)

// รับ imagePath ("/path/to/your_folder/your_img.png") เพื่อแปลงเป็น base64
func ConvertIMGToBase64(imagePath string) (string, error) {
	// อ่านไฟล์จาก path ที่ได้รับมา
	imageBytes, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return "", err
	}

	// แปลงไฟล์รูปภาพเป็น base64
	base64Image := base64.StdEncoding.EncodeToString(imageBytes)

	return base64Image, nil
}
