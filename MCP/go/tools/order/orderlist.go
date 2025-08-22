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

func OrderlistHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if val, ok := args["phone"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("phone=%v", val))
		}
		if val, ok := args["order_status"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("order_status=%v", val))
		}
		if val, ok := args["order_status_ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("order_status_ids=%v", val))
		}
		if val, ok := args["start"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("start=%v", val))
		}
		if val, ok := args["count"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("count=%v", val))
		}
		if val, ok := args["page_cursor"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page_cursor=%v", val))
		}
		if val, ok := args["sort_by"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort_by=%v", val))
		}
		if val, ok := args["sort_direction"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sort_direction=%v", val))
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
		if val, ok := args["basket_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("basket_id=%v", val))
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
		if val, ok := args["skip_order_ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("skip_order_ids=%v", val))
		}
		if val, ok := args["since_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("since_id=%v", val))
		}
		if val, ok := args["is_deleted"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("is_deleted=%v", val))
		}
		if val, ok := args["shipping_country_iso3"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("shipping_country_iso3=%v", val))
		}
		if val, ok := args["enable_cache"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("enable_cache=%v", val))
		}
		if val, ok := args["delivery_method"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("delivery_method=%v", val))
		}
		if val, ok := args["ship_node_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ship_node_type=%v", val))
		}
		if val, ok := args["currency_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("currency_id=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/order.list.json%s", cfg.BaseURL, queryString)
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
		var result models.ModelResponseOrderList
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

func CreateOrderlistTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_order_list_json",
		mcp.WithDescription("Get list of orders from store."),
		mcp.WithString("customer_id", mcp.Description("Retrieves orders specified by customer id")),
		mcp.WithString("customer_email", mcp.Description("Retrieves orders specified by customer email")),
		mcp.WithString("phone", mcp.Description("Filter orders by customer's phone number")),
		mcp.WithString("order_status", mcp.Description("Retrieves orders specified by order status")),
		mcp.WithArray("order_status_ids", mcp.Description("Retrieves orders specified by order statuses")),
		mcp.WithNumber("start", mcp.Description("This parameter sets the number from which you want to get entities")),
		mcp.WithNumber("count", mcp.Description("This parameter sets the entity amount that has to be retrieved. Max allowed count=250")),
		mcp.WithString("page_cursor", mcp.Description("Used to retrieve orders via cursor-based pagination (it can't be used with any other filtering parameter)")),
		mcp.WithString("sort_by", mcp.Description("Set field to sort by")),
		mcp.WithString("sort_direction", mcp.Description("Set sorting direction")),
		mcp.WithString("params", mcp.Description("Set this parameter in order to choose which entity fields you want to retrieve")),
		mcp.WithString("response_fields", mcp.Description("Set this parameter in order to choose which entity fields you want to retrieve")),
		mcp.WithString("exclude", mcp.Description("Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all")),
		mcp.WithString("created_to", mcp.Description("Retrieve entities to their creation date")),
		mcp.WithString("created_from", mcp.Description("Retrieve entities from their creation date")),
		mcp.WithString("modified_to", mcp.Description("Retrieve entities to their modification date")),
		mcp.WithString("modified_from", mcp.Description("Retrieve entities from their modification date")),
		mcp.WithString("store_id", mcp.Description("Store Id")),
		mcp.WithString("ids", mcp.Description("Retrieves orders specified by ids")),
		mcp.WithString("order_ids", mcp.Description("Retrieves orders specified by order ids")),
		mcp.WithString("ebay_order_status", mcp.Description("Retrieves orders specified by order status")),
		mcp.WithString("basket_id", mcp.Description("Retrieves order’s info specified by basket id.")),
		mcp.WithString("financial_status", mcp.Description("Retrieves orders specified by financial status")),
		mcp.WithString("fulfillment_status", mcp.Description("Create order with fulfillment status")),
		mcp.WithString("shipping_method", mcp.Description("Retrieve entities according to shipping method")),
		mcp.WithString("skip_order_ids", mcp.Description("Skipped orders by ids")),
		mcp.WithNumber("since_id", mcp.Description("Retrieve entities starting from the specified id.")),
		mcp.WithBoolean("is_deleted", mcp.Description("Filter deleted orders")),
		mcp.WithString("shipping_country_iso3", mcp.Description("Retrieve entities according to shipping country")),
		mcp.WithBoolean("enable_cache", mcp.Description("If the value is 'true', we will cache orders for a 15 minutes in order to increase speed and reduce requests throttling for some methods and shoping platforms (for example order.shipment.add)")),
		mcp.WithString("delivery_method", mcp.Description("Retrieves order with delivery method")),
		mcp.WithString("ship_node_type", mcp.Description("Retrieves order with ship node type")),
		mcp.WithString("currency_id", mcp.Description("Currency Id")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    OrderlistHandler(cfg),
	}
}
