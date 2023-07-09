package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/webview/webview"
	"github.com/xuri/excelize/v2"
)

func main() {
	isDebug := os.Getenv("ENVIRONMENT") == "DEBUG"

	w := webview.New(isDebug)
	defer w.Destroy()

	w.SetTitle("Excel Date Overlapping")
	w.SetSize(480, 320, webview.HintNone)

	if isDebug {
		w.Navigate("http://localhost:5173")
	} else {
		fileSystem := os.DirFS("../frontend/dist")
		htmlContent, err := fs.ReadFile(fileSystem, "index.html")
		if err != nil {
			log.Fatal(err)
			return
		}
		w.SetHtml(string(htmlContent))
	}

	// Commands
	w.Bind("scan", func(filePath string, dateCol string, startTimeCol string, endTimeCol string) ([]string, error) {
		f, err := excelize.OpenFile("your-file.xlsx")
		if err != nil {
			message := fmt.Sprintf("Error opening the file: %s", err)
			fmt.Println(message)
			return nil, errors.New(message)
		}

		sheetList := f.GetSheetList()
		for _, sheetName := range sheetList {
			fmt.Println(sheetName)
		}

		return sheetList, nil
	})

	w.Run()
}
