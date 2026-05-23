package handler

import (
	"net/http"
	"wae/config"
	"wae/pkg/restful"

	"github.com/gin-gonic/gin"
)

type OIDCConfigResponse struct {
	Provider string `json:"provider"`
	ClientID string `json:"client_id"`
}

func GetOIDCConfig(c *gin.Context) {
	c.JSON(http.StatusOK, restful.Restful[OIDCConfigResponse]{
		Code:    restful.RequestSuccess,
		Message: "success",
		Data: OIDCConfigResponse{
			Provider: config.Config.OIDCConfig.Provider,
			ClientID: config.Config.OIDCConfig.ClientID,
		},
	})
}
