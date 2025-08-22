package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ProductAdd represents the ProductAdd schema from the OpenAPI specification
type ProductAdd struct {
	Isbn string `json:"isbn,omitempty"` // International Standard Book Number. An ISBN is a unique identifier for books.
	Store_id string `json:"store_id,omitempty"` // Store Id
	Manufacturer string `json:"manufacturer,omitempty"` // Defines product's manufacturer
	Tax_class_id string `json:"tax_class_id,omitempty"` // Defines tax classes where entity has to be added
	Manage_stock bool `json:"manage_stock,omitempty"` // Defines inventory tracking for product
	Cost_price float64 `json:"cost_price,omitempty"` // Defines new product's cost price
	Sprice_create string `json:"sprice_create,omitempty"` // Defines the date of special price creation
	Backorder_status string `json:"backorder_status,omitempty"` // Set backorder status
	Files []map[string]interface{} `json:"files,omitempty"` // File Url
	Gtin string `json:"gtin,omitempty"` // Global Trade Item Number. An GTIN is an identifier for trade items.
	Available_for_view bool `json:"available_for_view,omitempty"` // Specifies the set of visible/invisible products for users
	Weight float64 `json:"weight,omitempty"` // Weight
	Attribute_name string `json:"attribute_name,omitempty"` // Defines product’s attribute name separated with a comma in Magento
	Created_at string `json:"created_at,omitempty"` // Defines the date of entity creation
	Shipping_template_id int `json:"shipping_template_id,omitempty"` // The numeric ID of the shipping template associated with the products in Etsy.
	Special_price float64 `json:"special_price,omitempty"` // Defines product's model that has to be added
	Category_id string `json:"category_id,omitempty"` // Defines product add that is specified by category id
	Paypal_email string `json:"paypal_email,omitempty"` // Valid PayPal email address for the PayPal account that the seller will use if they offer PayPal as a payment method for the listing.
	Attribute_set_name string `json:"attribute_set_name,omitempty"` // Defines product’s attribute set name in Magento
	Condition string `json:"condition,omitempty"` // The human-readable label for the condition (e.g., "New").
	Seo_url string `json:"seo_url,omitempty"` // Defines unique URL for SEO
	Return_accepted bool `json:"return_accepted,omitempty"` // Indicates whether the seller allows the buyer to return the item.
	Meta_keywords string `json:"meta_keywords,omitempty"` // Defines unique meta keywords for each entity
	Length float64 `json:"length,omitempty"` // Defines product's length
	Seller_profiles map[string]interface{} `json:"seller_profiles,omitempty"` // If the seller is subscribed to "Business Policies", use the seller_profiles instead of the shipping_details, payment_methods and return_accepted params.<hr><div style="font-style:normal">Param structure:<div style="margin-left: 2%;"><code style="padding:0; background-color:#ffffff;font-size:85%;font-family:monospace;">seller_profiles[<b>shipping_profile_id</b>] = integer</br>seller_profiles[<b>payment_profile_id</b>] = integer</br>seller_profiles[<b>return_profile_id</b>] = integer</br></code></div></div>
	Image_url string `json:"image_url,omitempty"` // Image Url
	Meta_description string `json:"meta_description,omitempty"` // Defines unique meta description of a entity
	Short_description string `json:"short_description,omitempty"` // Defines short description
	Price float64 `json:"price"` // Defines product's price that has to be added
	Quantity float64 `json:"quantity,omitempty"` // Defines product's quantity that has to be added
	TypeField string `json:"type,omitempty"` // Defines product's type
	Shipping_details []map[string]interface{} `json:"shipping_details,omitempty"` // The shipping details, including flat and calculated shipping costs and shipping insurance costs. Look at cart.info method response for allowed values.<hr><div style="font-style:normal">Param structure:<div style="margin-left: 2%;"><code style="padding:0; background-color:#ffffff;font-size:85%;font-family:monospace;">shipping_details[0][<b>shipping_type</b>] = string </br>shipping_details[0][<b>shipping_service</b>] = string</br>shipping_details[0][<b>shipping_cost</b>] = decimal</br>shipping_details[1][<b>shipping_type</b>] = string </br>shipping_details[1][<b>shipping_service</b>] = string</br>shipping_details[1][<b>shipping_cost</b>] = decimal</br></code></div></div>
	Url string `json:"url,omitempty"` // Defines unique product's URL
	Specifics []string `json:"specifics,omitempty"` // An array of Item Specific Name/Value pairs used by the seller to provide descriptive details of an item in a structured manner. <hr> <div style="font-style:normal">Param structure: <div style="margin-left: 2%;"> <code style="padding:0; background-color:#ffffff;font-size:85%;font-family:monospace;"> specifics[int][<b>name</b>] = string</br> specifics[int][<b>value</b>] = string</br> </code> </div> </div>
	Harmonized_system_code string `json:"harmonized_system_code,omitempty"` // Harmonized System Code. An HSC is a 6-digit identifier that allows participating countries to classify traded goods on a common basis for customs purposes
	Image_name string `json:"image_name,omitempty"` // Defines image's name
	Listing_duration string `json:"listing_duration,omitempty"` // Describes the number of days the seller wants the listing to be active. Look at cart.info method response for allowed values.
	Sprice_modified string `json:"sprice_modified,omitempty"` // Defines the date of special price modification
	Height float64 `json:"height,omitempty"` // Defines product's height
	Search_keywords string `json:"search_keywords,omitempty"` // Defines unique search keywords
	Mpn string `json:"mpn,omitempty"` // Manufacturer Part Number. A MPN is an identifier of a particular part design or material used.
	Product_class string `json:"product_class,omitempty"` // A categorization for the product
	Country_of_origin string `json:"country_of_origin,omitempty"` // The country where the inventory item was made
	Avail_from string `json:"avail_from,omitempty"` // Allows to schedule a time in the future that the item becomes available. The value should be greater than the current date and time.
	Listing_type string `json:"listing_type,omitempty"` // Indicates the selling format of the marketplace listing.
	Viewed_count int `json:"viewed_count,omitempty"` // Specifies the number of product's reviews
	Payment_methods []string `json:"payment_methods,omitempty"` // Identifies the payment method (such as PayPal) that the seller will accept when the buyer pays for the item. Look at cart.info method response for allowed values.<hr><div style="font-style:normal">Param structure:<div style="margin-left: 2%;"><code style="padding:0; background-color:#ffffff;font-size:85%;font-family:monospace;">payment_methods[0] = string</br>payment_methods[1] = string</br></code></div></div>
	Package_details []string `json:"package_details,omitempty"` // If the seller is subscribed to "Business Policies", use the seller_profiles instead of the shipping_details, payment_methods and return_accepted params.<hr><div style="font-style:normal">Param structure:<div style="margin-left: 2%;"><code style="padding:0; background-color:#ffffff;font-size:85%;font-family:monospace;">package_details[<b>measure_unit</b>] = string</br> Allowed measure_unit values: [English or Metric] </br> Default: Metric</br>package_details[<b>weigh_unit</b>] = string</br> Allowed weigh_unit values: [kg, g, lbs, oz]</br>package_details[<b>package_depth</b>] = decimal</br>package_details[<b>package_length</b>] = decimal</br>package_details[<b>package_width</b>] = decimal</br>package_details[<b>weight_major</b>] = decimal</br>package_details[<b>weight_minor</b>] = decimal</br>package_details[<b>shipping_package</b>] = string</br> See cart.info method, param `eBay_shipping_package_details`</code></div></div>
	Tags string `json:"tags,omitempty"` // Product tags
	Old_price float64 `json:"old_price,omitempty"` // Defines product's old price
	Categories_ids string `json:"categories_ids,omitempty"` // Defines product add that is specified by comma-separated categories id
	Sales_tax []string `json:"sales_tax,omitempty"` // Percent of an item's price to be charged as the sales tax for the order. Look at cart.info method response for allowed values.<hr><div style="font-style:normal">Param structure:<div style="margin-left: 2%;"><code style="padding:0; background-color:#ffffff;font-size:85%;font-family:monospace;">sales_tax[<b>tax_percent</b>] = decimal (##.###)</br>sales_tax[<b>tax_state</b>] = string</br>sales_tax[<b>shipping_inc_in_tax</b>] = bool</br></code></div></div>
	Weight_unit string `json:"weight_unit,omitempty"` // Weight Unit
	Description string `json:"description"` // Defines product's description that has to be added
	Brand_name string `json:"brand_name,omitempty"` // Retrieves brands specified by brand name
	Group_prices []map[string]interface{} `json:"group_prices,omitempty"` // Defines product's group prices
	Upc string `json:"upc,omitempty"` // Universal Product Code. A UPC (UPC-A) is a commonly used identifer for many different products.
	Meta_title string `json:"meta_title,omitempty"` // Defines unique meta title for each entity
	Available_for_sale bool `json:"available_for_sale,omitempty"` // Specifies the set of visible/invisible products for sale
	Model string `json:"model"` // Defines product's model that has to be added
	Width float64 `json:"width,omitempty"` // Defines product's width
	Lang_id string `json:"lang_id,omitempty"` // Language id
	Taxable bool `json:"taxable,omitempty"` // Specifies whether a tax is charged
	Sprice_expire string `json:"sprice_expire,omitempty"` // Defines the term of special price offer duration
	Ean string `json:"ean,omitempty"` // European Article Number. An EAN is a unique 8 or 13-digit identifier that many industries (such as book publishers) use to identify products.
	Marketplace_item_properties string `json:"marketplace_item_properties,omitempty"` // String containing the JSON representation of the supplied data
	Status string `json:"status,omitempty"` // Defines product's status
	Ordered_count int `json:"ordered_count,omitempty"` // Defines how many times the product was ordered
	Warehouse_id string `json:"warehouse_id,omitempty"` // This parameter is used for selecting a warehouse where you need to set/modify a product quantity.
	Best_offer []string `json:"best_offer,omitempty"` // The price at which Best Offers are automatically accepted.<hr><div style="font-style:normal">Param structure:<div style="margin-left: 2%;"><code style="padding:0; background-color:#ffffff;font-size:85%;font-family:monospace;">best_offer[<b>minimum_offer_price</b>] = decimal</br>best_offer[<b>auto_accept_price</b>] = decimal</br></code></div></div>
	Name string `json:"name"` // Defines product's name that has to be added
	Clear_cache bool `json:"clear_cache,omitempty"` // Is cache clear required
	Tier_prices []map[string]interface{} `json:"tier_prices,omitempty"` // Defines product's tier prices
	Stores_ids string `json:"stores_ids,omitempty"` // Assign product to the stores that is specified by comma-separated stores' id
	Downloadable bool `json:"downloadable,omitempty"` // Defines whether the product is downloadable
	Sku string `json:"sku,omitempty"` // Defines product's sku that has to be added
	Barcode string `json:"barcode,omitempty"` // A barcode is a unique code composed of numbers used as a product identifier.
	Visible string `json:"visible,omitempty"` // Set visibility status
	Wholesale_price float64 `json:"wholesale_price,omitempty"` // Defines product's sale price
}

// Image represents the Image schema from the OpenAPI specification
type Image struct {
	Alt string `json:"alt,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Create_at A2CDateTime `json:"create_at,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Id string `json:"id,omitempty"`
	Modified_at A2CDateTime `json:"modified_at,omitempty"`
	Size int `json:"size,omitempty"`
	TypeField string `json:"type,omitempty"`
	Avail bool `json:"avail,omitempty"`
	Http_path string `json:"http_path,omitempty"`
	Mime_type string `json:"mime-type,omitempty"`
	Sort_order int `json:"sort_order,omitempty"`
	File_name string `json:"file_name,omitempty"`
}

// OrderShipmentTrackingAdd represents the OrderShipmentTrackingAdd schema from the OpenAPI specification
type OrderShipmentTrackingAdd struct {
	Tracking_link string `json:"tracking_link,omitempty"` // Defines custom tracking link
	Tracking_number string `json:"tracking_number"` // Defines tracking number
	Tracking_provider string `json:"tracking_provider,omitempty"` // Defines name of the company which provides shipment tracking
	Carrier_id string `json:"carrier_id,omitempty"` // Defines tracking carrier id
	Order_id string `json:"order_id,omitempty"` // Defines the order id
	Send_notifications bool `json:"send_notifications,omitempty"` // Send notifications to customer after tracking was created
	Shipment_id string `json:"shipment_id"` // Shipment id indicates the number of delivery
	Store_id string `json:"store_id,omitempty"` // Store Id
}

// ProductOptionItem represents the ProductOptionItem schema from the OpenAPI specification
type ProductOptionItem struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Sku string `json:"sku,omitempty"`
	Sort_order int `json:"sort_order,omitempty"`
	Type_price string `json:"type_price,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Price string `json:"price,omitempty"`
	Product_option_item_id string `json:"product_option_item_id,omitempty"`
	Quantity int `json:"quantity,omitempty"`
	Weight string `json:"weight,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
}

