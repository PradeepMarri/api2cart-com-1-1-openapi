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

func OrderpreestimateshippinglistHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.OrderPreestimateShippingList
		
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
		url := fmt.Sprintf("%s/order.preestimate_shipping.list.json", cfg.BaseURL)
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
		var result models.ModelResponseOrderPreestimateShippingList
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

func CreateOrderpreestimateshippinglistTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_order_preestimate_shipping_list_json",
		mcp.WithDescription("Retrieve list of order preestimated shipping methods"),
		mcp.WithArray("order_item", mcp.Required(), mcp.Description("")),
		mcp.WithString("params", mcp.Description("Input parameter: Set this parameter in order to choose which entity fields you want to retrieve")),
		mcp.WithString("warehouse_id", mcp.Description("Input parameter: This parameter is used for selecting a warehouse where you need to set/modify a product quantity.")),
		mcp.WithString("customer_email", mcp.Description("Input parameter: Retrieves orders specified by customer email")),
		mcp.WithString("customer_id", mcp.Description("Input parameter: Retrieves orders specified by customer id")),
		mcp.WithString("shipp_city", mcp.Description("Input parameter: Specifies shipping city")),
		mcp.WithString("store_id", mcp.Description("Input parameter: Store Id")),
		mcp.WithString("exclude", mcp.Description("Input parameter: Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all")),
		mcp.WithString("shipp_address_1", mcp.Description("Input parameter: Specifies first shipping address")),
		mcp.WithString("shipp_country", mcp.Required(), mcp.Description("Input parameter: Specifies shipping country code")),
		mcp.WithString("shipp_postcode", mcp.Description("Input parameter: Specifies shipping postcode")),
		mcp.WithString("shipp_state", mcp.Description("Input parameter: Specifies shipping state code")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    OrderpreestimateshippinglistHandler(cfg),
	}
}
