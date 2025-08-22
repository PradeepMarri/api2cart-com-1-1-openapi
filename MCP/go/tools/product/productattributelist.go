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

func ProductattributelistHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["variant_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("variant_id=%v", val))
		}
		if val, ok := args["page_cursor"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page_cursor=%v", val))
		}
		if val, ok := args["start"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("start=%v", val))
		}
		if val, ok := args["count"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("count=%v", val))
		}
		if val, ok := args["attribute_group_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("attribute_group_id=%v", val))
		}
		if val, ok := args["set_name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("set_name=%v", val))
		}
		if val, ok := args["lang_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("lang_id=%v", val))
		}
		if val, ok := args["store_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("store_id=%v", val))
		}
		if val, ok := args["sort_by"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort_by=%v", val))
		}
		if val, ok := args["sort_direction"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort_direction=%v", val))
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
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/product.attribute.list.json%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.ModelResponseProductAttributeList
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

func CreateProductattributelistTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_product_attribute_list_json",
		mcp.WithDescription("Get list of attributes and values."),
		mcp.WithString("product_id", mcp.Required(), mcp.Description("Retrieves attributes specified by product id")),
		mcp.WithString("attribute_id", mcp.Description("Retrieves info for specified attribute_id")),
		mcp.WithString("variant_id", mcp.Description("Defines product's variants specified by variant id")),
		mcp.WithString("page_cursor", mcp.Description("Used to retrieve entities via cursor-based pagination (it can't be used with any other filtering parameter)")),
		mcp.WithNumber("start", mcp.Description("This parameter sets the number from which you want to get entities")),
		mcp.WithNumber("count", mcp.Description("This parameter sets the entity amount that has to be retrieved. Max allowed count=250")),
		mcp.WithString("attribute_group_id", mcp.Description("Filter by attribute_group_id")),
		mcp.WithString("set_name", mcp.Description("Retrieves attributes specified by set_name in Magento")),
		mcp.WithString("lang_id", mcp.Description("Retrieves attributes specified by language id")),
		mcp.WithString("store_id", mcp.Description("Retrieves attributes specified by store id")),
		mcp.WithString("sort_by", mcp.Description("Set field to sort by")),
		mcp.WithString("sort_direction", mcp.Description("Set sorting direction")),
		mcp.WithString("params", mcp.Description("Set this parameter in order to choose which entity fields you want to retrieve")),
		mcp.WithString("response_fields", mcp.Description("Set this parameter in order to choose which entity fields you want to retrieve")),
		mcp.WithString("exclude", mcp.Description("Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ProductattributelistHandler(cfg),
	}
}
