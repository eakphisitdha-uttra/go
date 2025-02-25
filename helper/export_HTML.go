package helpers

import (
	"html/template"
	"os"
	"path/filepath"
)

// export HTML report
func ExportHTML(data interface{}, reportName, HTMLTemplatePath string) (interface{}, error) {
	filePath := os.Getenv("GENARAL_HTML")
	templateName := os.Getenv(HTMLTemplatePath)

	templateGen, err := template.New(filepath.Base(templateName)).ParseFiles(templateName)
	if err != nil {

		return nil, err
	}

	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {

		return nil, err
	}

	fileName := filePath + reportName + ".html"
	fileWritter, err := os.Create(fileName)

	if err != nil {

		return nil, err
	}

	if err := templateGen.Execute(fileWritter, data); err != nil {

		return nil, err
	}

	if err := fileWritter.Close(); err != nil {

		return nil, err
	}

	url := map[string]interface{}{
		"url": os.Getenv("STORAGE_IP") + fileName,
	}

	return url, nil
}
