package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/swagger-api2cart/mcp-server/config"
	"github.com/swagger-api2cart/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func OrderrefundaddHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.OrderRefundAdd
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/order.refund.add.json", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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

func CreateOrderrefundaddTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_order_refund_add_json",
		mcp.WithDescription("Add a refund to the order."),
		mcp.WithString("shipping_price", mcp.Description("Input parameter: Defines refund shipping amount.")),
		mcp.WithString("date", mcp.Description("Input parameter: Specifies an order creation date in format Y-m-d H:i:s")),
		mcp.WithBoolean("item_restock", mcp.Description("Input parameter: Boolean, whether or not to add the line items back to the store inventory.")),
		mcp.WithArray("items", mcp.Description("Input parameter: Defines items in the order that will be refunded")),
		mcp.WithBoolean("send_notifications", mcp.Description("Input parameter: Send notifications to customer after refund was created")),
		mcp.WithString("message", mcp.Description("Input parameter: Refund reason, or some else message which assigned to refund.")),
		mcp.WithString("total_price", mcp.Description("Input parameter: Defines order refund amount.")),
		mcp.WithString("fee_price", mcp.Description("Input parameter: Specifies refund's fee price")),
		mcp.WithBoolean("is_online", mcp.Description("Input parameter: Indicates whether refund type is online")),
		mcp.WithString("order_id", mcp.Description("Input parameter: Defines the order for which the refund will be created.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    OrderrefundaddHandler(cfg),
	}
}
