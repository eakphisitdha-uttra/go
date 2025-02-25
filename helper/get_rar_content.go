package helpers

import (
	"fmt"
	"io"

	"github.com/nwaples/rardecode"
)

// อ่านรายการช้อมูลในไฟล์ RAR โดยการส่ง RARfilePath ("/path/to/your/archive.rar")
func GetRARContents(RARfilePath string) (string, error) {

	rr, err := rardecode.OpenReader(RARfilePath, "")

	if err != nil {
		return "", fmt.Errorf("read: failed to create reader: %v", err)
	}

	header, err := rr.Next()
	if err == io.EOF {
		return "", fmt.Errorf("archive is empty: %v", err)
	}
	return header.Name, nil
}
