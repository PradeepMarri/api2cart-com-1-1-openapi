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

func OrdercountHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["customer_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("customer_id=%v", val))
		}
		if val, ok := args["customer_email"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("customer_email=%v", val))
		}
		if val, ok := args["order_status"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("order_status=%v", val))
		}
		if val, ok := args["order_status_ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("order_status_ids=%v", val))
		}
		if val, ok := args["created_to"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("created_to=%v", val))
		}
		if val, ok := args["created_from"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("created_from=%v", val))
		}
		if val, ok := args["modified_to"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("modified_to=%v", val))
		}
		if val, ok := args["modified_from"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("modified_from=%v", val))
		}
		if val, ok := args["store_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("store_id=%v", val))
		}
		if val, ok := args["ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ids=%v", val))
		}
		if val, ok := args["order_ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("order_ids=%v", val))
		}
		if val, ok := args["ebay_order_status"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ebay_order_status=%v", val))
		}
		if val, ok := args["financial_status"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("financial_status=%v", val))
		}
		if val, ok := args["fulfillment_status"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("fulfillment_status=%v", val))
		}
		if val, ok := args["shipping_method"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("shipping_method=%v", val))
		}
		if val, ok := args["delivery_method"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("delivery_method=%v", val))
		}
		if val, ok := args["ship_node_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ship_node_type=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/order.count.json%s", cfg.BaseURL, queryString)
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

func CreateOrdercountTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_order_count_json",
		mcp.WithDescription("Count orders in store"),
		mcp.WithString("customer_id", mcp.Description("Counts orders quantity specified by customer id")),
		mcp.WithString("customer_email", mcp.Description("Counts orders quantity specified by customer email")),
		mcp.WithString("order_status", mcp.Description("Counts orders quantity specified by order status")),
		mcp.WithArray("order_status_ids", mcp.Description("Retrieves orders specified by order statuses")),
		mcp.WithString("created_to", mcp.Description("Retrieve entities to their creation date")),
		mcp.WithString("created_from", mcp.Description("Retrieve entities from their creation date")),
		mcp.WithString("modified_to", mcp.Description("Retrieve entities to their modification date")),
		mcp.WithString("modified_from", mcp.Description("Retrieve entities from their modification date")),
		mcp.WithString("store_id", mcp.Description("Counts orders quantity specified by store id")),
		mcp.WithString("ids", mcp.Description("Counts orders specified by ids")),
		mcp.WithString("order_ids", mcp.Description("Counts orders specified by order ids")),
		mcp.WithString("ebay_order_status", mcp.Description("Counts orders quantity specified by order status")),
		mcp.WithString("financial_status", mcp.Description("Counts orders quantity specified by financial status")),
		mcp.WithString("fulfillment_status", mcp.Description("Create order with fulfillment status")),
		mcp.WithString("shipping_method", mcp.Description("Retrieve entities according to shipping method")),
		mcp.WithString("delivery_method", mcp.Description("Retrieves order with delivery method")),
		mcp.WithString("ship_node_type", mcp.Description("Retrieves order with ship node type")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    OrdercountHandler(cfg),
	}
}
