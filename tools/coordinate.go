package tools

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func QueryByCoordinate(db *sql.DB, ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

func NewCoordinateTool() mcp.Tool {
	return mcp.NewTool("query_by_coordinate",
		mcp.WithDescription("Query fish stocking records by coordinate"),
		mcp.WithString("coordinate",
			mcp.Required(),
			mcp.Description("The coordinate to query"),
		),
	)
}
