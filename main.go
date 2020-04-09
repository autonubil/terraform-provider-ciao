package main

import (
	"github.com/hashicorp/terraform/plugin"
	"gitlab.autonubil.net/terraform/terraform-provider-ciao/ciao"
)

// Version set during build
var Version string

// Commit hash of build
var Commit string

// BuildDate date of build
var BuildDate string

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ciao.Provider,
	})
}
