package cmd

import (
	"context"
	"log"

	"github.com/joshgav/az-profiles/go/hybrid/cmd/internal/clients"
	"github.com/joshgav/az-profiles/go/hybrid/cmd/internal/env"

	"github.com/Azure/azure-sdk-for-go/profiles/2017-03-09/resources/mgmt/resources"
	// "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/Azure/go-autorest/autorest/to"
)

func CreateGroup(ctx context.Context, name string, config resources.Group) (*resources.Group, error) {
	c, err := clients.GetGroupsClient()
	if err != nil {
		log.Fatalf("could not get groups client: %s\n", err)
		return nil, err
	}

	group, err := c.CreateOrUpdate(ctx, name, config)
	if err != nil {
		log.Fatalf("could not create group: %s\n", err)
		return nil, err
	}
	return &group, nil
}

var groupConfig = resources.Group{
	Location: to.StringPtr(env.Location()),
}
