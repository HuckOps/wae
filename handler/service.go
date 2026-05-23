package handler

import (
	"net/http"
	"wae/config"
	"wae/model"
	"wae/pkg/restful"
	"wae/schema"
	"wae/service"

	"github.com/gin-gonic/gin"
)

func GetServices(c *gin.Context) {
	var params schema.GetServiceRequestParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, restful.Restful[any]{
			Code:    restful.ParamsError,
			Message: err.Error(),
		})
		return
	}
	total, items, err := service.GetService(c, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, restful.Restful[any]{
			Code:    restful.ServerError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, restful.Restful[any]{
		Code:    restful.RequestSuccess,
		Message: "success",
		Data: restful.Pagination[model.Service]{
			Items:    items,
			Total:    total,
			Page:     params.Page,
			PageSize: params.PageSize,
		},
	})
}

func CreateService(c *gin.Context) {
	var req schema.CreateServiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, restful.Restful[any]{
			Code:    restful.ParamsError,
			Message: err.Error(),
		})
		return
	}
	if err := service.CreateService(c, req); err != nil {
		c.JSON(http.StatusInternalServerError, restful.Restful[any]{
			Code:    restful.ServerError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, restful.Restful[any]{
		Code:    restful.RequestSuccess,
		Message: "success",
	})
}

func GetClusters(c *gin.Context) {
	clusters := make([]schema.ClusterInfo, len(config.Config.KubeConfigs))
	for i, k := range config.Config.KubeConfigs {
		clusters[i] = schema.ClusterInfo{Name: k.Name, Tags: k.Tags}
	}
	c.JSON(http.StatusOK, restful.Restful[any]{
		Code:    restful.RequestSuccess,
		Message: "success",
		Data:    clusters,
	})
}
