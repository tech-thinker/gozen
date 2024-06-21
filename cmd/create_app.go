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
    helper helper.CommonHelper
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
    cmd.helper.Write("templates/gitignore.tpl", appPath+"/.gitignore", cmd.project)
    cmd.helper.Write("templates/Makefile.tpl", appPath+"/Makefile", cmd.project)
    cmd.helper.Write("templates/go.tpl", appPath+"/go.mod", cmd.project)
    cmd.helper.Write("templates/main.tpl", appPath+"/main.go", cmd.project)
    
    cmd.helper.Write("templates/config/config.tpl", appPath+"/config/config.go", cmd.project)
    cmd.helper.Write("templates/constants/app.tpl", appPath+"/constants/app.go", cmd.project)
    cmd.helper.Write("templates/instance/instance.tpl", appPath+"/instance/instance.go", cmd.project)
    cmd.helper.Write("templates/logger/logger.tpl", appPath+"/logger/logger.go", cmd.project)
    cmd.helper.Write("templates/models/model_registry.tpl", appPath+"/models/model_registry.go", cmd.project)
    cmd.helper.Write("templates/utils/utils.tpl", appPath+"/utils/utils.go", cmd.project)
    
	return nil
}

func NewAppCmd(
    project models.Project,
    helper helper.CommonHelper,
) AppCmd {
    return &appCmd{
        project: project,
        helper: helper,
    }
}
