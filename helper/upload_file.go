package helpers

import (
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UploadFileStruct struct {
	FullPath string
	Name     string
	Type     string
	Size     int64
}

// อัพโหลดไฟล์ โดยการส่ง *gin.Context และ folderPath ("/path/to/your_folder")
func UploadFile(c *gin.Context, folderPath string) (UploadFileStruct, error) {

	var fileHeader *multipart.FileHeader
	if err := c.ShouldBind(&fileHeader); err != nil {
		return UploadFileStruct{}, err
	}

	dateStr := time.Now().Format("2006/01/02")
	storages := folderPath + "/" + dateStr + "/"

	if _, err := os.Stat(storages); os.IsNotExist(err) {
		err := os.MkdirAll(storages, 0755)
		if err != nil {
			return UploadFileStruct{}, err
		}
	}

	uid := uuid.NewString()
	fullPath := filepath.Join(storages, uid+fileHeader.Filename)
	if err := c.SaveUploadedFile(fileHeader, fullPath); err != nil {
		return UploadFileStruct{}, err
	}

	resp := UploadFileStruct{
		FullPath: fullPath,
		Name:     uid + fileHeader.Filename,
		Type:     filepath.Ext(fileHeader.Filename),
		Size:     fileHeader.Size,
	}

	return resp, nil
}
