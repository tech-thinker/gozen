package repository

import (
	"fmt"

	"github.com/tech-thinker/gozen/models"
	"github.com/tech-thinker/gozen/wrappers"
)

type ProjectRepo interface {
	Create(project models.Project) error
}

type projectRepo struct {
	fileSystemWrapper wrappers.FileSystemWrapper
}

func (p *projectRepo) Create(project models.Project) error {
	appPath := project.GetAppDir()
	fmt.Println("Project created at location: " + appPath)
	err := p.fileSystemWrapper.CreateDirectory(appPath)
	if err != nil {
		return err
	}
	// Write project config into file
	if err := project.WriteToJsonFile(); err != nil {
		return err
	}
	return nil
}

func NewProjectRepo(
	fileSystemWrapper wrappers.FileSystemWrapper,
) ProjectRepo {
	return &projectRepo{
		fileSystemWrapper: fileSystemWrapper,
	}
}
