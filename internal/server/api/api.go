package api

import (
	"withoutforget/cider/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type API struct {
	Redis *redis.Client
	Cfg   *config.Config
}

func NewAPI() *API {
	return &API{Redis: nil}
}

func (api *API) Setup(eng *gin.Engine) {
	eng.GET("/api/v1/healthcheck", api.HealthCheck)
	eng.POST("/api/v1/auth", api.Auth)
}
