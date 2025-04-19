package tools

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func QueryByYear(db *sql.DB, ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
}

func NewYearTool() mcp.Tool {
	return mcp.NewTool("query_by_year",
		mcp.WithDescription("Query fish stocking records by year"),
		mcp.WithNumber("year",
			mcp.Required(),
			mcp.Description("The year to query"),
		),
	)
}
