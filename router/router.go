package router

import (
	"wae/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(e *gin.Engine) {
	ServicesRouter(e)
	OIDCRouter(e)
}

func OIDCRouter(e *gin.Engine) {
	OIDCRouterMap := e.Group("/api/v1/oidc")
	{
		OIDCRouterMap.GET("/config", handler.GetOIDCConfig)
	}
}

func ServicesRouter(e *gin.Engine) {
	ServiceRouterMap := e.Group("/api/v1/services")
	ServiceRouterMap.Use(handler.OIDCMiddleware())
	{
		ServiceRouterMap.GET("", handler.GetServices)
		ServiceRouterMap.POST("", handler.CreateService)
	}
	ClustersRouterMap := e.Group("/api/v1/clusters")
	{
		ClustersRouterMap.GET("", handler.GetClusters)
	}
}
