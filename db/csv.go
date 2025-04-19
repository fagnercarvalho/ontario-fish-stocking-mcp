package db

import (
	"bytes"
	"database/sql"
	_ "embed"
	"encoding/csv"
	"fmt"
	"log"
	"strconv"
)

func LoadDataFromCSV(db *sql.DB, csvFile []byte) error {
	reader := csv.NewReader(bytes.NewReader(csvFile))
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

		species := record[4]
		location := record[5]
		latitude, err := strconv.ParseFloat(record[12], 2)
		if err != nil {
			log.Printf("Error converting latitude: %v", err)
			continue
		}

		longitude, err := strconv.ParseFloat(record[13], 2)
		if err != nil {
			log.Printf("Error converting longitude: %v", err)
			continue
		}

		err = InsertData(db, species, location, year, latitude, longitude)
		if err != nil {
			log.Printf("Error inserting record: %v", err)
			continue
		}
	}

	return nil
}