// BasketItemOption represents the BasketItemOption schema from the OpenAPI specification
type BasketItemOption struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Used_in_combination bool `json:"used_in_combination,omitempty"`
	Value string `json:"value,omitempty"`
	Value_id string `json:"value_id,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
}

// Order represents the Order schema from the OpenAPI specification
type Order struct {
	Warehouses_ids []string `json:"warehouses_ids,omitempty"`
	Create_at A2CDateTime `json:"create_at,omitempty"`
	Shipping_methods []OrderShippingMethod `json:"shipping_methods,omitempty"`
	Customer BaseCustomer `json:"customer,omitempty"`
	Comment string `json:"comment,omitempty"`
	Shipping_address CustomerAddress `json:"shipping_address,omitempty"`
	Basket_id string `json:"basket_id,omitempty"`
	Bundles []OrderItem `json:"bundles,omitempty"`
	Currency Currency `json:"currency,omitempty"`
	Order_id string `json:"order_id,omitempty"`
	Order_products []OrderItem `json:"order_products,omitempty"`
	Shipping_method OrderShippingMethod `json:"shipping_method,omitempty"`
	Discounts []OrderTotalsNewDiscount `json:"discounts,omitempty"`
	Id string `json:"id,omitempty"`
	Order_details_url string `json:"order_details_url,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Payment_method OrderPaymentMethod `json:"payment_method,omitempty"`
	Status OrderStatus `json:"status,omitempty"`
	Store_id string `json:"store_id,omitempty"`
	Total OrderTotal `json:"total,omitempty"`
	Gift_message string `json:"gift_message,omitempty"`
	Channel_id string `json:"channel_id,omitempty"`
	Refunds []OrderRefund `json:"refunds,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Finished_time A2CDateTime `json:"finished_time,omitempty"`
	Billing_address CustomerAddress `json:"billing_address,omitempty"`
	Modified_at A2CDateTime `json:"modified_at,omitempty"`
	Totals OrderTotals `json:"totals,omitempty"`
}

// OrderShipmentAdd represents the OrderShipmentAdd schema from the OpenAPI specification
type OrderShipmentAdd struct {
	Enable_cache bool `json:"enable_cache,omitempty"` // If the value is 'true' and order exist in our cache, we will use order.info from cache to prepare shipment items.
	Is_shipped bool `json:"is_shipped,omitempty"` // Defines shipment's status
	Shipment_provider string `json:"shipment_provider,omitempty"` // Defines company name that provide tracking of shipment
	Tracking_link string `json:"tracking_link,omitempty"` // Defines custom tracking link
	Tracking_numbers []map[string]interface{} `json:"tracking_numbers,omitempty"` // Defines shipment's tracking numbers that have to be added</br> How set tracking numbers to appropriate carrier:<ul><li>tracking_numbers[]=a2c.demo1,a2c.demo2 - set default carrier</li><li>tracking_numbers[<b>carrier_id</b>]=a2c.demo - set appropriate carrier</li></ul>To get the list of carriers IDs that are available in your store, use the <a href = "https://api2cart.com/docs/#/cart/CartInfo">cart.info</a > method
	Store_id string `json:"store_id,omitempty"` // Store Id
	Warehouse_id string `json:"warehouse_id,omitempty"` // This parameter is used for selecting a warehouse where you need to set/modify a product quantity.
	Send_notifications bool `json:"send_notifications,omitempty"` // Send notifications to customer after shipment was created
	Shipping_method string `json:"shipping_method,omitempty"` // Define shipping method
	Adjust_stock bool `json:"adjust_stock,omitempty"` // This parameter is used for adjust stock.
	Items []map[string]interface{} `json:"items,omitempty"` // Defines items in the order that will be shipped
	Order_id string `json:"order_id,omitempty"` // Defines the order for which the shipment will be created
}

// Info represents the Info schema from the OpenAPI specification
type Info struct {
	City string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Owner string `json:"owner,omitempty"`
	State_code string `json:"state_code,omitempty"`
	Email string `json:"email,omitempty"`
	Street_address_line_2 string `json:"street_address_line_2,omitempty"`
	Zip_code string `json:"zip_code,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	State string `json:"state,omitempty"`
	Street_address string `json:"street_address,omitempty"`
	Phone string `json:"phone,omitempty"`
}

// OrderPreestimateShipping represents the OrderPreestimateShipping schema from the OpenAPI specification
type OrderPreestimateShipping struct {
	Description string `json:"description,omitempty"`
	Method_code string `json:"method_code,omitempty"`
	Carrier_code string `json:"carrier_code,omitempty"`
	Method_name string `json:"method_name,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Carrier_name string `json:"carrier_name,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Delivery_time string `json:"delivery_time,omitempty"`
	Price float64 `json:"price,omitempty"`
	Price_inc_tax float64 `json:"price_inc_tax,omitempty"`
}

// CouponAction represents the CouponAction schema from the OpenAPI specification
type CouponAction struct {
	Apply_to string `json:"apply_to,omitempty"`
	Scope string `json:"scope,omitempty"`
	Discount_quantity_step int `json:"discount_quantity_step,omitempty"`
	Discounted_quantity float64 `json:"discounted_quantity,omitempty"`
	Include_tax bool `json:"include_tax,omitempty"`
	Amount float64 `json:"amount,omitempty"`
	Conditions []CouponCondition `json:"conditions,omitempty"`
	Currency_code string `json:"currency_code,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Logic_operator string `json:"logic_operator,omitempty"`
	TypeField string `json:"type,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
}

// ModelResponseOrderPreestimateShippingList represents the ModelResponseOrderPreestimateShippingList schema from the OpenAPI specification
type ModelResponseOrderPreestimateShippingList struct {
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
	Result ResponseOrderPreestimateShippingListResult `json:"result,omitempty"`
	Return_code int `json:"return_code,omitempty"`
	Return_message string `json:"return_message,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
}

// CartChannel represents the CartChannel schema from the OpenAPI specification
type CartChannel struct {
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Id string `json:"id,omitempty"`
	Is_enabled bool `json:"is_enabled,omitempty"`
	Name string `json:"name,omitempty"`
	Platform string `json:"platform,omitempty"`
}

// ModelResponseOrderAbandonedList represents the ModelResponseOrderAbandonedList schema from the OpenAPI specification
type ModelResponseOrderAbandonedList struct {
	Result ResponseOrderAbandonedListResult `json:"result,omitempty"`
	Return_code int `json:"return_code,omitempty"`
	Return_message string `json:"return_message,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
}

// CartMetaData represents the CartMetaData schema from the OpenAPI specification
type CartMetaData struct {
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Id string `json:"id,omitempty"`
	Key string `json:"key,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Value string `json:"value,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
}

// ProductAttribute represents the ProductAttribute schema from the OpenAPI specification
type ProductAttribute struct {
	Store_id string `json:"store_id,omitempty"`
	Attribute_group_id string `json:"attribute_group_id,omitempty"`
	Position int `json:"position,omitempty"`
	Required bool `json:"required,omitempty"`
	Lang_id string `json:"lang_id,omitempty"`
	Visible bool `json:"visible,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Name string `json:"name,omitempty"`
	Variant_id string `json:"variant_id,omitempty"`
	Code string `json:"code,omitempty"`
	Product_id string `json:"product_id,omitempty"`
	TypeField string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
	Attribute_id string `json:"attribute_id,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
}

// ModelResponseOrderShipmentList represents the ModelResponseOrderShipmentList schema from the OpenAPI specification
type ModelResponseOrderShipmentList struct {
	Return_message string `json:"return_message,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
	Result ResponseOrderShipmentListResult `json:"result,omitempty"`
	Return_code int `json:"return_code,omitempty"`
}

// ProductVariantPriceAdd represents the ProductVariantPriceAdd schema from the OpenAPI specification
type ProductVariantPriceAdd struct {
	Id string `json:"id,omitempty"` // Defines the variant to which the price has to be added
	Product_id string `json:"product_id,omitempty"` // Product id
	Group_prices []map[string]interface{} `json:"group_prices"` // Defines variants's group prices
}

// CustomerAttribute represents the CustomerAttribute schema from the OpenAPI specification
type CustomerAttribute struct {
	Name string `json:"name,omitempty"`
	TypeField string `json:"type,omitempty"`
	Values []CustomerAttributeValue `json:"values,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Attribute_id string `json:"attribute_id,omitempty"`
	Code string `json:"code,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
}

// ResponseOrderListResult represents the ResponseOrderListResult schema from the OpenAPI specification
type ResponseOrderListResult struct {
	Order []Order `json:"order,omitempty"`
	Orders_count int `json:"orders_count,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
}

// ProductPriceAdd represents the ProductPriceAdd schema from the OpenAPI specification
type ProductPriceAdd struct {
	Group_prices []map[string]interface{} `json:"group_prices,omitempty"` // Defines product's group prices
	Product_id string `json:"product_id,omitempty"` // Defines the product to which the price has to be added
}

// OrderAdd represents the OrderAdd schema from the OpenAPI specification
type OrderAdd struct {
	Store_id string `json:"store_id,omitempty"` // Defines store id where the order should be assigned
	Tax_price float64 `json:"tax_price,omitempty"` // The value of tax cost for order
	Create_invoice bool `json:"create_invoice,omitempty"` // Defines whether the invoice is created automatically along with the order
	Bill_company string `json:"bill_company,omitempty"` // Specifies billing company
	Coupon_discount float64 `json:"coupon_discount,omitempty"` // Specifies order's coupon discount
	Bill_postcode string `json:"bill_postcode"` // Specifies billing postcode
	Transaction_id string `json:"transaction_id,omitempty"` // Payment transaction id
	Admin_comment string `json:"admin_comment,omitempty"` // Specifies admin's order comment
	Date_modified string `json:"date_modified,omitempty"` // Specifies order's modification date
	Shipp_fax string `json:"shipp_fax,omitempty"` // Specifies shipping fax
	Currency string `json:"currency,omitempty"` // Currency code of order
	Bill_last_name string `json:"bill_last_name"` // Specifies billing last name
	Send_notifications bool `json:"send_notifications,omitempty"` // Send notifications to customer after order was created
	Shipp_first_name string `json:"shipp_first_name,omitempty"` // Specifies shipping first name
	Inventory_behaviour string `json:"inventory_behaviour,omitempty"` // The behaviour to use when updating inventory.<hr><div style="font-style:normal">Values description:<div style="margin-left: 2%; padding-top: 2%"><div style="font-size:85%"><b>bypass</b> = Do not claim inventory </br></br><b>decrement_ignoring_policy</b> = Ignore the product's </br> inventory policy and claim amounts</br></br><b>decrement_obeying_policy</b> = Obey the product's </br> inventory policy.</br></br></div></div></div>
	Customer_email string `json:"customer_email"` // Defines the customer specified by email for whom order has to be created
	Send_admin_notifications bool `json:"send_admin_notifications,omitempty"` // Notify admin when new order was created.
	Bill_address_1 string `json:"bill_address_1"` // Specifies first billing address
	Order_shipping_method string `json:"order_shipping_method,omitempty"` // Defines order shipping method
	Order_status string `json:"order_status"` // Defines order status.
	Prices_inc_tax bool `json:"prices_inc_tax,omitempty"` // Indicates whether prices and subtotal includes tax.
	Shipp_address_1 string `json:"shipp_address_1,omitempty"` // Specifies first shipping address
	Bill_country string `json:"bill_country"` // Specifies billing country code
	Customer_fax string `json:"customer_fax,omitempty"` // Specifies customer’s fax
	Date string `json:"date,omitempty"` // Specifies an order creation date in format Y-m-d H:i:s
	External_source string `json:"external_source,omitempty"` // Identifying the system used to generate the order
	Bill_fax string `json:"bill_fax,omitempty"` // Specifies billing fax
	Fulfillment_status string `json:"fulfillment_status,omitempty"` // Create order with fulfillment status
	Total_weight int `json:"total_weight,omitempty"` // Defines the sum of all line item weights in grams for the order
	Shipp_address_2 string `json:"shipp_address_2,omitempty"` // Specifies second address line of a shipping street address
	Customer_first_name string `json:"customer_first_name,omitempty"` // Specifies customer's first name
	Bill_first_name string `json:"bill_first_name"` // Specifies billing first name
	Total_paid float64 `json:"total_paid,omitempty"` // Defines total paid amount for the order
	Order_item []map[string]interface{} `json:"order_item"`
	Customer_last_name string `json:"customer_last_name,omitempty"` // Specifies customer’s last name
	Customer_birthday string `json:"customer_birthday,omitempty"` // Specifies customer’s birthday
	Shipp_company string `json:"shipp_company,omitempty"` // Specifies shipping company
	Channel_id string `json:"channel_id,omitempty"` // Channel ID
	Subtotal_price float64 `json:"subtotal_price,omitempty"` // Total price of all ordered products multiplied by their number, excluding tax, shipping price and discounts
	Tags string `json:"tags,omitempty"` // Order tags
	Shipp_phone string `json:"shipp_phone,omitempty"` // Specifies shipping phone
	Financial_status string `json:"financial_status,omitempty"` // Create order with financial status
	Customer_phone string `json:"customer_phone,omitempty"` // Specifies customer’s phone
	Order_payment_method string `json:"order_payment_method,omitempty"` // Defines order payment method.<br/>Setting order_payment_method on Shopify will also change financial_status field value to 'paid'
	Id string `json:"id,omitempty"` // Defines order's id
	Comment string `json:"comment,omitempty"` // Specifies order comment
	Coupons []string `json:"coupons,omitempty"` // Coupons that will be applied to order
	Shipp_city string `json:"shipp_city,omitempty"` // Specifies shipping city
	Shipp_country string `json:"shipp_country,omitempty"` // Specifies shipping country code
	Shipp_state string `json:"shipp_state,omitempty"` // Specifies shipping state code
	Bill_state string `json:"bill_state"` // Specifies billing state code
	Shipp_postcode string `json:"shipp_postcode,omitempty"` // Specifies shipping postcode
	Bill_city string `json:"bill_city"` // Specifies billing city
	Bill_address_2 string `json:"bill_address_2,omitempty"` // Specifies second billing address
	Order_id string `json:"order_id,omitempty"` // Defines the order id if it is supported by the cart
	Admin_private_comment string `json:"admin_private_comment,omitempty"` // Specifies private admin's order comment
	Shipping_tax float64 `json:"shipping_tax,omitempty"` // Specifies order's shipping price tax
	Shipp_last_name string `json:"shipp_last_name,omitempty"` // Specifies shipping last name
	Clear_cache bool `json:"clear_cache,omitempty"` // Is cache clear required
	Total_price float64 `json:"total_price,omitempty"` // Defines order's total price
	Bill_phone string `json:"bill_phone,omitempty"` // Specifies billing phone
	Discount float64 `json:"discount,omitempty"` // Specifies order's discount
	Gift_certificate_discount float64 `json:"gift_certificate_discount,omitempty"` // Discounts for order with gift certificates
	Date_finished string `json:"date_finished,omitempty"` // Specifies order's finished date
	Note_attributes []map[string]interface{} `json:"note_attributes,omitempty"` // Defines note attributes
	Shipping_price float64 `json:"shipping_price,omitempty"` // Specifies order's shipping price
}

