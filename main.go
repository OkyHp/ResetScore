package main

import (
	"resetscore/internal"

	"github.com/untrustedmodders/go-plugify"
)

const pluginName = "ResetScore"

func init() {
	plugin := internal.NewResetScorePlugin()
	plugin.Plugin = plugify.NewPlugin(pluginName, plugin.OnPluginStart, plugin.OnPluginUpdate, plugin.OnPluginEnd)
}

func main() {}
