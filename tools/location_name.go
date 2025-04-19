package tools

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"ontario-fish-stocking-mcp/db"

	"github.com/mark3labs/mcp-go/mcp"
)

func QueryByLocationName(dbConn *sql.DB, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	locationName, ok := request.Params.Arguments["location_name"].(string)
	if !ok {
		return nil, fmt.Errorf("location_name must be a string")
	}

	results, err := db.GetByLocationName(dbConn, locationName)
	if err != nil {
		return nil, err
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
