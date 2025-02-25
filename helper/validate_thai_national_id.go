package helpers

import (
	"errors"
	"strconv"
)

// ตรวจสอบความถูกค้องของรหัสบัตรประจำตัวประชาชน
func ThaiNationalIDValidator(id string) error {
	if len(id) != 13 {
		return errors.New("invalid length")
	}

	sum := 0
	for i := 0; i < 12; i++ {
		num, err := strconv.Atoi(string(id[i]))
		if err != nil {
			return errors.New("invalid Thai National ID")
		}
		sum += num * (13 - i)
	}

	remainder := sum % 11
	checkDigit := (11 - remainder) % 10

	lastDigit, err := strconv.Atoi(string(id[12]))
	if err != nil {
		return errors.New("invalid Thai National ID")
	}

	if checkDigit != lastDigit {
		return errors.New("invalid Thai National ID")
	}

	return nil
}
