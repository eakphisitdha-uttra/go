package helpers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"os"
	"strings"
	"time"
)

type UploadBase64Struct struct {
	FullPath string
	Name     string
	Type     string
	Size     int64
}

// อัพโหลดรูปภาพจาก Base64 โดยการส่ง base64Code, ชื่อไฟล์ที่ต้องการ และ folderPath ("/path/to/your_folder")
func UploadBase64(base64Code, fileName, folderPath string) (UploadBase64Struct, error) {

	base64Data := strings.Split(base64Code, ",")
	if len(base64Data) > 1 {
		base64Code = base64Data[1]
	}

	imgBytes, err := base64.StdEncoding.DecodeString(strings.TrimSpace(base64Code))
	if err != nil {
		return UploadBase64Struct{}, err
	}
	_, imageType, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		return UploadBase64Struct{}, err
	}

	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err := os.MkdirAll(folderPath, 0755)
		if err != nil {
			return UploadBase64Struct{}, err
		}
	}

	dateStr := time.Now().Format("2006/01/02")
	imgPath := fmt.Sprintf("%s/%s/%s.%s", folderPath, dateStr, fileName, imageType)

	// Create the image file.
	imgFile, err := os.Create(imgPath)
	if err != nil {
		return UploadBase64Struct{}, err
	}
	defer imgFile.Close()

	// Write the bytes into the image file.
	_, err = imgFile.Write(imgBytes)
	if err != nil {
		return UploadBase64Struct{}, err
	}
	// Get the file size
	fileInfo, err := imgFile.Stat()
	if err != nil {
		return UploadBase64Struct{}, err
	}
	fileSize := fileInfo.Size()

	resp := UploadBase64Struct{
		FullPath: imgPath,
		Name:     fileName,
		Type:     imageType,
		Size:     fileSize,
	}

	return resp, nil
}
