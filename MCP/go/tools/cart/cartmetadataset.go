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

func CartmetadatasetHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["entity_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("entity_id=%v", val))
		}
		if val, ok := args["entity"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("entity=%v", val))
		}
		if val, ok := args["store_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("store_id=%v", val))
		}
		if val, ok := args["lang_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("lang_id=%v", val))
		}
		if val, ok := args["key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("key=%v", val))
		}
		if val, ok := args["value"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("value=%v", val))
		}
		if val, ok := args["namespace"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("namespace=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/cart.meta_data.set.json%s", cfg.BaseURL, queryString)
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

func CreateCartmetadatasetTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_cart_meta_data_set_json",
		mcp.WithDescription("Set meta data for a specific entity"),
		mcp.WithString("entity_id", mcp.Required(), mcp.Description("Entity Id")),
		mcp.WithString("entity", mcp.Description("Entity")),
		mcp.WithString("store_id", mcp.Description("Store Id")),
		mcp.WithString("lang_id", mcp.Description("Language id")),
		mcp.WithString("key", mcp.Required(), mcp.Description("Key")),
		mcp.WithString("value", mcp.Required(), mcp.Description("Value")),
		mcp.WithString("namespace", mcp.Required(), mcp.Description("Metafield namespace")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CartmetadatasetHandler(cfg),
	}
}
