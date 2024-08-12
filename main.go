package main

import (
	"github.com/hashicorp/packer-plugin-sdk/plugin"
	"github.com/mkaczanowski/packer-builder-arm/builder"
)

func main() {
    pluginSet := plugin.NewSet()
	pluginSet.RegisterBuilder(plugin.DEFAULT_NAME, builder.NewBuilder())
    //pluginSet.RegisterPostProcessor("flasher", flasher.newPostProcessor())
    err := pluginSet.Run()
    if err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        os.Exit(1)
    }
}
