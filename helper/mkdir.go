package helpers

import (
	"fmt"
	"os"
)

// สร้าง folder โดยการส่ง path ("/root/myDirectory")
//
// dirMode
//
// 0755: เจ้าของสามารถอ่าน เขียน และรันได้ แต่กลุ่มและผู้อื่นสามารถอ่านและรันได้เท่านั้น (rwxr-xr-x)
//
// 0777: เจ้าของ, กลุ่ม, และผู้อื่นสามารถอ่าน เขียน และรันได้ (rwxrwxrwx)
//
// 0644: เจ้าของสามารถอ่านและเขียนได้ แต่กลุ่มและผู้อื่นสามารถอ่านได้เท่านั้น (rw-r--r--)
func Mkdir(path string, dirMode os.FileMode) error {
	err := os.MkdirAll(path, dirMode)
	if err != nil {
		return fmt.Errorf("%s: creating directory: %v", path, err)
	}
	return nil
}
