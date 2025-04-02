package cmd

import (
	"github.com/tech-thinker/gozen/cmd/service"
	"github.com/urfave/cli/v2"
)

type App interface {
	CreateProject() *cli.Command
	GenerateModel() *cli.Command
}

type app struct {
	appService service.AppService
}

func NewApp(
	appService service.AppService,
) App {
	return &app{
		appService: appService,
	}
}
