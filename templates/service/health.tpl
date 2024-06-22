package service

import (
	"context"

	"{{.PackageName}}/models"
	"{{.PackageName}}/repository"
)

type HealthSvc interface {
	Ping(ctx context.Context) (*models.Health, error)
}

type healthSvc struct {
    healthRepo repository.HealthRepo
}

func (svc *healthSvc) Ping(ctx context.Context) (*models.Health, error) {

	// groupError := "PING_HEALTH_SERVICE"
    // logger.Log.WithError(err).Error(groupError)
    doc := models.Health{
        Success: true,
        Message: "Service is up and running",
    }

	return &doc, nil
}

func NewHealthSvc(
    healthRepo repository.HealthRepo,
) HealthSvc {
	return &healthSvc{
        healthRepo: healthRepo,
	}
}
