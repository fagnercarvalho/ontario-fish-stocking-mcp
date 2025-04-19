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
