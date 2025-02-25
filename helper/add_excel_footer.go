package helpers

import (
	"fmt"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

// add footer in Excel report
func AddExcelFooter(f *excelize.File, sheetName string) {
	currentTime := time.Now().AddDate(543, 0, 0).Format("2 Jan 2006")
	month := time.Now().Format("Jan")
	monthThai := map[string]string{
		"Jan": "ม.ค.",
		"Feb": "ก.พ.",
		"Mar": "มี.ค.",
		"Apr": "เม.ย.",
		"May": "พ.ค.",
		"Jun": "มิ.ย.",
		"Jul": "ก.ค.",
		"Aug": "ส.ค.",
		"Sep": "ก.ย.",
		"Oct": "ต.ค.",
		"Nov": "พ.ย.",
		"Dec": "ธ.ค.",
	}
	currentTime = strings.Replace(currentTime, month, monthThai[month], 1)

	footerText := fmt.Sprintf("วันที่พิมพ์ %s ", currentTime)

	btm := 0.25
	f.SetPageMargins(sheetName, &excelize.PageLayoutMarginsOptions{
		Footer: &btm,
	})
	f.SetHeaderFooter(sheetName, &excelize.HeaderFooterOptions{
		//AlignWithMargins: true,
		OddFooter: `&R` + `&"TH SarabunPSK"` + footerText + "หน้า " + `&P` + "/" + `&N` + "\u00A0\u00A0\u00A0\u00A0\u00A0",
	})
}
