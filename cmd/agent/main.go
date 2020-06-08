package main

import (
	"os"

	"github.com/qlcchain/go-lsobus/cmd/agent/commands"
)

func main() {
	_ = os.Setenv("SWAGGER_DEBUG", "true")
	_ = os.Setenv("DEBUG", "true")

	commands.InitCmd()

	commands.Execute(os.Args)
}
