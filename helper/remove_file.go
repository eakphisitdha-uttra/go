package helpers

import "os"

// ลบไฟล์หรือโฟลเดอร์ว่าง โดยการส่ง path ("/path/to/your/archive.rar") หรือ ("/path/to/your_folder")
func RemoveEmpty(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

// ลบไฟล์หรือโฟลเดอร์ทั้งหมด (รวมถึงไฟล์และโฟลเดอร์ย่อยภายใน) โดยการส่ง path ("/path/to/your/archive.rar") หรือ ("/path/to/your_folder")
func RemoveAll(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}
