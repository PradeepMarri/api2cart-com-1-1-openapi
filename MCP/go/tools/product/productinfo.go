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

func ProductinfoHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("id=%v", val))
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
		if val, ok := args["store_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("store_id=%v", val))
		}
		if val, ok := args["lang_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("lang_id=%v", val))
		}
		if val, ok := args["currency_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("currency_id=%v", val))
		}
		if val, ok := args["report_request_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("report_request_id=%v", val))
		}
		if val, ok := args["disable_report_cache"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("disable_report_cache=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/product.info.json%s", cfg.BaseURL, queryString)
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

func CreateProductinfoTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_product_info_json",
		mcp.WithDescription("Get product info about product ID *** or specify other product ID."),
		mcp.WithString("id", mcp.Required(), mcp.Description("Retrieves product's info specified by product id")),
		mcp.WithString("params", mcp.Description("Set this parameter in order to choose which entity fields you want to retrieve")),
		mcp.WithString("response_fields", mcp.Description("Set this parameter in order to choose which entity fields you want to retrieve")),
		mcp.WithString("exclude", mcp.Description("Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all")),
		mcp.WithString("store_id", mcp.Description("Retrieves product info specified by store id")),
		mcp.WithString("lang_id", mcp.Description("Retrieves product info specified by language id")),
		mcp.WithString("currency_id", mcp.Description("Currency Id")),
		mcp.WithString("report_request_id", mcp.Description("Report request id")),
		mcp.WithBoolean("disable_report_cache", mcp.Description("Disable report cache for current request")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ProductinfoHandler(cfg),
	}
}
