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

func ProductcurrencyaddHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["iso3"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("iso3=%v", val))
		}
		if val, ok := args["rate"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("rate=%v", val))
		}
		if val, ok := args["name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("name=%v", val))
		}
		if val, ok := args["avail"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("avail=%v", val))
		}
		if val, ok := args["symbol_left"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("symbol_left=%v", val))
		}
		if val, ok := args["symbol_right"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("symbol_right=%v", val))
		}
		if val, ok := args["default"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("default=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/product.currency.add.json%s", cfg.BaseURL, queryString)
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

func CreateProductcurrencyaddTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_product_currency_add_json",
		mcp.WithDescription("Add currency and/or set default in store"),
		mcp.WithString("iso3", mcp.Required(), mcp.Description("Specifies standardized currency code")),
		mcp.WithString("rate", mcp.Required(), mcp.Description("Defines the numerical identifier against to the major currency")),
		mcp.WithString("name", mcp.Description("Defines currency's name")),
		mcp.WithBoolean("avail", mcp.Description("Specifies whether the currency is available")),
		mcp.WithString("symbol_left", mcp.Description("Defines the symbol that is located before the currency")),
		mcp.WithString("symbol_right", mcp.Description("Defines the symbol that is located after the currency")),
		mcp.WithBoolean("default", mcp.Description("Specifies currency's default meaning")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ProductcurrencyaddHandler(cfg),
	}
}
