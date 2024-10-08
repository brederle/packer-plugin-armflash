package main

import (
	"github.com/hashicorp/packer-plugin-sdk/plugin"
	"github.com/brederle/packer-plugin-armflash/builder"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	if err := server.RegisterBuilder(builder.NewBuilder()); err != nil {
		panic(err)
	}
	server.Serve()
}
