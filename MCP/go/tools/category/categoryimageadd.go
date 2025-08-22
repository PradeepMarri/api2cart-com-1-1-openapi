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

func CategoryimageaddHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["category_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("category_id=%v", val))
		}
		if val, ok := args["image_name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("image_name=%v", val))
		}
		if val, ok := args["url"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("url=%v", val))
		}
		if val, ok := args["label"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("label=%v", val))
		}
		if val, ok := args["mime"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("mime=%v", val))
		}
		if val, ok := args["type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("type=%v", val))
		}
		if val, ok := args["position"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("position=%v", val))
		}
		if val, ok := args["store_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("store_id=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/category.image.add.json%s", cfg.BaseURL, queryString)
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

func CreateCategoryimageaddTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_category_image_add_json",
		mcp.WithDescription("Add image to category"),
		mcp.WithString("category_id", mcp.Required(), mcp.Description("Defines category id where the image should be added")),
		mcp.WithString("image_name", mcp.Required(), mcp.Description("Defines image's name")),
		mcp.WithString("url", mcp.Required(), mcp.Description("Defines URL of the image that has to be added")),
		mcp.WithString("label", mcp.Description("Defines alternative text that has to be attached to the picture")),
		mcp.WithString("mime", mcp.Description("Mime type of image http://en.wikipedia.org/wiki/Internet_media_type.")),
		mcp.WithString("type", mcp.Required(), mcp.Description("Defines image's types that are specified by comma-separated list")),
		mcp.WithNumber("position", mcp.Description("Defines image’s position in the list")),
		mcp.WithString("store_id", mcp.Description("Store Id")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CategoryimageaddHandler(cfg),
	}
}
