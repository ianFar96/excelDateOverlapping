package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/golang-module/carbon/v2"
)

func GetCarbonDates(date string, startTime string, endTime string) (*carbon.Carbon, *carbon.Carbon, error) {
	splittedDate := strings.Split(date, "/")

	year, err := strconv.ParseInt(splittedDate[2], 10, 64)
	if err != nil {
		message := fmt.Sprintf("Wrong date format, expected dd/mm/yyyy given %s: %s", date, err)
		fmt.Println(message)
		return nil, nil, errors.New(message)
	}

	month, err := strconv.ParseInt(splittedDate[1], 10, 64)
	if err != nil {
		message := fmt.Sprintf("Wrong date format, expected dd/mm/yyyy given %s: %s", date, err)
		fmt.Println(message)
		return nil, nil, errors.New(message)
	}

	day, err := strconv.ParseInt(splittedDate[0], 10, 64)
	if err != nil {
		message := fmt.Sprintf("Wrong date format, expected dd/mm/yyyy given %s: %s", date, err)
		fmt.Println(message)
		return nil, nil, errors.New(message)
	}

	splittedStartTime := strings.Split(startTime, ":")

	hour, err := strconv.ParseInt(splittedStartTime[0], 10, 64)
	if err != nil {
		message := fmt.Sprintf("Wrong start time format, expected hh:mm given %s: %s", startTime, err)
		fmt.Println(message)
		return nil, nil, errors.New(message)
	}

	minutes, err := strconv.ParseInt(splittedStartTime[0], 10, 64)
	if err != nil {
		message := fmt.Sprintf("Wrong start time format, expected hh:mm given %s: %s", startTime, err)
		fmt.Println(message)
		return nil, nil, errors.New(message)
	}

	startDate := carbon.CreateFromDateTime(int(year), int(month), int(day), int(hour), int(minutes), 0)

	splittedEndTime := strings.Split(endTime, ":")

	hour, err = strconv.ParseInt(splittedEndTime[0], 10, 64)
	if err != nil {
		message := fmt.Sprintf("Wrong end time format, expected hh:mm given %s: %s", endTime, err)
		fmt.Println(message)
		return nil, nil, errors.New(message)
	}

	minutes, err = strconv.ParseInt(splittedEndTime[0], 10, 64)
	if err != nil {
		message := fmt.Sprintf("Wrong end time format, expected hh:mm given %s: %s", endTime, err)
		fmt.Println(message)
		return nil, nil, errors.New(message)
	}

	endDate := carbon.CreateFromDateTime(int(year), int(month), int(day), int(hour), int(minutes), 0)

	return &startDate, &endDate, nil
}