// CartStoreInfo represents the CartStoreInfo schema from the OpenAPI specification
type CartStoreInfo struct {
	Store_owner_info Info `json:"store_owner_info,omitempty"`
	Default_warehouse_id string `json:"default_warehouse_id,omitempty"`
	Multi_store_url string `json:"multi_store_url,omitempty"`
	Active bool `json:"active,omitempty"`
	Language string `json:"language,omitempty"`
	Store_languages []Language `json:"store_languages,omitempty"`
	Root_category_id string `json:"root_category_id,omitempty"`
	Country string `json:"country,omitempty"`
	Store_currencies []Currency `json:"store_currencies,omitempty"`
	Channels []CartChannel `json:"channels,omitempty"`
	Dimension_unit string `json:"dimension_unit,omitempty"`
	Prices_include_tax bool `json:"prices_include_tax,omitempty"`
	Store_id string `json:"store_id,omitempty"`
	Name string `json:"name,omitempty"`
	Carrier_info []Carrier `json:"carrier_info,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Weight_unit string `json:"weight_unit,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Timezone string `json:"timezone,omitempty"`
	Currency Currency `json:"currency,omitempty"`
}

// ProductPriceUpdate represents the ProductPriceUpdate schema from the OpenAPI specification
type ProductPriceUpdate struct {
	Group_prices []map[string]interface{} `json:"group_prices,omitempty"` // Defines product's group prices
	Product_id string `json:"product_id,omitempty"` // Defines the product where the price has to be updated
}

// Category represents the Category schema from the OpenAPI specification
type Category struct {
	Stores_ids []string `json:"stores_ids,omitempty"`
	Parent_id string `json:"parent_id,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Id string `json:"id,omitempty"`
	Created_at A2CDateTime `json:"created_at,omitempty"`
	Images []Image `json:"images,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Meta_description string `json:"meta_description,omitempty"`
	Avail bool `json:"avail,omitempty"`
	Path string `json:"path,omitempty"`
	Sort_order int `json:"sort_order,omitempty"`
	Meta_title string `json:"meta_title,omitempty"`
	Description string `json:"description,omitempty"`
	Keywords string `json:"keywords,omitempty"`
	Seo_url string `json:"seo_url,omitempty"`
	Modified_at A2CDateTime `json:"modified_at,omitempty"`
	Name string `json:"name,omitempty"`
	Short_description string `json:"short_description,omitempty"`
}

// ResponseProductListResult represents the ResponseProductListResult schema from the OpenAPI specification
type ResponseProductListResult struct {
	Products_count int `json:"products_count,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Product []Product `json:"product,omitempty"`
}

// Plugin represents the Plugin schema from the OpenAPI specification
type Plugin struct {
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Name string `json:"name,omitempty"`
	Active bool `json:"active,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
}

// GiftCard represents the GiftCard schema from the OpenAPI specification
type GiftCard struct {
	Initial_amount float64 `json:"initial_amount,omitempty"`
	Issuer_name string `json:"issuer_name,omitempty"`
	Name string `json:"name,omitempty"`
	Recipient_email string `json:"recipient_email,omitempty"`
	Free_product_ids string `json:"free_product_ids,omitempty"`
	Recipient_name string `json:"recipient_name,omitempty"`
	Usage_history []CouponHistory `json:"usage_history,omitempty"`
	Code string `json:"code,omitempty"`
	Status string `json:"status,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Message string `json:"message,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Issuer_email string `json:"issuer_email,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Id string `json:"id,omitempty"`
	TypeField string `json:"type,omitempty"`
	Amount float64 `json:"amount,omitempty"`
	Avail_to string `json:"avail_to,omitempty"`
	Currency_code string `json:"currency_code,omitempty"`
}

// ResponseCartCatalogPriceRulesListResult represents the ResponseCartCatalogPriceRulesListResult schema from the OpenAPI specification
type ResponseCartCatalogPriceRulesListResult struct {
	Catalog_price_rules_count int `json:"catalog_price_rules_count,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Catalog_price_rules []CatalogPriceRule `json:"catalog_price_rules,omitempty"`
}

// OrderStatusHistoryItem represents the OrderStatusHistoryItem schema from the OpenAPI specification
type OrderStatusHistoryItem struct {
	Modified_time A2CDateTime `json:"modified_time,omitempty"`
	Name string `json:"name,omitempty"`
	Notify bool `json:"notify,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Comment string `json:"comment,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Id string `json:"id,omitempty"`
}

// ProductVariantImageAdd represents the ProductVariantImageAdd schema from the OpenAPI specification
type ProductVariantImageAdd struct {
	TypeField string `json:"type"` // Defines image's types that are specified by comma-separated list
	Option_id string `json:"option_id,omitempty"` // Defines option id of the product variant for which the image will be added
	Url string `json:"url,omitempty"` // Defines URL of the image that has to be added
	Label string `json:"label,omitempty"` // Defines alternative text that has to be attached to the picture
	Position int `json:"position,omitempty"` // Defines image’s position in the list
	Product_id string `json:"product_id,omitempty"` // Defines product id where the variant image has to be added
	Product_variant_id int `json:"product_variant_id"` // Defines product's variants specified by variant id
	Content string `json:"content,omitempty"` // Content(body) encoded in base64 of image file
	Image_name string `json:"image_name"` // Defines image's name
	Mime string `json:"mime,omitempty"` // Mime type of image http://en.wikipedia.org/wiki/Internet_media_type.
	Store_id string `json:"store_id,omitempty"` // Store Id
}

// CatalogPriceRule represents the CatalogPriceRule schema from the OpenAPI specification
type CatalogPriceRule struct {
	Description string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
	Actions []CatalogPriceRuleAction `json:"actions,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Avail bool `json:"avail,omitempty"`
	Date_start A2CDateTime `json:"date_start,omitempty"`
	Uses_per_order_limit int `json:"uses_per_order_limit,omitempty"`
	Date_end A2CDateTime `json:"date_end,omitempty"`
	Id string `json:"id,omitempty"`
	Short_description string `json:"short_description,omitempty"`
	TypeField string `json:"type,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Conditions []CouponCondition `json:"conditions,omitempty"`
	Usage_count float64 `json:"usage_count,omitempty"`
	Gid string `json:"gid,omitempty"`
	Created_time A2CDateTime `json:"created_time,omitempty"`
}

// ProductInventory represents the ProductInventory schema from the OpenAPI specification
type ProductInventory struct {
	In_stock bool `json:"in_stock,omitempty"`
	Priority int `json:"priority,omitempty"`
	Quantity float64 `json:"quantity,omitempty"`
	Warehouse_id string `json:"warehouse_id,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
}

// CustomerWishList represents the CustomerWishList schema from the OpenAPI specification
type CustomerWishList struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Is_public string `json:"is_public,omitempty"`
	Modified_at A2CDateTime `json:"modified_at,omitempty"`
	Products []CustomerWishListItem `json:"products,omitempty"`
	Created_at A2CDateTime `json:"created_at,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
}

// ProductGroupItem represents the ProductGroupItem schema from the OpenAPI specification
type ProductGroupItem struct {
	Is_qty_in_pack_fixed bool `json:"is_qty_in_pack_fixed,omitempty"`
	Product_id string `json:"product_id,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Child_item_id string `json:"child_item_id,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Default_qty_in_pack string `json:"default_qty_in_pack,omitempty"`
}

// StoreAttributeGroup represents the StoreAttributeGroup schema from the OpenAPI specification
type StoreAttributeGroup struct {
	Assigned_attribute_ids []string `json:"assigned_attribute_ids,omitempty"`
	Attribute_set_id string `json:"attribute_set_id,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Position int `json:"position,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
}

// Brand represents the Brand schema from the OpenAPI specification
type Brand struct {
	Meta_title string `json:"meta_title,omitempty"`
	Modified_time string `json:"modified_time,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Images []Image `json:"images,omitempty"`
	Meta_keywords string `json:"meta_keywords,omitempty"`
	Created_time string `json:"created_time,omitempty"`
	Id string `json:"id,omitempty"`
	Stores_ids []string `json:"stores_ids,omitempty"`
	Url string `json:"url,omitempty"`
	Active bool `json:"active,omitempty"`
	Meta_description string `json:"meta_description,omitempty"`
	Short_description string `json:"short_description,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Full_description string `json:"full_description,omitempty"`
	Name string `json:"name,omitempty"`
}

// ProductChildItemCombination represents the ProductChildItemCombination schema from the OpenAPI specification
type ProductChildItemCombination struct {
	Option_id string `json:"option_id,omitempty"`
	Option_value_id string `json:"option_value_id,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
}

// OrderPaymentMethod represents the OrderPaymentMethod schema from the OpenAPI specification
type OrderPaymentMethod struct {
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Name string `json:"name,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
}

// ModelResponseOrderList represents the ModelResponseOrderList schema from the OpenAPI specification
type ModelResponseOrderList struct {
	Pagination Pagination `json:"pagination,omitempty"`
	Result ResponseOrderListResult `json:"result,omitempty"`
	Return_code int `json:"return_code,omitempty"`
	Return_message string `json:"return_message,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
}

// CustomerAddress represents the CustomerAddress schema from the OpenAPI specification
type CustomerAddress struct {
	Company string `json:"company,omitempty"`
	Region string `json:"region,omitempty"`
	Last_name string `json:"last_name,omitempty"`
	Id string `json:"id,omitempty"`
	Address2 string `json:"address2,omitempty"`
	State State `json:"state,omitempty"`
	City string `json:"city,omitempty"`
	DefaultField bool `json:"default,omitempty"`
	Website string `json:"website,omitempty"`
	Address1 string `json:"address1,omitempty"`
	Phone_mobile string `json:"phone_mobile,omitempty"`
	Tax_id string `json:"tax_id,omitempty"`
	TypeField string `json:"type,omitempty"`
	Fax string `json:"fax,omitempty"`
	Identification_number string `json:"identification_number,omitempty"`
	Postcode string `json:"postcode,omitempty"`
	First_name string `json:"first_name,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Country Country `json:"country,omitempty"`
	Gender string `json:"gender,omitempty"`
	Phone string `json:"phone,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
}

// ResponseProductAttributeListResult represents the ResponseProductAttributeListResult schema from the OpenAPI specification
type ResponseProductAttributeListResult struct {
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Attribute []ProductAttribute `json:"attribute,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
}

// ProductTaxAdd represents the ProductTaxAdd schema from the OpenAPI specification
type ProductTaxAdd struct {
	Tax_rates []map[string]interface{} `json:"tax_rates"` // Defines tax rates of specified tax classes
	Name string `json:"name"` // Defines tax class name where tax has to be added
	Product_id string `json:"product_id,omitempty"` // Defines products specified by product id
}

// ModelResponseCartScriptList represents the ModelResponseCartScriptList schema from the OpenAPI specification
type ModelResponseCartScriptList struct {
	Return_code int `json:"return_code,omitempty"`
	Return_message string `json:"return_message,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
	Result ResponseCartScriptListResult `json:"result,omitempty"`
}

// ModelResponseCustomerGroupList represents the ModelResponseCustomerGroupList schema from the OpenAPI specification
type ModelResponseCustomerGroupList struct {
	Result ResponseCustomerGroupListResult `json:"result,omitempty"`
	Return_code int `json:"return_code,omitempty"`
	Return_message string `json:"return_message,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
}

// ResponseProductChildItemListResult represents the ResponseProductChildItemListResult schema from the OpenAPI specification
type ResponseProductChildItemListResult struct {
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Children []Child `json:"children,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Total_count int `json:"total_count,omitempty"`
}

// OrderTotals represents the OrderTotals schema from the OpenAPI specification
type OrderTotals struct {
	Subtotal float64 `json:"subtotal,omitempty"`
	Tax float64 `json:"tax,omitempty"`
	Total float64 `json:"total,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Discount float64 `json:"discount,omitempty"`
	Shipping float64 `json:"shipping,omitempty"`
}

// ProductOption represents the ProductOption schema from the OpenAPI specification
type ProductOption struct {
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Id string `json:"id,omitempty"`
	Option_items []ProductOptionItem `json:"option_items,omitempty"`
	Product_option_id string `json:"product_option_id,omitempty"`
	Sort_order int `json:"sort_order,omitempty"`
	Available bool `json:"available,omitempty"`
	Required bool `json:"required,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Description string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
	TypeField string `json:"type,omitempty"`
	Used_in_combination bool `json:"used_in_combination,omitempty"`
}

// ModelResponseCategoryList represents the ModelResponseCategoryList schema from the OpenAPI specification
type ModelResponseCategoryList struct {
	Return_message string `json:"return_message,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
	Result ResponseCategoryListResult `json:"result,omitempty"`
	Return_code int `json:"return_code,omitempty"`
}

