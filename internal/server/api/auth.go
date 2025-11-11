package api

import (
	"context"
	"net/http"
	"withoutforget/cider/internal/usecase/auth"

	"github.com/gin-gonic/gin"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (api *API) Auth(c *gin.Context) {
	var request AuthRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, InvalidInputResponse)
		return
	}

	device := c.GetHeader("User-Agent")

	u := auth.NewAuthUsecase(api.deps)

	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	resp := u.CreateSession(ctx, auth.CreateSessionRequest{
		Username: request.Username,
		Password: request.Password,
		Device:   device,
	})

	c.JSON(http.StatusOK, resp)
}

type ValidateAuthRequest struct {
	Token string `json:"token"`
}

func (api *API) ValidateAuth(c *gin.Context) {
	var request ValidateAuthRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, InvalidInputResponse)
		return
	}
	u := auth.NewAuthUsecase(api.deps)

	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	resp := u.ValidateSession(ctx, auth.ValidateSessionRequest{Token: request.Token})

	c.JSON(http.StatusOK, resp)
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (api *API) Register(c *gin.Context) {
	var request RegisterRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, InvalidInputResponse)
		return
	}

	u := auth.NewAuthUsecase(api.deps)

	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	resp := u.RegisterUser(ctx, auth.RegisterUserRequest{
		Username: request.Username,
		Password: request.Password,
	})

	if resp.Error != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

type RevokeRequest struct {
	Token string `json:"token"`
}

func (api *API) Revoke(c *gin.Context) {
	var request RevokeRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, InvalidInputResponse)
		return
	}

	u := auth.NewAuthUsecase(api.deps)

	ctx, cancel := context.WithCancel(c.Request.Context())
	defer cancel()

	resp := u.RevokeSession(ctx, auth.RevokeSessionRequest{
		Token: request.Token,
	})

	if resp.Error != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}
