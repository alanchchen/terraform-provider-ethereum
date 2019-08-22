package main

import (
	"github.com/hashicorp/terraform/plugin"

	"github.com/alanchchen/terraform-provider-ethereum/ethereum"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ethereum.Provider})
}
