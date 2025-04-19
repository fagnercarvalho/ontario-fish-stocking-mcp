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

	// Insert data into the database
	for i, record := range records {
		if i == 0 {
			continue // Skip header row
		}

		year, err := strconv.Atoi(record[3])
		if err != nil {
			log.Printf("Error converting year: %v", err)
			continue
		}

		err = InsertData(db, record[0], record[1], record[2], year)
		if err != nil {
			log.Printf("Error inserting record: %v", err)
			continue
		}
	}

	return nil
}
