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

func CartcouponcountHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["store_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("store_id=%v", val))
		}
		if val, ok := args["date_start_from"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("date_start_from=%v", val))
		}
		if val, ok := args["date_start_to"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("date_start_to=%v", val))
		}
		if val, ok := args["date_end_from"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("date_end_from=%v", val))
		}
		if val, ok := args["date_end_to"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("date_end_to=%v", val))
		}
		if val, ok := args["avail"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("avail=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/cart.coupon.count.json%s", cfg.BaseURL, queryString)
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

func CreateCartcouponcountTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_cart_coupon_count_json",
		mcp.WithDescription("Get cart coupons count."),
		mcp.WithString("store_id", mcp.Description("Store Id")),
		mcp.WithString("date_start_from", mcp.Description("Filter entity by date_start (greater or equal)")),
		mcp.WithString("date_start_to", mcp.Description("Filter entity by date_start (less or equal)")),
		mcp.WithString("date_end_from", mcp.Description("Filter entity by date_end (greater or equal)")),
		mcp.WithString("date_end_to", mcp.Description("Filter entity by date_end (less or equal)")),
		mcp.WithBoolean("avail", mcp.Description("Defines category's visibility status")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CartcouponcountHandler(cfg),
	}
}
