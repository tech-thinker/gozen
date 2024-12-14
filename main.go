package main

import (
	"embed"
	"os"

	"github.com/tech-thinker/gozen/cmd"
	"github.com/tech-thinker/gozen/cmd/helpers"
	"github.com/tech-thinker/gozen/cmd/repository"
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

	shellWrapper := wrappers.NewShellWrapper()
	fileSystemWrapper := wrappers.NewFileSystemWrapper()

	systemRepo := repository.NewSystemRepo(templatesFS, shellWrapper, fileSystemWrapper)
	projectRepo := repository.NewProjectRepo(fileSystemWrapper)

	projectHelper := helpers.NewProjectHelper(systemRepo)

	appSvc := service.NewAppService(systemRepo, projectRepo, projectHelper)
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
