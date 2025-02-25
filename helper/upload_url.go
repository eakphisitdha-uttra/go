package helpers

import (
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

type UploadURLStruct struct {
	FullPath string
	Name     string
	Type     string
	Size     int64
}

// อัพโหลดไฟล์ โดยการรับ URL และ folderPath ("/path/to/your_folder")
func UploadURL(url, folderPath string) (UploadURLStruct, error) {

	urlSplit := strings.Split(url, "/")
	name := urlSplit[len(urlSplit)-1]
	name = uuid.NewString() + name
	ext := filepath.Ext(name)

	dateStr := time.Now().Format("2006/01/02")
	storages := folderPath + "/" + dateStr + "/"
	if _, err := os.Stat(storages); os.IsNotExist(err) {
		err := os.MkdirAll(storages, 0755)
		if err != nil {
			return UploadURLStruct{}, err
		}
	}
	fullPath := path.Join(storages, name)

	resp, err := http.Get(url)
	if err != nil {
		return UploadURLStruct{}, err
	}
	defer resp.Body.Close()

	out, err := os.Create(fullPath)
	if err != nil {
		return UploadURLStruct{}, err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return UploadURLStruct{}, err
	}

	result := UploadURLStruct{
		FullPath: fullPath,
		Name:     name,
		Type:     ext,
		Size:     resp.ContentLength,
	}

	return result, nil
}
