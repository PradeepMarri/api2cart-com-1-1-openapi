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

func CustomeraddHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.CustomerAdd
		
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
		url := fmt.Sprintf("%s/customer.add.json", cfg.BaseURL)
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

func CreateCustomeraddTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_customer_add_json",
		mcp.WithDescription("Add customer into store."),
		mcp.WithString("website", mcp.Description("Input parameter: Link to customer website")),
		mcp.WithString("modified_time", mcp.Description("Input parameter: Entity's date modification")),
		mcp.WithString("first_name", mcp.Required(), mcp.Description("Input parameter: Defines customer's first name")),
		mcp.WithString("status", mcp.Description("Input parameter: Defines customer's status")),
		mcp.WithString("created_time", mcp.Description("Input parameter: Entity's date creation")),
		mcp.WithBoolean("news_letter_subscription", mcp.Description("Input parameter: Defines whether the newsletter subscription is available for the user")),
		mcp.WithString("phone", mcp.Description("Input parameter: Defines customer's phone number")),
		mcp.WithString("company", mcp.Description("Input parameter: Defines customer's company")),
		mcp.WithString("last_name", mcp.Required(), mcp.Description("Input parameter: Defines customer's last name")),
		mcp.WithString("birth_day", mcp.Description("Input parameter: Defines customer's birthday")),
		mcp.WithString("fax", mcp.Description("Input parameter: Defines customer's fax")),
		mcp.WithString("group", mcp.Description("Input parameter: Defines the group where the customer")),
		mcp.WithString("last_login", mcp.Description("Input parameter: Defines customer's last login time")),
		mcp.WithArray("address", mcp.Description("")),
		mcp.WithString("login", mcp.Description("Input parameter: Specifies customer's login name")),
		mcp.WithString("email", mcp.Required(), mcp.Description("Input parameter: Defines customer's email")),
		mcp.WithString("gender", mcp.Description("Input parameter: Defines customer's gender")),
		mcp.WithString("password", mcp.Description("Input parameter: Defines customer's unique password")),
		mcp.WithString("store_id", mcp.Description("Input parameter: Store Id")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CustomeraddHandler(cfg),
	}
}
