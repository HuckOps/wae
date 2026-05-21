package router

import (
	"wae/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	ServicesRouter(e)
}

func ServicesRouter(e *gin.Engine) {
	ServiceRouterMap := e.Group("/api/v1/services")
	ServiceRouterMap.Use(handler.OIDCMiddleware())
	{
		ServiceRouterMap.POST("")
	}
}
