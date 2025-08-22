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

func OrderaddHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.OrderAdd
		
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
		url := fmt.Sprintf("%s/order.add.json", cfg.BaseURL)
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

func CreateOrderaddTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_order_add_json",
		mcp.WithDescription("Add a new order to the cart."),
		mcp.WithArray("order_item", mcp.Required(), mcp.Description("")),
		mcp.WithString("customer_last_name", mcp.Description("Input parameter: Specifies customer’s last name")),
		mcp.WithString("customer_birthday", mcp.Description("Input parameter: Specifies customer’s birthday")),
		mcp.WithString("shipp_company", mcp.Description("Input parameter: Specifies shipping company")),
		mcp.WithString("channel_id", mcp.Description("Input parameter: Channel ID")),
		mcp.WithString("subtotal_price", mcp.Description("Input parameter: Total price of all ordered products multiplied by their number, excluding tax, shipping price and discounts")),
		mcp.WithString("tags", mcp.Description("Input parameter: Order tags")),
		mcp.WithString("shipp_phone", mcp.Description("Input parameter: Specifies shipping phone")),
		mcp.WithString("financial_status", mcp.Description("Input parameter: Create order with financial status")),
		mcp.WithString("customer_phone", mcp.Description("Input parameter: Specifies customer’s phone")),
		mcp.WithString("order_payment_method", mcp.Description("Input parameter: Defines order payment method.<br/>Setting order_payment_method on Shopify will also change financial_status field value to 'paid'")),
		mcp.WithString("id", mcp.Description("Input parameter: Defines order's id")),
		mcp.WithString("comment", mcp.Description("Input parameter: Specifies order comment")),
		mcp.WithArray("coupons", mcp.Description("Input parameter: Coupons that will be applied to order")),
		mcp.WithString("shipp_city", mcp.Description("Input parameter: Specifies shipping city")),
		mcp.WithString("shipp_country", mcp.Description("Input parameter: Specifies shipping country code")),
		mcp.WithString("shipp_state", mcp.Description("Input parameter: Specifies shipping state code")),
		mcp.WithString("bill_state", mcp.Required(), mcp.Description("Input parameter: Specifies billing state code")),
		mcp.WithString("shipp_postcode", mcp.Description("Input parameter: Specifies shipping postcode")),
		mcp.WithString("bill_city", mcp.Required(), mcp.Description("Input parameter: Specifies billing city")),
		mcp.WithString("bill_address_2", mcp.Description("Input parameter: Specifies second billing address")),
		mcp.WithString("order_id", mcp.Description("Input parameter: Defines the order id if it is supported by the cart")),
		mcp.WithString("admin_private_comment", mcp.Description("Input parameter: Specifies private admin's order comment")),
		mcp.WithString("shipping_tax", mcp.Description("Input parameter: Specifies order's shipping price tax")),
		mcp.WithString("shipp_last_name", mcp.Description("Input parameter: Specifies shipping last name")),
		mcp.WithBoolean("clear_cache", mcp.Description("Input parameter: Is cache clear required")),
		mcp.WithString("total_price", mcp.Description("Input parameter: Defines order's total price")),
		mcp.WithString("bill_phone", mcp.Description("Input parameter: Specifies billing phone")),
		mcp.WithString("discount", mcp.Description("Input parameter: Specifies order's discount")),
		mcp.WithString("gift_certificate_discount", mcp.Description("Input parameter: Discounts for order with gift certificates")),
		mcp.WithString("date_finished", mcp.Description("Input parameter: Specifies order's  finished date")),
		mcp.WithArray("note_attributes", mcp.Description("Input parameter: Defines note attributes")),
		mcp.WithString("shipping_price", mcp.Description("Input parameter: Specifies order's shipping price")),
		mcp.WithString("store_id", mcp.Description("Input parameter: Defines store id where the order should be assigned")),
		mcp.WithString("tax_price", mcp.Description("Input parameter: The value of tax cost for order")),
		mcp.WithBoolean("create_invoice", mcp.Description("Input parameter: Defines whether the invoice is created automatically along with the order")),
		mcp.WithString("bill_company", mcp.Description("Input parameter: Specifies billing company")),
		mcp.WithString("coupon_discount", mcp.Description("Input parameter: Specifies order's coupon discount")),
		mcp.WithString("bill_postcode", mcp.Required(), mcp.Description("Input parameter: Specifies billing postcode")),
		mcp.WithString("transaction_id", mcp.Description("Input parameter: Payment transaction id")),
		mcp.WithString("admin_comment", mcp.Description("Input parameter: Specifies admin's order comment")),
		mcp.WithString("date_modified", mcp.Description("Input parameter: Specifies order's  modification date")),
		mcp.WithString("shipp_fax", mcp.Description("Input parameter: Specifies shipping fax")),
		mcp.WithString("currency", mcp.Description("Input parameter: Currency code of order")),
		mcp.WithString("bill_last_name", mcp.Required(), mcp.Description("Input parameter: Specifies billing last name")),
		mcp.WithBoolean("send_notifications", mcp.Description("Input parameter: Send notifications to customer after order was created")),
		mcp.WithString("shipp_first_name", mcp.Description("Input parameter: Specifies shipping first name")),
		mcp.WithString("inventory_behaviour", mcp.Description("Input parameter: The behaviour to use when updating inventory.<hr><div style=\"font-style:normal\">Values description:<div style=\"margin-left: 2%; padding-top: 2%\"><div style=\"font-size:85%\"><b>bypass</b> = Do not claim inventory </br></br><b>decrement_ignoring_policy</b> = Ignore the product's </br> inventory policy and claim amounts</br></br><b>decrement_obeying_policy</b> =  Obey the product's </br> inventory policy.</br></br></div></div></div>")),
		mcp.WithString("customer_email", mcp.Required(), mcp.Description("Input parameter: Defines the customer specified by email for whom order has to be created")),
		mcp.WithBoolean("send_admin_notifications", mcp.Description("Input parameter: Notify admin when new order was created.")),
		mcp.WithString("bill_address_1", mcp.Required(), mcp.Description("Input parameter: Specifies first billing address")),
		mcp.WithString("order_shipping_method", mcp.Description("Input parameter: Defines order shipping method")),
		mcp.WithString("order_status", mcp.Required(), mcp.Description("Input parameter: Defines order status.")),
		mcp.WithBoolean("prices_inc_tax", mcp.Description("Input parameter: Indicates whether prices and subtotal includes tax.")),
		mcp.WithString("shipp_address_1", mcp.Description("Input parameter: Specifies first shipping address")),
		mcp.WithString("bill_country", mcp.Required(), mcp.Description("Input parameter: Specifies billing country code")),
		mcp.WithString("customer_fax", mcp.Description("Input parameter: Specifies customer’s fax")),
		mcp.WithString("date", mcp.Description("Input parameter: Specifies an order creation date in format Y-m-d H:i:s")),
		mcp.WithString("external_source", mcp.Description("Input parameter: Identifying the system used to generate the order")),
		mcp.WithString("bill_fax", mcp.Description("Input parameter: Specifies billing fax")),
		mcp.WithString("fulfillment_status", mcp.Description("Input parameter: Create order with fulfillment status")),
		mcp.WithNumber("total_weight", mcp.Description("Input parameter: Defines the sum of all line item weights in grams for the order")),
		mcp.WithString("shipp_address_2", mcp.Description("Input parameter: Specifies second address line of a shipping street address")),
		mcp.WithString("customer_first_name", mcp.Description("Input parameter: Specifies customer's first name")),
		mcp.WithString("bill_first_name", mcp.Required(), mcp.Description("Input parameter: Specifies billing first name")),
		mcp.WithString("total_paid", mcp.Description("Input parameter: Defines total paid amount for the order")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    OrderaddHandler(cfg),
	}
}
