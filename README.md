```
                                 ..:=+++++.                                
                            ..-=+#########-                                
                  ...:-==+**##############*...                             
              ..:-+##########################*=-:...         ..:-:..       
           ..:+####################################+:.......:+###-.        
         ..-#############*#############################+::+######:.        
         :*######*:.=#####*#####################################*:.        
         .=###############+*############################==######*:.        
          .:*############**########################*-:.. ..-+####-.        
           ..:+#########**####################+--:..        ..-=*=..       
              ..:=+*###################*===:..                  ...        
                   ...:-=++++*####+...                                     
                             .:+**..                                       
                                                                           
                                                                         
```

# Ontario Fish Stocking Data

This project is an MCP (Model Context Protocol) server that fetches fish stocking data from Ontario, Canada.

## Capabilities

This MCP server provides the following tools:

*   **Coordinates Tool:** Queries fish stocking data by coordinate.
*   **Species Tool:** Queries fish stocking data by species.
*   **Location Name Tool:** Queries fish stocking data by location name.
*   **Year Tool:** Queries fish stocking data by year.

## How to Run

1.  Clone the repository.
2.  Run `go run main.go`.
3.  Or download the executable directly from the [Releases](https://github.com/fagnercarvalho/ontario-fish-stocking-mcp/releases) page. 

## Setup

In your tool of choice configure the MCP server like this:

```json
{
  "mcpServers": {
    "fish-stocking": {
      "command": "<path>/ontario-fish-stocking-mcp.exe"
    }
  }
}
```

## Example prompts

- Summary of fish stocking in 2024
- Compare 2024 with 2023 stocking
- Which species were stocked on Humber River in 2024?

## License

This project uses data from the Ontario government licensed under the [Open Government Licence - Ontario](https://www.ontario.ca/page/open-government-licence-ontario).