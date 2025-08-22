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

func ProductvariantlistHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["start"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("start=%v", val))
		}
		if val, ok := args["count"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("count=%v", val))
		}
		if val, ok := args["params"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("params=%v", val))
		}
		if val, ok := args["exclude"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("exclude=%v", val))
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
		if val, ok := args["category_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("category_id=%v", val))
		}
		if val, ok := args["product_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("product_id=%v", val))
		}
		if val, ok := args["store_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("store_id=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/product.variant.list.json%s", cfg.BaseURL, queryString)
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

func CreateProductvariantlistTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_product_variant_list_json",
		mcp.WithDescription("Get a list of variants. This method is deprecated, and its development is stopped. Please use "product.child_item.list" instead."),
		mcp.WithNumber("start", mcp.Description("This parameter sets the number from which you want to get entities")),
		mcp.WithNumber("count", mcp.Description("This parameter sets the entity amount that has to be retrieved. Max allowed count=250")),
		mcp.WithString("params", mcp.Description("Set this parameter in order to choose which entity fields you want to retrieve")),
		mcp.WithString("exclude", mcp.Description("Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all")),
		mcp.WithString("created_from", mcp.Description("Retrieve entities from their creation date")),
		mcp.WithString("created_to", mcp.Description("Retrieve entities to their creation date")),
		mcp.WithString("modified_from", mcp.Description("Retrieve entities from their modification date")),
		mcp.WithString("modified_to", mcp.Description("Retrieve entities to their modification date")),
		mcp.WithString("category_id", mcp.Description("Retrieves products’ variants specified by category id")),
		mcp.WithString("product_id", mcp.Description("Retrieves products' variants specified by product id")),
		mcp.WithString("store_id", mcp.Description("Retrieves variants specified by store id")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ProductvariantlistHandler(cfg),
	}
}
