package api

import "github.com/gin-gonic/gin"

type API struct {
}

func (api *API) Setup(eng *gin.Engine) {
	eng.GET("/api/v1/healthcheck", api.HealthCheck)
}
