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

func OrderupdateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["order_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("order_id=%v", val))
		}
		if val, ok := args["store_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("store_id=%v", val))
		}
		if val, ok := args["order_status"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("order_status=%v", val))
		}
		if val, ok := args["comment"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("comment=%v", val))
		}
		if val, ok := args["admin_comment"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("admin_comment=%v", val))
		}
		if val, ok := args["admin_private_comment"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("admin_private_comment=%v", val))
		}
		if val, ok := args["date_modified"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("date_modified=%v", val))
		}
		if val, ok := args["date_finished"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("date_finished=%v", val))
		}
		if val, ok := args["financial_status"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("financial_status=%v", val))
		}
		if val, ok := args["fulfillment_status"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fulfillment_status=%v", val))
		}
		if val, ok := args["order_payment_method"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("order_payment_method=%v", val))
		}
		if val, ok := args["send_notifications"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("send_notifications=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/order.update.json%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("PUT", url, nil)
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

func CreateOrderupdateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_order_update_json",
		mcp.WithDescription("Update existing order."),
		mcp.WithString("order_id", mcp.Required(), mcp.Description("Defines the orders specified by order id")),
		mcp.WithString("store_id", mcp.Description("Defines store id where the order should be found")),
		mcp.WithString("order_status", mcp.Description("Defines new order's status")),
		mcp.WithString("comment", mcp.Description("Specifies order comment")),
		mcp.WithString("admin_comment", mcp.Description("Specifies admin's order comment")),
		mcp.WithString("admin_private_comment", mcp.Description("Specifies private admin's order comment")),
		mcp.WithString("date_modified", mcp.Description("Specifies order's  modification date")),
		mcp.WithString("date_finished", mcp.Description("Specifies order's  finished date")),
		mcp.WithString("financial_status", mcp.Description("Update order financial status to specified")),
		mcp.WithString("fulfillment_status", mcp.Description("Create order with fulfillment status")),
		mcp.WithString("order_payment_method", mcp.Description("Defines order payment method.<br/>Setting order_payment_method on Shopify will also change financial_status field value to 'paid'")),
		mcp.WithBoolean("send_notifications", mcp.Description("Send notifications to customer after order was created")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    OrderupdateHandler(cfg),
	}
}
