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

func CategoryupdateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("id=%v", val))
		}
		if val, ok := args["name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("name=%v", val))
		}
		if val, ok := args["parent_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("parent_id=%v", val))
		}
		if val, ok := args["stores_ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("stores_ids=%v", val))
		}
		if val, ok := args["avail"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("avail=%v", val))
		}
		if val, ok := args["sort_order"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort_order=%v", val))
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
		url := fmt.Sprintf("%s/category.update.json%s", cfg.BaseURL, queryString)
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

func CreateCategoryupdateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_category_update_json",
		mcp.WithDescription("Update category in store"),
		mcp.WithString("id", mcp.Required(), mcp.Description("Defines category update specified by category id")),
		mcp.WithString("name", mcp.Description("Defines new category’s name")),
		mcp.WithString("parent_id", mcp.Description("Defines new parent category id")),
		mcp.WithString("stores_ids", mcp.Description("Update category in the stores that is specified by comma-separated stores' id")),
		mcp.WithBoolean("avail", mcp.Description("Defines category's visibility status")),
		mcp.WithNumber("sort_order", mcp.Description("Sort number in the list")),
		mcp.WithString("modified_time", mcp.Description("Entity's date modification")),
		mcp.WithString("description", mcp.Description("Defines new category's description")),
		mcp.WithString("meta_title", mcp.Description("Defines unique meta title for each entity")),
		mcp.WithString("meta_description", mcp.Description("Defines unique meta description of a entity")),
		mcp.WithString("meta_keywords", mcp.Description("Defines unique meta keywords for each entity")),
		mcp.WithString("seo_url", mcp.Description("Defines unique category's URL for SEO")),
		mcp.WithString("lang_id", mcp.Description("Language id")),
		mcp.WithString("store_id", mcp.Description("Store Id")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CategoryupdateHandler(cfg),
	}
}
