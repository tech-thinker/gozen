package service

import (
	"github.com/tech-thinker/gozen/cmd/helpers"
	"github.com/tech-thinker/gozen/cmd/repository"
	"github.com/tech-thinker/gozen/models"
)

type AppService interface {
	CreateApp(project models.Project) error
}

type appService struct {
	systemRepo    repository.SystemRepo
	projectRepo   repository.ProjectRepo
	projectHelper helpers.ProjectHelper
}

func (cmd *appService) CreateApp(project models.Project) error {
	// Create project
	err := cmd.projectRepo.Create(project)
	if err != nil {
		return err
	}

	// Generating basic codes
	err = cmd.projectHelper.InitProject(project)
	if err != nil {
		return err
	}

	err = cmd.projectHelper.SetupEnv(project)
	if err != nil {
		return err
	}

	err = cmd.projectHelper.SetupDocker(project)
	if err != nil {
		return err
	}

	// Config
	err = cmd.projectHelper.SetupConfig(project)
	if err != nil {
		return err
	}
	// Constants
	err = cmd.projectHelper.SetupConstants(project)
	if err != nil {
		return err
	}
	// Logger
	err = cmd.projectHelper.SetupLogger(project)
	if err != nil {
		return err
	}
	// Models
	err = cmd.projectHelper.SetupModel(project)
	if err != nil {
		return err
	}
	// Repository
	err = cmd.projectHelper.SetupRepository(project)
	if err != nil {
		return err
	}

	err = cmd.projectHelper.SetupService(project)
	if err != nil {
		return err
	}
	// Utils
	err = cmd.projectHelper.SetupUtils(project)

	// Rest
	err = cmd.projectHelper.CreateRestAPI(project)
	if err != nil {
		return err
	}
	// gRPC
	err = cmd.projectHelper.CreategRPCAPI(project)
	if err != nil {
		return err
	}

	return nil
}

func NewAppService(
	systemRepo repository.SystemRepo,
	projectRepo repository.ProjectRepo,
	projectHelper helpers.ProjectHelper,
) AppService {
	return &appService{
		systemRepo:    systemRepo,
		projectRepo:   projectRepo,
		projectHelper: projectHelper,
	}
}
