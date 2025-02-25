package helpers

import (
	"context"
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

type ExportPDFStruct struct {
	FullPath string
	Width    float64
	Height   float64
}

// export PNG img from HTML file
func ExportPNG(data interface{}, HTMLTemplatePath string) (ExportPDFStruct, error) {
	filePath := os.Getenv("GENARAL_EXCEL")
	templateName := os.Getenv(HTMLTemplatePath)
	result, err := InitDataToHtml(templateName, data, filePath)
	if err != nil {
		return ExportPDFStruct{}, err
	}

	html, err := os.ReadFile(result)
	if err != nil {
		return ExportPDFStruct{}, err
	}

	e := os.Remove(result)
	if e != nil {
		return ExportPDFStruct{}, err
	}

	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)

	defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx, PrintToPNG(string(html), &buf, false)); err != nil {

		return ExportPDFStruct{}, err
	}

	unix := strconv.Itoa(int(time.Now().Unix()))

	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {

		return ExportPDFStruct{}, err
	}

	fullFilePath := filePath + unix + ".png"
	if err := ioutil.WriteFile(fullFilePath, buf, 0777); err != nil {

		return ExportPDFStruct{}, err
	}

	// img size checker
	file, err := os.Open(fmt.Sprint(fullFilePath))
	if err != nil {
		fmt.Println("can not open file", err)
	}
	defer file.Close()

	imageConfig, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Println("can not check image size", err)
	}

	resp := ExportPDFStruct{
		FullPath: fullFilePath,
		Width:    float64(imageConfig.Width),
		Height:   float64(imageConfig.Height),
	}

	return resp, nil
}

func PrintToPNG(html string, res *[]byte, isDelay bool) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate("about:blank"),
		chromedp.ActionFunc(func(ctx context.Context) error {

			lctx, cancel := context.WithCancel(ctx)
			defer cancel()

			var wg sync.WaitGroup
			wg.Add(1)

			chromedp.ListenTarget(lctx, func(ev interface{}) {
				if _, ok := ev.(*page.EventLoadEventFired); ok {
					cancel()
					wg.Done()
				}
			})
			frameTree, err := page.GetFrameTree().Do(ctx)
			if err != nil {
				return err
			}

			if err := page.SetDocumentContent(frameTree.Frame.ID, html).Do(ctx); err != nil {
				return err
			}
			delay := 5
			if isDelay {
				delay = 20
			}

			defer chromedp.Run(
				ctx,
				RunWithTimeOut(&ctx, time.Duration(delay), chromedp.Tasks{
					chromedp.WaitVisible("div#success"),
				}),
			)

			wg.Wait()
			return nil
		}),

		chromedp.ActionFunc(func(ctx context.Context) error {
			width, height := 455, 340
			if err := emulation.SetDeviceMetricsOverride(int64(width), int64(height), 1.0, false).Do(ctx); err != nil {
				return err
			}

			buf, err := page.CaptureScreenshot().Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}