// OrderStatusRefundItem represents the OrderStatusRefundItem schema from the OpenAPI specification
type OrderStatusRefundItem struct {
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Order_product_id string `json:"order_product_id,omitempty"`
	Product_id string `json:"product_id,omitempty"`
	Qty float64 `json:"qty,omitempty"`
	Refund float64 `json:"refund,omitempty"`
	Variant_id string `json:"variant_id,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
}

// Pagination represents the Pagination schema from the OpenAPI specification
type Pagination struct {
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Next string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
}

// ResponseCustomerListResult represents the ResponseCustomerListResult schema from the OpenAPI specification
type ResponseCustomerListResult struct {
	Customer []Customer `json:"customer,omitempty"`
	Customers_count int `json:"customers_count,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
}

// ProductReview represents the ProductReview schema from the OpenAPI specification
type ProductReview struct {
	Created_time A2CDateTime `json:"created_time,omitempty"`
	Id string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
	Status string `json:"status,omitempty"`
	Product_id string `json:"product_id,omitempty"`
	Rating float64 `json:"rating,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Nick_name string `json:"nick_name,omitempty"`
	Ratings []ProductReviewRating `json:"ratings,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Customer_id string `json:"customer_id,omitempty"`
	Email string `json:"email,omitempty"`
	Summary string `json:"summary,omitempty"`
}

// CustomerWishListItem represents the CustomerWishListItem schema from the OpenAPI specification
type CustomerWishListItem struct {
	Child_id string `json:"child_id,omitempty"`
	Created_time A2CDateTime `json:"created_time,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Id string `json:"id,omitempty"`
	Product_id string `json:"product_id,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
}

// OrderItemOption represents the OrderItemOption schema from the OpenAPI specification
type OrderItemOption struct {
	Option_id string `json:"option_id,omitempty"`
	Price float64 `json:"price,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Product_option_value_id string `json:"product_option_value_id,omitempty"`
	TypeField string `json:"type,omitempty"`
	Weight float64 `json:"weight,omitempty"`
	Name string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
}

// ShipmentTrackingNumber represents the ShipmentTrackingNumber schema from the OpenAPI specification
type ShipmentTrackingNumber struct {
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Carrier_id string `json:"carrier_id,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Tracking_number string `json:"tracking_number,omitempty"`
}

// ModelResponseCartMetaDataList represents the ModelResponseCartMetaDataList schema from the OpenAPI specification
type ModelResponseCartMetaDataList struct {
	Return_message string `json:"return_message,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
	Result ResponseCartMetaDataListResult `json:"result,omitempty"`
	Return_code int `json:"return_code,omitempty"`
}

// ModelResponseCustomerList represents the ModelResponseCustomerList schema from the OpenAPI specification
type ModelResponseCustomerList struct {
	Pagination Pagination `json:"pagination,omitempty"`
	Result ResponseCustomerListResult `json:"result,omitempty"`
	Return_code int `json:"return_code,omitempty"`
	Return_message string `json:"return_message,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
}

// Script represents the Script schema from the OpenAPI specification
type Script struct {
	Modified_time A2CDateTime `json:"modified_time,omitempty"`
	Scope string `json:"scope,omitempty"`
	Src string `json:"src,omitempty"`
	Load_method string `json:"load_method,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Event string `json:"event,omitempty"`
	Id string `json:"id,omitempty"`
	Created_time A2CDateTime `json:"created_time,omitempty"`
	Description string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
	Html string `json:"html,omitempty"`
}

// ProductUpdate represents the ProductUpdate schema from the OpenAPI specification
type ProductUpdate struct {
	Cost_price float64 `json:"cost_price,omitempty"` // Defines new product's cost price
	Model string `json:"model,omitempty"` // Defines product model that has to be updated
	Seo_url string `json:"seo_url,omitempty"` // Defines unique URL for SEO
	Tags string `json:"tags,omitempty"` // Product tags
	Clear_cache bool `json:"clear_cache,omitempty"` // Is cache clear required
	Visible string `json:"visible,omitempty"` // Set visibility status
	Search_keywords string `json:"search_keywords,omitempty"` // Defines unique search keywords
	Short_description string `json:"short_description,omitempty"` // Defines short description
	Retail_price float64 `json:"retail_price,omitempty"` // Defines new product's retail price
	Reindex bool `json:"reindex,omitempty"` // Is reindex required
	Barcode string `json:"barcode,omitempty"` // A barcode is a unique code composed of numbers used as a product identifier.
	Product_class string `json:"product_class,omitempty"` // A categorization for the product
	Height float64 `json:"height,omitempty"` // Defines product's height
	Status string `json:"status,omitempty"` // Defines product's status
	Taxable bool `json:"taxable,omitempty"` // Specifies whether a tax is charged
	Meta_title string `json:"meta_title,omitempty"` // Defines unique meta title for each entity
	Price float64 `json:"price,omitempty"` // Defines new product's price
	Sprice_expire string `json:"sprice_expire,omitempty"` // Defines the term of special price offer duration
	Categories_ids string `json:"categories_ids,omitempty"` // Defines product add that is specified by comma-separated categories id
	Id string `json:"id,omitempty"` // Defines product id that has to be updated
	Warehouse_id string `json:"warehouse_id,omitempty"` // This parameter is used for selecting a warehouse where you need to set/modify a product quantity.
	Quantity float64 `json:"quantity,omitempty"` // Defines new product's quantity
	Country_of_origin string `json:"country_of_origin,omitempty"` // The country where the inventory item was made
	Disable_report_cache bool `json:"disable_report_cache,omitempty"` // Disable report cache for current request
	Sprice_create string `json:"sprice_create,omitempty"` // Defines the date of special price creation
	Store_id string `json:"store_id,omitempty"` // Defines store id where the product should be found
	Weight float64 `json:"weight,omitempty"` // Weight
	Old_price float64 `json:"old_price,omitempty"` // Defines product's old price
	Sku string `json:"sku,omitempty"` // Defines new product's sku
	Manufacturer_id string `json:"manufacturer_id,omitempty"` // Defines product's manufacturer by manufacturer_id
	Meta_keywords string `json:"meta_keywords,omitempty"` // Defines unique meta keywords for each entity
	Meta_description string `json:"meta_description,omitempty"` // Defines unique meta description of a entity
	Description string `json:"description,omitempty"` // Defines new product's description
	Increase_quantity float64 `json:"increase_quantity,omitempty"` // Defines the incremental changes in product quantity
	Manage_stock bool `json:"manage_stock,omitempty"` // Defines inventory tracking for product
	Backorder_status string `json:"backorder_status,omitempty"` // Set backorder status
	Reserve_quantity float64 `json:"reserve_quantity,omitempty"` // This parameter allows to reserve/unreserve product quantity.
	Width float64 `json:"width,omitempty"` // Defines product's width
	Lang_id string `json:"lang_id,omitempty"` // Language id
	Gtin string `json:"gtin,omitempty"` // Global Trade Item Number. An GTIN is an identifier for trade items.
	Reduce_quantity float64 `json:"reduce_quantity,omitempty"` // Defines the decrement changes in product quantity
	Harmonized_system_code string `json:"harmonized_system_code,omitempty"` // Harmonized System Code. An HSC is a 6-digit identifier that allows participating countries to classify traded goods on a common basis for customs purposes
	Length float64 `json:"length,omitempty"` // Defines product's length
	Name string `json:"name,omitempty"` // Defines product's name that has to be updated
	In_stock bool `json:"in_stock,omitempty"` // Set stock status
	Special_price float64 `json:"special_price,omitempty"` // Defines new product's special price
	Manufacturer string `json:"manufacturer,omitempty"` // Defines product's manufacturer
	Report_request_id string `json:"report_request_id,omitempty"` // Report request id
}

// AccountCartAdd represents the AccountCartAdd schema from the OpenAPI specification
type AccountCartAdd struct {
	Amazon_sp_aws_region string `json:"amazon_sp_aws_region"` // Amazon AWS Region
	Bigcommerceapi_api_key string `json:"bigcommerceapi_api_key,omitempty"` // Bigcommerce API Key
	Shopify_shared_secret string `json:"shopify_shared_secret,omitempty"` // Shared secret
	Shopware_api_secret string `json:"shopware_api_secret,omitempty"` // Shopware client secret access key
	Volusion_login string `json:"volusion_login,omitempty"` // It's a Volusion account for which API is enabled
	Magento_token_secret string `json:"magento_token_secret,omitempty"` // Magento Token Secret
	Volusion_password string `json:"volusion_password,omitempty"` // Volusion API Password
	Walmart_client_id string `json:"walmart_client_id,omitempty"` // Walmart client ID
	Wix_refresh_token string `json:"wix_refresh_token,omitempty"` // Wix refresh token
	Ebay_client_id string `json:"ebay_client_id,omitempty"` // Application ID (AppID).
	Magento_access_token string `json:"magento_access_token,omitempty"` // Magento Access Token
	Shopware_api_key string `json:"shopware_api_key,omitempty"` // Shopware api key
	Zid_refresh_token string `json:"zid_refresh_token,omitempty"` // Zid refresh token
	Magento_consumer_secret string `json:"magento_consumer_secret,omitempty"` // Magento Consumer Secret
	Amazon_sp_refresh_token string `json:"amazon_sp_refresh_token"` // Amazon SP API OAuth refresh token
	Etsy_access_token string `json:"etsy_access_token,omitempty"` // Access token authorizing the app to access resources on behalf of a user
	Ebay_runame string `json:"ebay_runame,omitempty"` // The RuName value that eBay assigns to your application.
	Field3dcart_access_token string `json:"3dcart_access_token,omitempty"` // 3DCart Token
	Bigcommerceapi_client_id string `json:"bigcommerceapi_client_id,omitempty"` // Client ID of the requesting app
	Ecwid_acess_token string `json:"ecwid_acess_token,omitempty"` // Access token authorizing the app to access resources on behalf of a user
	Lightspeed_api_secret string `json:"lightspeed_api_secret,omitempty"` // LightSpeed api secret
	Ebay_refresh_token string `json:"ebay_refresh_token,omitempty"` // Used to renew the access token.
	Shopify_api_key string `json:"shopify_api_key,omitempty"` // Shopify API Key
	Neto_api_key string `json:"neto_api_key,omitempty"` // Neto API Key
	Etsy_shared_secret string `json:"etsy_shared_secret,omitempty"` // Etsy shared secret
	Amazon_sp_aws_role_arn string `json:"amazon_sp_aws_role_arn"` // Amazon AWS Role ARN
	Aspdotnetstorefront_api_user string `json:"aspdotnetstorefront_api_user,omitempty"` // It's a AspDotNetStorefront account for which API is available
	Db_tables_prefix string `json:"db_tables_prefix,omitempty"` // DB tables prefix
	Amazon_access_token string `json:"amazon_access_token,omitempty"` // MWS Auth Token. Access token authorizing the app to access resources on behalf of a user
	Ebay_client_secret string `json:"ebay_client_secret,omitempty"` // Shared Secret from eBay application
	Amazon_sp_client_id string `json:"amazon_sp_client_id"` // Amazon SP API app client id
	Ebay_access_token string `json:"ebay_access_token,omitempty"` // Used to authenticate API requests.
	Walmart_client_secret string `json:"walmart_client_secret,omitempty"` // Walmart client secret
	Wix_app_id string `json:"wix_app_id,omitempty"` // Wix App ID
	Etsy_token_secret string `json:"etsy_token_secret,omitempty"` // Secret token authorizing the app to access resources on behalf of a user
	Ecwid_store_id string `json:"ecwid_store_id,omitempty"` // Store Id
	Hybris_password string `json:"hybris_password,omitempty"` // User password
	Amazon_sp_aws_user_secret string `json:"amazon_sp_aws_user_secret"` // Amazon AWS user secret access key
	Ftp_user string `json:"ftp_user,omitempty"` // FTP User
	Magento_consumer_key string `json:"magento_consumer_key,omitempty"` // Magento Consumer Key
	Verify bool `json:"verify,omitempty"` // Enables or disables cart's verification
	Amazon_access_key_id string `json:"amazon_access_key_id,omitempty"` // Amazon Secret Key Id
	Zid_client_id int `json:"zid_client_id,omitempty"` // Zid Client ID
	Ftp_port int `json:"ftp_port,omitempty"` // FTP Port
	Walmart_environment string `json:"walmart_environment,omitempty"` // Walmart environment
	Bigcommerceapi_context string `json:"bigcommerceapi_context,omitempty"` // API Path section unique to the store
	Hybris_username string `json:"hybris_username,omitempty"` // User Name
	Aspdotnetstorefront_api_pass string `json:"aspdotnetstorefront_api_pass,omitempty"` // AspDotNetStorefront API Password
	Ftp_password string `json:"ftp_password,omitempty"` // FTP Password
	Amazon_seller_id string `json:"amazon_seller_id,omitempty"` // Amazon Seller ID (Merchant token)
	Commercehq_api_key string `json:"commercehq_api_key,omitempty"` // CommerceHQ api key
	Bigcommerceapi_access_token string `json:"bigcommerceapi_access_token,omitempty"` // Access token authorizing the app to access resources on behalf of a user
	Field3dcart_private_key string `json:"3dcart_private_key,omitempty"` // 3DCart Private Key
	Amazon_marketplaces_ids string `json:"amazon_marketplaces_ids,omitempty"` // Amazon Marketplace IDs comma separated string
	Store_url string `json:"store_url"` // A web address of a store that you would like to connect to API2Cart
	Bigcommerceapi_api_path string `json:"bigcommerceapi_api_path,omitempty"` // BigCommerce API URL
	Ebay_environment string `json:"ebay_environment,omitempty"` // eBay environment
	Shopify_access_token string `json:"shopify_access_token,omitempty"` // Access token authorizing the app to access resources on behalf of a user
	Mercado_libre_refresh_token string `json:"mercado_libre_refresh_token,omitempty"` // Mercado Libre Refresh Token
	Store_key string `json:"store_key,omitempty"` // Set this parameter if bridge is already uploaded to store
	Demandware_user_name string `json:"demandware_user_name,omitempty"` // Demandware user name
	Neto_api_username string `json:"neto_api_username,omitempty"` // Neto User Name
	Wc_consumer_secret string `json:"wc_consumer_secret,omitempty"` // Woocommerce consumer secret
	Zid_access_token string `json:"zid_access_token,omitempty"` // Zid Access Token
	Bigcommerceapi_admin_account string `json:"bigcommerceapi_admin_account,omitempty"` // It's a BigCommerce account for which API is enabled
	Bridge_url string `json:"bridge_url,omitempty"` // This parameter allows to set up store with custom bridge url (also you must use store_root parameter if a bridge folder is not in the root folder of the store)
	Shopify_api_password string `json:"shopify_api_password,omitempty"` // Shopify API Password
	Demandware_api_password string `json:"demandware_api_password,omitempty"` // Demandware api password
	Mercado_libre_app_secret_key string `json:"mercado_libre_app_secret_key,omitempty"` // Mercado Libre App Secret Key
	Amazon_secret_key string `json:"amazon_secret_key,omitempty"` // Amazon Secret Key
	Ebay_site_id int `json:"ebay_site_id,omitempty"` // eBay global ID
	Hybris_client_secret string `json:"hybris_client_secret,omitempty"` // Omni Commerce Connector Client Secret
	Store_root string `json:"store_root,omitempty"` // Absolute path to the store root directory (used with "bridge_url" parameter)
	Amazon_sp_aws_user_key_id string `json:"amazon_sp_aws_user_key_id"` // Amazon AWS user access key ID
	Cart_id string `json:"cart_id"` // Store’s identifier which you can get from cart_list method
	Lightspeed_api_key string `json:"lightspeed_api_key,omitempty"` // LightSpeed api key
	Etsy_keystring string `json:"etsy_keystring,omitempty"` // Etsy keystring
	Zid_client_secret string `json:"zid_client_secret,omitempty"` // Zid Client Secret
	Prestashop_webservice_key string `json:"prestashop_webservice_key,omitempty"` // Prestashop webservice key
	Ftp_store_dir string `json:"ftp_store_dir,omitempty"` // FTP Store dir
	Amazon_sp_api_environment string `json:"amazon_sp_api_environment,omitempty"` // Amazon SP API environment
	Amazon_sp_client_secret string `json:"amazon_sp_client_secret"` // Amazon SP API app client secret
	Wix_app_secret_key string `json:"wix_app_secret_key,omitempty"` // Wix App Secret Key
	Demandware_user_password string `json:"demandware_user_password,omitempty"` // Demandware user password
	Demandware_client_id string `json:"demandware_client_id,omitempty"` // Demandware client id
	Mercado_libre_app_id string `json:"mercado_libre_app_id,omitempty"` // Mercado Libre App ID
	Etsy_client_id string `json:"etsy_client_id"` // Etsy Client Id
	Ftp_host string `json:"ftp_host,omitempty"` // FTP connection host
	Shopware_access_key string `json:"shopware_access_key,omitempty"` // Shopware access key
	Validate_version bool `json:"validate_version,omitempty"` // Specify if api2cart should validate cart version
	Field3dcartapi_api_key string `json:"3dcartapi_api_key,omitempty"` // 3DCart API Key
	Hybris_websites []map[string]interface{} `json:"hybris_websites,omitempty"` // Websites to stores mapping data
	Wc_consumer_key string `json:"wc_consumer_key,omitempty"` // Woocommerce consumer key
	Walmart_channel_type string `json:"walmart_channel_type,omitempty"` // Walmart WM_CONSUMER.CHANNEL.TYPE header
	Commercehq_api_password string `json:"commercehq_api_password,omitempty"` // CommerceHQ api password
	Etsy_refresh_token string `json:"etsy_refresh_token"` // Etsy Refresh token
	Squarespace_api_key string `json:"squarespace_api_key,omitempty"` // Squarespace API Key
	Zid_authorization string `json:"zid_authorization,omitempty"` // Zid Authorization
	Hybris_client_id string `json:"hybris_client_id,omitempty"` // Omni Commerce Connector Client ID
}

// CartShippingZone represents the CartShippingZone schema from the OpenAPI specification
type CartShippingZone struct {
	Country_iso2_codes []string `json:"country_iso2_codes,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Code string `json:"code,omitempty"`
	Country string `json:"country,omitempty"`
}

