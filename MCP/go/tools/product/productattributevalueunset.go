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

func ProductattributevalueunsetHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["product_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("product_id=%v", val))
		}
		if val, ok := args["attribute_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("attribute_id=%v", val))
		}
		if val, ok := args["store_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("store_id=%v", val))
		}
		if val, ok := args["include_default"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("include_default=%v", val))
		}
		if val, ok := args["reindex"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("reindex=%v", val))
		}
		if val, ok := args["clear_cache"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("clear_cache=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/product.attribute.value.unset.json%s", cfg.BaseURL, queryString)
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

func CreateProductattributevalueunsetTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_product_attribute_value_unset_json",
		mcp.WithDescription("Removes attribute value for a product."),
		mcp.WithString("product_id", mcp.Required(), mcp.Description("Product id")),
		mcp.WithString("attribute_id", mcp.Required(), mcp.Description("Attribute Id")),
		mcp.WithString("store_id", mcp.Description("Store Id")),
		mcp.WithBoolean("include_default", mcp.Description("Boolean, whether or not to unset default value of the attribute, if applicable")),
		mcp.WithBoolean("reindex", mcp.Description("Is reindex required")),
		mcp.WithBoolean("clear_cache", mcp.Description("Is cache clear required")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ProductattributevalueunsetHandler(cfg),
	}
}
