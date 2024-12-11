package main

import (
	"embed"
	"os"

	"github.com/tech-thinker/gozen/cmd"
	"github.com/tech-thinker/gozen/cmd/helper"
	"github.com/tech-thinker/gozen/cmd/service"
	"github.com/tech-thinker/gozen/wrappers"
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

	shellRepo := wrappers.NewShellWrapper()
	fileSystemRepo := wrappers.NewFileSystemWrapper()

	commonHelper := helper.NewCommonHelper(templatesFS, shellRepo, fileSystemRepo)
	appSvc := service.NewAppService(fileSystemRepo, commonHelper)
	app := cmd.NewApp(appSvc)

	cliApp := cli.NewApp()
	cliApp.Name = "gozen"
	cliApp.Version = AppVersion
	cliApp.Commands = []*cli.Command{
		app.CreateProject(),
	}
	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}
}
