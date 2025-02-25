package helpers

import (
	"fmt"
	"time"
)

// format type
//
// 1: 31 January 2024 00:00:00 AM
//
// 2: 31 January 2024 00:00 AM
//
// 3: 31 January 2024
//
// 4: 31 JAN 2024 00:00:00 AM
//
// 5: 31 JAN 2024 00:00 AM
//
// 6: 31 JAN 2024
//
// 7: 31 JAN 24 00:00:00 AM
//
// 8: 31 JAN 24 00:00 AM
//
// 9: 31 JAN 24
//
// 10: 31/01/2024 00:00:00 AM
//
// 11: 31/01/2024 00:00 AM
//
// 12: 31/01/2024
//
// 13: 31/01/24 00:00:00 AM
//
// 14: 31/01/24 00:00 AM
//
// 15: 31/01/24
func ConvertTimeENFormat(t time.Time, formatType int) string {
	year := t.Year()
	shortYear := t.Year() % 100
	monthFull := t.Month().String()
	monthShort := t.Format("Jan")
	date := t.Day()
	hour := t.Hour()

	period := "AM"
	if hour >= 12 {
		period = "PM"
	}
	if hour > 12 {
		hour -= 12
	}

	var formattedTime string

	switch formatType {
	case 1:
		formattedTime = fmt.Sprintf("%02d %s %d %02d:%02d:%02d %s", date, monthFull, year, hour, t.Minute(), t.Second(), period)
	case 2:
		formattedTime = fmt.Sprintf("%02d %s %d %02d:%02d %s", date, monthFull, year, hour, t.Minute(), period)
	case 3:
		formattedTime = fmt.Sprintf("%02d %s %d", date, monthFull, year)
	case 4:
		formattedTime = fmt.Sprintf("%02d %s %d %02d:%02d:%02d %s", date, monthShort, year, hour, t.Minute(), t.Second(), period)
	case 5:
		formattedTime = fmt.Sprintf("%02d %s %d %02d:%02d %s", date, monthShort, year, hour, t.Minute(), period)
	case 6:
		formattedTime = fmt.Sprintf("%02d %s %d", date, monthShort, year)
	case 7:
		formattedTime = fmt.Sprintf("%02d %s %02d %02d:%02d:%02d %s", date, monthShort, shortYear, hour, t.Minute(), t.Second(), period)
	case 8:
		formattedTime = fmt.Sprintf("%02d %s %02d %02d:%02d %s", date, monthShort, shortYear, hour, t.Minute(), period)
	case 9:
		formattedTime = fmt.Sprintf("%02d %s %02d", date, monthShort, shortYear)
	case 10:
		formattedTime = fmt.Sprintf("%02d/%02d/%d %02d:%02d:%02d %s", date, t.Month(), year, hour, t.Minute(), t.Second(), period)
	case 11:
		formattedTime = fmt.Sprintf("%02d/%02d/%d %02d:%02d %s", date, t.Month(), year, hour, t.Minute(), period)
	case 12:
		formattedTime = fmt.Sprintf("%02d/%02d/%d", date, t.Month(), year)
	case 13:
		formattedTime = fmt.Sprintf("%02d/%02d/%02d %02d:%02d:%02d %s", date, t.Month(), shortYear, hour, t.Minute(), t.Second(), period)
	case 14:
		formattedTime = fmt.Sprintf("%02d/%02d/%02d %02d:%02d %s", date, t.Month(), shortYear, hour, t.Minute(), period)
	case 15:
		formattedTime = fmt.Sprintf("%02d/%02d/%02d", date, t.Month(), shortYear)
	default:
		formattedTime = "Invalid format type"
	}

	return formattedTime
}
