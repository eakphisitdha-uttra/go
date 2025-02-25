package helpers

import (
	"fmt"
)

// ตรวจสอบความถูกค้องของเบอร์โทรศัพท์
func ValidateTel(tel string) error {

	if len(tel) == 0 {
		return fmt.Errorf("zero value")
	}

	if len(tel) < 10 { // Assuming 10 is the minimum length
		return fmt.Errorf("less than min")
	}

	// Add more conditions as needed

	return nil
}
