package router

import (
	"github.com/gin-gonic/gin"
	"{{.PackageName}}/app/controllers"
	"{{.PackageName}}/app/initializer"
)

// Init sets router
func Init(svc initializer.Services) *gin.Engine {
	router := gin.Default()

	health := controllers.NewHealthController(svc)
	router.GET("/ping", health.Ping)

	return router
}