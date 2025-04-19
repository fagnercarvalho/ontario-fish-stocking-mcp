package db

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func LoadDataFromCSV(db *sql.DB, csvFilePath string) error {
	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		return fmt.Errorf("failed to open csv file: %w", err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read all records: %w", err)
	}

	for i, record := range records {
		if i == 0 {
			continue
		}

		year, err := strconv.Atoi(record[3])
		if err != nil {
			log.Printf("Error converting year: %v", err)
			continue
		}

		latitude := record[12]
		longitude := record[13]
		coordinate := latitude + "," + longitude

		err = InsertData(db, coordinate, record[4], record[5], year)
		if err != nil {
			log.Printf("Error inserting record: %v", err)
			continue
		}
	}

	return nil
}
