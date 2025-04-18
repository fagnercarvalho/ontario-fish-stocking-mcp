package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	_ "modernc.org/sqlite"
)

// FishStockingRecord represents a fish stocking record.
type FishStockingRecord struct {
	ID           int
	Coordinate   string
	Species      string
	LocationName string
	Year         int
}

func main() {
	// Initialize SQLite database
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS fish_stocking (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			coordinate TEXT,
			species TEXT,
			location_name TEXT,
			year INTEGER
		)`)
	if err != nil {
		log.Fatal(err)
	}

	// Load data from CSV
	csvFile, err := os.Open("Fish_Stocking_Data_for_Recreational_Purposes.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
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

		_, err = db.Exec(`
			INSERT INTO fish_stocking (coordinate, species, location_name, year)
			VALUES (?, ?, ?, ?)
		`, record[0], record[1], record[2], year)
		if err != nil {
			log.Printf("Error inserting record: %v", err)
			continue
		}
	}

	// Create MCP server
	srv := server.NewMCPServer("ontario-fish-stocking", "1.0.0")

	// Add tool to query by coordinate
	tool := mcp.NewTool("query_by_coordinate",
		mcp.WithDescription("Query fish stocking records by coordinate"),
		mcp.WithString("coordinate",
			mcp.Required(),
			mcp.Description("The coordinate to query"),
		),
	)
	srv.AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		coordinate, ok := request.Params.Arguments["coordinate"].(string)
		if !ok {
			return nil, fmt.Errorf("coordinate must be a string")
		}

		rows, err := db.Query(`
			SELECT coordinate, species, location_name, year
			FROM fish_stocking
			WHERE coordinate = ?
		`, coordinate)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var results []map[string]interface{}
		for rows.Next() {
			var coord string
			var spec string
			var locName string
			var yr int

			err = rows.Scan(&coord, &spec, &locName, &yr)
			if err != nil {
				return nil, err
			}

			results = append(results, map[string]interface{}{
				"coordinate":   coord,
				"species":      spec,
				"locationName": locName,
				"year":         yr,
			})
		}

		jsonResults, err := json.Marshal(results)
		if err != nil {
			return nil, err
		}

		return mcp.NewToolResultText(string(jsonResults)), nil
	})

	// Add tool to query by coordinate
	tool = mcp.NewTool("query_by_coordinate",
		mcp.WithDescription("Query fish stocking records by coordinate"),
		mcp.WithString("coordinate",
			mcp.Required(),
			mcp.Description("The coordinate to query"),
		),
	)
	srv.AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		coordinate, ok := request.Params.Arguments["coordinate"].(string)
		if !ok {
			return nil, fmt.Errorf("coordinate must be a string")
		}

		rows, err := db.Query(`
			SELECT coordinate, species, location_name, year
			FROM fish_stocking
			WHERE coordinate = ?
		`, coordinate)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var results []map[string]interface{}
		for rows.Next() {
			var coordinate string
			var species string
			var locationName string
			var year int

			err = rows.Scan(&coordinate, &species, &locationName, &year)
			if err != nil {
				return nil, err
			}

			results = append(results, map[string]interface{}{
				"coordinate":   coordinate,
				"species":      species,
				"locationName": locationName,
				"year":         year,
			})
		}

		jsonResults, err := json.Marshal(results)
		if err != nil {
			return nil, err
		}

		return mcp.NewToolResultText(string(jsonResults)), nil
	})

	// Add tool to query by species
	tool = mcp.NewTool("query_by_species",
		mcp.WithDescription("Query fish stocking records by species"),
		mcp.WithString("species",
			mcp.Required(),
			mcp.Description("The species to query"),
		),
	)
	srv.AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		species, ok := request.Params.Arguments["species"].(string)
		if !ok {
			return nil, fmt.Errorf("species must be a string")
		}

		rows, err := db.Query(`
			SELECT coordinate, species, location_name, year
			FROM fish_stocking
			WHERE species = ?
		`, species)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var results []map[string]interface{}
		for rows.Next() {
			var coordinate string
			var species string
			var locationName string
			var year int

			err = rows.Scan(&coordinate, &species, &locationName, &year)
			if err != nil {
				return nil, err
			}

			results = append(results, map[string]interface{}{
				"coordinate":   coordinate,
				"species":      species,
				"locationName": locationName,
				"year":         year,
			})
		}

		jsonResults, err := json.Marshal(results)
		if err != nil {
			return nil, err
		}

		return mcp.NewToolResultText(string(jsonResults)), nil
	})

	// Add tool to query by location name
	tool = mcp.NewTool("query_by_location_name",
		mcp.WithDescription("Query fish stocking records by location name"),
		mcp.WithString("location_name",
			mcp.Required(),
			mcp.Description("The location name to query"),
		),
	)
	srv.AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		locationName, ok := request.Params.Arguments["location_name"].(string)
		if !ok {
			return nil, fmt.Errorf("location_name must be a string")
		}

		rows, err := db.Query(`
			SELECT coordinate, species, location_name, year
			FROM fish_stocking
			WHERE location_name = ?
		`, locationName)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var results []map[string]interface{}
		for rows.Next() {
			var coordinate string
			var species string
			var locationName string
			var year int

			err = rows.Scan(&coordinate, &species, &locationName, &year)
			if err != nil {
				return nil, err
			}

			results = append(results, map[string]interface{}{
				"coordinate":   coordinate,
				"species":      species,
				"locationName": locationName,
				"year":         year,
			})
		}

		jsonResults, err := json.Marshal(results)
		if err != nil {
			return nil, err
		}

		return mcp.NewToolResultText(string(jsonResults)), nil
	})

	// Add tool to query by year
	tool = mcp.NewTool("query_by_year",
		mcp.WithDescription("Query fish stocking records by year"),
		mcp.WithNumber("year",
			mcp.Required(),
			mcp.Description("The year to query"),
		),
	)
	srv.AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		yearFloat, ok := request.Params.Arguments["year"].(float64)
		if !ok {
			return nil, fmt.Errorf("invalid year format")
		}
		year := int(yearFloat)

		rows, err := db.Query(`
			SELECT coordinate, species, location_name, year
			FROM fish_stocking
			WHERE year = ?
		`, year)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var results []map[string]interface{}
		for rows.Next() {
			var coordinate string
			var species string
			var locationName string
			var year int

			err = rows.Scan(&coordinate, &species, &locationName, &year)
			if err != nil {
				return nil, err
			}

			results = append(results, map[string]interface{}{
				"coordinate":   coordinate,
				"species":      species,
				"locationName": locationName,
				"year":         year,
			})
		}

		jsonResults, err := json.Marshal(results)
		if err != nil {
			return nil, err
		}

		return mcp.NewToolResultText(string(jsonResults)), nil
	})

	// Start MCP server
	fmt.Println("Starting MCP server...")

	// Start the server
	// Start the stdio server
	if err := server.ServeStdio(srv); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
