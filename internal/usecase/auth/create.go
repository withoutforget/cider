package auth

import (
	"context"
	"withoutforget/cider/internal/infra/repository/session"
)

type CreateSessionRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Device   string `json:"device"`
}
type CreateSessionResponse struct {
	Token *string `json:"token,omitempty"`
	Error *string `json:"error,omitempty"`
}

func (u *AuthUsecase) CreateSession(ctx context.Context, r CreateSessionRequest) CreateSessionResponse {
	var token string

	user, err := u.userRepository.GetUserByUsername(ctx, r.Username)
	if err != nil {
		err_v := "invalid credentials (username)"
		return CreateSessionResponse{Token: nil, Error: &err_v}
	}
	if !u.hasher.IsPasswordValid(user.PasswordHash, r.Password) {
		err_v := "invalid credentials (password)"
		return CreateSessionResponse{Token: nil, Error: &err_v}
	}

	res, err := u.sessionRepository.Create(ctx,
		session.CreateSessionModel{
			UserID:   user.ID,
			Username: r.Username,
			Device:   r.Device,
		})

	token = res

	if err != nil {
		err_v := err.Error()
		return CreateSessionResponse{Token: nil, Error: &err_v}
	}

	return CreateSessionResponse{Token: &token, Error: nil}
}
