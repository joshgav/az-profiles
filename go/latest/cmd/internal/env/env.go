package env

import (
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/gobuffalo/envy"
)

var clientID, clientSecret, tenantID, subscriptionID, location, resourceURL, groupName string
var authorizationServerURL string

func init() {
	envy.Load()                                                  // load project-local .env
	azureEnv, _ := azure.EnvironmentFromName("AzurePublicCloud") // shouldn't fail

	clientID, _ = envy.MustGet("AZURE_CLIENT_ID")
	clientSecret, _ = envy.MustGet("AZURE_CLIENT_SECRET")
	tenantID, _ = envy.MustGet("AZURE_TENANT_ID")
	subscriptionID, _ = envy.MustGet("AZURE_SUBSCRIPTION_ID")
	// assert: all these vars are set

	resourceURL = envy.Get("AZURE_RESOURCE_URL", azureEnv.ResourceManagerEndpoint)
	groupName = envy.Get("AZURE_GROUP_NAME", "default-group-name-01")
	location = envy.Get("AZURE_LOCATION_DEFAULT", "westus2")

	authorizationServerURL = azureEnv.ActiveDirectoryEndpoint
}

func ClientID() string {
	return clientID
}

func ClientSecret() string {
	return clientSecret
}

func TenantID() string {
	return tenantID
}

func SubscriptionID() string {
	return subscriptionID
}

func ResourceURL() string {
	return resourceURL
}

func GroupName() string {
	return groupName
}

func Location() string {
	return location
}

func AuthorizationServerURL() string {
	return authorizationServerURL
}
