package repository

import (
    "context"
    "gorm.io/gorm"

	"{{.PackageName}}/logger"
	"{{.PackageName}}/models"
)

type HealthRepo interface {
	Ping(ctx context.Context) (bool, error)
}

type healthRepo struct {
	db *gorm.DB
}

func (r *healthRepo) Ping(ctx context.Context) (bool, error) {
    doc := models.Health{
		Success: true,
		Message: "Service is up and running.",
	}
	logger.Log.Info(doc)
	return true, nil
}

func NewHealthRepo(db *gorm.DB) HealthRepo {
	return &healthRepo{
		db: db,
	}
}
