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

func AttributeaddHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("type=%v", val))
		}
		if val, ok := args["code"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("code=%v", val))
		}
		if val, ok := args["name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("name=%v", val))
		}
		if val, ok := args["store_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("store_id=%v", val))
		}
		if val, ok := args["lang_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("lang_id=%v", val))
		}
		if val, ok := args["visible"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("visible=%v", val))
		}
		if val, ok := args["required"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("required=%v", val))
		}
		if val, ok := args["position"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("position=%v", val))
		}
		if val, ok := args["attribute_group_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("attribute_group_id=%v", val))
		}
		if val, ok := args["is_global"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("is_global=%v", val))
		}
		if val, ok := args["is_searchable"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("is_searchable=%v", val))
		}
		if val, ok := args["is_filterable"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("is_filterable=%v", val))
		}
		if val, ok := args["is_comparable"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("is_comparable=%v", val))
		}
		if val, ok := args["is_html_allowed_on_front"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("is_html_allowed_on_front=%v", val))
		}
		if val, ok := args["is_filterable_in_search"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("is_filterable_in_search=%v", val))
		}
		if val, ok := args["is_configurable"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("is_configurable=%v", val))
		}
		if val, ok := args["is_visible_in_advanced_search"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("is_visible_in_advanced_search=%v", val))
		}
		if val, ok := args["is_used_for_promo_rules"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("is_used_for_promo_rules=%v", val))
		}
		if val, ok := args["used_in_product_listing"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("used_in_product_listing=%v", val))
		}
		if val, ok := args["used_for_sort_by"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("used_for_sort_by=%v", val))
		}
		if val, ok := args["apply_to"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("apply_to=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/attribute.add.json%s", cfg.BaseURL, queryString)
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

func CreateAttributeaddTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_attribute_add_json",
		mcp.WithDescription("Add new attribute"),
		mcp.WithString("type", mcp.Required(), mcp.Description("Defines attribute's type")),
		mcp.WithString("code", mcp.Description("Entity code")),
		mcp.WithString("name", mcp.Required(), mcp.Description("Defines attributes's name")),
		mcp.WithString("store_id", mcp.Description("Store Id")),
		mcp.WithString("lang_id", mcp.Description("Language id")),
		mcp.WithBoolean("visible", mcp.Description("Set visibility status")),
		mcp.WithBoolean("required", mcp.Description("Defines if the option is required")),
		mcp.WithNumber("position", mcp.Description("Attribute`s position")),
		mcp.WithString("attribute_group_id", mcp.Description("Filter by attribute_group_id")),
		mcp.WithString("is_global", mcp.Description("Attribute saving scope")),
		mcp.WithBoolean("is_searchable", mcp.Description("Use attribute in Quick Search")),
		mcp.WithString("is_filterable", mcp.Description("Use In Layered Navigation")),
		mcp.WithBoolean("is_comparable", mcp.Description("Comparable on Front-end")),
		mcp.WithBoolean("is_html_allowed_on_front", mcp.Description("Allow HTML Tags on Frontend")),
		mcp.WithBoolean("is_filterable_in_search", mcp.Description("Use In Search Results Layered Navigation")),
		mcp.WithBoolean("is_configurable", mcp.Description("Use To Create Configurable Product")),
		mcp.WithBoolean("is_visible_in_advanced_search", mcp.Description("Use in Advanced Search")),
		mcp.WithBoolean("is_used_for_promo_rules", mcp.Description("Use for Promo Rule Conditions")),
		mcp.WithBoolean("used_in_product_listing", mcp.Description("Used in Product Listing")),
		mcp.WithBoolean("used_for_sort_by", mcp.Description("Used for Sorting in Product Listing")),
		mcp.WithString("apply_to", mcp.Description("Types of products which can have this attribute")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    AttributeaddHandler(cfg),
	}
}
