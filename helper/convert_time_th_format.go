package helpers

import (
	"fmt"
	"time"
)

// map เดือนเป็นภาษาไทย
var ThaiMonths = []string{
	"มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน",
	"กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม",
}

var ThaiMths = []string{
	"ม.ค.", "ก.พ.", "มี.ค.", "เม.ย.", "พ.ค.", "มิ.ย.",
	"ก.ค.", "ส.ค.", "ก.ย.", "ต.ค.", "พ.ย.", "ธ.ค.",
}

// format type
//
// 1: 31 มกราคม 2567 00:00:00 น.
//
// 2: 31 มกราคม 2567 00:00 น.
//
// 3: 31 มกราคม 2567
//
// 4: 31 ม.ค. 2567 00:00:00 น.
//
// 5: 31 ม.ค. 2567 00:00 น.
//
// 6: 31 ม.ค. 2567
//
// 7: 31 ม.ค. 67 00:00:00 น.
//
// 8: 31 ม.ค. 67 00:00 น.
//
// 9: 31 ม.ค. 67
//
// 10: 31/01/2567 00:00:00 น.
//
// 11: 31/01/2567 00:00 น.
//
// 12: 31/01/2567
//
// 13: 31/01/67 00:00:00 น.
//
// 14: 31/01/67 00:00 น.
//
// 15: 31/01/67
func ConvertTimeTHFormat(t time.Time, formatType int) string {
	yearTh := t.Year() + 543    // ปี พ.ศ.
	shortYearTh := yearTh % 100 // ปี พ.ศ. แบบสองหลัก
	monthFull := ThaiMonths[t.Month()]
	monthShort := ThaiMths[t.Month()]
	date := t.Day()

	var formattedTime string

	switch formatType {
	case 1:
		formattedTime = fmt.Sprintf("%02d %s %d %02d:%02d:%02d น.", date, monthFull, yearTh, t.Hour(), t.Minute(), t.Second())
	case 2:
		formattedTime = fmt.Sprintf("%02d %s %d %02d:%02d น.", date, monthFull, yearTh, t.Hour(), t.Minute())
	case 3:
		formattedTime = fmt.Sprintf("%02d %s %d", date, monthFull, yearTh)
	case 4:
		formattedTime = fmt.Sprintf("%02d %s %d %02d:%02d:%02d น.", date, monthShort, yearTh, t.Hour(), t.Minute(), t.Second())
	case 5:
		formattedTime = fmt.Sprintf("%02d %s %d %02d:%02d น.", date, monthShort, yearTh, t.Hour(), t.Minute())
	case 6:
		formattedTime = fmt.Sprintf("%02d %s %d", date, monthShort, yearTh)
	case 7:
		formattedTime = fmt.Sprintf("%02d %s %02d %02d:%02d:%02d น.", date, monthShort, shortYearTh, t.Hour(), t.Minute(), t.Second())
	case 8:
		formattedTime = fmt.Sprintf("%02d %s %02d %02d:%02d น.", date, monthShort, shortYearTh, t.Hour(), t.Minute())
	case 9:
		formattedTime = fmt.Sprintf("%02d %s %02d", date, monthShort, shortYearTh)
	case 10:
		formattedTime = fmt.Sprintf("%02d/%02d/%d %02d:%02d:%02d น.", date, t.Month(), yearTh, t.Hour(), t.Minute(), t.Second())
	case 11:
		formattedTime = fmt.Sprintf("%02d/%02d/%d %02d:%02d น.", date, t.Month(), yearTh, t.Hour(), t.Minute())
	case 12:
		formattedTime = fmt.Sprintf("%02d/%02d/%d", date, t.Month(), yearTh)
	case 13:
		formattedTime = fmt.Sprintf("%02d/%02d/%02d %02d:%02d:%02d น.", date, t.Month(), shortYearTh, t.Hour(), t.Minute(), t.Second())
	case 14:
		formattedTime = fmt.Sprintf("%02d/%02d/%02d %02d:%02d น.", date, t.Month(), shortYearTh, t.Hour(), t.Minute())
	case 15:
		formattedTime = fmt.Sprintf("%02d/%02d/%02d", date, t.Month(), shortYearTh)
	default:
		formattedTime = "Invalid format type"
	}

	return formattedTime
}
