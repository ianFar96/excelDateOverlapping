package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/golang-module/carbon/v2"
	"github.com/webview/webview"
	"github.com/xuri/excelize/v2"
)

type rowRecap struct {
	row       []string
	startDate carbon.Carbon
	endDate   carbon.Carbon
}

type conflict struct {
	First  []string
	Second []string
}

type fileConflict struct {
	First  string
	Second []string
	Third  []conflict
}

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

		log.Println(htmlContent)

		if err != nil {
			log.Fatal(err)
			return
		}
		w.SetHtml(string(htmlContent))
	}

	// Commands
	w.Bind("scan", func(filePath string, dateCol string, startTimeCol string, endTimeCol string) ([]fileConflict, error) {
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

		var fileConflicts []fileConflict

		sheetList := f.GetSheetList()
		for _, sheetName := range sheetList {
			rows, err := f.GetRows(sheetName)
			if err != nil {
				message := fmt.Sprintf("Unexpected error, cannot read rows from file: %s", err)
				fmt.Println(message)
				return nil, errors.New(message)
			}

			// Delete first element as they are the titles
			header := rows[0:1][0]
			rows = rows[1:]

			// Get all the dates with carbon format from the rows
			var rowRecaps []rowRecap
			for _, row := range rows {
				date := row[dateColIndex-1]
				startTime := row[startTimeColIndex-1]
				endTime := row[endTimeColIndex-1]

				startDate, endDate, err := GetCarbonDates(date, startTime, endTime)
				if err != nil {
					fmt.Println(err)
					return nil, err
				}

				rowRecaps = append(rowRecaps, rowRecap{
					row:       row,
					startDate: *startDate,
					endDate:   *endDate,
				})
			}

			var conflicts []conflict
			for i := 0; i < len(rowRecaps); i++ {
				for j := i + 1; j < len(rowRecaps); j++ {

					isStartDateBetweenDates := rowRecaps[j].startDate.Between(rowRecaps[i].startDate, rowRecaps[i].endDate)
					isEndDateBetweenDates := rowRecaps[j].endDate.Between(rowRecaps[i].startDate, rowRecaps[i].endDate)
					if isStartDateBetweenDates || isEndDateBetweenDates {
						conflicts = append(conflicts, conflict{
							rowRecaps[i].row,
							rowRecaps[j].row,
						})
					}
				}
			}

			fileConflicts = append(fileConflicts, fileConflict{
				sheetName,
				header,
				conflicts,
			})
		}

		return fileConflicts, nil
	})

	w.Run()
}
