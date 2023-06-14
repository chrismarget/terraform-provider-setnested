package main

import (
	"context"
	"flag"
	"github.com/chrismarget/terraform-provider-setnested/setnested"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"log"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	err := providerserver.Serve(context.Background(), func() provider.Provider { return new(setnested.Provider) }, providerserver.ServeOpts{
		Address: "registry.terraform.io/chrismarget/setnested",
		Debug:   debug,
	})
	if err != nil {
		log.Fatal(err)
	}
}