// ResponseOrderTransactionListResult represents the ResponseOrderTransactionListResult schema from the OpenAPI specification
type ResponseOrderTransactionListResult struct {
	Transactions_count int `json:"transactions_count,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Transactions []OrderTransaction `json:"transactions,omitempty"`
}

// ModelResponseCartCatalogPriceRulesList represents the ModelResponseCartCatalogPriceRulesList schema from the OpenAPI specification
type ModelResponseCartCatalogPriceRulesList struct {
	Return_message string `json:"return_message,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
	Result ResponseCartCatalogPriceRulesListResult `json:"result,omitempty"`
	Return_code int `json:"return_code,omitempty"`
}

// Language represents the Language schema from the OpenAPI specification
type Language struct {
	Id string `json:"id,omitempty"`
	Iso_code string `json:"iso_code,omitempty"`
	Name string `json:"name,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
}

// OrderPreestimateShippingList represents the OrderPreestimateShippingList schema from the OpenAPI specification
type OrderPreestimateShippingList struct {
	Order_item []map[string]interface{} `json:"order_item"`
	Params string `json:"params,omitempty"` // Set this parameter in order to choose which entity fields you want to retrieve
	Warehouse_id string `json:"warehouse_id,omitempty"` // This parameter is used for selecting a warehouse where you need to set/modify a product quantity.
	Customer_email string `json:"customer_email,omitempty"` // Retrieves orders specified by customer email
	Customer_id string `json:"customer_id,omitempty"` // Retrieves orders specified by customer id
	Shipp_city string `json:"shipp_city,omitempty"` // Specifies shipping city
	Store_id string `json:"store_id,omitempty"` // Store Id
	Exclude string `json:"exclude,omitempty"` // Set this parameter in order to choose which entity fields you want to ignore. Works only if parameter `params` equal force_all
	Shipp_address_1 string `json:"shipp_address_1,omitempty"` // Specifies first shipping address
	Shipp_country string `json:"shipp_country"` // Specifies shipping country code
	Shipp_postcode string `json:"shipp_postcode,omitempty"` // Specifies shipping postcode
	Shipp_state string `json:"shipp_state,omitempty"` // Specifies shipping state code
}

// ModelResponseProductAttributeList represents the ModelResponseProductAttributeList schema from the OpenAPI specification
type ModelResponseProductAttributeList struct {
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
	Result ResponseProductAttributeListResult `json:"result,omitempty"`
	Return_code int `json:"return_code,omitempty"`
	Return_message string `json:"return_message,omitempty"`
}

// CatalogPriceRuleAction represents the CatalogPriceRuleAction schema from the OpenAPI specification
type CatalogPriceRuleAction struct {
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Include_tax bool `json:"include_tax,omitempty"`
	TypeField string `json:"type,omitempty"`
	Apply_to string `json:"apply_to,omitempty"`
	Value float64 `json:"value,omitempty"`
	Conditions []CouponCondition `json:"conditions,omitempty"`
	Currency_code string `json:"currency_code,omitempty"`
	Quantity float64 `json:"quantity,omitempty"`
	Scope string `json:"scope,omitempty"`
}

// Shipment represents the Shipment schema from the OpenAPI specification
type Shipment struct {
	Tracking_numbers []ShipmentTrackingNumber `json:"tracking_numbers,omitempty"`
	Id string `json:"id,omitempty"`
	Shipment_provider string `json:"shipment_provider,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Items []ShipmentItem `json:"items,omitempty"`
	Modified_time A2CDateTime `json:"modified_time,omitempty"`
	Warehouse_id string `json:"warehouse_id,omitempty"`
	Order_id string `json:"order_id,omitempty"`
	Created_at A2CDateTime `json:"created_at,omitempty"`
	Is_shipped bool `json:"is_shipped,omitempty"`
}

// BasketLiveShippingService represents the BasketLiveShippingService schema from the OpenAPI specification
type BasketLiveShippingService struct {
	Enabled_on_store bool `json:"enabled_on_store,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Callback string `json:"callback,omitempty"`
	Callback_err_cnt int `json:"callback_err_cnt,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
}

// Subscriber represents the Subscriber schema from the OpenAPI specification
type Subscriber struct {
	First_name string `json:"first_name,omitempty"`
	Modified_time string `json:"modified_time,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Id string `json:"id,omitempty"`
	Last_name string `json:"last_name,omitempty"`
	Customer_id string `json:"customer_id,omitempty"`
	Created_time string `json:"created_time,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Email string `json:"email,omitempty"`
	Subscribed bool `json:"subscribed,omitempty"`
}

// ProductGroupPrice represents the ProductGroupPrice schema from the OpenAPI specification
type ProductGroupPrice struct {
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Price float64 `json:"price,omitempty"`
	Start_time string `json:"start_time,omitempty"`
	Store_id string `json:"store_id,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Expire_time string `json:"expire_time,omitempty"`
	Id string `json:"id,omitempty"`
	Group_id string `json:"group_id,omitempty"`
	Quantity float64 `json:"quantity,omitempty"`
}

// Coupon represents the Coupon schema from the OpenAPI specification
type Coupon struct {
	Date_start A2CDateTime `json:"date_start,omitempty"`
	Priority int `json:"priority,omitempty"`
	Description string `json:"description,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Logic_operator string `json:"logic_operator,omitempty"`
	Usage_history []CouponHistory `json:"usage_history,omitempty"`
	Used_times int `json:"used_times,omitempty"`
	Actions []CouponAction `json:"actions,omitempty"`
	Conditions []CouponCondition `json:"conditions,omitempty"`
	Date_end A2CDateTime `json:"date_end,omitempty"`
	Usage_limit_per_customer int `json:"usage_limit_per_customer,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Code string `json:"code,omitempty"`
	Usage_limit int `json:"usage_limit,omitempty"`
	Avail bool `json:"avail,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Codes []CouponCode `json:"codes,omitempty"`
}

// CustomerUpdate represents the CustomerUpdate schema from the OpenAPI specification
type CustomerUpdate struct {
	Group_id string `json:"group_id,omitempty"` // Customer group_id
	Group_ids string `json:"group_ids,omitempty"` // Groups that will be assigned to a customer
	First_name string `json:"first_name,omitempty"` // Defines customer's first name
	News_letter_subscription bool `json:"news_letter_subscription,omitempty"` // Defines whether the newsletter subscription is available for the user
	Phone string `json:"phone,omitempty"` // Defines customer's phone number
	Tags string `json:"tags,omitempty"` // Customer tags
	Address []map[string]interface{} `json:"address,omitempty"`
	Email string `json:"email,omitempty"` // Defines customer's email
	Id string `json:"id,omitempty"` // Entity id
	Last_name string `json:"last_name,omitempty"` // Defines customer's last name
}

