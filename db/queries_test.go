package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "modernc.org/sqlite"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatalf("failed to open test db: %v", err)
	}
	testDB = db

	CreateTable(testDB)

	os.Exit(m.Run())
}

func TestGetFishStockingRecordsByCoordinate(t *testing.T) {
	// Insert test data
	err := InsertData(testDB, "43.7001,-79.4163", "Rainbow Trout", "Test Location", 2023)
	if err != nil {
		t.Fatalf("failed to insert test data: %v", err)
	}

	// Call the function
	records, err := GetFishStockingRecordsByCoordinate(testDB, "43.7001,-79.4163")
	if err != nil {
		t.Fatalf("failed to get records: %v", err)
	}

	// Assert that the records are correct
	if len(records) != 1 {
		t.Fatalf("expected 1 record, got %d", len(records))
	}

	if records[0]["coordinate"] != "43.7001,-79.4163" {
		t.Errorf("expected coordinate to be '43.7001,-79.4163', got %s", records[0]["coordinate"])
	}
}

func TestGetFishStockingRecordsBySpecies(t *testing.T) {
	// Insert test data
	err := InsertData(testDB, "43.7001,-79.4163", "Rainbow Trout", "Test Location", 2023)
	if err != nil {
		t.Fatalf("failed to insert test data: %v", err)
	}

	// Call the function
	records, err := GetFishStockingRecordsBySpecies(testDB, "Rainbow Trout")
	if err != nil {
		t.Fatalf("failed to get records: %v", err)
	}

	// Assert that the records are correct
	if len(records) != 1 {
		t.Fatalf("expected 1 record, got %d", len(records))
	}

	if records[0]["species"] != "Rainbow Trout" {
		t.Errorf("expected species to be 'Rainbow Trout', got %s", records[0]["species"])
	}
}

func TestGetFishStockingRecordsByLocationName(t *testing.T) {
	// Insert test data
	err := InsertData(testDB, "43.7001,-79.4163", "Rainbow Trout", "Test Location", 2023)
	if err != nil {
		t.Fatalf("failed to insert test data: %v", err)
	}

	// Call the function
	records, err := GetFishStockingRecordsByLocationName(testDB, "Test Location")
	if err != nil {
		t.Fatalf("failed to get records: %v", err)
	}

	// Assert that the records are correct
	if len(records) != 1 {
		t.Fatalf("expected 1 record, got %d", len(records))
	}

	if records[0]["locationName"] != "Test Location" {
		t.Errorf("expected locationName to be 'Test Location', got %s", records[0]["locationName"])
	}
}

func TestGetFishStockingRecordsByYear(t *testing.T) {
	// Insert test data
	err := InsertData(testDB, "43.7001,-79.4163", "Rainbow Trout", "Test Location", 2023)
	if err != nil {
		t.Fatalf("failed to insert test data: %v", err)
	}

	// Call the function
	records, err := GetFishStockingRecordsByYear(testDB, 2023)
	if err != nil {
		t.Fatalf("failed to get records: %v", err)
	}

	// Assert that the records are correct
	if len(records) != 1 {
		t.Fatalf("expected 1 record, got %d", len(records))
	}

	if records[0]["year"] != 2023 {
		t.Errorf("expected year to be 2023, got %d", records[0]["year"])
	}
}
