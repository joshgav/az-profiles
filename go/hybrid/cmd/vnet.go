package cmd

import (
	"context"
	"log"
	"net/http"

	"github.com/joshgav/az-profiles/go/hybrid/cmd/internal/clients"
	"github.com/joshgav/az-profiles/go/hybrid/cmd/internal/env"

	"github.com/Azure/azure-sdk-for-go/profiles/2017-03-09/network/mgmt/network"
	// "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-04-01/network"
	"github.com/Azure/go-autorest/autorest/to"
)

func CreateVNet(ctx context.Context, name string, config network.VirtualNetwork) (*http.Response, error) {
	c, err := clients.GetVNetClient()
	if err != nil {
		log.Fatalf("could not get vnet client: %s\n", err)
		return nil, err
	}

	future, err := c.CreateOrUpdate(ctx, env.GroupName(), name, config)
	if err != nil {
		log.Fatalf("failed to start creation of vnet: %s\n", err)
	}

	if err = future.WaitForCompletion(ctx, c.Client); err != nil {
		log.Fatalf("error on waiting for completion: %s\n", err)
		return nil, err
	}

	response := future.Response()
	// if vnet, err := response.(network.VirtualNetwork); err != nil {
	//	log.Fatalf("could not parse response: %s\n", err)
	//	return nil, err
	// }
	// return vnet, nil
	return response, nil
}

var vnetConfig = network.VirtualNetwork{
	Location: to.StringPtr(env.Location()),
	VirtualNetworkPropertiesFormat: &network.VirtualNetworkPropertiesFormat{
		AddressSpace: &network.AddressSpace{
			AddressPrefixes: &[]string{"10.0.0.0/8"},
		},
		Subnets: &[]network.Subnet{
			{
				Name: to.StringPtr("subnet0"),
				SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
					AddressPrefix: to.StringPtr("10.0.0.0/16"),
					// NetworkSecurityGroup: &nsgConfig,
				},
			},
			{
				Name: to.StringPtr("subnet1"),
				SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
					AddressPrefix: to.StringPtr("10.1.0.0/16"),
					// NetworkSecurityGroup: &nsgConfig,
				},
			},
		},
	},
}

var nsgConfig = network.SecurityGroup{
	Location: to.StringPtr(env.Location()),
	SecurityGroupPropertiesFormat: &network.SecurityGroupPropertiesFormat{
		SecurityRules: &[]network.SecurityRule{
			{
				Name: to.StringPtr("allow_ssh"),
				SecurityRulePropertiesFormat: &network.SecurityRulePropertiesFormat{
					//Protocol:              network.SecurityRuleProtocolTCP,
					//Access:                network.SecurityRuleAccessAllow,
					//Direction:             network.SecurityRuleDirectionInbound,
					Protocol:                 network.TCP,
					Access:                   network.Allow,
					Direction:                network.Inbound,
					SourceAddressPrefix:      to.StringPtr("0.0.0.0/0"),
					SourcePortRange:          to.StringPtr("1-65535"),
					DestinationAddressPrefix: to.StringPtr("0.0.0.0/0"),
					DestinationPortRange:     to.StringPtr("22"),
					Priority:                 to.Int32Ptr(100),
				},
			},
			{
				Name: to.StringPtr("allow_https"),
				SecurityRulePropertiesFormat: &network.SecurityRulePropertiesFormat{
					//Protocol:                 network.SecurityRuleProtocolTCP,
					//Access:                   network.SecurityRuleAccessAllow,
					//Direction:                network.SecurityRuleDirectionInbound,
					Protocol:                 network.TCP,
					Access:                   network.Allow,
					Direction:                network.Inbound,
					SourceAddressPrefix:      to.StringPtr("0.0.0.0/0"),
					SourcePortRange:          to.StringPtr("1-65535"),
					DestinationAddressPrefix: to.StringPtr("0.0.0.0/0"),
					DestinationPortRange:     to.StringPtr("443"),
					Priority:                 to.Int32Ptr(200),
				},
			},
		},
	},
}
