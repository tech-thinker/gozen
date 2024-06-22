package initializer

import (
	"{{.PackageName}}/config"
	"{{.PackageName}}/instance"
	"{{.PackageName}}/repository"
	"{{.PackageName}}/service"
)

// Services is interface for all service entrypoint
type Services interface {
    HealthService() service.HealthSvc
	// TODO: add service dependency here
}

type services struct {
    healthSvc service.HealthSvc
	// TODO: add service dependency here
}

func (svc *services) HealthService() service.HealthSvc {
    return svc.healthSvc
}


// Init initializes services repo
func Init(cfg config.Configuration, instance instance.Instance) Services {
    // Object init
    db := instance.DB()
    
    // Repository init
    healthRepo := repository.NewHealthRepo(db)
    
    // Service init
    healthSvc := service.NewHealthSvc(healthRepo)
	
    return &services{
        healthSvc: healthSvc,
    }
}
