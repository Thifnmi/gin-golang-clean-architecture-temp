package example

import (
	"github.com/gin-gonic/gin"
	"github.com/thifnmi/gin-golang-clean-architecture-temp/pkg/interfaces/http/handlers"
)

func SetupRouter(router *gin.Engine, exampleHandler handlers.ExampleHandler) {
	api := router.Group("/examples")
	{
		api.GET("", exampleHandler.GetAll)
		api.GET("/:uuid", exampleHandler.GetByID)
	}
}
