package helpers

import (
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

// Add icon img to cell in Excel report (accept only .jpeg and .png)
func AddIconToCell(sheetName, filePath, column, row string, mergeCount int, f *excelize.File) error {

	newRow, _ := strconv.Atoi(row)

	// trim img path
	trimmedURL := strings.TrimPrefix(fmt.Sprint(filePath), os.Getenv("STORAGE_IP")+"/")

	// img size checker
	file, err := os.Open(trimmedURL)
	if err != nil {
		fmt.Println("can not open file", err)
		return err
	}
	defer file.Close()

	imageConfig, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Println("can not check image size", err)
		return err
	}

	columnNumber, err := strconv.Atoi(column)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	picPosition := fmt.Sprintf("%c", columnNumber-int((mergeCount-1)/2)) + fmt.Sprint(newRow+1)

	cellHight, _ := f.GetRowHeight("sheetName", newRow)

	if cellHight < 20 {
		cellHight = 20
	}

	setY := -1 * int((cellHight*1.3-19)/2+19.5) // 1.3 offset = 1 cell hight

	if mergeCount%2 == 0 {

		if imageConfig.Height == 19 && imageConfig.Width == 19 {
			f.AddPicture("sheetName", picPosition, trimmedURL, &excelize.GraphicOptions{OffsetX: -9, OffsetY: setY})

		} else {
			f.AddPicture("sheetName", picPosition, trimmedURL, &excelize.GraphicOptions{OffsetX: -9, OffsetY: setY, ScaleX: 19.00 / float64(imageConfig.Height), ScaleY: 19.00 / float64(imageConfig.Width)})
		}

	} else {
		cellWidth, _ := f.GetColWidth("sheetName", column)

		if imageConfig.Height == 19 && imageConfig.Width == 19 {
			f.AddPicture("sheetName", picPosition, trimmedURL, &excelize.GraphicOptions{OffsetX: int((cellWidth * 3) + 0.5), OffsetY: setY})

		} else {
			f.AddPicture("sheetName", picPosition, trimmedURL, &excelize.GraphicOptions{OffsetX: int((cellWidth * 3) + 0.5), OffsetY: setY, ScaleX: 19.00 / float64(imageConfig.Height), ScaleY: 19.00 / float64(imageConfig.Width)})

		}
	}
	return nil
}
