package tools

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"ontario-fish-stocking-mcp/db"

	"github.com/mark3labs/mcp-go/mcp"
)

func QueryByYear(dbConn *sql.DB, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	yearFloat, ok := request.Params.Arguments["year"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid year format")
	}
	year := int(yearFloat)

	results, err := db.GetByYear(dbConn, year)
	if err != nil {
		return nil, err
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
