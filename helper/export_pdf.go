package helpers

import (
	"context"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/google/uuid"
)

// export PDF report
func ExportPDF(data interface{}, reportName, HTMLTemplatePath string) (interface{}, error) {
	filePath := os.Getenv("GENARAL_PDF")
	templateName := os.Getenv(HTMLTemplatePath)

	result, err := InitDataToHtml(templateName, data, filePath)
	if err != nil {
		return nil, err
	}

	html, err := os.ReadFile(result)
	if err != nil {
		return nil, err
	}

	e := os.Remove(result)
	if e != nil {
		return nil, err
	}

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoSandbox,
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)

	defer cancel()

	var buf []byte

	if err := chromedp.Run(ctx, PrintToPDF(string(html), &buf)); err != nil {

		return nil, err
	}

	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {

		return nil, err
	}

	fullFilePath := filePath + reportName + ".pdf"
	if err := ioutil.WriteFile(fullFilePath, buf, 0777); err != nil {

		return nil, err
	}
	url := map[string]interface{}{
		"url": os.Getenv("STORAGE_IP") + fullFilePath,
	}

	return url, nil
}

func InitDataToHtml(templateName string, data interface{}, filePath string) (string, error) {

	templateGen, err := template.New(filepath.Base(templateName)).ParseFiles(templateName)
	if err != nil {
		return "", err
	}
	filePath = filePath + "temp/"
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		return "", err
	}

	fileName := filePath + uuid.New().String() + ".html"
	fileWritter, err := os.Create(fileName)
	if err != nil {
		return "", err
	}

	if err := templateGen.Execute(fileWritter, data); err != nil {
		return "", err
	}
	if err := fileWritter.Close(); err != nil {
		return "", err
	}
	return fileName, nil
}

func PrintToPDF(html string, res *[]byte) chromedp.Tasks {
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

			defer chromedp.Run(
				ctx,
				chromedp.WaitReady("template#success-pagejs"),
				// RunWithTimeOut(&ctx, 40, chromedp.Tasks{
				// 	chromedp.WaitVisible("div#success-pagejs"),
				// }),
			)

			wg.Wait()
			return nil
		}),

		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().
				WithPrintBackground(true).
				WithPreferCSSPageSize(true).
				WithDisplayHeaderFooter(true).
				Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}

func RunWithTimeOut(ctx *context.Context, timeout time.Duration, tasks chromedp.Tasks) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		timeoutContext, cancel := context.WithTimeout(ctx, timeout*time.Second)
		defer cancel()
		return tasks.Do(timeoutContext)
	}
}
