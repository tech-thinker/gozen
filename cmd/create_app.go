package cmd

import (
	"fmt"

	"github.com/tech-thinker/gozen/cmd/helper"
	"github.com/tech-thinker/gozen/constants"
	"github.com/tech-thinker/gozen/models"
	"github.com/tech-thinker/gozen/utils"
)

type AppCmd interface {
	CreateApp() error
}

type appCmd struct {
	project models.Project
	helper  helper.CommonHelper
}

func (cmd *appCmd) CreateApp() error {
	workDir := constants.CURRENT_WORKING_DIRECTORY
	if len(cmd.project.WorkingDir) > 1 {
		workDir = cmd.project.WorkingDir
	}
	appPath := fmt.Sprintf(`%s/%s`, workDir, cmd.project.AppName)
	fmt.Println("Project created at location: " + appPath)
	utils.CreateDirectory(appPath)
	// Write project config into file
	if err := cmd.project.WriteToJsonFile(); err != nil {
		return err
	}

	// Generating basic codes
	cmd.helper.Write("templates/env.sample.tmpl", appPath+"/.env", cmd.project)
	cmd.helper.Write("templates/env.sample.tmpl", appPath+"/.env.sample", cmd.project)
	cmd.helper.Write("templates/env.sample.tmpl", appPath+"/docker/.env", cmd.project)
	cmd.helper.Write("templates/env.sample.tmpl", appPath+"/docker/.env.sample", cmd.project)
	cmd.helper.Write("templates/docker/Dockerfile.debug", appPath+"/docker/Dockerfile.debug", cmd.project)
	cmd.helper.Write("templates/docker/Dockerfile.dev", appPath+"/docker/Dockerfile.dev", cmd.project)
	cmd.helper.Write("templates/docker/Dockerfile.prod", appPath+"/docker/Dockerfile.prod", cmd.project)
	cmd.helper.Write("templates/docker/docker-compose-debug.yml", appPath+"/docker/docker-compose-debug.yml", cmd.project)
	cmd.helper.Write("templates/docker/docker-compose.yml", appPath+"/docker/docker-compose.yml", cmd.project)
	cmd.helper.Write("templates/docker/modd-debug.conf", appPath+"/docker/modd-debug.conf", cmd.project)
	cmd.helper.Write("templates/docker/modd-dev.conf", appPath+"/docker/modd-dev.conf", cmd.project)

	cmd.helper.Write("templates/gitignore.tmpl", appPath+"/.gitignore", cmd.project)
	cmd.helper.Write("templates/Makefile.tmpl", appPath+"/Makefile", cmd.project)
	cmd.helper.Write("templates/go.tmpl", appPath+"/go.mod", cmd.project)
	cmd.helper.Write("templates/main.tmpl", appPath+"/main.go", cmd.project)

	cmd.helper.Write("templates/app/init.tmpl", appPath+"/app/init.go", cmd.project)

	cmd.helper.Write("templates/app/rest/controllers/health.tmpl", appPath+"/app/rest/controllers/health.go", cmd.project)
	cmd.helper.Write("templates/app/rest/router/router.tmpl", appPath+"/app/rest/router/router.go", cmd.project)

	cmd.helper.Write("templates/app/grpc/handlers/health.tmpl", appPath+"/app/grpc/handlers/health.go", cmd.project)
	cmd.helper.Write("templates/app/grpc/proto/health.proto", appPath+"/app/grpc/proto/health.proto", cmd.project)
	cmd.helper.Write("templates/app/grpc/proto/health.pb.tmpl", appPath+"/app/grpc/proto/health.pb.go", cmd.project)
	cmd.helper.Write("templates/app/grpc/proto/health_grpc.pb.tmpl", appPath+"/app/grpc/proto/health_hrpc.pb.go", cmd.project)
	cmd.helper.Write("templates/app/grpc/router/router.tmpl", appPath+"/app/grpc/router/router.go", cmd.project)

	cmd.helper.Write("templates/config/config.tmpl", appPath+"/config/config.go", cmd.project)
	cmd.helper.Write("templates/constants/app.tmpl", appPath+"/constants/app.go", cmd.project)
	cmd.helper.Write("templates/instance/instance.tmpl", appPath+"/instance/instance.go", cmd.project)
	cmd.helper.Write("templates/logger/logger.tmpl", appPath+"/logger/logger.go", cmd.project)
	cmd.helper.Write("templates/models/model_registry.tmpl", appPath+"/models/model_registry.go", cmd.project)
	cmd.helper.Write("templates/models/health.tmpl", appPath+"/models/health.go", cmd.project)
	cmd.helper.Write("templates/instance/registry/models.tmpl", appPath+"/instance/registry/models.go", cmd.project)
	cmd.helper.Write("templates/repository/health.tmpl", appPath+"/repository/health.go", cmd.project)

	cmd.helper.Write("templates/runner/api.tmpl", appPath+"/runner/api.go", cmd.project)
	cmd.helper.Write("templates/runner/grpc.tmpl", appPath+"/runner/grpc.go", cmd.project)

	cmd.helper.Write("templates/service/health.tmpl", appPath+"/service/health.go", cmd.project)
	cmd.helper.Write("templates/utils/utils.tmpl", appPath+"/utils/utils.go", cmd.project)

	return nil
}

func NewAppCmd(
	project models.Project,
	helper helper.CommonHelper,
) AppCmd {
	return &appCmd{
		project: project,
		helper:  helper,
	}
}
