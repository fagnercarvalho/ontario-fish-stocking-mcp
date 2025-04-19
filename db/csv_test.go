package db

import (
	"database/sql"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func TestLoadDataFromCSV(t *testing.T) {
	file, err := os.ReadFile("../Fish_Stocking_Data_for_Recreational_Purposes.csv")
	require.NoError(t, err)

	lines := strings.Split(string(file), "\n")
	if len(lines) < 3 {
		t.Fatalf("CSV file must have at least 3 lines (header + 2 data rows)")
	}
	csvData := strings.Join(lines[:3], "\n")

	db, err := sql.Open("sqlite", ":memory:")
	require.NoError(t, err)
	defer db.Close()

	err = CreateTable(db)
	require.NoError(t, err)

	err = LoadDataFromCSV(db, []byte(csvData))
	require.NoError(t, err)

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM fish_stocking").Scan(&count)
	require.NoError(t, err)
	require.Equal(t, 2, count, "Expected 2 rows, got %d")

	var species, locationName, latitude, longitude string
	var year int

	err = db.QueryRow("SELECT species, location_name, year, latitude, longitude FROM fish_stocking WHERE location_name = 'Credit River'").Scan(&species, &locationName, &year, &latitude, &longitude)
	require.NoError(t, err)

	require.Equal(t, "Rainbow Trout", species, "Expected species 'Rainbow Trout', got %s")
	require.Equal(t, "Credit River", locationName, "Expected location_name 'Credit River', got %s")
	require.Equal(t, "43.54", latitude, "Expected latitude '43.54', got %s")
	require.Equal(t, "-79.58", longitude, "Expected longitude '-79.58', got %s")
	require.Equal(t, 2019, year, "Expected year 2019, got %d")
}
