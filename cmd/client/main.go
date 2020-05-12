package main

import (
	"os"

	cmd "github.com/iixlabs/virtual-lsobus/cmd/client/commands"
)

func main() {
	cmd.Execute(os.Args)
}
