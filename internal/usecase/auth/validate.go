package auth

import (
	"context"
	"withoutforget/cider/internal/infra/repository/session"
)

type ValidateSessionRequest struct {
	Token string `json:"token"`
}

type ValidateSessionResponse struct {
	Session *session.SessionModel `json:"session,omitempty"`
	Error   *string               `json:"error,omitempty"`
}

func (u *AuthUsecase) ValidateSession(ctx context.Context, r ValidateSessionRequest) ValidateSessionResponse {
	res, err := u.sessionRepository.Validate(ctx, r.Token)
	if err != nil {
		err_v := err.Error()
		return ValidateSessionResponse{Session: nil, Error: &err_v}
	}
	return ValidateSessionResponse{Session: res, Error: nil}
}
