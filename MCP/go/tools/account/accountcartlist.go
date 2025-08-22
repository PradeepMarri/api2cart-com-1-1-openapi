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

func AccountcartlistHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["params"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("params=%v", val))
		}
		if val, ok := args["exclude"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("exclude=%v", val))
		}
		if val, ok := args["request_from_date"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("request_from_date=%v", val))
		}
		if val, ok := args["request_to_date"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("request_to_date=%v", val))
		}
		if val, ok := args["store_url"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("store_url=%v", val))
		}
		if val, ok := args["store_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("store_key=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/account.cart.list.json%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("x-api-key", cfg.APIKey)
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

func CreateAccountcartlistTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_account_cart_list_json",
		mcp.WithDescription("Get list of carts."),
		mcp.WithString("params", mcp.Description("Set this parameter in order to choose which entity fields you want to retrieve")),
		mcp.WithString("exclude", mcp.Description("Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all")),
		mcp.WithString("request_from_date", mcp.Description("Retrieve entities from their creation date")),
		mcp.WithString("request_to_date", mcp.Description("Retrieve entities to their creation date")),
		mcp.WithString("store_url", mcp.Description("A web address of a store")),
		mcp.WithString("store_key", mcp.Description("Find store by store key")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    AccountcartlistHandler(cfg),
	}
}
