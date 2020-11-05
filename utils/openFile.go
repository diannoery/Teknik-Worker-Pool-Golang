package utils

import (
	"encoding/csv"
	"log"
	"os"
)

const csvFile = "majestic_million.csv"

func OpenCsvFile() (*csv.Reader, *os.File, error) {
	log.Println("=> open csv file")

	f, err := os.Open(csvFile)
	if err != nil {
		return nil, nil, err
	}

	reader := csv.NewReader(f)
	return reader, f, nil
}
