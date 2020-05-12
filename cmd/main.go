package main

import (
	"fmt"
	"os"

	client "github.com/iixlabs/virtual-lsobus/cmd/client/commands"
	server "github.com/iixlabs/virtual-lsobus/cmd/server/commands"
	"github.com/iixlabs/virtual-lsobus/services/version"
)

func main() {
	fmt.Println(version.ShortVersion())
	args := os.Args
	if len(args) > 1 && (args[1] == "-i" || args[1] == "--endpoint") {
		client.Execute(os.Args)
	} else {
		server.Execute()
	}
}
