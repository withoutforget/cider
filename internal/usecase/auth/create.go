package auth

import (
	"context"
	"withoutforget/cider/internal/infra/repository/session"
)

type CreateSessionRequest struct {
	Username string `json:"username"`
	Device   string `json:"device"`
}

type CreateSessionResponse struct {
	Token *string `json:"token,omitempty"`
	Error *string `json:"error,omitempty"`
}

func (u *AuthUsecase) CreateSession(ctx context.Context, r CreateSessionRequest) CreateSessionResponse {
	res, err := u.session_repository.Create(ctx,
		session.CreateSessionModel{
			UserID:   0,
			Username: r.Username,
			Device:   r.Device,
		})
	if err != nil {
		err_v := err.Error()
		return CreateSessionResponse{Token: nil, Error: &err_v}
	}
	return CreateSessionResponse{Token: &res, Error: nil}
}
