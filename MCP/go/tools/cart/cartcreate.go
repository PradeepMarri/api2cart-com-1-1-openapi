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

func CartcreateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["cart_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("cart_id=%v", val))
		}
		if val, ok := args["store_url"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("store_url=%v", val))
		}
		if val, ok := args["bridge_url"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("bridge_url=%v", val))
		}
		if val, ok := args["store_root"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("store_root=%v", val))
		}
		if val, ok := args["store_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("store_key=%v", val))
		}
		if val, ok := args["shared_secret"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("shared_secret=%v", val))
		}
		if val, ok := args["validate_version"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("validate_version=%v", val))
		}
		if val, ok := args["verify"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("verify=%v", val))
		}
		if val, ok := args["db_tables_prefix"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("db_tables_prefix=%v", val))
		}
		if val, ok := args["ftp_host"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ftp_host=%v", val))
		}
		if val, ok := args["ftp_user"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ftp_user=%v", val))
		}
		if val, ok := args["ftp_password"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ftp_password=%v", val))
		}
		if val, ok := args["ftp_port"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ftp_port=%v", val))
		}
		if val, ok := args["ftp_store_dir"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ftp_store_dir=%v", val))
		}
		if val, ok := args["apiKey_3dcart"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("apiKey_3dcart=%v", val))
		}
		if val, ok := args["AdminAccount"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("AdminAccount=%v", val))
		}
		if val, ok := args["ApiPath"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ApiPath=%v", val))
		}
		if val, ok := args["ApiKey"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ApiKey=%v", val))
		}
		if val, ok := args["client_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("client_id=%v", val))
		}
		if val, ok := args["accessToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("accessToken=%v", val))
		}
		if val, ok := args["context"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("context=%v", val))
		}
		if val, ok := args["access_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("access_token=%v", val))
		}
		if val, ok := args["apiKey_shopify"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("apiKey_shopify=%v", val))
		}
		if val, ok := args["apiPassword"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("apiPassword=%v", val))
		}
		if val, ok := args["accessToken_shopify"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("accessToken_shopify=%v", val))
		}
		if val, ok := args["apiKey"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("apiKey=%v", val))
		}
		if val, ok := args["apiUsername"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("apiUsername=%v", val))
		}
		if val, ok := args["EncryptedPassword"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("EncryptedPassword=%v", val))
		}
		if val, ok := args["Login"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("Login=%v", val))
		}
		if val, ok := args["apiUser_adnsf"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("apiUser_adnsf=%v", val))
		}
		if val, ok := args["apiPass"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("apiPass=%v", val))
		}
		if val, ok := args["privateKey"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("privateKey=%v", val))
		}
		if val, ok := args["appToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("appToken=%v", val))
		}
		if val, ok := args["etsy_keystring"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("etsy_keystring=%v", val))
		}
		if val, ok := args["etsy_shared_secret"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("etsy_shared_secret=%v", val))
		}
		if val, ok := args["tokenSecret"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("tokenSecret=%v", val))
		}
		if val, ok := args["etsy_client_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("etsy_client_id=%v", val))
		}
		if val, ok := args["etsy_refresh_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("etsy_refresh_token=%v", val))
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
		if val, ok := args["dw_client_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("dw_client_id=%v", val))
		}
		if val, ok := args["dw_api_pass"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("dw_api_pass=%v", val))
		}
		if val, ok := args["demandware_user_name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("demandware_user_name=%v", val))
		}
		if val, ok := args["demandware_user_password"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("demandware_user_password=%v", val))
		}
		if val, ok := args["store_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("store_id=%v", val))
		}
		if val, ok := args["seller_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("seller_id=%v", val))
		}
		if val, ok := args["amazon_secret_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("amazon_secret_key=%v", val))
		}
		if val, ok := args["amazon_access_key_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("amazon_access_key_id=%v", val))
		}
		if val, ok := args["marketplaces_ids"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("marketplaces_ids=%v", val))
		}
		if val, ok := args["environment"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("environment=%v", val))
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
		if val, ok := args["lightspeed_api_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("lightspeed_api_key=%v", val))
		}
		if val, ok := args["lightspeed_api_secret"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("lightspeed_api_secret=%v", val))
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
		if val, ok := args["commercehq_api_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("commercehq_api_key=%v", val))
		}
		if val, ok := args["commercehq_api_password"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("commercehq_api_password=%v", val))
		}
		if val, ok := args["3dcart_private_key"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("3dcart_private_key=%v", val))
		}
		if val, ok := args["3dcart_access_token"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("3dcart_access_token=%v", val))
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
		url := fmt.Sprintf("%s/cart.create.json%s", cfg.BaseURL, queryString)
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

func CreateCartcreateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_cart_create_json",
		mcp.WithDescription("Add store to the account"),
		mcp.WithString("cart_id", mcp.Required(), mcp.Description("Store’s identifier which you can get from cart_list method")),
		mcp.WithString("store_url", mcp.Required(), mcp.Description("A web address of a store that you would like to connect to API2Cart")),
		mcp.WithString("bridge_url", mcp.Description("This parameter allows to set up store with custom bridge url (also you must use store_root parameter if a bridge folder is not in the root folder of the store)")),
		mcp.WithString("store_root", mcp.Description("Absolute path to the store root directory (used with \"bridge_url\" parameter)")),
		mcp.WithString("store_key", mcp.Description("Set this parameter if bridge is already uploaded to store")),
		mcp.WithString("shared_secret", mcp.Description("Shared secret")),
		mcp.WithBoolean("validate_version", mcp.Description("Specify if api2cart should validate cart version")),
		mcp.WithBoolean("verify", mcp.Description("Enables or disables cart's verification")),
		mcp.WithString("db_tables_prefix", mcp.Description("DB tables prefix")),
		mcp.WithString("ftp_host", mcp.Description("FTP connection host")),
		mcp.WithString("ftp_user", mcp.Description("FTP User")),
		mcp.WithString("ftp_password", mcp.Description("FTP Password")),
		mcp.WithNumber("ftp_port", mcp.Description("FTP Port")),
		mcp.WithString("ftp_store_dir", mcp.Description("FTP Store dir")),
		mcp.WithString("apiKey_3dcart", mcp.Description("3DCart API Key")),
		mcp.WithString("AdminAccount", mcp.Description("It's a BigCommerce account for which API is enabled")),
		mcp.WithString("ApiPath", mcp.Description("BigCommerce API URL")),
		mcp.WithString("ApiKey", mcp.Description("Bigcommerce API Key")),
		mcp.WithString("client_id", mcp.Description("Client ID of the requesting app")),
		mcp.WithString("accessToken", mcp.Description("Access token authorizing the app to access resources on behalf of a user")),
		mcp.WithString("context", mcp.Description("API Path section unique to the store")),
		mcp.WithString("access_token", mcp.Description("Access token authorizing the app to access resources on behalf of a user")),
		mcp.WithString("apiKey_shopify", mcp.Description("Shopify API Key")),
		mcp.WithString("apiPassword", mcp.Description("Shopify API Password")),
		mcp.WithString("accessToken_shopify", mcp.Description("Access token authorizing the app to access resources on behalf of a user")),
		mcp.WithString("apiKey", mcp.Description("Neto API Key")),
		mcp.WithString("apiUsername", mcp.Description("Neto User Name")),
		mcp.WithString("EncryptedPassword", mcp.Description("Volusion API Password")),
		mcp.WithString("Login", mcp.Description("It's a Volusion account for which API is enabled")),
		mcp.WithString("apiUser_adnsf", mcp.Description("It's a AspDotNetStorefront account for which API is available")),
		mcp.WithString("apiPass", mcp.Description("AspDotNetStorefront API Password")),
		mcp.WithString("privateKey", mcp.Description("3DCart Application Private Key")),
		mcp.WithString("appToken", mcp.Description("3DCart Token from Application")),
		mcp.WithString("etsy_keystring", mcp.Description("Etsy keystring")),
		mcp.WithString("etsy_shared_secret", mcp.Description("Etsy shared secret")),
		mcp.WithString("tokenSecret", mcp.Description("Secret token authorizing the app to access resources on behalf of a user")),
		mcp.WithString("etsy_client_id", mcp.Required(), mcp.Description("Etsy Client Id")),
		mcp.WithString("etsy_refresh_token", mcp.Required(), mcp.Description("Etsy Refresh token")),
		mcp.WithString("ebay_client_id", mcp.Description("Application ID (AppID).")),
		mcp.WithString("ebay_client_secret", mcp.Description("Shared Secret from eBay application")),
		mcp.WithString("ebay_runame", mcp.Description("The RuName value that eBay assigns to your application.")),
		mcp.WithString("ebay_access_token", mcp.Description("Used to authenticate API requests.")),
		mcp.WithString("ebay_refresh_token", mcp.Description("Used to renew the access token.")),
		mcp.WithString("ebay_environment", mcp.Description("eBay environment")),
		mcp.WithNumber("ebay_site_id", mcp.Description("eBay global ID")),
		mcp.WithString("dw_client_id", mcp.Description("Demandware client id")),
		mcp.WithString("dw_api_pass", mcp.Description("Demandware api password")),
		mcp.WithString("demandware_user_name", mcp.Description("Demandware user name")),
		mcp.WithString("demandware_user_password", mcp.Description("Demandware user password")),
		mcp.WithString("store_id", mcp.Required(), mcp.Description("Store Id")),
		mcp.WithString("seller_id", mcp.Description("Seller Id")),
		mcp.WithString("amazon_secret_key", mcp.Description("Amazon Secret Key")),
		mcp.WithString("amazon_access_key_id", mcp.Description("Amazon Secret Key Id")),
		mcp.WithString("marketplaces_ids", mcp.Description("Comma separated marketplaces ids")),
		mcp.WithString("environment", mcp.Description("")),
		mcp.WithString("hybris_client_id", mcp.Description("Omni Commerce Connector Client ID")),
		mcp.WithString("hybris_client_secret", mcp.Description("Omni Commerce Connector Client Secret")),
		mcp.WithString("hybris_username", mcp.Description("User Name")),
		mcp.WithString("hybris_password", mcp.Description("User password")),
		mcp.WithArray("hybris_websites", mcp.Description("Websites to stores mapping data")),
		mcp.WithString("walmart_client_id", mcp.Description("Walmart client ID")),
		mcp.WithString("walmart_client_secret", mcp.Description("Walmart client secret")),
		mcp.WithString("walmart_environment", mcp.Description("Walmart environment")),
		mcp.WithString("walmart_channel_type", mcp.Description("Walmart WM_CONSUMER.CHANNEL.TYPE header")),
		mcp.WithString("lightspeed_api_key", mcp.Description("LightSpeed api key")),
		mcp.WithString("lightspeed_api_secret", mcp.Description("LightSpeed api secret")),
		mcp.WithString("shopware_access_key", mcp.Description("Shopware access key")),
		mcp.WithString("shopware_api_key", mcp.Description("Shopware api key")),
		mcp.WithString("shopware_api_secret", mcp.Description("Shopware client secret access key")),
		mcp.WithString("commercehq_api_key", mcp.Description("CommerceHQ api key")),
		mcp.WithString("commercehq_api_password", mcp.Description("CommerceHQ api password")),
		mcp.WithString("3dcart_private_key", mcp.Description("3DCart Private Key")),
		mcp.WithString("3dcart_access_token", mcp.Description("3DCart Token")),
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
		Handler:    CartcreateHandler(cfg),
	}
}
