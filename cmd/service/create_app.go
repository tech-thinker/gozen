package service

import (
	"fmt"

	"github.com/tech-thinker/gozen/cmd/helper"
	"github.com/tech-thinker/gozen/constants"
	"github.com/tech-thinker/gozen/models"
	"github.com/tech-thinker/gozen/wrappers"
)

type AppService interface {
	CreateApp(project models.Project) error
}

type appService struct {
	fileSystemRepo wrappers.FileSystemWrapper
	helper         helper.CommonHelper
}

func (cmd *appService) CreateApp(project models.Project) error {
	workDir := constants.CURRENT_WORKING_DIRECTORY
	if len(project.WorkingDir) > 1 {
		workDir = project.WorkingDir
	}
	appPath := fmt.Sprintf(`%s/%s`, workDir, project.AppName)
	fmt.Println("Project created at location: " + appPath)
	cmd.fileSystemRepo.CreateDirectory(appPath)
	// Write project config into file
	if err := project.WriteToJsonFile(); err != nil {
		return err
	}

	// Generating basic codes
	cmd.helper.Write("templates/env.sample.tmpl", appPath+"/.env", project)
	cmd.helper.Write("templates/env.sample.tmpl", appPath+"/.env.sample", project)
	cmd.helper.Write("templates/env.sample.tmpl", appPath+"/docker/.env", project)
	cmd.helper.Write("templates/env.sample.tmpl", appPath+"/docker/.env.sample", project)
	cmd.helper.Write("templates/docker/Dockerfile.debug", appPath+"/docker/Dockerfile.debug", project)
	cmd.helper.Write("templates/docker/Dockerfile.dev", appPath+"/docker/Dockerfile.dev", project)
	cmd.helper.Write("templates/docker/Dockerfile.prod", appPath+"/docker/Dockerfile.prod", project)
	cmd.helper.Write("templates/docker/docker-compose-debug.yml", appPath+"/docker/docker-compose-debug.yml", project)
	cmd.helper.Write("templates/docker/docker-compose.yml", appPath+"/docker/docker-compose.yml", project)
	cmd.helper.Write("templates/docker/modd-debug.conf", appPath+"/docker/modd-debug.conf", project)
	cmd.helper.Write("templates/docker/modd-dev.conf", appPath+"/docker/modd-dev.conf", project)

	cmd.helper.Write("templates/gitignore.tmpl", appPath+"/.gitignore", project)
	cmd.helper.Write("templates/Makefile.tmpl", appPath+"/Makefile", project)
	cmd.helper.Write("templates/go.tmpl", appPath+"/go.mod", project)
	cmd.helper.Write("templates/main.tmpl", appPath+"/main.go", project)

	cmd.helper.Write("templates/app/init.tmpl", appPath+"/app/init.go", project)

	cmd.helper.Write("templates/app/rest/controllers/health.tmpl", appPath+"/app/rest/controllers/health.go", project)
	cmd.helper.Write("templates/app/rest/router/router.tmpl", appPath+"/app/rest/router/router.go", project)

	cmd.helper.Write("templates/app/grpc/handlers/health.tmpl", appPath+"/app/grpc/handlers/health.go", project)
	cmd.helper.Write("templates/app/grpc/proto/health.proto", appPath+"/app/grpc/proto/health.proto", project)
	cmd.helper.Write("templates/app/grpc/proto/health.pb.tmpl", appPath+"/app/grpc/proto/health.pb.go", project)
	cmd.helper.Write("templates/app/grpc/proto/health_grpc.pb.tmpl", appPath+"/app/grpc/proto/health_hrpc.pb.go", project)
	cmd.helper.Write("templates/app/grpc/router/router.tmpl", appPath+"/app/grpc/router/router.go", project)

	cmd.helper.Write("templates/config/config.tmpl", appPath+"/config/config.go", project)
	cmd.helper.Write("templates/constants/app.tmpl", appPath+"/constants/app.go", project)
	cmd.helper.Write("templates/instance/instance.tmpl", appPath+"/instance/instance.go", project)
	cmd.helper.Write("templates/logger/logger.tmpl", appPath+"/logger/logger.go", project)
	cmd.helper.Write("templates/models/model_registry.tmpl", appPath+"/models/model_registry.go", project)
	cmd.helper.Write("templates/models/health.tmpl", appPath+"/models/health.go", project)
	cmd.helper.Write("templates/instance/registry/models.tmpl", appPath+"/instance/registry/models.go", project)
	cmd.helper.Write("templates/repository/health.tmpl", appPath+"/repository/health.go", project)

	cmd.helper.Write("templates/runner/api.tmpl", appPath+"/runner/api.go", project)
	cmd.helper.Write("templates/runner/grpc.tmpl", appPath+"/runner/grpc.go", project)

	cmd.helper.Write("templates/service/health.tmpl", appPath+"/service/health.go", project)
	cmd.helper.Write("templates/utils/utils.tmpl", appPath+"/utils/utils.go", project)

	return nil
}

func NewAppService(
	fileSystemRepo wrappers.FileSystemWrapper,
	helper helper.CommonHelper,
) AppService {
	return &appService{
		fileSystemRepo: fileSystemRepo,
		helper:         helper,
	}
}
