package main

import (
	"bufio"
	"encoding/csv"
	"os"
)

type TimeZone struct {
	csv [][]string
}

func (zone *TimeZone) GetTimeList() []string {
	var zoneList []string
	for _, row := range zone.csv {
		zoneList = append(zoneList, row[2])
	}
	return zoneList
}

func readTimezoneFromFS() (timezone TimeZone, err error) {
	file, err := os.Open("zones.csv")

	csvFile, err := csv.NewReader(bufio.NewReader(file)).ReadAll()

	timezone = TimeZone{
		csv: csvFile,
	}
	return
}
