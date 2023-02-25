// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"terraform-provider-lynx/internal/provider"
	"terraform-provider-lynx/sdk"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate -provider-name scaffolding

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary.
	version string = "dev"

	// goreleaser can pass other information to the main package, such as the specific commit
	// https://goreleaser.com/cookbooks/using-main.version/
)

func main() {
	var debug bool

	client := sdk.NewClient(sdk.LocalApiServer, "bd11a454-a694-49c8-b3da-0fe6cf48a27d")

	usr, err := client.CreateUser(sdk.User{
		Name:     "Selena",
		Email:    "selena@clivern.com",
		Role:     sdk.RegularUser,
		Password: "$123456789$",
	})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%v", usr)
	}

	client.DeleteUser("84c1e6ec-7e6a-4ab8-967e-666ebcc27198")

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		// TODO: Update this string with the published name of your provider.
		// Also update the tfplugindocs generate command to either remove the
		// -provider-name flag or set its value to the updated provider name.
		Address: "registry.terraform.io/clivern/lynx",
		Debug:   debug,
	}

	err = providerserver.Serve(context.Background(), provider.New(version), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
