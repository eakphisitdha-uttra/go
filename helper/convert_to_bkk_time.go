package helpers

import "time"

// รับ time.Time และแปลงเป็นเวลา กรุงเทพฯ ประเทศไทย
func ConvertToBKKTime(t time.Time) time.Time {
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return time.Time{}
	}
	return t.In(location) // return current time in Bangkok
}
