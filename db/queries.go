package db

import (
	"database/sql"
	"fmt"
)

func GetByCoordinate(db *sql.DB, coordinate string) ([]map[string]interface{}, error) {
	rows, err := db.Query(`
		SELECT coordinate, species, location_name, year
		FROM fish_stocking
		WHERE coordinate = ?
	`, coordinate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseRows(rows)
}

func GetBySpecies(db *sql.DB, species string) ([]map[string]interface{}, error) {
	rows, err := db.Query(`
		SELECT coordinate, species, location_name, year
		FROM fish_stocking
		WHERE species = ?
	`, species)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseRows(rows)
}

func GetByLocationName(db *sql.DB, locationName string) ([]map[string]interface{}, error) {
	rows, err := db.Query(`
		SELECT coordinate, species, location_name, year
		FROM fish_stocking
		WHERE location_name = ?
	`, locationName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseRows(rows)
}

func GetByYear(db *sql.DB, year int) ([]map[string]interface{}, error) {
	rows, err := db.Query(`
		SELECT coordinate, species, location_name, year
		FROM fish_stocking
		WHERE year = ?
	`, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseRows(rows)
}

func CreateTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS fish_stocking (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			coordinate TEXT,
			species TEXT,
			location_name TEXT,
			year INTEGER
		)`)

	return err
}

func InsertData(db *sql.DB, coordinate string, species string, locationName string, year int) error {
	_, err := db.Exec(`
			INSERT INTO fish_stocking (coordinate, species, location_name, year)
			VALUES (?, ?, ?, ?)
		`, coordinate, species, locationName, year)
	if err != nil {
		return fmt.Errorf("error inserting record: %w", err)
	}
	return nil
}

func deleteData(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM fish_stocking")
	if err != nil {
		return fmt.Errorf("error deleting records: %w", err)
	}
	return nil
}

func parseRows(rows *sql.Rows) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	for rows.Next() {
		var c string
		var s string
		var l string
		var y int

		err := rows.Scan(&c, &s, &l, &y)
		if err != nil {
			return nil, err
		}

		results = append(results, map[string]interface{}{
			"coordinate":   c,
			"species":      s,
			"locationName": l,
			"year":         y,
		})
	}
	return results, nil
}
