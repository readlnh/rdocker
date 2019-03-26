package main

import (
	"github.com/urfave/cli"
	log "github.com/sirupsen/logrus"
	"os"
)


const usage = `readlnh mydocker`

func main()  {
	app := cli.NewApp()
	app.Name = "rdocker"
	app.Usage = usage

	app.Commands = []cli.Command {
		initCommand,
		runCommand,
	}

	app.Before = func(context *cli.Context) error {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(os.Stdout)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
