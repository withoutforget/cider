package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckRequest struct {
	Echo string `form:"echo"`
}

type HealthCheckResponse struct {
	Status string `json:"status"`
	Echo   string `json:"echo,omitempty"`
}

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}

func (api *API) HealthCheck(c *gin.Context) {
	var request HealthCheckRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{ErrorMessage: "Invalid input"})
		return
	}
	c.JSON(http.StatusOK, HealthCheckResponse{Status: "ok", Echo: request.Echo})
}
