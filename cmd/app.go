package cmd

import (
	"fmt"

	"github.com/tech-thinker/gozen/cmd/service"
	"github.com/tech-thinker/gozen/models"
	"github.com/urfave/cli/v2"
)

type App interface {
	CreateProject() *cli.Command
}

type app struct {
	appService service.AppService
}

func (c *app) CreateProject() *cli.Command {

	var packageName, outputDir, driver string

	return &cli.Command{
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

			return c.appService.CreateApp(project)
		},
	}
}

func NewApp(
	appService service.AppService,
) App {
	return &app{
		appService: appService,
	}
}
