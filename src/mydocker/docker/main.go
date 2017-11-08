package main

import (
	"github.com/urfave/cli"
)

const usage = "mydocker is a simple container runtime implementation. The Purpose" +
	" of this project is to learn how docker works and how to write a docker by ourselves" +
	" Enjoy it, just for fun."

func main()  {

	app := cli.NewApp()
	app.Name = "mydocker"
	app.Usage = usage

	app.Commands = []cli.Command{initCommand, runCommand}
}
