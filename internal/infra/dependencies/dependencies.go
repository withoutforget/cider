package dependencies

import (
	"withoutforget/cider/internal/config"

	"github.com/redis/go-redis/v9"
)

type Dependencies struct {
	Redis  *redis.Client
	Config *config.Config
}

func NewDependencies(
	config *config.Config,
) *Dependencies {
	return &Dependencies{
		Config: config,
		Redis: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}), // TODO: use real creds from config
	}
}
