package main

import (
	"context"
	"os"
	"sync"

    "github.com/urfave/cli"
	"github.com/spf13/viper"
	"{{.PackageName}}/config"
	"{{.PackageName}}/instance"
	"{{.PackageName}}/runner"
)

func main() {
	v := viper.New()
	cfg := config.Init(v)

	instance := instance.Init(cfg)

	clientApp := cli.NewApp()
	clientApp.Name = "{{.AppName}}"
	clientApp.Version = cfg.Version()
	clientApp.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "Start the service",
			Action: func(c *cli.Context) error {
				
				return nil
			},
		},
	}
	if err := clientApp.Run(os.Args); err != nil {
		panic(err)
	}
}
