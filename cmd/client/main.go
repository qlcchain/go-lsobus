package main

import (
	"os"

	cmd "github.com/qlcchain/go-lsobus/cmd/client/commands"
)

func main() {
	cmd.Execute(os.Args)
}
