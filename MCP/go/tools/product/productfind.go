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

func ProductfindHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["find_value"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("find_value=%v", val))
		}
		if val, ok := args["find_where"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("find_where=%v", val))
		}
		if val, ok := args["find_params"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("find_params=%v", val))
		}
		if val, ok := args["find_what"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("find_what=%v", val))
		}
		if val, ok := args["lang_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("lang_id=%v", val))
		}
		if val, ok := args["store_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("store_id=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/product.find.json%s", cfg.BaseURL, queryString)
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

func CreateProductfindTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_product_find_json",
		mcp.WithDescription("Search product in store catalog. "Apple" is specified here by default."),
		mcp.WithString("find_value", mcp.Required(), mcp.Description("Entity search that is specified by some value")),
		mcp.WithString("find_where", mcp.Description("Entity search that is specified by the comma-separated unique fields")),
		mcp.WithString("find_params", mcp.Description("Entity search that is specified by comma-separated parameters")),
		mcp.WithString("find_what", mcp.Description("Parameter's value specifies the entity that has to be found")),
		mcp.WithString("lang_id", mcp.Description("Search products specified by language id")),
		mcp.WithString("store_id", mcp.Description("Store Id")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ProductfindHandler(cfg),
	}
}
