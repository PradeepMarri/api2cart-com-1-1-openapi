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

func CategoryaddHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("name=%v", val))
		}
		if val, ok := args["parent_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("parent_id=%v", val))
		}
		if val, ok := args["stores_ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("stores_ids=%v", val))
		}
		if val, ok := args["store_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("store_id=%v", val))
		}
		if val, ok := args["lang_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("lang_id=%v", val))
		}
		if val, ok := args["avail"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("avail=%v", val))
		}
		if val, ok := args["sort_order"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort_order=%v", val))
		}
		if val, ok := args["created_time"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("created_time=%v", val))
		}
		if val, ok := args["modified_time"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("modified_time=%v", val))
		}
		if val, ok := args["description"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("description=%v", val))
		}
		if val, ok := args["meta_title"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("meta_title=%v", val))
		}
		if val, ok := args["meta_description"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("meta_description=%v", val))
		}
		if val, ok := args["meta_keywords"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("meta_keywords=%v", val))
		}
		if val, ok := args["seo_url"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("seo_url=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/category.add.json%s", cfg.BaseURL, queryString)
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

func CreateCategoryaddTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_category_add_json",
		mcp.WithDescription("Add new category in store"),
		mcp.WithString("name", mcp.Required(), mcp.Description("Defines category's name that has to be added")),
		mcp.WithString("parent_id", mcp.Description("Adds categories specified by parent id")),
		mcp.WithString("stores_ids", mcp.Description("Create category in the stores that is specified by comma-separated stores' id")),
		mcp.WithString("store_id", mcp.Description("Store Id")),
		mcp.WithString("lang_id", mcp.Description("Language id")),
		mcp.WithBoolean("avail", mcp.Description("Defines category's visibility status")),
		mcp.WithNumber("sort_order", mcp.Description("Sort number in the list")),
		mcp.WithString("created_time", mcp.Description("Entity's date creation")),
		mcp.WithString("modified_time", mcp.Description("Entity's date modification")),
		mcp.WithString("description", mcp.Description("Defines category's description")),
		mcp.WithString("meta_title", mcp.Description("Defines unique meta title for each entity")),
		mcp.WithString("meta_description", mcp.Description("Defines unique meta description of a entity")),
		mcp.WithString("meta_keywords", mcp.Description("Defines unique meta keywords for each entity")),
		mcp.WithString("seo_url", mcp.Description("Defines unique category's URL for SEO")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CategoryaddHandler(cfg),
	}
}
