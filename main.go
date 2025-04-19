package main

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	_ "modernc.org/sqlite"
	"ontario-fish-stocking-mcp/db"
	"ontario-fish-stocking-mcp/tools"
)

//go:embed Fish_Stocking_Data_for_Recreational_Purposes.csv
var csvFile []byte

func main() {
	dbConn, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	err = db.CreateTable(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.LoadDataFromCSV(dbConn, csvFile)
	if err != nil {
		log.Fatal(err)
	}

	srv := server.NewMCPServer("ontario-fish-stocking", "1.0.0")

	srv.AddTool(tools.NewCoordinateTool(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return tools.QueryByCoordinate(dbConn, request)
	})
	srv.AddTool(tools.NewSpeciesTool(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return tools.QueryBySpecies(dbConn, request)
	})
	srv.AddTool(tools.NewLocationNameTool(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return tools.QueryByLocationName(dbConn, request)
	})
	srv.AddTool(tools.NewYearTool(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return tools.QueryByYear(dbConn, request)
	})

	fmt.Println("Starting MCP server...")

	if err := server.ServeStdio(srv); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
