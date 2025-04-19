package db

import (
	"database/sql"
	"fmt"
)

func GetByCoordinate(db *sql.DB, latMin float64, latMax float64, lonMin float64, lonMax float64) ([]map[string]interface{}, error) {
	rows, err := db.Query(`
		SELECT species, location_name, year, latitude, longitude
		FROM fish_stocking
		WHERE latitude BETWEEN ? AND ?
		AND longitude BETWEEN ? AND ?
	`, latMin, latMax, lonMin, lonMax)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return parseRows(rows)
}

func GetBySpecies(db *sql.DB, species string) ([]map[string]interface{}, error) {
	rows, err := db.Query(`
		SELECT species, location_name, year, latitude, longitude
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
		SELECT species, location_name, year, latitude, longitude
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
		SELECT species, location_name, year, latitude, longitude
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
			species TEXT,
			location_name TEXT,
			year INTEGER,
			latitude REAL,
			longitude REAL
		)`)

	return err
}

func InsertData(db *sql.DB, species string, locationName string, year int, latitude float64, longitude float64) error {
	_, err := db.Exec(`
			INSERT INTO fish_stocking (species, location_name, year, latitude, longitude)
			VALUES (?, ?, ?, ?, ?)
		`, species, locationName, year, latitude, longitude)
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
		var s string
		var l string
		var y int
		var lat float64
		var lon float64

		err := rows.Scan(&s, &l, &y, &lat, &lon)
		if err != nil {
			return nil, err
		}

		results = append(results, map[string]interface{}{
			"species":      s,
			"locationName": l,
			"year":         y,
			"latitude":     lat,
			"longitude":    lon,
		})
	}
	return results, nil
}
