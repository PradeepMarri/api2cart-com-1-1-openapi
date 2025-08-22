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

func ProductvariantupdateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.ProductVariantUpdate
		
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
		url := fmt.Sprintf("%s/product.variant.update.json", cfg.BaseURL)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyBytes))
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

func CreateProductvariantupdateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_product_variant_update_json",
		mcp.WithDescription("Update variant."),
		mcp.WithBoolean("reindex", mcp.Description("Input parameter: Is reindex required")),
		mcp.WithString("warehouse_id", mcp.Description("Input parameter: This parameter is used for selecting a warehouse where you need to set/modify a product quantity.")),
		mcp.WithString("height", mcp.Description("Input parameter: Defines product's height")),
		mcp.WithString("cost_price", mcp.Description("Input parameter: Defines new product's cost price")),
		mcp.WithString("meta_keywords", mcp.Description("Input parameter: Defines unique meta keywords for each entity")),
		mcp.WithString("short_description", mcp.Description("Input parameter: Defines short description")),
		mcp.WithString("quantity", mcp.Description("Input parameter: Defines new products' variants quantity")),
		mcp.WithString("backorder_status", mcp.Description("Input parameter: Set backorder status")),
		mcp.WithString("length", mcp.Description("Input parameter: Defines product's length")),
		mcp.WithString("status", mcp.Description("Input parameter: Defines product variant's status")),
		mcp.WithString("meta_description", mcp.Description("Input parameter: Defines unique meta description of a entity")),
		mcp.WithString("sprice_expire", mcp.Description("Input parameter: Defines the term of special price offer duration")),
		mcp.WithString("barcode", mcp.Description("Input parameter: A barcode is a unique code composed of numbers used as a product identifier.")),
		mcp.WithString("visible", mcp.Description("Input parameter: Set visibility status")),
		mcp.WithString("meta_title", mcp.Description("Input parameter: Defines unique meta title for each entity")),
		mcp.WithString("reduce_quantity", mcp.Description("Input parameter: Defines the decrement changes in product quantity")),
		mcp.WithBoolean("available_for_sale", mcp.Description("Input parameter: Specifies the set of visible/invisible product's variants for sale")),
		mcp.WithString("lang_id", mcp.Description("Input parameter: Language id")),
		mcp.WithString("reserve_quantity", mcp.Description("Input parameter: This parameter allows to reserve/unreserve product variants quantity.")),
		mcp.WithString("sku", mcp.Description("Input parameter: Defines new product's variant sku")),
		mcp.WithString("price", mcp.Description("Input parameter: Defines new product's variant price")),
		mcp.WithBoolean("manage_stock", mcp.Description("Input parameter: Defines inventory tracking for product variant")),
		mcp.WithString("harmonized_system_code", mcp.Description("Input parameter: Harmonized System Code. An HSC is a 6-digit identifier that allows participating countries to classify traded goods on a common basis for customs purposes")),
		mcp.WithArray("options", mcp.Description("Input parameter: Defines variant's options list")),
		mcp.WithString("special_price", mcp.Description("Input parameter: Defines new product's variant special price")),
		mcp.WithString("weight", mcp.Description("Input parameter: Weight")),
		mcp.WithString("old_price", mcp.Description("Input parameter: Defines product's old price")),
		mcp.WithString("description", mcp.Description("Input parameter: Specifies variant's description")),
		mcp.WithString("gtin", mcp.Description("Input parameter: Global Trade Item Number. An GTIN is an identifier for trade items.")),
		mcp.WithString("id", mcp.Description("Input parameter: Defines variant update specified by variant id")),
		mcp.WithString("retail_price", mcp.Description("Input parameter: Defines new product's retail price")),
		mcp.WithString("model", mcp.Description("Input parameter: Specifies variant's model that has to be added")),
		mcp.WithBoolean("clear_cache", mcp.Description("Input parameter: Is cache clear required")),
		mcp.WithString("sprice_create", mcp.Description("Input parameter: Defines the date of special price creation")),
		mcp.WithString("increase_quantity", mcp.Description("Input parameter: Defines the incremental changes in product quantity")),
		mcp.WithBoolean("taxable", mcp.Description("Input parameter: Specifies whether a tax is charged")),
		mcp.WithString("country_of_origin", mcp.Description("Input parameter: The country where the inventory item was made")),
		mcp.WithString("product_id", mcp.Description("Input parameter: Defines product's id where the variant has to be updated")),
		mcp.WithString("store_id", mcp.Description("Input parameter: Defines store id where the variant should be found")),
		mcp.WithBoolean("in_stock", mcp.Description("Input parameter: Set stock status")),
		mcp.WithString("width", mcp.Description("Input parameter: Defines product's width")),
		mcp.WithString("name", mcp.Description("Input parameter: Defines variant's name that has to be updated")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ProductvariantupdateHandler(cfg),
	}
}
