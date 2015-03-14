package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/foostan/ants/command"
)

func main() {
	app := cli.NewApp()
	app.Name = "ants"
	app.Version = Version
	app.Usage = ""
	app.Author = "foostan"
	app.Email = "ks@fstn.jp"
	app.Commands = Commands

	app.Run(os.Args)
}

var Commands = []cli.Command{
	cli.Command{
		Name:        "run",
		Usage:       "",
		Description: "",
		Flags:       command.RunFlags,
		Action:      command.RunCommand,
	},
}
