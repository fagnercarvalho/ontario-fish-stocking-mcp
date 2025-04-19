package tools

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/mark3labs/mcp-go/mcp"
	"ontario-fish-stocking-mcp/db"
)

func QueryByCoordinate(dbConn *sql.DB, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	latMinStr, ok := request.Params.Arguments["latMin"].(string)
	if !ok {
		return nil, fmt.Errorf("latMin must be a string")
	}
	latMaxStr, ok := request.Params.Arguments["latMax"].(string)
	if !ok {
		return nil, fmt.Errorf("latMax must be a string")
	}
	lonMinStr, ok := request.Params.Arguments["lonMin"].(string)
	if !ok {
		return nil, fmt.Errorf("lonMin must be a string")
	}
	lonMaxStr, ok := request.Params.Arguments["lonMax"].(string)
	if !ok {
		return nil, fmt.Errorf("lonMax must be a string")
	}

	latMin, err := strconv.ParseFloat(latMinStr, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid latMin: %w", err)
	}
	latMax, err := strconv.ParseFloat(latMaxStr, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid latMax: %w", err)
	}
	lonMin, err := strconv.ParseFloat(lonMinStr, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid lonMin: %w", err)
	}
	lonMax, err := strconv.ParseFloat(lonMaxStr, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid lonMax: %w", err)
	}

	results, err := db.GetByCoordinate(dbConn, latMin, latMax, lonMin, lonMax)
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
		mcp.WithString("latMin",
			mcp.Required(),
			mcp.Description("The minimum latitude"),
		),
		mcp.WithString("latMax",
			mcp.Required(),
			mcp.Description("The maximum latitude"),
		),
		mcp.WithString("lonMin",
			mcp.Required(),
			mcp.Description("The minimum longitude"),
		),
		mcp.WithString("lonMax",
			mcp.Required(),
			mcp.Description("The maximum longitude"),
		),
	)
}
