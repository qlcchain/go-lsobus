package main

import (
	"fmt"
	"os"

	client "github.com/qlcchain/go-virtual-lsobus/cmd/client/commands"
	server "github.com/qlcchain/go-virtual-lsobus/cmd/server/commands"
	"github.com/qlcchain/go-virtual-lsobus/services/version"
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
