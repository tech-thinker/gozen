package helpers

import (
	"fmt"

	"github.com/tech-thinker/gozen/cmd/repository"
	"github.com/tech-thinker/gozen/models"
)

type CodeHelper interface {
	GenerateModel(doc models.Generator) error
}

type codeHelper struct {
	systemRepo repository.SystemRepo
}

func (h *codeHelper) GenerateModel(doc models.Generator) error {

	modelPath := fmt.Sprintf(`/models/%s.go`, doc.StructName)
	repoPath := fmt.Sprintf(`/repository/%s.go`, doc.StructName)
	svcPath := fmt.Sprintf(`/service/%s.go`, doc.StructName)
	ctrlPath := fmt.Sprintf(`/app/rest/controllers/%s.go`, doc.StructName)

	fileConfigs := []models.FileConfig{
		{TemplatePath: "templates/models/standard.tmpl", Destination: modelPath},
		{TemplatePath: "templates/repository/standard.tmpl", Destination: repoPath},
		{TemplatePath: "templates/service/standard.tmpl", Destination: svcPath},
		{TemplatePath: "templates/app/rest/controllers/standard.tmpl", Destination: ctrlPath},
	}

	fmt.Println("Need to add New Service to Service registry.")

	return h.systemRepo.WriteAll(doc.Project.GetCWD(), fileConfigs, doc)
}

func NewCodeHelper(
	systemRepo repository.SystemRepo,
) CodeHelper {
	return &codeHelper{
		systemRepo: systemRepo,
	}
}
