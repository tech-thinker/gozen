package main

import (
	"embed"
	"testing"

	"github.com/tech-thinker/gozen/cmd"
	"github.com/tech-thinker/gozen/cmd/helpers"
	"github.com/tech-thinker/gozen/cmd/repository"
	"github.com/tech-thinker/gozen/cmd/service"
	"github.com/tech-thinker/gozen/wrappers"
)

//go:embed templates/*
var testTemplatesFS embed.FS

func TestMainDISetup(t *testing.T) {
	shellWrapper := wrappers.NewShellWrapper()
	fileSystemWrapper := wrappers.NewFileSystemWrapper()

	systemRepo := repository.NewSystemRepo(testTemplatesFS, shellWrapper, fileSystemWrapper)
	projectRepo := repository.NewProjectRepo(fileSystemWrapper)

	projectHelper := helpers.NewProjectHelper(systemRepo)
	codeHelper := helpers.NewCodeHelper(systemRepo)

	appSvc := service.NewAppService(projectRepo, projectHelper, codeHelper)
	app := cmd.NewApp(appSvc)

	if app == nil {
		t.Error("cmd.NewApp() returned nil")
	}

	if appSvc == nil {
		t.Error("service.NewAppService() returned nil")
	}

	if projectHelper == nil {
		t.Error("helpers.NewProjectHelper() returned nil")
	}

	if codeHelper == nil {
		t.Error("helpers.NewCodeHelper() returned nil")
	}

	if projectRepo == nil {
		t.Error("repository.NewProjectRepo() returned nil")
	}

	if systemRepo == nil {
		t.Error("repository.NewSystemRepo() returned nil")
	}

	if shellWrapper == nil {
		t.Error("wrappers.NewShellWrapper() returned nil")
	}

	if fileSystemWrapper == nil {
		t.Error("wrappers.NewFileSystemWrapper() returned nil")
	}
}

func TestAppVersion(t *testing.T) {
	if AppVersion == "" {
		t.Error("AppVersion should not be empty")
	}
}

func TestCommitHash(t *testing.T) {
	if CommitHash == "" {
		t.Error("CommitHash should not be empty")
	}
}

func TestBuildDate(t *testing.T) {
	if BuildDate == "" {
		t.Error("BuildDate should not be empty")
	}
}