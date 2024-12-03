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
	appPath := fmt.Sprintf(`%s/%s`, constants.CURRENT_WORKING_DIRECTORY, cmd.project.AppName)
	fmt.Println("Project created at location: " + appPath)
	utils.CreateDirectory(appPath)
	// Write project config into file
	if err := cmd.project.WriteToJsonFile(); err != nil {
		return err
	}

	// Generating basic codes
	cmd.helper.Write("templates/env.sample.tpl", appPath+"/.env", cmd.project)
	cmd.helper.Write("templates/env.sample.tpl", appPath+"/.env.sample", cmd.project)
	cmd.helper.Write("templates/env.sample.tpl", appPath+"/docker/.env", cmd.project)
	cmd.helper.Write("templates/env.sample.tpl", appPath+"/docker/.env.sample", cmd.project)
    cmd.helper.Write("templates/docker/Dockerfile.debug", appPath+"/docker/Dockerfile.debug", cmd.project)
    cmd.helper.Write("templates/docker/Dockerfile.dev", appPath+"/docker/Dockerfile.dev", cmd.project)
    cmd.helper.Write("templates/docker/Dockerfile.prod", appPath+"/docker/Dockerfile.prod", cmd.project)
    cmd.helper.Write("templates/docker/docker-compose-debug.yml", appPath+"/docker/docker-compose-debug.yml", cmd.project)
    cmd.helper.Write("templates/docker/docker-compose.yml", appPath+"/docker/docker-compose.yml", cmd.project)
    cmd.helper.Write("templates/docker/entrypoint-debug.sh", appPath+"/docker/entrypoint-debug.sh", cmd.project)
    cmd.helper.Write("templates/docker/entrypoint-dev.sh", appPath+"/docker/entrypoint-dev.sh", cmd.project)
    cmd.helper.Write("templates/docker/modd-debug.conf", appPath+"/docker/modd-debug.conf", cmd.project)
    cmd.helper.Write("templates/docker/modd-dev.conf", appPath+"/docker/modd-dev.conf", cmd.project)

	cmd.helper.Write("templates/gitignore.tpl", appPath+"/.gitignore", cmd.project)
	cmd.helper.Write("templates/Makefile.tpl", appPath+"/Makefile", cmd.project)
	cmd.helper.Write("templates/go.tpl", appPath+"/go.mod", cmd.project)
	cmd.helper.Write("templates/main.tpl", appPath+"/main.go", cmd.project)

	cmd.helper.Write("templates/app/controllers/health.tpl", appPath+"/app/controllers/health.go", cmd.project)
	cmd.helper.Write("templates/app/initializer/service.tpl", appPath+"/app/initializer/service.go", cmd.project)
	cmd.helper.Write("templates/app/router/router.tpl", appPath+"/app/router/router.go", cmd.project)
	cmd.helper.Write("templates/config/config.tpl", appPath+"/config/config.go", cmd.project)
	cmd.helper.Write("templates/constants/app.tpl", appPath+"/constants/app.go", cmd.project)
	cmd.helper.Write("templates/instance/instance.tpl", appPath+"/instance/instance.go", cmd.project)
	cmd.helper.Write("templates/logger/logger.tpl", appPath+"/logger/logger.go", cmd.project)
	cmd.helper.Write("templates/models/model_registry.tpl", appPath+"/models/model_registry.go", cmd.project)
	cmd.helper.Write("templates/models/health.tpl", appPath+"/models/health.go", cmd.project)
	cmd.helper.Write("templates/instance/registry/models.tpl", appPath+"/instance/registry/models.go", cmd.project)
	cmd.helper.Write("templates/repository/health.tpl", appPath+"/repository/health.go", cmd.project)
	cmd.helper.Write("templates/runner/api.tpl", appPath+"/runner/api.go", cmd.project)
	cmd.helper.Write("templates/service/health.tpl", appPath+"/service/health.go", cmd.project)
	cmd.helper.Write("templates/utils/utils.tpl", appPath+"/utils/utils.go", cmd.project)

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
