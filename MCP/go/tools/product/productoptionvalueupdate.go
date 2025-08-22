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

func ProductoptionvalueupdateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["option_value_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("option_value_id=%v", val))
		}
		if val, ok := args["option_value"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("option_value=%v", val))
		}
		if val, ok := args["price"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("price=%v", val))
		}
		if val, ok := args["quantity"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("quantity=%v", val))
		}
		if val, ok := args["clear_cache"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("clear_cache=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/product.option.value.update.json%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("PUT", url, nil)
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

func CreateProductoptionvalueupdateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_product_option_value_update_json",
		mcp.WithDescription("Update product option item from option."),
		mcp.WithString("product_id", mcp.Required(), mcp.Description("Defines product id where the option value should be updated")),
		mcp.WithString("option_id", mcp.Required(), mcp.Description("Defines option id where the value has to be updated")),
		mcp.WithNumber("option_value_id", mcp.Required(), mcp.Description("Defines value id that has to be assigned")),
		mcp.WithString("option_value", mcp.Required(), mcp.Description("Defines option value that has to be added")),
		mcp.WithString("price", mcp.Description("Defines new product option price")),
		mcp.WithString("quantity", mcp.Description("Defines new products' options quantity")),
		mcp.WithBoolean("clear_cache", mcp.Description("Is cache clear required")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ProductoptionvalueupdateHandler(cfg),
	}
}
