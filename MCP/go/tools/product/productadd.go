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

func ProductaddHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.ProductAdd
		
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
		url := fmt.Sprintf("%s/product.add.json", cfg.BaseURL)
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

func CreateProductaddTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_product_add_json",
		mcp.WithDescription("Add new product to store."),
		mcp.WithString("seo_url", mcp.Description("Input parameter: Defines unique URL for SEO")),
		mcp.WithBoolean("return_accepted", mcp.Description("Input parameter: Indicates whether the seller allows the buyer to return the item.")),
		mcp.WithString("meta_keywords", mcp.Description("Input parameter: Defines unique meta keywords for each entity")),
		mcp.WithString("length", mcp.Description("Input parameter: Defines product's length")),
		mcp.WithObject("seller_profiles", mcp.Description("Input parameter: If the seller is subscribed to \"Business Policies\", use the seller_profiles instead of the shipping_details, payment_methods and return_accepted params.<hr><div style=\"font-style:normal\">Param structure:<div style=\"margin-left: 2%;\"><code style=\"padding:0; background-color:#ffffff;font-size:85%;font-family:monospace;\">seller_profiles[<b>shipping_profile_id</b>] = integer</br>seller_profiles[<b>payment_profile_id</b>] = integer</br>seller_profiles[<b>return_profile_id</b>] = integer</br></code></div></div>")),
		mcp.WithString("image_url", mcp.Description("Input parameter: Image Url")),
		mcp.WithString("meta_description", mcp.Description("Input parameter: Defines unique meta description of a entity")),
		mcp.WithString("short_description", mcp.Description("Input parameter: Defines short description")),
		mcp.WithString("price", mcp.Required(), mcp.Description("Input parameter: Defines product's price that has to be added")),
		mcp.WithString("quantity", mcp.Description("Input parameter: Defines product's quantity that has to be added")),
		mcp.WithString("type", mcp.Description("Input parameter: Defines product's type")),
		mcp.WithArray("shipping_details", mcp.Description("Input parameter: The shipping details, including flat and calculated shipping costs and shipping insurance costs. Look at cart.info method response for allowed values.<hr><div style=\"font-style:normal\">Param structure:<div style=\"margin-left: 2%;\"><code style=\"padding:0; background-color:#ffffff;font-size:85%;font-family:monospace;\">shipping_details[0][<b>shipping_type</b>] = string </br>shipping_details[0][<b>shipping_service</b>] = string</br>shipping_details[0][<b>shipping_cost</b>] = decimal</br>shipping_details[1][<b>shipping_type</b>] = string </br>shipping_details[1][<b>shipping_service</b>] = string</br>shipping_details[1][<b>shipping_cost</b>] = decimal</br></code></div></div>")),
		mcp.WithString("url", mcp.Description("Input parameter: Defines unique product's URL")),
		mcp.WithArray("specifics", mcp.Description("Input parameter: An array of Item Specific Name/Value pairs used by the seller to provide descriptive details of an item in a structured manner.\n        <hr>\n        <div style=\"font-style:normal\">Param structure:\n          <div style=\"margin-left: 2%;\">\n            <code style=\"padding:0; background-color:#ffffff;font-size:85%;font-family:monospace;\">\n              specifics[int][<b>name</b>] = string</br>\n              specifics[int][<b>value</b>] = string</br>\n            </code>\n          </div>\n        </div>")),
		mcp.WithString("harmonized_system_code", mcp.Description("Input parameter: Harmonized System Code. An HSC is a 6-digit identifier that allows participating countries to classify traded goods on a common basis for customs purposes")),
		mcp.WithString("image_name", mcp.Description("Input parameter: Defines image's name")),
		mcp.WithString("listing_duration", mcp.Description("Input parameter: Describes the number of days the seller wants the listing to be active. Look at cart.info method response for allowed values.")),
		mcp.WithString("sprice_modified", mcp.Description("Input parameter: Defines the date of special price modification")),
		mcp.WithString("height", mcp.Description("Input parameter: Defines product's height")),
		mcp.WithString("search_keywords", mcp.Description("Input parameter: Defines unique search keywords")),
		mcp.WithString("mpn", mcp.Description("Input parameter: Manufacturer Part Number. A MPN is an identifier of a particular part design or material used.")),
		mcp.WithString("product_class", mcp.Description("Input parameter: A categorization for the product")),
		mcp.WithString("country_of_origin", mcp.Description("Input parameter: The country where the inventory item was made")),
		mcp.WithString("avail_from", mcp.Description("Input parameter: Allows to schedule a time in the future that the item becomes available. The value should be greater than the current date and time.")),
		mcp.WithString("listing_type", mcp.Description("Input parameter: Indicates the selling format of the marketplace listing.")),
		mcp.WithNumber("viewed_count", mcp.Description("Input parameter: Specifies the number of product's reviews")),
		mcp.WithArray("payment_methods", mcp.Description("Input parameter: Identifies the payment method (such as PayPal) that the seller will accept when the buyer pays for the item. Look at cart.info method response for allowed values.<hr><div style=\"font-style:normal\">Param structure:<div style=\"margin-left: 2%;\"><code style=\"padding:0; background-color:#ffffff;font-size:85%;font-family:monospace;\">payment_methods[0] = string</br>payment_methods[1] = string</br></code></div></div>")),
		mcp.WithArray("package_details", mcp.Description("Input parameter: If the seller is subscribed to \"Business Policies\", use the seller_profiles instead of the shipping_details, payment_methods and return_accepted params.<hr><div style=\"font-style:normal\">Param structure:<div style=\"margin-left: 2%;\"><code style=\"padding:0; background-color:#ffffff;font-size:85%;font-family:monospace;\">package_details[<b>measure_unit</b>] = string</br> Allowed measure_unit values: [English or Metric] </br> Default: Metric</br>package_details[<b>weigh_unit</b>] = string</br> Allowed weigh_unit values: [kg, g, lbs, oz]</br>package_details[<b>package_depth</b>] = decimal</br>package_details[<b>package_length</b>] = decimal</br>package_details[<b>package_width</b>] = decimal</br>package_details[<b>weight_major</b>] = decimal</br>package_details[<b>weight_minor</b>] = decimal</br>package_details[<b>shipping_package</b>] = string</br> See cart.info method, param `eBay_shipping_package_details`</code></div></div>")),
		mcp.WithString("tags", mcp.Description("Input parameter: Product tags")),
		mcp.WithString("old_price", mcp.Description("Input parameter: Defines product's old price")),
		mcp.WithString("categories_ids", mcp.Description("Input parameter: Defines product add that is specified by comma-separated categories id")),
		mcp.WithArray("sales_tax", mcp.Description("Input parameter: Percent of an item's price to be charged as the sales tax for the order. Look at cart.info method response for allowed values.<hr><div style=\"font-style:normal\">Param structure:<div style=\"margin-left: 2%;\"><code style=\"padding:0; background-color:#ffffff;font-size:85%;font-family:monospace;\">sales_tax[<b>tax_percent</b>] = decimal (##.###)</br>sales_tax[<b>tax_state</b>] = string</br>sales_tax[<b>shipping_inc_in_tax</b>] = bool</br></code></div></div>")),
		mcp.WithString("weight_unit", mcp.Description("Input parameter: Weight Unit")),
		mcp.WithString("description", mcp.Required(), mcp.Description("Input parameter: Defines product's description that has to be added")),
		mcp.WithString("brand_name", mcp.Description("Input parameter: Retrieves brands specified by brand name")),
		mcp.WithArray("group_prices", mcp.Description("Input parameter: Defines product's group prices")),
		mcp.WithString("upc", mcp.Description("Input parameter: Universal Product Code. A UPC (UPC-A) is a commonly used identifer for many different products.")),
		mcp.WithString("meta_title", mcp.Description("Input parameter: Defines unique meta title for each entity")),
		mcp.WithBoolean("available_for_sale", mcp.Description("Input parameter: Specifies the set of visible/invisible products for sale")),
		mcp.WithString("model", mcp.Required(), mcp.Description("Input parameter: Defines product's model that has to be added")),
		mcp.WithString("width", mcp.Description("Input parameter: Defines product's width")),
		mcp.WithString("lang_id", mcp.Description("Input parameter: Language id")),
		mcp.WithBoolean("taxable", mcp.Description("Input parameter: Specifies whether a tax is charged")),
		mcp.WithString("sprice_expire", mcp.Description("Input parameter: Defines the term of special price offer duration")),
		mcp.WithString("ean", mcp.Description("Input parameter: European Article Number. An EAN is a unique 8 or 13-digit identifier that many industries (such as book publishers) use to identify products.")),
		mcp.WithString("marketplace_item_properties", mcp.Description("Input parameter: String containing the JSON representation of the supplied data")),
		mcp.WithString("status", mcp.Description("Input parameter: Defines product's status")),
		mcp.WithNumber("ordered_count", mcp.Description("Input parameter: Defines how many times the product was ordered")),
		mcp.WithString("warehouse_id", mcp.Description("Input parameter: This parameter is used for selecting a warehouse where you need to set/modify a product quantity.")),
		mcp.WithArray("best_offer", mcp.Description("Input parameter: The price at which Best Offers are automatically accepted.<hr><div style=\"font-style:normal\">Param structure:<div style=\"margin-left: 2%;\"><code style=\"padding:0; background-color:#ffffff;font-size:85%;font-family:monospace;\">best_offer[<b>minimum_offer_price</b>] = decimal</br>best_offer[<b>auto_accept_price</b>] = decimal</br></code></div></div>")),
		mcp.WithString("name", mcp.Required(), mcp.Description("Input parameter: Defines product's name that has to be added")),
		mcp.WithBoolean("clear_cache", mcp.Description("Input parameter: Is cache clear required")),
		mcp.WithArray("tier_prices", mcp.Description("Input parameter: Defines product's tier prices")),
		mcp.WithString("stores_ids", mcp.Description("Input parameter: Assign product to the stores that is specified by comma-separated stores' id")),
		mcp.WithBoolean("downloadable", mcp.Description("Input parameter: Defines whether the product is downloadable")),
		mcp.WithString("sku", mcp.Description("Input parameter: Defines product's sku that has to be added")),
		mcp.WithString("barcode", mcp.Description("Input parameter: A barcode is a unique code composed of numbers used as a product identifier.")),
		mcp.WithString("visible", mcp.Description("Input parameter: Set visibility status")),
		mcp.WithString("wholesale_price", mcp.Description("Input parameter: Defines product's sale price")),
		mcp.WithString("isbn", mcp.Description("Input parameter: International Standard Book Number. An ISBN is a unique identifier for books.")),
		mcp.WithString("store_id", mcp.Description("Input parameter: Store Id")),
		mcp.WithString("manufacturer", mcp.Description("Input parameter: Defines product's manufacturer")),
		mcp.WithString("tax_class_id", mcp.Description("Input parameter: Defines tax classes where entity has to be added")),
		mcp.WithBoolean("manage_stock", mcp.Description("Input parameter: Defines inventory tracking for product")),
		mcp.WithString("cost_price", mcp.Description("Input parameter: Defines new product's cost price")),
		mcp.WithString("sprice_create", mcp.Description("Input parameter: Defines the date of special price creation")),
		mcp.WithString("backorder_status", mcp.Description("Input parameter: Set backorder status")),
		mcp.WithArray("files", mcp.Description("Input parameter: File Url")),
		mcp.WithString("gtin", mcp.Description("Input parameter: Global Trade Item Number. An GTIN is an identifier for trade items.")),
		mcp.WithBoolean("available_for_view", mcp.Description("Input parameter: Specifies the set of visible/invisible products for users")),
		mcp.WithString("weight", mcp.Description("Input parameter: Weight")),
		mcp.WithString("attribute_name", mcp.Description("Input parameter: Defines product’s attribute name separated with a comma in Magento")),
		mcp.WithString("created_at", mcp.Description("Input parameter: Defines the date of entity creation")),
		mcp.WithNumber("shipping_template_id", mcp.Description("Input parameter: The numeric ID of the shipping template associated with the products in Etsy.")),
		mcp.WithString("special_price", mcp.Description("Input parameter: Defines product's model that has to be added")),
		mcp.WithString("category_id", mcp.Description("Input parameter: Defines product add that is specified by category id")),
		mcp.WithString("paypal_email", mcp.Description("Input parameter: Valid PayPal email address for the PayPal account that the seller will use if they offer PayPal as a payment method for the listing.")),
		mcp.WithString("attribute_set_name", mcp.Description("Input parameter: Defines product’s attribute set name in Magento")),
		mcp.WithString("condition", mcp.Description("Input parameter: The human-readable label for the condition (e.g., \"New\").")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ProductaddHandler(cfg),
	}
}
