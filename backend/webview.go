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
		f, err := excelize.OpenFile(filePath)
		if err != nil {
			message := fmt.Sprintf("Error opening the file: %s", err)
			fmt.Println(message)
			return nil, errors.New(message)
		}

		dateColIndex, err := excelize.ColumnNameToNumber(dateCol)
		if err != nil {
			message := fmt.Sprintf("Error getting the column for the date: %s", err)
			fmt.Println(message)
			return nil, errors.New(message)
		}

		startTimeColIndex, err := excelize.ColumnNameToNumber(startTimeCol)
		if err != nil {
			message := fmt.Sprintf("Error getting the column for the start time: %s", err)
			fmt.Println(message)
			return nil, errors.New(message)
		}

		endTimeColIndex, err := excelize.ColumnNameToNumber(endTimeCol)
		if err != nil {
			message := fmt.Sprintf("Error getting the column for the end time: %s", err)
			fmt.Println(message)
			return nil, errors.New(message)
		}

		sheetList := f.GetSheetList()
		for _, sheetName := range sheetList {
			rows, err := f.GetRows(sheetName)
			if err != nil {
				message := fmt.Sprintf("Unexpected error, cannot read rows from file: %s", err)
				fmt.Println(message)
				return nil, errors.New(message)
			}

			// Delete first element as they are the titles
			rows = rows[1:]

			for _, row := range rows {
				date := row[dateColIndex-1]
				startTime := row[startTimeColIndex-1]
				endTime := row[endTimeColIndex-1]

				fmt.Println("date", date)
				fmt.Println("startTime", startTime)
				fmt.Println("endTime", endTime)
			}
		}

		return sheetList, nil
	})

	w.Run()
}
