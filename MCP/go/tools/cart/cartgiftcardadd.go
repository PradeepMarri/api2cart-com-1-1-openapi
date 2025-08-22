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

func CartgiftcardaddHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["amount"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("amount=%v", val))
		}
		if val, ok := args["code"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("code=%v", val))
		}
		if val, ok := args["owner_email"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("owner_email=%v", val))
		}
		if val, ok := args["recipient_email"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("recipient_email=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/cart.giftcard.add.json%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("POST", url, nil)
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

func CreateCartgiftcardaddTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_cart_giftcard_add_json",
		mcp.WithDescription("Create new gift card"),
		mcp.WithString("amount", mcp.Required(), mcp.Description("Defines the gift card amount value.")),
		mcp.WithString("code", mcp.Description("Gift card code")),
		mcp.WithString("owner_email", mcp.Description("Gift card owner email")),
		mcp.WithString("recipient_email", mcp.Description("Gift card recipient email")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CartgiftcardaddHandler(cfg),
	}
}
