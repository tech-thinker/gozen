package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/tech-thinker/gozen/cmd"
	"github.com/tech-thinker/gozen/cmd/helper"
	"github.com/tech-thinker/gozen/models"
	"github.com/urfave/cli/v2"
)

var (
	AppVersion = "v0.0.0"
	CommitHash = "unknown"
	BuildDate  = "unknown"
)

//go:embed templates/*
var templatesFS embed.FS

func main() {
	// Declaring flags
	var packageName string
	var outputDir string
	var driver string

	// Initialize common helper
	helper := helper.NewCommonHelper(templatesFS)

	clientApp := cli.NewApp()
	clientApp.Name = "gozen"
	clientApp.Version = AppVersion
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
					Name:        "output",
					Aliases:     []string{"o"},
					Value:       ".",
					Usage:       "Output directory for new project.",
					Destination: &outputDir,
				},
				&cli.StringFlag{
					Name:        "driver",
					Aliases:     []string{"d"},
					Value:       "sqlite",
					Usage:       "Database driver for new project. eg. [sqlite, mysql, postgres]",
					Destination: &driver,
				},
			},
			Action: func(ctx *cli.Context) error {
				project := models.Project{
					AppName:     ctx.Args().Get(0),
					PackageName: packageName,
					Driver:      driver,
					WorkingDir:  outputDir,
				}

				err := project.Validate()
				if err != nil {
					fmt.Println(err)
					return nil
				}

				project.AutoFixes()

				app := cmd.NewAppCmd(project, helper)

				return app.CreateApp()
			},
		},
	}
	if err := clientApp.Run(os.Args); err != nil {
		panic(err)
	}
}
