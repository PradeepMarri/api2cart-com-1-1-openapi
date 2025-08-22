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

func ProductvariantaddHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.ProductVariantAdd
		
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
		url := fmt.Sprintf("%s/product.variant.add.json", cfg.BaseURL)
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

func CreateProductvariantaddTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_product_variant_add_json",
		mcp.WithDescription("Add variant to product."),
		mcp.WithString("warehouse_id", mcp.Description("Input parameter: This parameter is used for selecting a warehouse where you need to set/modify a product quantity.")),
		mcp.WithString("meta_keywords", mcp.Description("Input parameter: Defines unique meta keywords for each entity")),
		mcp.WithString("meta_title", mcp.Description("Input parameter: Defines unique meta title for each entity")),
		mcp.WithString("created_at", mcp.Description("Input parameter: Defines the date of entity creation")),
		mcp.WithString("country_of_origin", mcp.Description("Input parameter: The country where the inventory item was made")),
		mcp.WithString("short_description", mcp.Description("Input parameter: Defines short description")),
		mcp.WithBoolean("available_for_view", mcp.Description("Input parameter: Specifies the set of visible/invisible product's variants for users")),
		mcp.WithBoolean("available_for_sale", mcp.Description("Input parameter: Specifies the set of visible/invisible product's variants for sale")),
		mcp.WithString("description", mcp.Description("Input parameter: Specifies variant's description")),
		mcp.WithString("harmonized_system_code", mcp.Description("Input parameter: Harmonized System Code. An HSC is a 6-digit identifier that allows participating countries to classify traded goods on a common basis for customs purposes")),
		mcp.WithString("height", mcp.Description("Input parameter: Defines product's height")),
		mcp.WithString("tax_class_id", mcp.Description("Input parameter: Defines tax classes where entity has to be added")),
		mcp.WithBoolean("taxable", mcp.Description("Input parameter: Specifies whether a tax is charged")),
		mcp.WithString("special_price", mcp.Description("Input parameter: Specifies variant's model that has to be added")),
		mcp.WithString("price", mcp.Description("Input parameter: Defines new product's variant price")),
		mcp.WithString("cost_price", mcp.Description("Input parameter: Defines new product's cost price")),
		mcp.WithString("width", mcp.Description("Input parameter: Defines product's width")),
		mcp.WithString("sku", mcp.Description("Input parameter: Defines variant's sku that has to be added")),
		mcp.WithString("sprice_create", mcp.Description("Input parameter: Defines the date of special price creation")),
		mcp.WithArray("attributes", mcp.Description("Input parameter: Defines variant's attributes list")),
		mcp.WithString("quantity", mcp.Description("Input parameter: Defines product variant's quantity that has to be added")),
		mcp.WithString("product_id", mcp.Description("Input parameter: Defines product's id where the variant has to be added")),
		mcp.WithString("name", mcp.Description("Input parameter: Defines variant's name that has to be added")),
		mcp.WithString("url", mcp.Description("Input parameter: Defines unique product variant's URL")),
		mcp.WithString("sprice_modified", mcp.Description("Input parameter: Defines the date of special price modification")),
		mcp.WithString("weight_unit", mcp.Description("Input parameter: Weight Unit")),
		mcp.WithString("store_id", mcp.Description("Input parameter: Add variants specified by store id")),
		mcp.WithString("sprice_expire", mcp.Description("Input parameter: Defines the term of special price offer duration")),
		mcp.WithString("barcode", mcp.Description("Input parameter: A barcode is a unique code composed of numbers used as a product identifier.")),
		mcp.WithString("length", mcp.Description("Input parameter: Defines product's length")),
		mcp.WithString("meta_description", mcp.Description("Input parameter: Defines unique meta description of a entity")),
		mcp.WithString("lang_id", mcp.Description("Input parameter: Language id")),
		mcp.WithString("weight", mcp.Description("Input parameter: Weight")),
		mcp.WithBoolean("manage_stock", mcp.Description("Input parameter: Defines inventory tracking for product variant")),
		mcp.WithString("manufacturer", mcp.Description("Input parameter: Specifies the product variant's manufacturer")),
		mcp.WithBoolean("clear_cache", mcp.Description("Input parameter: Is cache clear required")),
		mcp.WithString("model", mcp.Required(), mcp.Description("Input parameter: Specifies variant's model that has to be added")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ProductvariantaddHandler(cfg),
	}
}
