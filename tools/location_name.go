package tools

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func QueryByLocationName(db *sql.DB, ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		var c string
		var s string
		var l string
		var y int

		err = rows.Scan(&c, &s, &l, &y)
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

	jsonResults, err := json.Marshal(results)
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(string(jsonResults)), nil
}

func NewLocationNameTool() mcp.Tool {
	return mcp.NewTool("query_by_location_name",
		mcp.WithDescription("Query fish stocking records by location name"),
		mcp.WithString("location_name",
			mcp.Required(),
			mcp.Description("The location name to query"),
		),
	)
}
