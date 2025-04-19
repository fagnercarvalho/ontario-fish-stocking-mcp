package tools

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"ontario-fish-stocking-mcp/db"

	"github.com/mark3labs/mcp-go/mcp"
)

func QueryBySpecies(dbConn *sql.DB, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	species, ok := request.Params.Arguments["species"].(string)
	if !ok {
		return nil, fmt.Errorf("species must be a string")
	}

	res, err := db.GetBySpecies(dbConn, species)
	if err != nil {
		return nil, err
	}

	jsonResults, err := json.Marshal(res)
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
