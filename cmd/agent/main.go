package main

import (
	"github.com/qlcchain/go-lsobus/cmd/agent/commands"
	"os"
)

func main() {
	_ = os.Setenv("SWAGGER_DEBUG", "true")
	_ = os.Setenv("DEBUG", "true")

	commands.InitCmd()

	commands.Execute(os.Args)
}
