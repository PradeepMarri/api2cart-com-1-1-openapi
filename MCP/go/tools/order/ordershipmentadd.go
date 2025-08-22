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

func OrdershipmentaddHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.OrderShipmentAdd
		
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
		url := fmt.Sprintf("%s/order.shipment.add.json", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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

func CreateOrdershipmentaddTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_order_shipment_add_json",
		mcp.WithDescription("Add a shipment to the order."),
		mcp.WithString("store_id", mcp.Description("Input parameter: Store Id")),
		mcp.WithString("warehouse_id", mcp.Description("Input parameter: This parameter is used for selecting a warehouse where you need to set/modify a product quantity.")),
		mcp.WithBoolean("send_notifications", mcp.Description("Input parameter: Send notifications to customer after shipment was created")),
		mcp.WithString("shipping_method", mcp.Description("Input parameter: Define shipping method")),
		mcp.WithBoolean("adjust_stock", mcp.Description("Input parameter: This parameter is used for adjust stock.")),
		mcp.WithArray("items", mcp.Description("Input parameter: Defines items in the order that will be shipped")),
		mcp.WithString("order_id", mcp.Description("Input parameter: Defines the order for which the shipment will be created")),
		mcp.WithBoolean("enable_cache", mcp.Description("Input parameter: If the value is 'true' and order exist in our cache, we will use order.info from cache to prepare shipment items.")),
		mcp.WithBoolean("is_shipped", mcp.Description("Input parameter: Defines shipment's status")),
		mcp.WithString("shipment_provider", mcp.Description("Input parameter: Defines company name that provide tracking of shipment")),
		mcp.WithString("tracking_link", mcp.Description("Input parameter: Defines custom tracking link")),
		mcp.WithArray("tracking_numbers", mcp.Description("Input parameter: Defines shipment's tracking numbers that have to be added</br> How set tracking numbers to appropriate carrier:<ul><li>tracking_numbers[]=a2c.demo1,a2c.demo2 - set default carrier</li><li>tracking_numbers[<b>carrier_id</b>]=a2c.demo - set appropriate carrier</li></ul>To get the list of carriers IDs that are available in your store, use the <a href = \"https://api2cart.com/docs/#/cart/CartInfo\">cart.info</a > method")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    OrdershipmentaddHandler(cfg),
	}
}
