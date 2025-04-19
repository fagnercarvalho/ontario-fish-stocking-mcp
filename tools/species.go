package tools

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func QueryBySpecies(db *sql.DB, ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
}

func NewSpeciesTool() mcp.Tool {
	return mcp.NewTool("query_by_species",
		mcp.WithDescription("Query fish stocking records by species"),
		mcp.WithString("species",
			mcp.Required(),
			mcp.Description("The species to query"),
		),
	)
}
