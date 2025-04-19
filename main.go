package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	_ "modernc.org/sqlite"
	"ontario-fish-stocking-mcp/db"
	"ontario-fish-stocking-mcp/tools"
)

func main() {
	// Initialize SQLite database
	dbConn, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	// Create table
	err = db.CreateTable(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	// Load data from CSV
	err = db.LoadDataFromCSV(dbConn, "Fish_Stocking_Data_for_Recreational_Purposes.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Create MCP server
	srv := server.NewMCPServer("ontario-fish-stocking", "1.0.0")

	// Add tools
	srv.AddTool(tools.NewCoordinateTool(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return tools.QueryByCoordinate(dbConn, ctx, request)
	})
	srv.AddTool(tools.NewSpeciesTool(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return tools.QueryBySpecies(dbConn, ctx, request)
	})
	srv.AddTool(tools.NewLocationNameTool(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return tools.QueryByLocationName(dbConn, ctx, request)
	})
	srv.AddTool(tools.NewYearTool(), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		return tools.QueryByYear(dbConn, ctx, request)
	})

	// Start MCP server
	fmt.Println("Starting MCP server...")

	// Start the server
	// Start the stdio server
	if err := server.ServeStdio(srv); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
