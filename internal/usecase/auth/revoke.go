package auth

import (
	"context"
)

type RevokeSessionRequest struct {
	Token string `json:"token"`
}
type RevokeSessionResponse struct {
	Error *string `json:"error,omitempty"`
}

func (u *AuthUsecase) RevokeSession(ctx context.Context, r RevokeSessionRequest) RevokeSessionResponse {
	err := u.sessionRepository.Revoke(ctx, r.Token)

	var err_str *string
	if err != nil {
		t := err.Error()
		err_str = &t
	}

	return RevokeSessionResponse{Error: err_str}
}
