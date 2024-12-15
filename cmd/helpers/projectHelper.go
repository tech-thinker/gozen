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
	err := h.systemRepo.Write("templates/gitignore.tmpl", project.GetAppDir()+"/.gitignore", project)
	if err != nil {
		return err
	}
	err = h.systemRepo.Write("templates/Makefile.tmpl", project.GetAppDir()+"/Makefile", project)
	if err != nil {
		return err
	}
	err = h.systemRepo.Write("templates/go.tmpl", project.GetAppDir()+"/go.mod", project)
	if err != nil {
		return err
	}
	err = h.systemRepo.Write("templates/main.tmpl", project.GetAppDir()+"/main.go", project)
	if err != nil {
		return err
	}
	err = h.systemRepo.Write("templates/instance/instance.tmpl", project.GetAppDir()+"/instance/instance.go", project)
	if err != nil {
		return err
	}
	err = h.systemRepo.Write("templates/app/init.tmpl", project.GetAppDir()+"/app/init.go", project)
	if err != nil {
		return err
	}
	return nil
}

func (h *projectHelper) SetupEnv(project models.Project) error {
	err := h.systemRepo.Write("templates/env.sample.tmpl", project.GetAppDir()+"/.env", project)
	if err != nil {
		return err
	}
	err = h.systemRepo.Write("templates/env.sample.tmpl", project.GetAppDir()+"/.env.sample", project)
	if err != nil {
		return err
	}
	err = h.systemRepo.Write("templates/env.sample.tmpl", project.GetAppDir()+"/docker/.env", project)
	if err != nil {
		return err
	}
	err = h.systemRepo.Write("templates/env.sample.tmpl", project.GetAppDir()+"/docker/.env.sample", project)
	if err != nil {
		return err
	}
	return nil
}

func (h *projectHelper) SetupDocker(project models.Project) error {
	err := h.systemRepo.Write("templates/docker/Dockerfile.debug", project.GetAppDir()+"/docker/Dockerfile.debug", project)
	if err != nil {
		return err
	}
	err = h.systemRepo.Write("templates/docker/Dockerfile.dev", project.GetAppDir()+"/docker/Dockerfile.dev", project)
	if err != nil {
		return err
	}
	err = h.systemRepo.Write("templates/docker/Dockerfile.prod", project.GetAppDir()+"/docker/Dockerfile.prod", project)
	if err != nil {
		return err
	}
	err = h.systemRepo.Write("templates/docker/docker-compose-debug.yml", project.GetAppDir()+"/docker/docker-compose-debug.yml", project)
	if err != nil {
		return err
	}
	err = h.systemRepo.Write("templates/docker/docker-compose.yml", project.GetAppDir()+"/docker/docker-compose.yml", project)
	if err != nil {
		return err
	}
	err = h.systemRepo.Write("templates/docker/modd-debug.conf", project.GetAppDir()+"/docker/modd-debug.conf", project)
	if err != nil {
		return err
	}
	err = h.systemRepo.Write("templates/docker/modd-dev.conf", project.GetAppDir()+"/docker/modd-dev.conf", project)
	if err != nil {
		return err
	}
	return nil
}

func (h *projectHelper) SetupUtils(project models.Project) error {
	return h.systemRepo.Write("templates/utils/utils.tmpl", project.GetAppDir()+"/utils/utils.go", project)
}

func (h *projectHelper) SetupRepository(project models.Project) error {
	return h.systemRepo.Write("templates/repository/health.tmpl", project.GetAppDir()+"/repository/health.go", project)
}

func (h *projectHelper) SetupService(project models.Project) error {
	return h.systemRepo.Write("templates/service/health.tmpl", project.GetAppDir()+"/service/health.go", project)
}

func (h *projectHelper) SetupModel(project models.Project) error {
	err := h.systemRepo.Write("templates/instance/registry/models.tmpl", project.GetAppDir()+"/instance/registry/models.go", project)
	if err != nil {
		return err
	}
	err = h.systemRepo.Write("templates/models/health.tmpl", project.GetAppDir()+"/models/health.go", project)
	if err != nil {
		return err
	}

	return nil
}

func (h *projectHelper) SetupLogger(project models.Project) error {
	return h.systemRepo.Write("templates/logger/logger.tmpl", project.GetAppDir()+"/logger/logger.go", project)
}

func (h *projectHelper) SetupConfig(project models.Project) error {
	return h.systemRepo.Write("templates/config/config.tmpl", project.GetAppDir()+"/config/config.go", project)
}

func (h *projectHelper) SetupConstants(project models.Project) error {
	return h.systemRepo.Write("templates/constants/app.tmpl", project.GetAppDir()+"/constants/app.go", project)
}

func (h *projectHelper) CreateRestAPI(project models.Project) error {
	err := h.systemRepo.Write("templates/app/rest/controllers/health.tmpl", project.GetAppDir()+"/app/rest/controllers/health.go", project)
	if err != nil {
		return err
	}
	err = h.systemRepo.Write("templates/app/rest/router/router.tmpl", project.GetAppDir()+"/app/rest/router/router.go", project)
	if err != nil {
		return err
	}
	err = h.systemRepo.Write("templates/runner/api.tmpl", project.GetAppDir()+"/runner/api.go", project)
	if err != nil {
		return err
	}
	return nil
}

func (h *projectHelper) CreategRPCAPI(project models.Project) error {
	err := h.systemRepo.Write("templates/app/grpc/handlers/health.tmpl", project.GetAppDir()+"/app/grpc/handlers/health.go", project)
	if err != nil {
		return err
	}
	err = h.systemRepo.Write("templates/app/grpc/proto/health.proto", project.GetAppDir()+"/app/grpc/proto/health.proto", project)
	if err != nil {
		return err
	}
	err = h.systemRepo.Write("templates/app/grpc/proto/health.pb.tmpl", project.GetAppDir()+"/app/grpc/proto/health.pb.go", project)
	if err != nil {
		return err
	}
	err = h.systemRepo.Write("templates/app/grpc/proto/health_grpc.pb.tmpl", project.GetAppDir()+"/app/grpc/proto/health_hrpc.pb.go", project)
	if err != nil {
		return err
	}
	err = h.systemRepo.Write("templates/app/grpc/router/router.tmpl", project.GetAppDir()+"/app/grpc/router/router.go", project)
	if err != nil {
		return err
	}
	err = h.systemRepo.Write("templates/runner/grpc.tmpl", project.GetAppDir()+"/runner/grpc.go", project)
	if err != nil {
		return err
	}
	return nil
}

func NewProjectHelper(
	systemRepo repository.SystemRepo,
) ProjectHelper {
	return &projectHelper{
		systemRepo: systemRepo,
	}
}
