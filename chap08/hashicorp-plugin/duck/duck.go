package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/omihirofumi/cloud-native-go-book/chap08/hashicorp-plugin/commons"
)

type Duck struct{}

func (g *Duck) Says() string {
	return "クワッ!"
}

func main() {
	sayer := &Duck{}

	var pluginMap = map[string]plugin.Plugin{
		"sayer": &commons.SayerPlugin{Impl: sayer},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: commons.HandshakeConfig,
		Plugins:         pluginMap,
	})
}
