package cmd

import (
	"embed"
	"fmt"
	"path/filepath"

	"github.com/tech-thinker/gozen/constants"
	"github.com/tech-thinker/gozen/models"
	"github.com/tech-thinker/gozen/utils"
)

type AppCmd interface {
    CreateApp() error
}

type appCmd struct {
    templatesFS embed.FS
    project models.Project
}

func (cmd *appCmd) CreateApp() error {
	appPath := fmt.Sprintf(`%s/%s`, constants.CURRENT_WORKING_DIRECTORY, cmd.project.AppName)
	fmt.Println("Project created at location: " + appPath)
	utils.CreateDirectory(appPath)
    // Write project config into file
    if err := cmd.project.WriteToJsonFile(); err != nil {
        return err
    }
    
    cmd.write("templates/env.sample.tpl", appPath+"/.env", cmd.project)
    cmd.write("templates/env.sample.tpl", appPath+"/.env.sample", cmd.project)
    cmd.write("templates/gitignore.tpl", appPath+"/.gitignore", cmd.project)
    cmd.write("templates/Makefile.tpl", appPath+"/Makefile", cmd.project)
    cmd.write("templates/go.tpl", appPath+"/go.mod", cmd.project)
    cmd.write("templates/main.tpl", appPath+"/main.go", cmd.project)
    
    cmd.write("templates/config/config.tpl", appPath+"/config/config.go", cmd.project)
    cmd.write("templates/constants/app.tpl", appPath+"/constants/app.go", cmd.project)
    cmd.write("templates/instance/instance.tpl", appPath+"/instance/instance.go", cmd.project)
    cmd.write("templates/logger/logger.tpl", appPath+"/logger/logger.go", cmd.project)
    cmd.write("templates/models/model_registry.tpl", appPath+"/models/model_registry.go", cmd.project)
    cmd.write("templates/utils/utils.tpl", appPath+"/utils/utils.go", cmd.project)
    


	return nil
}

func (cmd *appCmd) write(templatePath string, outputPath string, data interface{}) error {
    baseDir := filepath.Dir(outputPath)
    utils.CreateDirectory(baseDir)
    tpl, _ := utils.GenerateCode(cmd.templatesFS, templatePath, data)
    return utils.WriteFile(outputPath, tpl)
}

func NewAppCmd(tpl embed.FS, project models.Project) AppCmd {
    return &appCmd{
        templatesFS: tpl,
        project: project,
    }
}
