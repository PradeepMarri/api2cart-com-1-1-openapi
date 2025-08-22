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

func ProductlistHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["page_cursor"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page_cursor=%v", val))
		}
		if val, ok := args["start"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("start=%v", val))
		}
		if val, ok := args["count"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("count=%v", val))
		}
		if val, ok := args["params"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("params=%v", val))
		}
		if val, ok := args["response_fields"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("response_fields=%v", val))
		}
		if val, ok := args["exclude"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("exclude=%v", val))
		}
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
		if val, ok := args["currency_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("currency_id=%v", val))
		}
		if val, ok := args["product_ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("product_ids=%v", val))
		}
		if val, ok := args["since_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("since_id=%v", val))
		}
		if val, ok := args["report_request_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("report_request_id=%v", val))
		}
		if val, ok := args["disable_report_cache"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("disable_report_cache=%v", val))
		}
		if val, ok := args["sort_by"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort_by=%v", val))
		}
		if val, ok := args["sort_direction"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort_direction=%v", val))
		}
		if val, ok := args["sku"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sku=%v", val))
		}
		if val, ok := args["disable_cache"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("disable_cache=%v", val))
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
		url := fmt.Sprintf("%s/product.list.json%s", cfg.BaseURL, queryString)
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
		var result models.ModelResponseProductList
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

func CreateProductlistTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_product_list_json",
		mcp.WithDescription("Get list of products from your store. Returns 10 products by default."),
		mcp.WithString("page_cursor", mcp.Description("Used to retrieve products via cursor-based pagination (it can't be used with any other filtering parameter)")),
		mcp.WithNumber("start", mcp.Description("This parameter sets the number from which you want to get entities")),
		mcp.WithNumber("count", mcp.Description("This parameter sets the entity amount that has to be retrieved. Max allowed count=250")),
		mcp.WithString("params", mcp.Description("Set this parameter in order to choose which entity fields you want to retrieve")),
		mcp.WithString("response_fields", mcp.Description("Set this parameter in order to choose which entity fields you want to retrieve")),
		mcp.WithString("exclude", mcp.Description("Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all")),
		mcp.WithString("category_id", mcp.Description("Retrieves products specified by category id")),
		mcp.WithString("created_from", mcp.Description("Retrieve entities from their creation date")),
		mcp.WithString("created_to", mcp.Description("Retrieve entities to their creation date")),
		mcp.WithString("modified_from", mcp.Description("Retrieve entities from their modification date")),
		mcp.WithString("modified_to", mcp.Description("Retrieve entities to their modification date")),
		mcp.WithBoolean("avail_view", mcp.Description("Specifies the set of visible/invisible products")),
		mcp.WithBoolean("avail_sale", mcp.Description("Specifies the set of available/not available products for sale")),
		mcp.WithString("store_id", mcp.Description("Retrieves products specified by store id")),
		mcp.WithString("lang_id", mcp.Description("Retrieves products specified by language id")),
		mcp.WithString("currency_id", mcp.Description("Currency Id")),
		mcp.WithString("product_ids", mcp.Description("Retrieves products specified by product ids")),
		mcp.WithNumber("since_id", mcp.Description("Retrieve entities starting from the specified id.")),
		mcp.WithString("report_request_id", mcp.Description("Report request id")),
		mcp.WithBoolean("disable_report_cache", mcp.Description("Disable report cache for current request")),
		mcp.WithString("sort_by", mcp.Description("Set field to sort by")),
		mcp.WithString("sort_direction", mcp.Description("Set sorting direction")),
		mcp.WithString("sku", mcp.Description("Filter by product's sku")),
		mcp.WithBoolean("disable_cache", mcp.Description("Disable cache for current request")),
		mcp.WithString("brand_name", mcp.Description("Retrieves brands specified by brand name")),
		mcp.WithArray("product_attributes", mcp.Description("Defines product attributes")),
		mcp.WithString("status", mcp.Description("Defines product's status")),
		mcp.WithString("type", mcp.Description("Defines products's type")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ProductlistHandler(cfg),
	}
}
