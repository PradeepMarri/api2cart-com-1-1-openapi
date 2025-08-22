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

func AccountcartaddHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.AccountCartAdd
		
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
		url := fmt.Sprintf("%s/account.cart.add.json", cfg.BaseURL)
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

func CreateAccountcartaddTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_account_cart_add_json",
		mcp.WithDescription("Add store to the account"),
		mcp.WithString("etsy_keystring", mcp.Description("Input parameter: Etsy keystring")),
		mcp.WithString("zid_client_secret", mcp.Description("Input parameter: Zid Client Secret")),
		mcp.WithString("prestashop_webservice_key", mcp.Description("Input parameter: Prestashop webservice key")),
		mcp.WithString("ftp_store_dir", mcp.Description("Input parameter: FTP Store dir")),
		mcp.WithString("amazon_sp_api_environment", mcp.Description("Input parameter: Amazon SP API environment")),
		mcp.WithString("amazon_sp_client_secret", mcp.Required(), mcp.Description("Input parameter: Amazon SP API app client secret")),
		mcp.WithString("wix_app_secret_key", mcp.Description("Input parameter: Wix App Secret Key")),
		mcp.WithString("demandware_user_password", mcp.Description("Input parameter: Demandware user password")),
		mcp.WithString("demandware_client_id", mcp.Description("Input parameter: Demandware client id")),
		mcp.WithString("mercado_libre_app_id", mcp.Description("Input parameter: Mercado Libre App ID")),
		mcp.WithString("etsy_client_id", mcp.Required(), mcp.Description("Input parameter: Etsy Client Id")),
		mcp.WithString("ftp_host", mcp.Description("Input parameter: FTP connection host")),
		mcp.WithString("shopware_access_key", mcp.Description("Input parameter: Shopware access key")),
		mcp.WithBoolean("validate_version", mcp.Description("Input parameter: Specify if api2cart should validate cart version")),
		mcp.WithString("3dcartapi_api_key", mcp.Description("Input parameter: 3DCart API Key")),
		mcp.WithArray("hybris_websites", mcp.Description("Input parameter: Websites to stores mapping data")),
		mcp.WithString("wc_consumer_key", mcp.Description("Input parameter: Woocommerce consumer key")),
		mcp.WithString("walmart_channel_type", mcp.Description("Input parameter: Walmart WM_CONSUMER.CHANNEL.TYPE header")),
		mcp.WithString("commercehq_api_password", mcp.Description("Input parameter: CommerceHQ api password")),
		mcp.WithString("etsy_refresh_token", mcp.Required(), mcp.Description("Input parameter: Etsy Refresh token")),
		mcp.WithString("squarespace_api_key", mcp.Description("Input parameter: Squarespace API Key")),
		mcp.WithString("zid_authorization", mcp.Description("Input parameter: Zid Authorization")),
		mcp.WithString("hybris_client_id", mcp.Description("Input parameter: Omni Commerce Connector Client ID")),
		mcp.WithString("amazon_sp_aws_region", mcp.Required(), mcp.Description("Input parameter: Amazon AWS Region")),
		mcp.WithString("bigcommerceapi_api_key", mcp.Description("Input parameter: Bigcommerce API Key")),
		mcp.WithString("shopify_shared_secret", mcp.Description("Input parameter: Shared secret")),
		mcp.WithString("shopware_api_secret", mcp.Description("Input parameter: Shopware client secret access key")),
		mcp.WithString("volusion_login", mcp.Description("Input parameter: It's a Volusion account for which API is enabled")),
		mcp.WithString("magento_token_secret", mcp.Description("Input parameter: Magento Token Secret")),
		mcp.WithString("volusion_password", mcp.Description("Input parameter: Volusion API Password")),
		mcp.WithString("walmart_client_id", mcp.Description("Input parameter: Walmart client ID")),
		mcp.WithString("wix_refresh_token", mcp.Description("Input parameter: Wix refresh token")),
		mcp.WithString("ebay_client_id", mcp.Description("Input parameter: Application ID (AppID).")),
		mcp.WithString("magento_access_token", mcp.Description("Input parameter: Magento Access Token")),
		mcp.WithString("shopware_api_key", mcp.Description("Input parameter: Shopware api key")),
		mcp.WithString("zid_refresh_token", mcp.Description("Input parameter: Zid refresh token")),
		mcp.WithString("magento_consumer_secret", mcp.Description("Input parameter: Magento Consumer Secret")),
		mcp.WithString("amazon_sp_refresh_token", mcp.Required(), mcp.Description("Input parameter: Amazon SP API OAuth refresh token")),
		mcp.WithString("etsy_access_token", mcp.Description("Input parameter: Access token authorizing the app to access resources on behalf of a user")),
		mcp.WithString("ebay_runame", mcp.Description("Input parameter: The RuName value that eBay assigns to your application.")),
		mcp.WithString("3dcart_access_token", mcp.Description("Input parameter: 3DCart Token")),
		mcp.WithString("bigcommerceapi_client_id", mcp.Description("Input parameter: Client ID of the requesting app")),
		mcp.WithString("ecwid_acess_token", mcp.Description("Input parameter: Access token authorizing the app to access resources on behalf of a user")),
		mcp.WithString("lightspeed_api_secret", mcp.Description("Input parameter: LightSpeed api secret")),
		mcp.WithString("ebay_refresh_token", mcp.Description("Input parameter: Used to renew the access token.")),
		mcp.WithString("shopify_api_key", mcp.Description("Input parameter: Shopify API Key")),
		mcp.WithString("neto_api_key", mcp.Description("Input parameter: Neto API Key")),
		mcp.WithString("etsy_shared_secret", mcp.Description("Input parameter: Etsy shared secret")),
		mcp.WithString("amazon_sp_aws_role_arn", mcp.Required(), mcp.Description("Input parameter: Amazon AWS Role ARN")),
		mcp.WithString("aspdotnetstorefront_api_user", mcp.Description("Input parameter: It's a AspDotNetStorefront account for which API is available")),
		mcp.WithString("db_tables_prefix", mcp.Description("Input parameter: DB tables prefix")),
		mcp.WithString("amazon_access_token", mcp.Description("Input parameter: MWS Auth Token. Access token authorizing the app to access resources on behalf of a user")),
		mcp.WithString("ebay_client_secret", mcp.Description("Input parameter: Shared Secret from eBay application")),
		mcp.WithString("amazon_sp_client_id", mcp.Required(), mcp.Description("Input parameter: Amazon SP API app client id")),
		mcp.WithString("ebay_access_token", mcp.Description("Input parameter: Used to authenticate API requests.")),
		mcp.WithString("walmart_client_secret", mcp.Description("Input parameter: Walmart client secret")),
		mcp.WithString("wix_app_id", mcp.Description("Input parameter: Wix App ID")),
		mcp.WithString("etsy_token_secret", mcp.Description("Input parameter: Secret token authorizing the app to access resources on behalf of a user")),
		mcp.WithString("ecwid_store_id", mcp.Description("Input parameter: Store Id")),
		mcp.WithString("hybris_password", mcp.Description("Input parameter: User password")),
		mcp.WithString("amazon_sp_aws_user_secret", mcp.Required(), mcp.Description("Input parameter: Amazon AWS user secret access key")),
		mcp.WithString("ftp_user", mcp.Description("Input parameter: FTP User")),
		mcp.WithString("magento_consumer_key", mcp.Description("Input parameter: Magento Consumer Key")),
		mcp.WithBoolean("verify", mcp.Description("Input parameter: Enables or disables cart's verification")),
		mcp.WithString("amazon_access_key_id", mcp.Description("Input parameter: Amazon Secret Key Id")),
		mcp.WithNumber("zid_client_id", mcp.Description("Input parameter: Zid Client ID")),
		mcp.WithNumber("ftp_port", mcp.Description("Input parameter: FTP Port")),
		mcp.WithString("walmart_environment", mcp.Description("Input parameter: Walmart environment")),
		mcp.WithString("bigcommerceapi_context", mcp.Description("Input parameter: API Path section unique to the store")),
		mcp.WithString("hybris_username", mcp.Description("Input parameter: User Name")),
		mcp.WithString("aspdotnetstorefront_api_pass", mcp.Description("Input parameter: AspDotNetStorefront API Password")),
		mcp.WithString("ftp_password", mcp.Description("Input parameter: FTP Password")),
		mcp.WithString("amazon_seller_id", mcp.Description("Input parameter: Amazon Seller ID (Merchant token)")),
		mcp.WithString("commercehq_api_key", mcp.Description("Input parameter: CommerceHQ api key")),
		mcp.WithString("bigcommerceapi_access_token", mcp.Description("Input parameter: Access token authorizing the app to access resources on behalf of a user")),
		mcp.WithString("3dcart_private_key", mcp.Description("Input parameter: 3DCart Private Key")),
		mcp.WithString("amazon_marketplaces_ids", mcp.Description("Input parameter: Amazon Marketplace IDs comma separated string")),
		mcp.WithString("store_url", mcp.Required(), mcp.Description("Input parameter: A web address of a store that you would like to connect to API2Cart")),
		mcp.WithString("bigcommerceapi_api_path", mcp.Description("Input parameter: BigCommerce API URL")),
		mcp.WithString("ebay_environment", mcp.Description("Input parameter: eBay environment")),
		mcp.WithString("shopify_access_token", mcp.Description("Input parameter: Access token authorizing the app to access resources on behalf of a user")),
		mcp.WithString("mercado_libre_refresh_token", mcp.Description("Input parameter: Mercado Libre Refresh Token")),
		mcp.WithString("store_key", mcp.Description("Input parameter: Set this parameter if bridge is already uploaded to store")),
		mcp.WithString("demandware_user_name", mcp.Description("Input parameter: Demandware user name")),
		mcp.WithString("neto_api_username", mcp.Description("Input parameter: Neto User Name")),
		mcp.WithString("wc_consumer_secret", mcp.Description("Input parameter: Woocommerce consumer secret")),
		mcp.WithString("zid_access_token", mcp.Description("Input parameter: Zid Access Token")),
		mcp.WithString("bigcommerceapi_admin_account", mcp.Description("Input parameter: It's a BigCommerce account for which API is enabled")),
		mcp.WithString("bridge_url", mcp.Description("Input parameter: This parameter allows to set up store with custom bridge url (also you must use store_root parameter if a bridge folder is not in the root folder of the store)")),
		mcp.WithString("shopify_api_password", mcp.Description("Input parameter: Shopify API Password")),
		mcp.WithString("demandware_api_password", mcp.Description("Input parameter: Demandware api password")),
		mcp.WithString("mercado_libre_app_secret_key", mcp.Description("Input parameter: Mercado Libre App Secret Key")),
		mcp.WithString("amazon_secret_key", mcp.Description("Input parameter: Amazon Secret Key")),
		mcp.WithNumber("ebay_site_id", mcp.Description("Input parameter: eBay global ID")),
		mcp.WithString("hybris_client_secret", mcp.Description("Input parameter: Omni Commerce Connector Client Secret")),
		mcp.WithString("store_root", mcp.Description("Input parameter: Absolute path to the store root directory (used with \"bridge_url\" parameter)")),
		mcp.WithString("amazon_sp_aws_user_key_id", mcp.Required(), mcp.Description("Input parameter: Amazon AWS user access key ID")),
		mcp.WithString("cart_id", mcp.Required(), mcp.Description("Input parameter: Store’s identifier which you can get from cart_list method")),
		mcp.WithString("lightspeed_api_key", mcp.Description("Input parameter: LightSpeed api key")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    AccountcartaddHandler(cfg),
	}
}
