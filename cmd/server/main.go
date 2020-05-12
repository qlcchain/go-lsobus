package main

import (
	"os"

	"github.com/iixlabs/virtual-lsobus/cmd/server/commands"
)

func main() {
	_ = os.Setenv("SWAGGER_DEBUG", "true")
	_ = os.Setenv("DEBUG", "true")

	commands.Execute()
}
