package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatalf("failed to open test db: %v", err)
	}
	testDB = db

	err = CreateTable(testDB)
	if err != nil {
		log.Fatalf("failed to create test db: %v", err)
	}

	os.Exit(m.Run())
}

func TestGetFishStockingRecordsByCoordinate(t *testing.T) {
	// Insert test data
	err := InsertData(testDB, "43.7001,-79.4163", "Rainbow Trout", "Test Location", 2023)
	require.NoError(t, err)

	// Call the function
	records, err := GetFishStockingRecordsByCoordinate(testDB, "43.7001,-79.4163")
	if err != nil {
		t.Fatalf("failed to get records: %v", err)
	}

	// Assert that the records are correct
	assert.Len(t, records, 1, "expected 1 record")
	assert.Equal(t, "43.7001,-79.4163", records[0]["coordinate"], "expected coordinate to be '43.7001,-79.4163'")
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
	assert.Len(t, records, 1, "expected 1 record")
	assert.Equal(t, "Rainbow Trout", records[0]["species"], "expected species to be 'Rainbow Trout'")
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
	assert.Len(t, records, 1, "expected 1 record")
	assert.Equal(t, "Test Location", records[0]["locationName"], "expected locationName to be 'Test Location'")
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
	assert.Len(t, records, 1, "expected 1 record")
	assert.Equal(t, 2023, records[0]["year"], "expected year to be 2023")
}
