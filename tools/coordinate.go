package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"ontario-fish-stocking-mcp/db"
	"database/sql"

	"github.com/mark3labs/mcp-go/mcp"
)

func QueryByCoordinate(dbConn *sql.DB, ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	coordinate, ok := request.Params.Arguments["coordinate"].(string)
	if !ok {
		return nil, fmt.Errorf("coordinate must be a string")
	}

	results, err := db.GetFishStockingRecordsByCoordinate(dbConn, coordinate)
	if err != nil {
		return nil, err
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
