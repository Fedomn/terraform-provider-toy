package main

import (
	"github.com/fedomn/terraform-provider-toy/toy"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{ProviderFunc: toy.Provider})
}
