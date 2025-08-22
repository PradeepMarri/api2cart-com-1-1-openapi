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

func ProductimageaddHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.ProductImageAdd
		
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
		url := fmt.Sprintf("%s/product.image.add.json", cfg.BaseURL)
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

func CreateProductimageaddTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_product_image_add_json",
		mcp.WithDescription("Add image to product"),
		mcp.WithString("lang_id", mcp.Description("Input parameter: Add product image on specified language id")),
		mcp.WithString("product_id", mcp.Description("Input parameter: Defines product id where the image should be added")),
		mcp.WithNumber("position", mcp.Description("Input parameter: Defines image’s position in the list")),
		mcp.WithString("store_id", mcp.Description("Input parameter: Store Id")),
		mcp.WithString("type", mcp.Required(), mcp.Description("Input parameter: Defines image's types that are specified by comma-separated list")),
		mcp.WithString("image_name", mcp.Required(), mcp.Description("Input parameter: Defines image's name")),
		mcp.WithString("label", mcp.Description("Input parameter: Defines alternative text that has to be attached to the picture")),
		mcp.WithString("mime", mcp.Description("Input parameter: Mime type of image http://en.wikipedia.org/wiki/Internet_media_type.")),
		mcp.WithNumber("product_variant_id", mcp.Description("Input parameter: Defines product's variants specified by variant id")),
		mcp.WithString("url", mcp.Description("Input parameter: Defines URL of the image that has to be added")),
		mcp.WithString("variant_ids", mcp.Description("Input parameter: Defines product's variants ids")),
		mcp.WithString("content", mcp.Description("Input parameter: Content(body) encoded in base64 of image file")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ProductimageaddHandler(cfg),
	}
}
