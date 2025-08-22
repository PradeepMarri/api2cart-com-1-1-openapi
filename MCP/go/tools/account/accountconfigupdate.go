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

func AccountconfigupdateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["new_store_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("new_store_key=%v", val))
		}
		if val, ok := args["bridge_url"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("bridge_url=%v", val))
		}
		if val, ok := args["store_root"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("store_root=%v", val))
		}
		if val, ok := args["db_tables_prefix"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("db_tables_prefix=%v", val))
		}
		if val, ok := args["3dcart_private_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("3dcart_private_key=%v", val))
		}
		if val, ok := args["3dcart_access_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("3dcart_access_token=%v", val))
		}
		if val, ok := args["3dcartapi_api_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("3dcartapi_api_key=%v", val))
		}
		if val, ok := args["amazon_sp_client_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("amazon_sp_client_id=%v", val))
		}
		if val, ok := args["amazon_sp_client_secret"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("amazon_sp_client_secret=%v", val))
		}
		if val, ok := args["amazon_sp_aws_user_key_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("amazon_sp_aws_user_key_id=%v", val))
		}
		if val, ok := args["amazon_sp_aws_user_secret"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("amazon_sp_aws_user_secret=%v", val))
		}
		if val, ok := args["amazon_sp_aws_region"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("amazon_sp_aws_region=%v", val))
		}
		if val, ok := args["amazon_sp_aws_role_arn"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("amazon_sp_aws_role_arn=%v", val))
		}
		if val, ok := args["amazon_sp_refresh_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("amazon_sp_refresh_token=%v", val))
		}
		if val, ok := args["amazon_sp_api_environment"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("amazon_sp_api_environment=%v", val))
		}
		if val, ok := args["amazon_access_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("amazon_access_token=%v", val))
		}
		if val, ok := args["amazon_seller_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("amazon_seller_id=%v", val))
		}
		if val, ok := args["amazon_marketplaces_ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("amazon_marketplaces_ids=%v", val))
		}
		if val, ok := args["amazon_secret_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("amazon_secret_key=%v", val))
		}
		if val, ok := args["amazon_access_key_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("amazon_access_key_id=%v", val))
		}
		if val, ok := args["aspdotnetstorefront_api_user"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("aspdotnetstorefront_api_user=%v", val))
		}
		if val, ok := args["aspdotnetstorefront_api_pass"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("aspdotnetstorefront_api_pass=%v", val))
		}
		if val, ok := args["bigcommerceapi_admin_account"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("bigcommerceapi_admin_account=%v", val))
		}
		if val, ok := args["bigcommerceapi_api_path"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("bigcommerceapi_api_path=%v", val))
		}
		if val, ok := args["bigcommerceapi_api_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("bigcommerceapi_api_key=%v", val))
		}
		if val, ok := args["bigcommerceapi_client_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("bigcommerceapi_client_id=%v", val))
		}
		if val, ok := args["bigcommerceapi_access_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("bigcommerceapi_access_token=%v", val))
		}
		if val, ok := args["bigcommerceapi_context"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("bigcommerceapi_context=%v", val))
		}
		if val, ok := args["demandware_client_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("demandware_client_id=%v", val))
		}
		if val, ok := args["demandware_api_password"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("demandware_api_password=%v", val))
		}
		if val, ok := args["demandware_user_name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("demandware_user_name=%v", val))
		}
		if val, ok := args["demandware_user_password"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("demandware_user_password=%v", val))
		}
		if val, ok := args["ebay_client_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ebay_client_id=%v", val))
		}
		if val, ok := args["ebay_client_secret"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ebay_client_secret=%v", val))
		}
		if val, ok := args["ebay_runame"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ebay_runame=%v", val))
		}
		if val, ok := args["ebay_access_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ebay_access_token=%v", val))
		}
		if val, ok := args["ebay_refresh_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ebay_refresh_token=%v", val))
		}
		if val, ok := args["ebay_environment"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ebay_environment=%v", val))
		}
		if val, ok := args["ebay_site_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ebay_site_id=%v", val))
		}
		if val, ok := args["ecwid_acess_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ecwid_acess_token=%v", val))
		}
		if val, ok := args["ecwid_store_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ecwid_store_id=%v", val))
		}
		if val, ok := args["etsy_keystring"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("etsy_keystring=%v", val))
		}
		if val, ok := args["etsy_shared_secret"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("etsy_shared_secret=%v", val))
		}
		if val, ok := args["etsy_access_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("etsy_access_token=%v", val))
		}
		if val, ok := args["etsy_token_secret"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("etsy_token_secret=%v", val))
		}
		if val, ok := args["etsy_client_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("etsy_client_id=%v", val))
		}
		if val, ok := args["etsy_refresh_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("etsy_refresh_token=%v", val))
		}
		if val, ok := args["neto_api_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("neto_api_key=%v", val))
		}
		if val, ok := args["neto_api_username"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("neto_api_username=%v", val))
		}
		if val, ok := args["shopify_api_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("shopify_api_key=%v", val))
		}
		if val, ok := args["shopify_api_password"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("shopify_api_password=%v", val))
		}
		if val, ok := args["shopify_shared_secret"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("shopify_shared_secret=%v", val))
		}
		if val, ok := args["shopify_access_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("shopify_access_token=%v", val))
		}
		if val, ok := args["shopware_access_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("shopware_access_key=%v", val))
		}
		if val, ok := args["shopware_api_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("shopware_api_key=%v", val))
		}
		if val, ok := args["shopware_api_secret"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("shopware_api_secret=%v", val))
		}
		if val, ok := args["volusion_login"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("volusion_login=%v", val))
		}
		if val, ok := args["volusion_password"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("volusion_password=%v", val))
		}
		if val, ok := args["walmart_client_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("walmart_client_id=%v", val))
		}
		if val, ok := args["walmart_client_secret"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("walmart_client_secret=%v", val))
		}
		if val, ok := args["walmart_environment"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("walmart_environment=%v", val))
		}
		if val, ok := args["walmart_channel_type"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("walmart_channel_type=%v", val))
		}
		if val, ok := args["squarespace_api_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("squarespace_api_key=%v", val))
		}
		if val, ok := args["hybris_client_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("hybris_client_id=%v", val))
		}
		if val, ok := args["hybris_client_secret"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("hybris_client_secret=%v", val))
		}
		if val, ok := args["hybris_username"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("hybris_username=%v", val))
		}
		if val, ok := args["hybris_password"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("hybris_password=%v", val))
		}
		if val, ok := args["hybris_websites"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("hybris_websites=%v", val))
		}
		if val, ok := args["lightspeed_api_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("lightspeed_api_key=%v", val))
		}
		if val, ok := args["lightspeed_api_secret"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("lightspeed_api_secret=%v", val))
		}
		if val, ok := args["commercehq_api_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("commercehq_api_key=%v", val))
		}
		if val, ok := args["commercehq_api_password"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("commercehq_api_password=%v", val))
		}
		if val, ok := args["wc_consumer_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("wc_consumer_key=%v", val))
		}
		if val, ok := args["wc_consumer_secret"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("wc_consumer_secret=%v", val))
		}
		if val, ok := args["magento_consumer_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("magento_consumer_key=%v", val))
		}
		if val, ok := args["magento_consumer_secret"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("magento_consumer_secret=%v", val))
		}
		if val, ok := args["magento_access_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("magento_access_token=%v", val))
		}
		if val, ok := args["magento_token_secret"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("magento_token_secret=%v", val))
		}
		if val, ok := args["prestashop_webservice_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("prestashop_webservice_key=%v", val))
		}
		if val, ok := args["wix_app_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("wix_app_id=%v", val))
		}
		if val, ok := args["wix_app_secret_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("wix_app_secret_key=%v", val))
		}
		if val, ok := args["wix_refresh_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("wix_refresh_token=%v", val))
		}
		if val, ok := args["mercado_libre_app_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("mercado_libre_app_id=%v", val))
		}
		if val, ok := args["mercado_libre_app_secret_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("mercado_libre_app_secret_key=%v", val))
		}
		if val, ok := args["mercado_libre_refresh_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("mercado_libre_refresh_token=%v", val))
		}
		if val, ok := args["zid_client_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("zid_client_id=%v", val))
		}
		if val, ok := args["zid_client_secret"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("zid_client_secret=%v", val))
		}
		if val, ok := args["zid_access_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("zid_access_token=%v", val))
		}
		if val, ok := args["zid_authorization"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("zid_authorization=%v", val))
		}
		if val, ok := args["zid_refresh_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("zid_refresh_token=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/account.config.update.json%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("PUT", url, nil)
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

func CreateAccountconfigupdateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_account_config_update_json",
		mcp.WithDescription("Update configs in the API2Cart database."),
		mcp.WithString("new_store_key", mcp.Description("Update store key")),
		mcp.WithString("bridge_url", mcp.Description("This parameter allows to set up store with custom bridge url (also you must use store_root parameter if a bridge folder is not in the root folder of the store)")),
		mcp.WithString("store_root", mcp.Description("Absolute path to the store root directory (used with \"bridge_url\" parameter)")),
		mcp.WithString("db_tables_prefix", mcp.Description("DB tables prefix")),
		mcp.WithString("3dcart_private_key", mcp.Description("3DCart Private Key")),
		mcp.WithString("3dcart_access_token", mcp.Description("3DCart Token")),
		mcp.WithString("3dcartapi_api_key", mcp.Description("3DCart API Key")),
		mcp.WithString("amazon_sp_client_id", mcp.Description("Amazon SP API app client id")),
		mcp.WithString("amazon_sp_client_secret", mcp.Description("Amazon SP API app client secret")),
		mcp.WithString("amazon_sp_aws_user_key_id", mcp.Description("Amazon AWS user access key ID")),
		mcp.WithString("amazon_sp_aws_user_secret", mcp.Description("Amazon AWS user secret access key")),
		mcp.WithString("amazon_sp_aws_region", mcp.Description("Amazon AWS Region")),
		mcp.WithString("amazon_sp_aws_role_arn", mcp.Description("Amazon AWS Role ARN")),
		mcp.WithString("amazon_sp_refresh_token", mcp.Description("Amazon SP API OAuth refresh token")),
		mcp.WithString("amazon_sp_api_environment", mcp.Description("Amazon SP API environment")),
		mcp.WithString("amazon_access_token", mcp.Description("MWS Auth Token. Access token authorizing the app to access resources on behalf of a user")),
		mcp.WithString("amazon_seller_id", mcp.Description("Amazon Seller ID (Merchant token)")),
		mcp.WithString("amazon_marketplaces_ids", mcp.Description("Amazon Marketplace IDs comma separated string")),
		mcp.WithString("amazon_secret_key", mcp.Description("Amazon Secret Key")),
		mcp.WithString("amazon_access_key_id", mcp.Description("Amazon Secret Key Id")),
		mcp.WithString("aspdotnetstorefront_api_user", mcp.Description("It's a AspDotNetStorefront account for which API is available")),
		mcp.WithString("aspdotnetstorefront_api_pass", mcp.Description("AspDotNetStorefront API Password")),
		mcp.WithString("bigcommerceapi_admin_account", mcp.Description("It's a BigCommerce account for which API is enabled")),
		mcp.WithString("bigcommerceapi_api_path", mcp.Description("BigCommerce API URL")),
		mcp.WithString("bigcommerceapi_api_key", mcp.Description("Bigcommerce API Key")),
		mcp.WithString("bigcommerceapi_client_id", mcp.Description("Client ID of the requesting app")),
		mcp.WithString("bigcommerceapi_access_token", mcp.Description("Access token authorizing the app to access resources on behalf of a user")),
		mcp.WithString("bigcommerceapi_context", mcp.Description("API Path section unique to the store")),
		mcp.WithString("demandware_client_id", mcp.Description("Demandware client id")),
		mcp.WithString("demandware_api_password", mcp.Description("Demandware api password")),
		mcp.WithString("demandware_user_name", mcp.Description("Demandware user name")),
		mcp.WithString("demandware_user_password", mcp.Description("Demandware user password")),
		mcp.WithString("ebay_client_id", mcp.Description("Application ID (AppID).")),
		mcp.WithString("ebay_client_secret", mcp.Description("Shared Secret from eBay application")),
		mcp.WithString("ebay_runame", mcp.Description("The RuName value that eBay assigns to your application.")),
		mcp.WithString("ebay_access_token", mcp.Description("Used to authenticate API requests.")),
		mcp.WithString("ebay_refresh_token", mcp.Description("Used to renew the access token.")),
		mcp.WithString("ebay_environment", mcp.Description("eBay environment")),
		mcp.WithNumber("ebay_site_id", mcp.Description("eBay global ID")),
		mcp.WithString("ecwid_acess_token", mcp.Description("Access token authorizing the app to access resources on behalf of a user")),
		mcp.WithString("ecwid_store_id", mcp.Description("Store Id")),
		mcp.WithString("etsy_keystring", mcp.Description("Etsy keystring")),
		mcp.WithString("etsy_shared_secret", mcp.Description("Etsy shared secret")),
		mcp.WithString("etsy_access_token", mcp.Description("Access token authorizing the app to access resources on behalf of a user")),
		mcp.WithString("etsy_token_secret", mcp.Description("Secret token authorizing the app to access resources on behalf of a user")),
		mcp.WithString("etsy_client_id", mcp.Description("Etsy Client Id")),
		mcp.WithString("etsy_refresh_token", mcp.Description("Etsy Refresh token")),
		mcp.WithString("neto_api_key", mcp.Description("Neto API Key")),
		mcp.WithString("neto_api_username", mcp.Description("Neto User Name")),
		mcp.WithString("shopify_api_key", mcp.Description("Shopify API Key")),
		mcp.WithString("shopify_api_password", mcp.Description("Shopify API Password")),
		mcp.WithString("shopify_shared_secret", mcp.Description("Shared secret")),
		mcp.WithString("shopify_access_token", mcp.Description("Access token authorizing the app to access resources on behalf of a user")),
		mcp.WithString("shopware_access_key", mcp.Description("Shopware access key")),
		mcp.WithString("shopware_api_key", mcp.Description("Shopware api key")),
		mcp.WithString("shopware_api_secret", mcp.Description("Shopware client secret access key")),
		mcp.WithString("volusion_login", mcp.Description("It's a Volusion account for which API is enabled")),
		mcp.WithString("volusion_password", mcp.Description("Volusion API Password")),
		mcp.WithString("walmart_client_id", mcp.Description("Walmart client ID")),
		mcp.WithString("walmart_client_secret", mcp.Description("Walmart client secret")),
		mcp.WithString("walmart_environment", mcp.Description("Walmart environment")),
		mcp.WithString("walmart_channel_type", mcp.Description("Walmart WM_CONSUMER.CHANNEL.TYPE header")),
		mcp.WithString("squarespace_api_key", mcp.Description("Squarespace API Key")),
		mcp.WithString("hybris_client_id", mcp.Description("Omni Commerce Connector Client ID")),
		mcp.WithString("hybris_client_secret", mcp.Description("Omni Commerce Connector Client Secret")),
		mcp.WithString("hybris_username", mcp.Description("User Name")),
		mcp.WithString("hybris_password", mcp.Description("User password")),
		mcp.WithArray("hybris_websites", mcp.Description("Websites to stores mapping data")),
		mcp.WithString("lightspeed_api_key", mcp.Description("LightSpeed api key")),
		mcp.WithString("lightspeed_api_secret", mcp.Description("LightSpeed api secret")),
		mcp.WithString("commercehq_api_key", mcp.Description("CommerceHQ api key")),
		mcp.WithString("commercehq_api_password", mcp.Description("CommerceHQ api password")),
		mcp.WithString("wc_consumer_key", mcp.Description("Woocommerce consumer key")),
		mcp.WithString("wc_consumer_secret", mcp.Description("Woocommerce consumer secret")),
		mcp.WithString("magento_consumer_key", mcp.Description("Magento Consumer Key")),
		mcp.WithString("magento_consumer_secret", mcp.Description("Magento Consumer Secret")),
		mcp.WithString("magento_access_token", mcp.Description("Magento Access Token")),
		mcp.WithString("magento_token_secret", mcp.Description("Magento Token Secret")),
		mcp.WithString("prestashop_webservice_key", mcp.Description("Prestashop webservice key")),
		mcp.WithString("wix_app_id", mcp.Description("Wix App ID")),
		mcp.WithString("wix_app_secret_key", mcp.Description("Wix App Secret Key")),
		mcp.WithString("wix_refresh_token", mcp.Description("Wix refresh token")),
		mcp.WithString("mercado_libre_app_id", mcp.Description("Mercado Libre App ID")),
		mcp.WithString("mercado_libre_app_secret_key", mcp.Description("Mercado Libre App Secret Key")),
		mcp.WithString("mercado_libre_refresh_token", mcp.Description("Mercado Libre Refresh Token")),
		mcp.WithNumber("zid_client_id", mcp.Description("Zid Client ID")),
		mcp.WithString("zid_client_secret", mcp.Description("Zid Client Secret")),
		mcp.WithString("zid_access_token", mcp.Description("Zid Access Token")),
		mcp.WithString("zid_authorization", mcp.Description("Zid Authorization")),
		mcp.WithString("zid_refresh_token", mcp.Description("Zid refresh token")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    AccountconfigupdateHandler(cfg),
	}
}