// TaxClassRate represents the TaxClassRate schema from the OpenAPI specification
type TaxClassRate struct {
	Cities []string `json:"cities,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Tax_type int `json:"tax_type,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Zip_codes []TaxClassZipCodes `json:"zip_codes,omitempty"`
	Countries []TaxClassCountries `json:"countries,omitempty"`
	Address []string `json:"address,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Tax float64 `json:"tax,omitempty"`
	Tax_based_on string `json:"tax_based_on,omitempty"`
}

// ProductVariantUpdate represents the ProductVariantUpdate schema from the OpenAPI specification
type ProductVariantUpdate struct {
	Cost_price float64 `json:"cost_price,omitempty"` // Defines new product's cost price
	Meta_keywords string `json:"meta_keywords,omitempty"` // Defines unique meta keywords for each entity
	Short_description string `json:"short_description,omitempty"` // Defines short description
	Quantity float64 `json:"quantity,omitempty"` // Defines new products' variants quantity
	Backorder_status string `json:"backorder_status,omitempty"` // Set backorder status
	Length float64 `json:"length,omitempty"` // Defines product's length
	Status string `json:"status,omitempty"` // Defines product variant's status
	Meta_description string `json:"meta_description,omitempty"` // Defines unique meta description of a entity
	Sprice_expire string `json:"sprice_expire,omitempty"` // Defines the term of special price offer duration
	Barcode string `json:"barcode,omitempty"` // A barcode is a unique code composed of numbers used as a product identifier.
	Visible string `json:"visible,omitempty"` // Set visibility status
	Meta_title string `json:"meta_title,omitempty"` // Defines unique meta title for each entity
	Reduce_quantity float64 `json:"reduce_quantity,omitempty"` // Defines the decrement changes in product quantity
	Available_for_sale bool `json:"available_for_sale,omitempty"` // Specifies the set of visible/invisible product's variants for sale
	Lang_id string `json:"lang_id,omitempty"` // Language id
	Reserve_quantity float64 `json:"reserve_quantity,omitempty"` // This parameter allows to reserve/unreserve product variants quantity.
	Sku string `json:"sku,omitempty"` // Defines new product's variant sku
	Price float64 `json:"price,omitempty"` // Defines new product's variant price
	Manage_stock bool `json:"manage_stock,omitempty"` // Defines inventory tracking for product variant
	Harmonized_system_code string `json:"harmonized_system_code,omitempty"` // Harmonized System Code. An HSC is a 6-digit identifier that allows participating countries to classify traded goods on a common basis for customs purposes
	Options []map[string]interface{} `json:"options,omitempty"` // Defines variant's options list
	Special_price float64 `json:"special_price,omitempty"` // Defines new product's variant special price
	Weight float64 `json:"weight,omitempty"` // Weight
	Old_price float64 `json:"old_price,omitempty"` // Defines product's old price
	Description string `json:"description,omitempty"` // Specifies variant's description
	Gtin string `json:"gtin,omitempty"` // Global Trade Item Number. An GTIN is an identifier for trade items.
	Id string `json:"id,omitempty"` // Defines variant update specified by variant id
	Retail_price float64 `json:"retail_price,omitempty"` // Defines new product's retail price
	Model string `json:"model,omitempty"` // Specifies variant's model that has to be added
	Clear_cache bool `json:"clear_cache,omitempty"` // Is cache clear required
	Sprice_create string `json:"sprice_create,omitempty"` // Defines the date of special price creation
	Increase_quantity float64 `json:"increase_quantity,omitempty"` // Defines the incremental changes in product quantity
	Taxable bool `json:"taxable,omitempty"` // Specifies whether a tax is charged
	Country_of_origin string `json:"country_of_origin,omitempty"` // The country where the inventory item was made
	Product_id string `json:"product_id,omitempty"` // Defines product's id where the variant has to be updated
	Store_id string `json:"store_id,omitempty"` // Defines store id where the variant should be found
	In_stock bool `json:"in_stock,omitempty"` // Set stock status
	Width float64 `json:"width,omitempty"` // Defines product's width
	Name string `json:"name,omitempty"` // Defines variant's name that has to be updated
	Reindex bool `json:"reindex,omitempty"` // Is reindex required
	Warehouse_id string `json:"warehouse_id,omitempty"` // This parameter is used for selecting a warehouse where you need to set/modify a product quantity.
	Height float64 `json:"height,omitempty"` // Defines product's height
}

// Customer represents the Customer schema from the OpenAPI specification
type Customer struct {
	Birth_day A2CDateTime `json:"birth_day,omitempty"`
	First_name string `json:"first_name,omitempty"`
	Orders_count int `json:"orders_count,omitempty"`
	Modified_time A2CDateTime `json:"modified_time,omitempty"`
	Email string `json:"email,omitempty"`
	Gender string `json:"gender,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Lang_id string `json:"lang_id,omitempty"`
	Last_order_id string `json:"last_order_id,omitempty"`
	News_letter_subscription bool `json:"news_letter_subscription,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Id string `json:"id,omitempty"`
	Last_login A2CDateTime `json:"last_login,omitempty"`
	Last_name string `json:"last_name,omitempty"`
	Created_time A2CDateTime `json:"created_time,omitempty"`
	Website string `json:"website,omitempty"`
	Ip_address string `json:"ip_address,omitempty"`
	Phone string `json:"phone,omitempty"`
	Stores_ids []string `json:"stores_ids,omitempty"`
	Address_book []CustomerAddress `json:"address_book,omitempty"`
	Group []CustomerGroup `json:"group,omitempty"`
	Login string `json:"login,omitempty"`
	Fax string `json:"fax,omitempty"`
	Company string `json:"company,omitempty"`
	Status string `json:"status,omitempty"`
}

// TaxClassZipCodesRange represents the TaxClassZipCodesRange schema from the OpenAPI specification
type TaxClassZipCodesRange struct {
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	From string `json:"from,omitempty"`
	To string `json:"to,omitempty"`
}

// ResponseCustomerGroupListResult represents the ResponseCustomerGroupListResult schema from the OpenAPI specification
type ResponseCustomerGroupListResult struct {
	Group []CustomerGroup `json:"group,omitempty"`
	Group_count int `json:"group_count,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
}

// CustomerAttributeValue represents the CustomerAttributeValue schema from the OpenAPI specification
type CustomerAttributeValue struct {
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Id string `json:"id,omitempty"`
	Value string `json:"value,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
}

// ProductAdvancedPrice represents the ProductAdvancedPrice schema from the OpenAPI specification
type ProductAdvancedPrice struct {
	Value float64 `json:"value,omitempty"`
	Avail bool `json:"avail,omitempty"`
	Expire_time A2CDateTime `json:"expire_time,omitempty"`
	Group_id string `json:"group_id,omitempty"`
	Id string `json:"id,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Quantity_from float64 `json:"quantity_from,omitempty"`
	Start_time A2CDateTime `json:"start_time,omitempty"`
}

// ProductVariantAdd represents the ProductVariantAdd schema from the OpenAPI specification
type ProductVariantAdd struct {
	Lang_id string `json:"lang_id,omitempty"` // Language id
	Weight float64 `json:"weight,omitempty"` // Weight
	Manage_stock bool `json:"manage_stock,omitempty"` // Defines inventory tracking for product variant
	Manufacturer string `json:"manufacturer,omitempty"` // Specifies the product variant's manufacturer
	Clear_cache bool `json:"clear_cache,omitempty"` // Is cache clear required
	Model string `json:"model"` // Specifies variant's model that has to be added
	Warehouse_id string `json:"warehouse_id,omitempty"` // This parameter is used for selecting a warehouse where you need to set/modify a product quantity.
	Meta_keywords string `json:"meta_keywords,omitempty"` // Defines unique meta keywords for each entity
	Meta_title string `json:"meta_title,omitempty"` // Defines unique meta title for each entity
	Created_at string `json:"created_at,omitempty"` // Defines the date of entity creation
	Country_of_origin string `json:"country_of_origin,omitempty"` // The country where the inventory item was made
	Short_description string `json:"short_description,omitempty"` // Defines short description
	Available_for_view bool `json:"available_for_view,omitempty"` // Specifies the set of visible/invisible product's variants for users
	Available_for_sale bool `json:"available_for_sale,omitempty"` // Specifies the set of visible/invisible product's variants for sale
	Description string `json:"description,omitempty"` // Specifies variant's description
	Harmonized_system_code string `json:"harmonized_system_code,omitempty"` // Harmonized System Code. An HSC is a 6-digit identifier that allows participating countries to classify traded goods on a common basis for customs purposes
	Height float64 `json:"height,omitempty"` // Defines product's height
	Tax_class_id string `json:"tax_class_id,omitempty"` // Defines tax classes where entity has to be added
	Taxable bool `json:"taxable,omitempty"` // Specifies whether a tax is charged
	Special_price float64 `json:"special_price,omitempty"` // Specifies variant's model that has to be added
	Price float64 `json:"price,omitempty"` // Defines new product's variant price
	Cost_price float64 `json:"cost_price,omitempty"` // Defines new product's cost price
	Width float64 `json:"width,omitempty"` // Defines product's width
	Sku string `json:"sku,omitempty"` // Defines variant's sku that has to be added
	Sprice_create string `json:"sprice_create,omitempty"` // Defines the date of special price creation
	Attributes []map[string]interface{} `json:"attributes,omitempty"` // Defines variant's attributes list
	Quantity float64 `json:"quantity,omitempty"` // Defines product variant's quantity that has to be added
	Product_id string `json:"product_id,omitempty"` // Defines product's id where the variant has to be added
	Name string `json:"name,omitempty"` // Defines variant's name that has to be added
	Url string `json:"url,omitempty"` // Defines unique product variant's URL
	Sprice_modified string `json:"sprice_modified,omitempty"` // Defines the date of special price modification
	Weight_unit string `json:"weight_unit,omitempty"` // Weight Unit
	Store_id string `json:"store_id,omitempty"` // Add variants specified by store id
	Sprice_expire string `json:"sprice_expire,omitempty"` // Defines the term of special price offer duration
	Barcode string `json:"barcode,omitempty"` // A barcode is a unique code composed of numbers used as a product identifier.
	Length float64 `json:"length,omitempty"` // Defines product's length
	Meta_description string `json:"meta_description,omitempty"` // Defines unique meta description of a entity
}

// OrderStatus represents the OrderStatus schema from the OpenAPI specification
type OrderStatus struct {
	History []OrderStatusHistoryItem `json:"history,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Refund_info OrderStatusRefund `json:"refund_info,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
}

// OrderStatusRefund represents the OrderStatusRefund schema from the OpenAPI specification
type OrderStatusRefund struct {
	Fee float64 `json:"fee,omitempty"`
	Time A2CDateTime `json:"time,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Refunded_items []OrderStatusRefundItem `json:"refunded_items,omitempty"`
	Shipping float64 `json:"shipping,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Comment string `json:"comment,omitempty"`
	Tax float64 `json:"tax,omitempty"`
	Total_refunded float64 `json:"total_refunded,omitempty"`
}

// ProductTierPrice represents the ProductTierPrice schema from the OpenAPI specification
type ProductTierPrice struct {
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Price float64 `json:"price,omitempty"`
	Qty float64 `json:"qty,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
}

// ModelResponseProductList represents the ModelResponseProductList schema from the OpenAPI specification
type ModelResponseProductList struct {
	Pagination Pagination `json:"pagination,omitempty"`
	Result ResponseProductListResult `json:"result,omitempty"`
	Return_code int `json:"return_code,omitempty"`
	Return_message string `json:"return_message,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
}

