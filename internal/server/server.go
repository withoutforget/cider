package server

import (
	"context"
	"strconv"
	"withoutforget/cider/internal/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	eng *gin.Engine
	ctx context.Context
	cfg *config.Config
}

func NewServer(ctx context.Context, cfg *config.Config) *Server {
	var srv Server

	srv.eng = gin.New()
	srv.ctx = ctx
	srv.cfg = cfg

	return &srv
}

func (s *Server) Run() error {
	return s.eng.Run(
		s.cfg.Server.Host + ":" + strconv.Itoa(s.cfg.Server.Port),
	)
}
