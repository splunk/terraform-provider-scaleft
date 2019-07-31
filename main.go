package main

import (
	"github.com/splunk/terraform-provider-scaleft/scaleft"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: scaleft.Provider})
}
