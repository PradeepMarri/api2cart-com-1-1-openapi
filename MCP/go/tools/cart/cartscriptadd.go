package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/swagger-api2cart/mcp-server/config"
	"github.com/swagger-api2cart/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CartscriptaddHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("name=%v", val))
		}
		if val, ok := args["description"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("description=%v", val))
		}
		if val, ok := args["html"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("html=%v", val))
		}
		if val, ok := args["src"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("src=%v", val))
		}
		if val, ok := args["load_method"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("load_method=%v", val))
		}
		if val, ok := args["scope"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("scope=%v", val))
		}
		if val, ok := args["store_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("store_id=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/cart.script.add.json%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("POST", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("x-store-key", cfg.APIKey)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateCartscriptaddTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_cart_script_add_json",
		mcp.WithDescription("Add new script to the storefront"),
		mcp.WithString("name", mcp.Description("The user-friendly script name")),
		mcp.WithString("description", mcp.Description("The user-friendly description")),
		mcp.WithString("html", mcp.Description("An html string containing exactly one `script` tag.")),
		mcp.WithString("src", mcp.Description("The URL of the remote script")),
		mcp.WithString("load_method", mcp.Description("The load method to use for the script")),
		mcp.WithString("scope", mcp.Description("The page or pages on the online store where the script should be included")),
		mcp.WithString("store_id", mcp.Description("Store Id")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CartscriptaddHandler(cfg),
	}
}
