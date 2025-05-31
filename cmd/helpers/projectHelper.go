package helpers

import (
	"github.com/tech-thinker/gozen/cmd/repository"
	"github.com/tech-thinker/gozen/models"
)

type ProjectHelper interface {
	InitProject(project models.Project) error
	SetupEnv(project models.Project) error
	SetupDocker(project models.Project) error
	SetupUtils(project models.Project) error
	SetupRepository(project models.Project) error
	SetupService(project models.Project) error
	SetupModel(project models.Project) error
	SetupLogger(project models.Project) error
	SetupConfig(project models.Project) error
	SetupConstants(project models.Project) error
	CreateRestAPI(project models.Project) error
	CreategRPCAPI(project models.Project) error
}

type projectHelper struct {
	systemRepo repository.SystemRepo
}

func (h *projectHelper) InitProject(project models.Project) error {
	fileConfigs := []models.FileConfig{
		{TemplatePath: "templates/gitignore.tmpl", Destination: "/.gitignore"},
		{TemplatePath: "templates/Makefile.tmpl", Destination: "/Makefile"},
		{TemplatePath: "templates/README.md", Destination: "/README.md"},
		{TemplatePath: "templates/go.tmpl", Destination: "/go.mod"},
		{TemplatePath: "templates/main.tmpl", Destination: "/main.go"},
		{TemplatePath: "templates/editorconfig.tmpl", Destination: "/.editorconfig"},
		{TemplatePath: "templates/instance/instance.tmpl", Destination: "/instance/instance.go"},
		{TemplatePath: "templates/app/init.tmpl", Destination: "/app/init.go"},
	}

	return h.systemRepo.WriteAll(project.GetAppDir(), fileConfigs, project)
}

func (h *projectHelper) SetupEnv(project models.Project) error {
	fileConfigs := []models.FileConfig{
		{TemplatePath: "templates/env.sample.tmpl", Destination: "/.env"},
		{TemplatePath: "templates/env.sample.tmpl", Destination: "/.env.sample"},
	}

	return h.systemRepo.WriteAll(project.GetAppDir(), fileConfigs, project)
}

func (h *projectHelper) SetupDocker(project models.Project) error {
	fileConfigs := []models.FileConfig{
		{TemplatePath: "templates/docker/Dockerfile.debug", Destination: "/docker/Dockerfile.debug"},
		{TemplatePath: "templates/docker/Dockerfile.dev", Destination: "/docker/Dockerfile.dev"},
		{TemplatePath: "templates/docker/modd-debug.conf", Destination: "/docker/modd-debug.conf"},
		{TemplatePath: "templates/docker/modd-dev.conf", Destination: "/docker/modd-dev.conf"},
		{TemplatePath: "templates/docker/Dockerfile.prod", Destination: "Dockerfile"},
		{TemplatePath: "templates/docker-compose-debug.yml.tmpl", Destination: "/docker-compose-debug.yml"},
		{TemplatePath: "templates/docker-compose.yml.tmpl", Destination: "/docker-compose.yml"},
	}

	return h.systemRepo.WriteAll(project.GetAppDir(), fileConfigs, project)
}

func (h *projectHelper) SetupUtils(project models.Project) error {
	return h.systemRepo.Write(project.GetAppDir(), "templates/utils/utils.tmpl", "/utils/utils.go", project)
}

func (h *projectHelper) SetupRepository(project models.Project) error {
	return h.systemRepo.Write(project.GetAppDir(), "templates/repository/health.tmpl", "/repository/health.go", project)
}

func (h *projectHelper) SetupService(project models.Project) error {
	return h.systemRepo.Write(project.GetAppDir(), "templates/service/health.tmpl", "/service/health.go", project)
}

func (h *projectHelper) SetupModel(project models.Project) error {
	fileConfigs := []models.FileConfig{
		{TemplatePath: "templates/instance/registry/models.tmpl", Destination: "/instance/registry/models.go"},
		{TemplatePath: "templates/models/health.tmpl", Destination: "/models/health.go"},
	}

	return h.systemRepo.WriteAll(project.GetAppDir(), fileConfigs, project)
}

func (h *projectHelper) SetupLogger(project models.Project) error {
	return h.systemRepo.Write(project.GetAppDir(), "templates/logger/logger.tmpl", "/logger/logger.go", project)
}

func (h *projectHelper) SetupConfig(project models.Project) error {
	return h.systemRepo.Write(project.GetAppDir(), "templates/config/config.tmpl", "/config/config.go", project)
}

func (h *projectHelper) SetupConstants(project models.Project) error {
	return h.systemRepo.Write(project.GetAppDir(), "templates/constants/app.tmpl", "/constants/app.go", project)
}

func (h *projectHelper) CreateRestAPI(project models.Project) error {
	fileConfigs := []models.FileConfig{
		{TemplatePath: "templates/app/rest/controllers/health.tmpl", Destination: "/app/rest/controllers/health.go"},
		{TemplatePath: "templates/app/rest/router/router.tmpl", Destination: "/app/rest/router/router.go"},
		{TemplatePath: "templates/runner/api.tmpl", Destination: "/runner/api.go"},
	}

	return h.systemRepo.WriteAll(project.GetAppDir(), fileConfigs, project)
}

func (h *projectHelper) CreategRPCAPI(project models.Project) error {
	fileConfigs := []models.FileConfig{
		{TemplatePath: "templates/app/grpc/handlers/health.tmpl", Destination: "/app/grpc/handlers/health.go"},
		{TemplatePath: "templates/app/grpc/proto/health.proto", Destination: "/app/grpc/proto/health.proto"},
		{TemplatePath: "templates/app/grpc/proto/health.pb.tmpl", Destination: "/app/grpc/proto/health.pb.go"},
		{TemplatePath: "templates/app/grpc/proto/health_grpc.pb.tmpl", Destination: "/app/grpc/proto/health_grpc.pb.go"},
		{TemplatePath: "templates/app/grpc/router/router.tmpl", Destination: "/app/grpc/router/router.go"},
		{TemplatePath: "templates/runner/grpc.tmpl", Destination: "/runner/grpc.go"},
	}

	return h.systemRepo.WriteAll(project.GetAppDir(), fileConfigs, project)
}

func NewProjectHelper(
	systemRepo repository.SystemRepo,
) ProjectHelper {
	return &projectHelper{
		systemRepo: systemRepo,
	}
}
