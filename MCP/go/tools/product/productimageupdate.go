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

func ProductimageupdateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["product_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("product_id=%v", val))
		}
		if val, ok := args["variant_ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("variant_ids=%v", val))
		}
		if val, ok := args["image_name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("image_name=%v", val))
		}
		if val, ok := args["type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("type=%v", val))
		}
		if val, ok := args["label"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("label=%v", val))
		}
		if val, ok := args["position"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("position=%v", val))
		}
		if val, ok := args["id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("id=%v", val))
		}
		if val, ok := args["store_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("store_id=%v", val))
		}
		if val, ok := args["lang_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("lang_id=%v", val))
		}
		if val, ok := args["hidden"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("hidden=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/product.image.update.json%s", cfg.BaseURL, queryString)
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

func CreateProductimageupdateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_product_image_update_json",
		mcp.WithDescription("Update details of image"),
		mcp.WithString("product_id", mcp.Required(), mcp.Description("Defines product id where the image should be updated")),
		mcp.WithString("variant_ids", mcp.Description("Defines product's variants ids")),
		mcp.WithString("image_name", mcp.Description("Defines image's name")),
		mcp.WithString("type", mcp.Description("Defines image's types that are specified by comma-separated list")),
		mcp.WithString("label", mcp.Description("Defines alternative text that has to be attached to the picture")),
		mcp.WithNumber("position", mcp.Description("Defines image’s position in the list")),
		mcp.WithString("id", mcp.Required(), mcp.Description("Defines image update specified by image id")),
		mcp.WithString("store_id", mcp.Description("Store Id")),
		mcp.WithString("lang_id", mcp.Description("Language id")),
		mcp.WithBoolean("hidden", mcp.Description("Define is hide image")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ProductimageupdateHandler(cfg),
	}
}
