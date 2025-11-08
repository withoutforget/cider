package api

import "github.com/gin-gonic/gin"

type HealthCheckRequest struct {
}

type HealthCheckResponse struct {
	Status string `json:"status"`
}

func (api *API) HealthCheck(c *gin.Context) {
	c.JSON(200, HealthCheckResponse{Status: "ok"})
}
