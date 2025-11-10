package dependencies

import (
	"database/sql"
	"withoutforget/cider/internal/config"
	"withoutforget/cider/internal/provider"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

type Dependencies struct {
	Redis    *redis.Client
	Postgres *sql.DB
	Config   *config.Config
	Hasher   *provider.PasswordHasher
}

func NewDependencies(
	config *config.Config,
) *Dependencies {
	pg, err := sql.Open(
		"postgres",
		config.Postgres.Dsn(),
	)
	if err != nil {
		panic(err)
	}
	return &Dependencies{
		Config:   config,
		Postgres: pg,
		Redis: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}), // TODO: use real creds from config
		Hasher: provider.NewPasswordHasher(),
	}
}
