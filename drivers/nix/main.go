package main

import (
	"github.com/hernad/nomad/drivers/nix/nix"

	"github.com/hashicorp/go-hclog"
	"github.com/hernad/nomad/plugins"
)

func main() {
	// Serve the plugin
	plugins.Serve(factory)
}

// factory returns a new instance of a nomad driver plugin
func factory(log hclog.Logger) interface{} {
	return nix.NewPlugin(log)
}