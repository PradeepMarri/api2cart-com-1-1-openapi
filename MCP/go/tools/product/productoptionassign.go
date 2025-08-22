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

func ProductoptionassignHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["product_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("product_id=%v", val))
		}
		if val, ok := args["option_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("option_id=%v", val))
		}
		if val, ok := args["required"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("required=%v", val))
		}
		if val, ok := args["sort_order"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort_order=%v", val))
		}
		if val, ok := args["option_values"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("option_values=%v", val))
		}
		if val, ok := args["clear_cache"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("clear_cache=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/product.option.assign.json%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("POST", url, nil)
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

func CreateProductoptionassignTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_product_option_assign_json",
		mcp.WithDescription("Assign option from product."),
		mcp.WithString("product_id", mcp.Required(), mcp.Description("Defines product id where the option should be assigned")),
		mcp.WithString("option_id", mcp.Required(), mcp.Description("Defines option id which has to be assigned")),
		mcp.WithBoolean("required", mcp.Description("Defines if the option is required")),
		mcp.WithNumber("sort_order", mcp.Description("Sort number in the list")),
		mcp.WithString("option_values", mcp.Description("Defines option values that has to be assigned")),
		mcp.WithBoolean("clear_cache", mcp.Description("Is cache clear required")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ProductoptionassignHandler(cfg),
	}
}