// ResponseCartCouponListResult represents the ResponseCartCouponListResult schema from the OpenAPI specification
type ResponseCartCouponListResult struct {
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Coupon []Coupon `json:"coupon,omitempty"`
	Coupon_count int `json:"coupon_count,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
}

// OrderRefundAdd represents the OrderRefundAdd schema from the OpenAPI specification
type OrderRefundAdd struct {
	Message string `json:"message,omitempty"` // Refund reason, or some else message which assigned to refund.
	Total_price float64 `json:"total_price,omitempty"` // Defines order refund amount.
	Fee_price float64 `json:"fee_price,omitempty"` // Specifies refund's fee price
	Is_online bool `json:"is_online,omitempty"` // Indicates whether refund type is online
	Order_id string `json:"order_id,omitempty"` // Defines the order for which the refund will be created.
	Shipping_price float64 `json:"shipping_price,omitempty"` // Defines refund shipping amount.
	Date string `json:"date,omitempty"` // Specifies an order creation date in format Y-m-d H:i:s
	Item_restock bool `json:"item_restock,omitempty"` // Boolean, whether or not to add the line items back to the store inventory.
	Items []map[string]interface{} `json:"items,omitempty"` // Defines items in the order that will be refunded
	Send_notifications bool `json:"send_notifications,omitempty"` // Send notifications to customer after refund was created
}

// ProductReviewRating represents the ProductReviewRating schema from the OpenAPI specification
type ProductReviewRating struct {
	Value int `json:"value,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// ModelResponseCustomerAttributeList represents the ModelResponseCustomerAttributeList schema from the OpenAPI specification
type ModelResponseCustomerAttributeList struct {
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
	Result ResponseCustomerAttributeListResult `json:"result,omitempty"`
	Return_code int `json:"return_code,omitempty"`
	Return_message string `json:"return_message,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
}

// Cart represents the Cart schema from the OpenAPI specification
type Cart struct {
	Shipping_zones []CartShippingZone `json:"shipping_zones,omitempty"`
	Stores_info []CartStoreInfo `json:"stores_info,omitempty"`
	Url string `json:"url,omitempty"`
	Version string `json:"version,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Warehouses []CartWarehouse `json:"warehouses,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Name string `json:"name,omitempty"`
	Db_prefix string `json:"db_prefix,omitempty"`
}

// TaxClassZipCodes represents the TaxClassZipCodes schema from the OpenAPI specification
type TaxClassZipCodes struct {
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Fields []TaxClassZipCodesRange `json:"fields,omitempty"`
	Is_range bool `json:"is_range,omitempty"`
	RangeField []string `json:"range,omitempty"`
}

// StoreAttribute represents the StoreAttribute schema from the OpenAPI specification
type StoreAttribute struct {
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Default_values []string `json:"default_values,omitempty"`
	Id string `json:"id,omitempty"`
	Lang_id string `json:"lang_id,omitempty"`
	TypeField string `json:"type,omitempty"`
	Values []string `json:"values,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Position int `json:"position,omitempty"`
	Visible bool `json:"visible,omitempty"`
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
	Required bool `json:"required,omitempty"`
	Store_id string `json:"store_id,omitempty"`
	System bool `json:"system,omitempty"`
}

// TaxClass represents the TaxClass schema from the OpenAPI specification
type TaxClass struct {
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Tax float64 `json:"tax,omitempty"`
	Tax_rates []TaxClassRate `json:"tax_rates,omitempty"`
	Tax_type int `json:"tax_type,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Avail bool `json:"avail,omitempty"`
}

// ResponseCartMetaDataListResult represents the ResponseCartMetaDataListResult schema from the OpenAPI specification
type ResponseCartMetaDataListResult struct {
	Items []CartMetaData `json:"items,omitempty"`
	Total_count int `json:"total_count,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
}

// TaxClassCountries represents the TaxClassCountries schema from the OpenAPI specification
type TaxClassCountries struct {
	Tax_type int `json:"tax_type,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Code string `json:"code,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	States []TaxClassStates `json:"states,omitempty"`
	Tax float64 `json:"tax,omitempty"`
}

// CouponHistory represents the CouponHistory schema from the OpenAPI specification
type CouponHistory struct {
	Amount float64 `json:"amount,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Order_id string `json:"order_id,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
}

// State represents the State schema from the OpenAPI specification
type State struct {
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Name string `json:"name,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Code string `json:"code,omitempty"`
}

// CartConfigUpdate represents the CartConfigUpdate schema from the OpenAPI specification
type CartConfigUpdate struct {
	Db_tables_prefix string `json:"db_tables_prefix,omitempty"` // This parameter is deprecated for this method. Please, use this parameter in method account.config.update
	Store_id string `json:"store_id,omitempty"` // Store Id
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"` // This parameter sets the list of params to the shopping cart.
}

// ModelResponseCartCouponList represents the ModelResponseCartCouponList schema from the OpenAPI specification
type ModelResponseCartCouponList struct {
	Result ResponseCartCouponListResult `json:"result,omitempty"`
	Return_code int `json:"return_code,omitempty"`
	Return_message string `json:"return_message,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
}

// Product represents the Product schema from the OpenAPI specification
type Product struct {
	In_stock bool `json:"in_stock,omitempty"`
	Modified_at A2CDateTime `json:"modified_at,omitempty"`
	Quantity float64 `json:"quantity,omitempty"`
	Description string `json:"description,omitempty"`
	Length float64 `json:"length,omitempty"`
	Related_products_ids []string `json:"related_products_ids,omitempty"`
	Meta_title string `json:"meta_title,omitempty"`
	Short_description string `json:"short_description,omitempty"`
	Tier_price []ProductTierPrice `json:"tier_price,omitempty"`
	U_model string `json:"u_model,omitempty"`
	Name string `json:"name,omitempty"`
	Weight_unit string `json:"weight_unit,omitempty"`
	Avail_view bool `json:"avail_view,omitempty"`
	U_mpn string `json:"u_mpn,omitempty"`
	U_sku string `json:"u_sku,omitempty"`
	Create_at A2CDateTime `json:"create_at,omitempty"`
	Meta_keywords string `json:"meta_keywords,omitempty"`
	Discounts []Discount `json:"discounts,omitempty"`
	Advanced_price []ProductAdvancedPrice `json:"advanced_price,omitempty"`
	Tax_class_id string `json:"tax_class_id,omitempty"`
	Group_items []ProductGroupItem `json:"group_items,omitempty"`
	Cost_price float64 `json:"cost_price,omitempty"`
	Height float64 `json:"height,omitempty"`
	Is_stock_managed bool `json:"is_stock_managed,omitempty"`
	Inventory []ProductInventory `json:"inventory,omitempty"`
	Seo_url string `json:"seo_url,omitempty"`
	Is_downloadable bool `json:"is_downloadable,omitempty"`
	Avail_sale bool `json:"avail_sale,omitempty"`
	Product_options []ProductOption `json:"product_options,omitempty"`
	TypeField string `json:"type,omitempty"`
	U_brand_id string `json:"u_brand_id,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Images []Image `json:"images,omitempty"`
	Special_price SpecialPrice `json:"special_price,omitempty"`
	U_upc string `json:"u_upc,omitempty"`
	Id string `json:"id,omitempty"`
	Group_price []ProductGroupPrice `json:"group_price,omitempty"`
	Categories_ids []string `json:"categories_ids,omitempty"`
	Backorders string `json:"backorders,omitempty"`
	Meta_description string `json:"meta_description,omitempty"`
	Sort_order int `json:"sort_order,omitempty"`
	U_brand string `json:"u_brand,omitempty"`
	Price float64 `json:"price,omitempty"`
	Dimensions_unit string `json:"dimensions_unit,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Manage_stock string `json:"manage_stock,omitempty"`
	Url string `json:"url,omitempty"`
	Weight float64 `json:"weight,omitempty"`
	Width float64 `json:"width,omitempty"`
	Is_virtual bool `json:"is_virtual,omitempty"`
	Stores_ids []string `json:"stores_ids,omitempty"`
}

// OrderShippingMethod represents the OrderShippingMethod schema from the OpenAPI specification
type OrderShippingMethod struct {
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Name string `json:"name,omitempty"`
}

// OrderAbandoned represents the OrderAbandoned schema from the OpenAPI specification
type OrderAbandoned struct {
	Created_at A2CDateTime `json:"created_at,omitempty"`
	Currency Currency `json:"currency,omitempty"`
	Customer BaseCustomer `json:"customer,omitempty"`
	Id string `json:"id,omitempty"`
	Basket_id string `json:"basket_id,omitempty"`
	Basket_url string `json:"basket_url,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Modified_at A2CDateTime `json:"modified_at,omitempty"`
	Order_products []OrderItem `json:"order_products,omitempty"`
	Totals OrderTotals `json:"totals,omitempty"`
}

// OrderTotal represents the OrderTotal schema from the OpenAPI specification
type OrderTotal struct {
	Total float64 `json:"total,omitempty"`
	Total_paid float64 `json:"total_paid,omitempty"`
	Wrapping_ex_tax float64 `json:"wrapping_ex_tax,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Total_tax float64 `json:"total_tax,omitempty"`
	Shipping_ex_tax float64 `json:"shipping_ex_tax,omitempty"`
	Subtotal_ex_tax float64 `json:"subtotal_ex_tax,omitempty"`
	Total_discount float64 `json:"total_discount,omitempty"`
}

// OrderTransaction represents the OrderTransaction schema from the OpenAPI specification
type OrderTransaction struct {
	Amount float64 `json:"amount,omitempty"`
	Created_time A2CDateTime `json:"created_time,omitempty"`
	Card_brand string `json:"card_brand,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Id string `json:"id,omitempty"`
	Avs_postal_resp_code string `json:"avs_postal_resp_code,omitempty"`
	Cvv_code string `json:"cvv_code,omitempty"`
	Cvv_message string `json:"cvv_message,omitempty"`
	Description string `json:"description,omitempty"`
	Is_test_mode bool `json:"is_test_mode,omitempty"`
	Reference_number string `json:"reference_number,omitempty"`
	Card_bin string `json:"card_bin,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Gateway string `json:"gateway,omitempty"`
	Settlement_created_time A2CDateTime `json:"settlement_created_time,omitempty"`
	Currency string `json:"currency,omitempty"`
	Order_id string `json:"order_id,omitempty"`
	Parent_id string `json:"parent_id,omitempty"`
	Settlement_amount float64 `json:"settlement_amount,omitempty"`
	Avs_street_resp_code string `json:"avs_street_resp_code,omitempty"`
	Avs_message string `json:"avs_message,omitempty"`
	Transaction_id string `json:"transaction_id,omitempty"`
	Card_last_four string `json:"card_last_four,omitempty"`
	Settlement_currency string `json:"settlement_currency,omitempty"`
	Status string `json:"status,omitempty"`
}

// ResponseOrderPreestimateShippingListResult represents the ResponseOrderPreestimateShippingListResult schema from the OpenAPI specification
type ResponseOrderPreestimateShippingListResult struct {
	Preestimate_shippings []OrderPreestimateShipping `json:"preestimate_shippings,omitempty"`
	Preestimate_shippings_count int `json:"preestimate_shippings_count,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
}

// A2CDateTime represents the A2CDateTime schema from the OpenAPI specification
type A2CDateTime struct {
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Format string `json:"format,omitempty"`
	Value string `json:"value,omitempty"`
}

// Carrier represents the Carrier schema from the OpenAPI specification
type Carrier struct {
	Active bool `json:"active,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Shipping_methods []OrderShippingMethod `json:"shipping_methods,omitempty"`
}

// CouponCondition represents the CouponCondition schema from the OpenAPI specification
type CouponCondition struct {
	Entity string `json:"entity,omitempty"`
	Operator string `json:"operator,omitempty"`
	Sub_conditions []CouponCondition `json:"sub-conditions,omitempty"`
	Key string `json:"key,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Id string `json:"id,omitempty"`
	Match_items string `json:"match_items,omitempty"`
	Logic_operator string `json:"logic_operator,omitempty"`
	Value string `json:"value,omitempty"`
}

// ResponseOrderShipmentListResult represents the ResponseOrderShipmentListResult schema from the OpenAPI specification
type ResponseOrderShipmentListResult struct {
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Shipment []Shipment `json:"shipment,omitempty"`
	Shipment_count int `json:"shipment_count,omitempty"`
}

// ModelResponseOrderTransactionList represents the ModelResponseOrderTransactionList schema from the OpenAPI specification
type ModelResponseOrderTransactionList struct {
	Return_code int `json:"return_code,omitempty"`
	Return_message string `json:"return_message,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
	Result ResponseOrderTransactionListResult `json:"result,omitempty"`
}

// OrderTotalsNewDiscount represents the OrderTotalsNewDiscount schema from the OpenAPI specification
type OrderTotalsNewDiscount struct {
	Value float64 `json:"value,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Code string `json:"code,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	TypeField string `json:"type,omitempty"`
}

// ShipmentItem represents the ShipmentItem schema from the OpenAPI specification
type ShipmentItem struct {
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Name string `json:"name,omitempty"`
	Order_product_id string `json:"order_product_id,omitempty"`
	Variant_id string `json:"variant_id,omitempty"`
	Price float64 `json:"price,omitempty"`
	Product_id string `json:"product_id,omitempty"`
	Model string `json:"model,omitempty"`
	Quantity float64 `json:"quantity,omitempty"`
}

// Child represents the Child schema from the OpenAPI specification
type Child struct {
	Images []Image `json:"images,omitempty"`
	List_price float64 `json:"list_price,omitempty"`
	Inventory_level float64 `json:"inventory_level,omitempty"`
	Modified_time A2CDateTime `json:"modified_time,omitempty"`
	Weight_unit string `json:"weight_unit,omitempty"`
	Length float64 `json:"length,omitempty"`
	Url string `json:"url,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Isbn string `json:"isbn,omitempty"`
	Wholesale_price float64 `json:"wholesale_price,omitempty"`
	Parent_id string `json:"parent_id,omitempty"`
	Upc string `json:"upc,omitempty"`
	Full_description string `json:"full_description,omitempty"`
	Meta_description string `json:"meta_description,omitempty"`
	Is_qty_in_pack_fixed bool `json:"is_qty_in_pack_fixed,omitempty"`
	Gtin string `json:"gtin,omitempty"`
	Advanced_price []ProductAdvancedPrice `json:"advanced_price,omitempty"`
	Default_qty_in_pack float64 `json:"default_qty_in_pack,omitempty"`
	Dimensions_unit string `json:"dimensions_unit,omitempty"`
	In_stock bool `json:"in_stock,omitempty"`
	Default_price float64 `json:"default_price,omitempty"`
	Width float64 `json:"width,omitempty"`
	Short_description string `json:"short_description,omitempty"`
	Weight float64 `json:"weight,omitempty"`
	Cost_price float64 `json:"cost_price,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Ean string `json:"ean,omitempty"`
	Discounts []Discount `json:"discounts,omitempty"`
	Inventory []ProductInventory `json:"inventory,omitempty"`
	Sku string `json:"sku,omitempty"`
	Min_quantity float64 `json:"min_quantity,omitempty"`
	Manage_stock bool `json:"manage_stock,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Sort_order int `json:"sort_order,omitempty"`
	Tax_class_id string `json:"tax_class_id,omitempty"`
	Meta_keywords string `json:"meta_keywords,omitempty"`
	Created_time A2CDateTime `json:"created_time,omitempty"`
	Meta_title string `json:"meta_title,omitempty"`
	Avail_for_sale bool `json:"avail_for_sale,omitempty"`
	Height float64 `json:"height,omitempty"`
	Mpn string `json:"mpn,omitempty"`
	Allow_backorders bool `json:"allow_backorders,omitempty"`
	Combination []ProductChildItemCombination `json:"combination,omitempty"`
}

// ResponseOrderAbandonedListResult represents the ResponseOrderAbandonedListResult schema from the OpenAPI specification
type ResponseOrderAbandonedListResult struct {
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Order []OrderAbandoned `json:"order,omitempty"`
}

// ModelResponseAttributeList represents the ModelResponseAttributeList schema from the OpenAPI specification
type ModelResponseAttributeList struct {
	Pagination Pagination `json:"pagination,omitempty"`
	Result ResponseAttributeListResult `json:"result,omitempty"`
	Return_code int `json:"return_code,omitempty"`
	Return_message string `json:"return_message,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
}

// CartWarehouse represents the CartWarehouse schema from the OpenAPI specification
type CartWarehouse struct {
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Description string `json:"description,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Stores_ids []string `json:"stores_ids,omitempty"`
	Address CustomerAddress `json:"address,omitempty"`
	Avail bool `json:"avail,omitempty"`
	Carriers_ids []string `json:"carriers_ids,omitempty"`
}

// OrderShipmentUpdate represents the OrderShipmentUpdate schema from the OpenAPI specification
type OrderShipmentUpdate struct {
	Is_shipped bool `json:"is_shipped,omitempty"` // Defines shipment's status
	Order_id string `json:"order_id,omitempty"` // Defines the order that will be updated
	Replace bool `json:"replace,omitempty"` // Allows rewrite tracking numbers
	Shipment_id string `json:"shipment_id"` // Shipment id indicates the number of delivery
	Store_id string `json:"store_id,omitempty"` // Store Id
	Tracking_link string `json:"tracking_link,omitempty"` // Defines custom tracking link
	Tracking_numbers []map[string]interface{} `json:"tracking_numbers,omitempty"` // Defines shipment's tracking numbers that have to be added</br> How set tracking numbers to appropriate carrier:<ul><li>tracking_numbers[]=a2c.demo1,a2c.demo2 - set default carrier</li><li>tracking_numbers[<b>carrier_id</b>]=a2c.demo - set appropriate carrier</li></ul>To get the list of carriers IDs that are available in your store, use the <a href = "https://api2cart.com/docs/#/cart/CartInfo">cart.info</a > method
}

// TaxClassStates represents the TaxClassStates schema from the OpenAPI specification
type TaxClassStates struct {
	Code string `json:"code,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Tax float64 `json:"tax,omitempty"`
	Tax_type int `json:"tax_type,omitempty"`
	Zip_codes []TaxClassZipCodes `json:"zip_codes,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
}

// SpecialPrice represents the SpecialPrice schema from the OpenAPI specification
type SpecialPrice struct {
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Expired_at A2CDateTime `json:"expired_at,omitempty"`
	Modified_at A2CDateTime `json:"modified_at,omitempty"`
	Value float64 `json:"value,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Avail bool `json:"avail,omitempty"`
	Created_at A2CDateTime `json:"created_at,omitempty"`
}

// ResponseCustomerAttributeListResult represents the ResponseCustomerAttributeListResult schema from the OpenAPI specification
type ResponseCustomerAttributeListResult struct {
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Items []CustomerAttribute `json:"items,omitempty"`
	Total_count int `json:"total_count,omitempty"`
}

// ResponseCategoryListResult represents the ResponseCategoryListResult schema from the OpenAPI specification
type ResponseCategoryListResult struct {
	Category []Category `json:"category,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Categories_count int `json:"categories_count,omitempty"`
}

// ResponseAttributeListResult represents the ResponseAttributeListResult schema from the OpenAPI specification
type ResponseAttributeListResult struct {
	Attributes_count int `json:"attributes_count,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Attribute []StoreAttribute `json:"attribute,omitempty"`
}

// OrderRefund represents the OrderRefund schema from the OpenAPI specification
type OrderRefund struct {
	Fee float64 `json:"fee,omitempty"`
	Id string `json:"id,omitempty"`
	Tax float64 `json:"tax,omitempty"`
	Total float64 `json:"total,omitempty"`
	Items []OrderStatusRefundItem `json:"items,omitempty"`
	Shipping float64 `json:"shipping,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Comment string `json:"comment,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Modified_time A2CDateTime `json:"modified_time,omitempty"`
}

// Basket represents the Basket schema from the OpenAPI specification
type Basket struct {
	Basket_url string `json:"basket_url,omitempty"`
	Id string `json:"id,omitempty"`
	Modified_at A2CDateTime `json:"modified_at,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Created_at A2CDateTime `json:"created_at,omitempty"`
	Currency Currency `json:"currency,omitempty"`
	Customer BaseCustomer `json:"customer,omitempty"`
	Basket_products []BasketItem `json:"basket_products,omitempty"`
}

// ModelResponseProductChildItemList represents the ModelResponseProductChildItemList schema from the OpenAPI specification
type ModelResponseProductChildItemList struct {
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
	Result ResponseProductChildItemListResult `json:"result,omitempty"`
	Return_code int `json:"return_code,omitempty"`
	Return_message string `json:"return_message,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
}

// ResponseCartGiftcardListResult represents the ResponseCartGiftcardListResult schema from the OpenAPI specification
type ResponseCartGiftcardListResult struct {
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Gift_card []GiftCard `json:"gift_card,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
}

// Currency represents the Currency schema from the OpenAPI specification
type Currency struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Iso3 string `json:"iso3,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Avail bool `json:"avail,omitempty"`
	DefaultField bool `json:"default,omitempty"`
	Rate float64 `json:"rate,omitempty"`
	Symbol_left string `json:"symbol_left,omitempty"`
	Symbol_right string `json:"symbol_right,omitempty"`
}

// Country represents the Country schema from the OpenAPI specification
type Country struct {
	Name string `json:"name,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Code2 string `json:"code2,omitempty"`
	Code3 string `json:"code3,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
}

// ResponseCartScriptListResult represents the ResponseCartScriptListResult schema from the OpenAPI specification
type ResponseCartScriptListResult struct {
	Scripts []Script `json:"scripts,omitempty"`
	Total_count int `json:"total_count,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
}

// PluginList represents the PluginList schema from the OpenAPI specification
type PluginList struct {
	Plugins []Plugin `json:"plugins,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	All_plugins int `json:"all_plugins,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
}

// BaseCustomer represents the BaseCustomer schema from the OpenAPI specification
type BaseCustomer struct {
	Phone string `json:"phone,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Email string `json:"email,omitempty"`
	First_name string `json:"first_name,omitempty"`
	Id string `json:"id,omitempty"`
	Last_name string `json:"last_name,omitempty"`
}

// CustomerGroup represents the CustomerGroup schema from the OpenAPI specification
type CustomerGroup struct {
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// ResponseCustomerWishListListResult represents the ResponseCustomerWishListListResult schema from the OpenAPI specification
type ResponseCustomerWishListListResult struct {
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Total_count int `json:"total_count,omitempty"`
	Wish_lists []CustomerWishList `json:"wish_lists,omitempty"`
}

// CustomerAdd represents the CustomerAdd schema from the OpenAPI specification
type CustomerAdd struct {
	Gender string `json:"gender,omitempty"` // Defines customer's gender
	Password string `json:"password,omitempty"` // Defines customer's unique password
	Store_id string `json:"store_id,omitempty"` // Store Id
	Website string `json:"website,omitempty"` // Link to customer website
	Modified_time string `json:"modified_time,omitempty"` // Entity's date modification
	First_name string `json:"first_name"` // Defines customer's first name
	Status string `json:"status,omitempty"` // Defines customer's status
	Created_time string `json:"created_time,omitempty"` // Entity's date creation
	News_letter_subscription bool `json:"news_letter_subscription,omitempty"` // Defines whether the newsletter subscription is available for the user
	Phone string `json:"phone,omitempty"` // Defines customer's phone number
	Company string `json:"company,omitempty"` // Defines customer's company
	Last_name string `json:"last_name"` // Defines customer's last name
	Birth_day string `json:"birth_day,omitempty"` // Defines customer's birthday
	Fax string `json:"fax,omitempty"` // Defines customer's fax
	Group string `json:"group,omitempty"` // Defines the group where the customer
	Last_login string `json:"last_login,omitempty"` // Defines customer's last login time
	Address []map[string]interface{} `json:"address,omitempty"`
	Login string `json:"login,omitempty"` // Specifies customer's login name
	Email string `json:"email"` // Defines customer's email
}

// ProductVariantPriceUpdate represents the ProductVariantPriceUpdate schema from the OpenAPI specification
type ProductVariantPriceUpdate struct {
	Id string `json:"id,omitempty"` // Defines the variant where the price has to be updated
	Product_id string `json:"product_id,omitempty"` // Product id
	Group_prices []map[string]interface{} `json:"group_prices"` // Defines variants's group prices
}

// ProductImageAdd represents the ProductImageAdd schema from the OpenAPI specification
type ProductImageAdd struct {
	Mime string `json:"mime,omitempty"` // Mime type of image http://en.wikipedia.org/wiki/Internet_media_type.
	Product_variant_id int `json:"product_variant_id,omitempty"` // Defines product's variants specified by variant id
	Url string `json:"url,omitempty"` // Defines URL of the image that has to be added
	Variant_ids string `json:"variant_ids,omitempty"` // Defines product's variants ids
	Content string `json:"content,omitempty"` // Content(body) encoded in base64 of image file
	Lang_id string `json:"lang_id,omitempty"` // Add product image on specified language id
	Product_id string `json:"product_id,omitempty"` // Defines product id where the image should be added
	Position int `json:"position,omitempty"` // Defines image’s position in the list
	Store_id string `json:"store_id,omitempty"` // Store Id
	TypeField string `json:"type"` // Defines image's types that are specified by comma-separated list
	Image_name string `json:"image_name"` // Defines image's name
	Label string `json:"label,omitempty"` // Defines alternative text that has to be attached to the picture
}

// CartCouponAdd represents the CartCouponAdd schema from the OpenAPI specification
type CartCouponAdd struct {
	Code string `json:"code"` // Coupon code
	Action_amount float64 `json:"action_amount"` // Defines the discount amount value.
	Name string `json:"name,omitempty"` // Coupon name
	Action_scope string `json:"action_scope"` // Specify how discount should be applied. If scope=matching_items, then discount will be applied to each of the items that match action conditions. Scope order means that discount will be applied once.
	Action_type string `json:"action_type"` // Coupon discount type
	Usage_limit int `json:"usage_limit,omitempty"` // Usage limit for coupon.
	Action_apply_to string `json:"action_apply_to"` // Defines where discount should be applied
	Date_end string `json:"date_end,omitempty"` // Defines when discount code will be expired.
	Usage_limit_per_customer int `json:"usage_limit_per_customer,omitempty"` // Usage limit per customer.
	Action_condition_operator string `json:"action_condition_operator,omitempty"` // Defines condition operator.
	Codes []string `json:"codes,omitempty"` // Entity codes
	Action_condition_entity string `json:"action_condition_entity,omitempty"` // Defines entity for action condition.
	Store_id string `json:"store_id,omitempty"` // Store Id
	Action_condition_key string `json:"action_condition_key,omitempty"` // Defines entity attribute code for action condition.
	Date_start string `json:"date_start,omitempty"` // Defines when discount code will be available.
	Action_condition_value string `json:"action_condition_value,omitempty"` // Defines condition attribute value/s. Can be comma separated string.
}

// Webhook represents the Webhook schema from the OpenAPI specification
type Webhook struct {
	Label string `json:"label,omitempty"`
	Store_id string `json:"store_id,omitempty"`
	Action string `json:"action,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Entity string `json:"entity,omitempty"`
	Active bool `json:"active,omitempty"`
	Callback string `json:"callback,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Id int `json:"id,omitempty"`
	Fields string `json:"fields,omitempty"`
}

// OrderItem represents the OrderItem schema from the OpenAPI specification
type OrderItem struct {
	Tax_value_after_discount float64 `json:"tax_value_after_discount,omitempty"`
	Variant_id string `json:"variant_id,omitempty"`
	Price float64 `json:"price,omitempty"`
	Price_inc_tax float64 `json:"price_inc_tax,omitempty"`
	Product_id string `json:"product_id,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Tax_value float64 `json:"tax_value,omitempty"`
	Barcode string `json:"barcode,omitempty"`
	Weight float64 `json:"weight,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Model string `json:"model,omitempty"`
	Weight_unit string `json:"weight_unit,omitempty"`
	Name string `json:"name,omitempty"`
	Quantity float64 `json:"quantity,omitempty"`
	Tax_percent float64 `json:"tax_percent,omitempty"`
	Total_price float64 `json:"total_price,omitempty"`
	Discount_amount float64 `json:"discount_amount,omitempty"`
	Parent_order_product_id string `json:"parent_order_product_id,omitempty"`
	Options []OrderItemOption `json:"options,omitempty"`
	Order_product_id string `json:"order_product_id,omitempty"`
}

// CouponCode represents the CouponCode schema from the OpenAPI specification
type CouponCode struct {
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Code string `json:"code,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Id string `json:"id,omitempty"`
	Used_times int `json:"used_times,omitempty"`
}

// ModelResponseCartGiftCardList represents the ModelResponseCartGiftCardList schema from the OpenAPI specification
type ModelResponseCartGiftCardList struct {
	Return_message string `json:"return_message,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
	Result ResponseCartGiftcardListResult `json:"result,omitempty"`
	Return_code int `json:"return_code,omitempty"`
}

// Discount represents the Discount schema from the OpenAPI specification
type Discount struct {
	To_time string `json:"to_time,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
	From_time string `json:"from_time,omitempty"`
	Value float64 `json:"value,omitempty"`
	Modifier_type string `json:"modifier_type,omitempty"`
	Name string `json:"name,omitempty"`
	Sort_order int `json:"sort_order,omitempty"`
	Customer_group_ids string `json:"customer_group_ids,omitempty"`
	Id string `json:"id,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
}

// BasketItem represents the BasketItem schema from the OpenAPI specification
type BasketItem struct {
	Name string `json:"name,omitempty"`
	Parent_id string `json:"parent_id,omitempty"`
	Quantity float64 `json:"quantity,omitempty"`
	Tax float64 `json:"tax,omitempty"`
	Weight float64 `json:"weight,omitempty"`
	Price float64 `json:"price,omitempty"`
	Weight_unit string `json:"weight_unit,omitempty"`
	Custom_fields map[string]interface{} `json:"custom_fields,omitempty"`
	Options []BasketItemOption `json:"options,omitempty"`
	Id string `json:"id,omitempty"`
	Product_id string `json:"product_id,omitempty"`
	Sku string `json:"sku,omitempty"`
	Variant_id string `json:"variant_id,omitempty"`
	Additional_fields map[string]interface{} `json:"additional_fields,omitempty"`
}
