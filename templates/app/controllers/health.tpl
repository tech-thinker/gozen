package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"{{.PackageName}}/app/initializer"
)

// HealthController is an interface for health controller
type HealthController interface {
	Ping(ctx *gin.Context)
}

type healthController struct {
	svc initializer.Services
}

func (c *healthController) Ping(ctx *gin.Context) {
    res, err := c.svc.HealthService().Ping(ctx)
	if err!=nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": res.Success,
		"message": res.Message,
	})
}

// NewHealthController initializes health controller with dependency
func NewHealthController(svc initializer.Services) HealthController {
	return &healthController{
		svc: svc,
	}
}
