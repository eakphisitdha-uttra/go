package helpers

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"

	"github.com/nwaples/rardecode"
)

// แตกไฟล์ RAR โดยการส่ง RARfilePath (/path/to/your/archive.rar) และ extractDestination (/destination/folder)
func ExtractRAR(RARfilePath string, extractDestination string) error {

	rr, err := rardecode.OpenReader(RARfilePath, "")

	if err != nil {
		return fmt.Errorf("read: failed to create reader: %v", err)
	}

	for {
		header, err := rr.Next()
		if err == io.EOF {
			break
		}

		if header.IsDir {
			err = Mkdir(filepath.Join(extractDestination, header.Name), 0755)
			if err != nil {
				return err
			}
			continue
		}
		err = Mkdir(filepath.Dir(filepath.Join(extractDestination, header.Name)), 0755)
		if err != nil {
			return err
		}

		err = writeNewFile(filepath.Join(extractDestination, header.Name), rr, header.Mode())
		if err != nil {
			return err
		}
	}
	return nil
}

func writeNewFile(path string, in io.Reader, mode os.FileMode) error {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return fmt.Errorf("%s: creating directory for file: %v", path, err)
	}

	out, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("%s: creating new file: %v", path, err)
	}
	defer out.Close()

	err = out.Chmod(mode)
	if err != nil && runtime.GOOS != "windows" {
		return fmt.Errorf("%s: changing file mode: %v", path, err)
	}

	_, err = io.Copy(out, in)
	if err != nil {
		return fmt.Errorf("%s: writing file: %v", path, err)
	}
	return nil
}
