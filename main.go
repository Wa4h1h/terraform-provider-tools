// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"flag"
	"log"

	"github.com/Wa4h1h/terraform-provider-tools/internal/provider"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// these will be set by the goreleaser configuration
// to appropriate values for the compiled binary.
var version string = "dev"

// goreleaser can pass other information to the main package, such as the specific commit
// https://goreleaser.com/cookbooks/using-main.version/

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug",
		true, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		// Also update the tfplugindocs generate command to either remove the
		// -provider-name flag or set its value to the updated provider name.
		Address: "registry.terraform.io/Wa4h1h/tools",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(),
		provider.NewToolsProvider(version), opts)
	if err != nil {
		log.Fatal(err.Error())
	}
}
