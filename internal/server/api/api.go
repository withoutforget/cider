package api

import (
	"withoutforget/cider/internal/dependencies"

	"github.com/gin-gonic/gin"
)

type API struct {
	deps *dependencies.Dependencies
}

func NewAPI(deps *dependencies.Dependencies) *API {
	return &API{deps: deps}
}

func (api *API) Setup(eng *gin.Engine) {
	eng.GET("/api/v1/healthcheck", api.HealthCheck)
	eng.POST("/api/v1/auth", api.Auth)
	eng.POST("/api/v1/auth/validate", api.ValidateAuth)
	eng.POST("/api/v1/auth/register", api.Register)
}
