package api

import (
	"context"
	"net/http"
	"withoutforget/cider/internal/infra/repository/session"
	"withoutforget/cider/internal/usecase/auth"

	"github.com/gin-gonic/gin"
)

type AuthRequest struct {
	Username string `json:"username"`
}

func (api *API) Auth(c *gin.Context) {
	var request AuthRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, InvalidInputResponse)
		return
	}

	device := c.GetHeader("User-Agent")

	u := auth.NewAuthUsecase(session.NewSessionRepository(api.Cfg.Session))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resp := u.CreateSession(ctx, auth.CreateSessionRequest{
		Username: request.Username,
		Device:   device,
	})

	c.JSON(http.StatusOK, resp)
}
