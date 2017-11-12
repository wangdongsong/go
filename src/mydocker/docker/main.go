package main

import (
	"github.com/urfave/cli"
	log "github.com/sirupsen/logrus"
	"mydocker/docker/container"
	"fmt"
	"os"
)

const usage = "mydocker is a simple container runtime implementation. The Purpose" +
	" of this project is to learn how docker works and how to write a docker by ourselves" +
	" Enjoy it, just for fun."

func main()  {

	app := cli.NewApp()
	app.Name = "mydocker"
	app.Usage = usage

	app.Commands = []cli.Command{initCommand, runCommand, }

	app.Before = func(context *cli.Context) error {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(os.Stdout)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

var runCommand = cli.Command{
	Name : "run",
	Usage: "Create a container with namespace and cgroups limit mydocker run -ti[comman]",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name: "ti",
			Usage: "enable tty",
		},
	},

	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container command")
		}

		cmd := context.Args().Get(0)
		tty := context.Bool("ti")
		Run(tty, cmd)
		return nil
	},

}


var initCommand = cli.Command{
	Name: "init",
	Usage: "Init container process run user's process in container. Do not call it outside",

	Action: func(context *cli.Context) error {
		log.Info("Init come on")
		cmd := context.Args().Get(0)
		log.Infof("command %s", cmd)
		err := container.RunContainerInitProcess(cmd, nil)
		return err
	},
}

func Run(tty bool, command string){
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error(err)
	}

	parent.Wait()
	os.Exit(-1)

}
