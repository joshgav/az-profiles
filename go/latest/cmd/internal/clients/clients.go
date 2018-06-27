package clients

import (
	"github.com/joshgav/az-profiles/go/latest/cmd/internal/env"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-04-01/network"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"

	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func GetVNetClient() (network.VirtualNetworksClient, error) {
	c := network.NewVirtualNetworksClient(env.SubscriptionID())
	// c := network.NewVirtualNetworksClient()
	authConfig := auth.ClientCredentialsConfig{
		ClientID:     env.ClientID(),
		ClientSecret: env.ClientSecret(),
		TenantID:     env.TenantID(),
		AADEndpoint:  env.AuthorizationServerURL(),
		Resource:     env.ResourceURL(),
	}
	a, err := authConfig.Authorizer()
	// a, err := auth.NewServicePrincipalAuthorizerFromEnvironment()
	if err != nil {
		c.Authorizer = nil
	} else {
		c.Authorizer = a
	}
	return c, err
}

func GetGroupsClient() (resources.GroupsClient, error) {
	c := resources.NewGroupsClient(env.SubscriptionID())
	authConfig := auth.ClientCredentialsConfig{
		ClientID:     env.ClientID(),
		ClientSecret: env.ClientSecret(),
		TenantID:     env.TenantID(),
		AADEndpoint:  env.AuthorizationServerURL(),
		Resource:     env.ResourceURL(),
	}
	a, err := authConfig.Authorizer()
	// a, err := auth.NewServicePrincipalAuthorizerFromEnvironment()
	if err != nil {
		c.Authorizer = nil
	} else {
		c.Authorizer = a
	}
	return c, err
}
