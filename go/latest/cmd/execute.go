package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/joshgav/az-profiles/go/latest/cmd/internal/env"
)

var (
	vnetName  = "azure-auth-vpc"
	groupName = env.GroupName()
)

func Execute() {
	ctx := context.Background()
	// defer ctx.Cancel()

	g, err := CreateGroup(ctx, groupName, groupConfig)
	if err != nil {
		log.Fatalf("could not create group: %s\n", err)
	}
	fmt.Printf("created group: %+v\n", g)

	response, err := CreateVNet(ctx, vnetName, vnetConfig)
	if err != nil {
		log.Fatalf("could not create vnet: %s\n", err)
	}
	fmt.Printf("created vnet: %+v\n", response)
}
