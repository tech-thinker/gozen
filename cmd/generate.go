package cmd

import (
	"strings"

	"github.com/tech-thinker/gozen/models"
	"github.com/tech-thinker/gozen/utils"
	"github.com/urfave/cli/v2"
)

func (c *app) GenerateModel() *cli.Command {
	return &cli.Command{
		Name:    "generate",
		Aliases: []string{"g"},
		Usage:   "Code generator.",
		Action: func(ctx *cli.Context) error {

			var project models.Project
			project.LoadFromJsonFile()

			// project := models.Project{
			// 	AppName:     ctx.Args().Get(0),
			// 	PackageName: "github.com/mrasif/demo1",
			// 	WorkingDir:  "./",
			// }

			// project.AutoFixes()

			generateType := ctx.Args().Get(0)
			name := ctx.Args().Get(1)

			interfaceName, structName := utils.TransformString(name)
			route := strings.ToLower(structName)

			generator := models.Generator{
				Project:       project,
				Type:          generateType,
				InterfaceName: interfaceName,
				StructName:    structName,
				RouteName:     route,
			}

			return c.appService.GenerateModel(generator)
		},
	}
}
