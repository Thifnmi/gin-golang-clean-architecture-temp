package router

import (
	"github.com/gin-gonic/gin"
	"github.com/thifnmi/gin-golang-clean-architecture-temp/pkg/interfaces/http/handlers"
)

func SetupRouterHealth(router *gin.Engine, healthcheck handlers.HealthHandler) {
	api := router.Group("")
	{
		api.GET("/", healthcheck.Info)
		api.GET("/healthcheck", healthcheck.Heathcheck)
	}
}
