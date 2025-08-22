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

func CartcouponaddHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.CartCouponAdd
		
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
		url := fmt.Sprintf("%s/cart.coupon.add.json", cfg.BaseURL)
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

func CreateCartcouponaddTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_cart_coupon_add_json",
		mcp.WithDescription("Create new coupon"),
		mcp.WithString("action_apply_to", mcp.Required(), mcp.Description("Input parameter: Defines where discount should be applied")),
		mcp.WithString("date_end", mcp.Description("Input parameter: Defines when discount code will be expired.")),
		mcp.WithNumber("usage_limit_per_customer", mcp.Description("Input parameter: Usage limit per customer.")),
		mcp.WithString("action_condition_operator", mcp.Description("Input parameter: Defines condition operator.")),
		mcp.WithArray("codes", mcp.Description("Input parameter: Entity codes")),
		mcp.WithString("action_condition_entity", mcp.Description("Input parameter: Defines entity for action condition.")),
		mcp.WithString("store_id", mcp.Description("Input parameter: Store Id")),
		mcp.WithString("action_condition_key", mcp.Description("Input parameter: Defines entity attribute code for action condition.")),
		mcp.WithString("date_start", mcp.Description("Input parameter: Defines when discount code will be available.")),
		mcp.WithString("action_condition_value", mcp.Description("Input parameter: Defines condition attribute value/s. Can be comma separated string.")),
		mcp.WithString("code", mcp.Required(), mcp.Description("Input parameter: Coupon code")),
		mcp.WithString("action_amount", mcp.Required(), mcp.Description("Input parameter: Defines the discount amount value.")),
		mcp.WithString("name", mcp.Description("Input parameter: Coupon name")),
		mcp.WithString("action_scope", mcp.Required(), mcp.Description("Input parameter: Specify how discount should be applied. If scope=matching_items, then discount will be applied to each of the items that match action conditions. Scope order means that discount will be applied once.")),
		mcp.WithString("action_type", mcp.Required(), mcp.Description("Input parameter: Coupon discount type")),
		mcp.WithNumber("usage_limit", mcp.Description("Input parameter: Usage limit for coupon.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CartcouponaddHandler(cfg),
	}
}
