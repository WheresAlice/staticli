package cmd

import (
	"github.com/skybet/cali"
)

var (
	// Define this here, then all other files in cmd can add subcommands to it
	cli = cali.NewCli("staticli")
)

func init() {
	cli.SetShort("Static site generator tooling")
	cli.SetLong("Static site generator tooling")
}

func Execute() {
	cli.Start()
}