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

func ProductoptionaddHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("name=%v", val))
		}
		if val, ok := args["type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("type=%v", val))
		}
		if val, ok := args["product_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("product_id=%v", val))
		}
		if val, ok := args["default_option_value"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("default_option_value=%v", val))
		}
		if val, ok := args["option_values"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("option_values=%v", val))
		}
		if val, ok := args["description"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("description=%v", val))
		}
		if val, ok := args["avail"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("avail=%v", val))
		}
		if val, ok := args["sort_order"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort_order=%v", val))
		}
		if val, ok := args["required"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("required=%v", val))
		}
		if val, ok := args["clear_cache"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("clear_cache=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/product.option.add.json%s", cfg.BaseURL, queryString)
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

func CreateProductoptionaddTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_product_option_add_json",
		mcp.WithDescription("Add product option from store."),
		mcp.WithString("name", mcp.Required(), mcp.Description("Defines option's name")),
		mcp.WithString("type", mcp.Required(), mcp.Description("Defines option's type that has to be added")),
		mcp.WithString("product_id", mcp.Description("Defines product id where the option should be added")),
		mcp.WithString("default_option_value", mcp.Description("Defines default option value that has to be added")),
		mcp.WithString("option_values", mcp.Description("Defines option values that has to be added")),
		mcp.WithString("description", mcp.Description("Defines option's description")),
		mcp.WithBoolean("avail", mcp.Description("Defines whether the option is available")),
		mcp.WithNumber("sort_order", mcp.Description("Sort number in the list")),
		mcp.WithBoolean("required", mcp.Description("Defines if the option is required")),
		mcp.WithBoolean("clear_cache", mcp.Description("Is cache clear required")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ProductoptionaddHandler(cfg),
	}
}
