package main

import (
	"os"

	"github.com/tech-thinker/gozen/cmd"
	"github.com/tech-thinker/gozen/models"
	"github.com/urfave/cli/v2"
)

func main() {
	var packageName string
	var outputDir string

	clientApp := cli.NewApp()
	clientApp.Name = "gozen"
	clientApp.Commands = []*cli.Command{
		{
			Name:  "create",
			Usage: "Create new Projects.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "pkg",
					Aliases:     []string{"p"},
					Value:       "",
					Usage:       "Package name for new project.",
					Destination: &packageName,
				},
				&cli.StringFlag{
					Name:        "dir",
					Aliases:     []string{"d"},
					Value:       ".",
					Usage:       "Output directory for new project.",
					Destination: &outputDir,
				},
			},
			Action: func(ctx *cli.Context) error {
                project := models.Project{
                    AppName: ctx.Args().Get(0),
                    PackageName:     packageName,
                }
                app := cmd.NewAppCmd(project)

				return app.CreateApp()
			},
		},
	}
	if err := clientApp.Run(os.Args); err != nil {
		panic(err)
	}
}
