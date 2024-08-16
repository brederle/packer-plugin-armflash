package main

import (
	"fmt"
	"os"
	"github.com/hashicorp/packer-plugin-sdk/plugin"
	"github.com/brederle/packer-plugin-armflash/builder"
	"github.com/brederle/packer-plugin-armflash/version"
)


func main() {
    pluginSet := plugin.NewSet()
	pluginSet.RegisterBuilder(plugin.DEFAULT_NAME, builder.NewBuilder())
    //pluginSet.RegisterPostProcessor("flasher", flasher.newPostProcessor())
	pluginSet.SetVersion(version.PluginVersion)
    err := pluginSet.Run()
    if err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        os.Exit(1)
    }
}
