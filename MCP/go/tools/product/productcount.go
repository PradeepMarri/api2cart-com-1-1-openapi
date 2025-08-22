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

func ProductcountHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["category_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("category_id=%v", val))
		}
		if val, ok := args["created_from"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("created_from=%v", val))
		}
		if val, ok := args["created_to"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("created_to=%v", val))
		}
		if val, ok := args["modified_from"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("modified_from=%v", val))
		}
		if val, ok := args["modified_to"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("modified_to=%v", val))
		}
		if val, ok := args["avail_view"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("avail_view=%v", val))
		}
		if val, ok := args["avail_sale"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("avail_sale=%v", val))
		}
		if val, ok := args["store_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("store_id=%v", val))
		}
		if val, ok := args["lang_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("lang_id=%v", val))
		}
		if val, ok := args["product_ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("product_ids=%v", val))
		}
		if val, ok := args["report_request_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("report_request_id=%v", val))
		}
		if val, ok := args["disable_report_cache"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("disable_report_cache=%v", val))
		}
		if val, ok := args["brand_name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("brand_name=%v", val))
		}
		if val, ok := args["product_attributes"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("product_attributes=%v", val))
		}
		if val, ok := args["status"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("status=%v", val))
		}
		if val, ok := args["type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("type=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/product.count.json%s", cfg.BaseURL, queryString)
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

func CreateProductcountTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_product_count_json",
		mcp.WithDescription("Count products in store."),
		mcp.WithString("category_id", mcp.Description("Counts products specified by category id")),
		mcp.WithString("created_from", mcp.Description("Retrieve entities from their creation date")),
		mcp.WithString("created_to", mcp.Description("Retrieve entities to their creation date")),
		mcp.WithString("modified_from", mcp.Description("Retrieve entities from their modification date")),
		mcp.WithString("modified_to", mcp.Description("Retrieve entities to their modification date")),
		mcp.WithBoolean("avail_view", mcp.Description("Specifies the set of visible/invisible products")),
		mcp.WithBoolean("avail_sale", mcp.Description("Specifies the set of available/not available products for sale")),
		mcp.WithString("store_id", mcp.Description("Counts products specified by store id")),
		mcp.WithString("lang_id", mcp.Description("Counts products specified by language id")),
		mcp.WithString("product_ids", mcp.Description("Counts products specified by product ids")),
		mcp.WithString("report_request_id", mcp.Description("Report request id")),
		mcp.WithBoolean("disable_report_cache", mcp.Description("Disable report cache for current request")),
		mcp.WithString("brand_name", mcp.Description("Retrieves brands specified by brand name")),
		mcp.WithArray("product_attributes", mcp.Description("Defines product attributes")),
		mcp.WithString("status", mcp.Description("Defines product's status")),
		mcp.WithString("type", mcp.Description("Defines products's type")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ProductcountHandler(cfg),
	}
}
