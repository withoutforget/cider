package server

import (
	"context"
	"log/slog"
	"net/http"
	"regexp"
	"strconv"
	"withoutforget/cider/internal/config"
	"withoutforget/cider/internal/server/api"

	"github.com/gin-contrib/cors"
	gslog "github.com/gin-contrib/slog"
	"github.com/gin-gonic/gin"
)

type Server struct {
	srv *http.Server
	eng *gin.Engine
	ctx context.Context
	cfg *config.Config
	api *api.API
}

func NewServer(ctx context.Context, cfg *config.Config) *Server {
	var srv Server

	srv.eng = makeGin(cfg)
	srv.ctx = ctx
	srv.cfg = cfg
	srv.srv = &http.Server{
		Addr:    cfg.Server.Host + ":" + strconv.Itoa(cfg.Server.Port),
		Handler: srv.eng,
	}
	srv.api = &api.API{}

	srv.api.Setup(srv.eng)

	return &srv
}

func (srv *Server) Run() error {
	return srv.srv.ListenAndServe()
}

func (srv *Server) Shutdown() error {
	return srv.srv.Shutdown(srv.ctx)
}
func makeGin(cfg *config.Config) *gin.Engine {
	eng := gin.New()
	eng.Use(cors.New(
		cors.Config{
			AllowAllOrigins:  cfg.Server.AllowAllOrigins,
			AllowOrigins:     cfg.Server.AllowOrigins,
			AllowMethods:     cfg.Server.AllowMethods,
			AllowHeaders:     cfg.Server.AllowHeaders,
			AllowCredentials: cfg.Server.AllowCredentials,
			AllowWildcard:    cfg.Server.AllowWildcard,
			AllowWebSockets:  cfg.Server.AllowWebSockets,
		}))
	eng.Use(gslog.SetLogger(
		gslog.WithUTC(true),
		gslog.WithSkipPathRegexps(regexp.MustCompile(`\.ico$`), regexp.MustCompile(`^/static/`)),
		gslog.WithMessage("Handled request"),
		gslog.WithLogger(func(c *gin.Context, l *slog.Logger) *slog.Logger {
			return slog.Default().With("request_id", c.GetString("request_id"))
		}),
	))
	return eng
}
