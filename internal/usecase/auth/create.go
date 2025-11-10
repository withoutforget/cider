package auth

import (
	"context"
	"errors"
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

	err := u.txManager.WithTransaction(ctx, func(txCtx context.Context) error {
		user, err := u.userRepository.GetUserByUsername(txCtx, r.Username)
		if err != nil {
			return errors.New("invalid credentials (username)")
		}
		if !u.hasher.IsPasswordValid(user.PasswordHash, r.Password) {
			return errors.New("invalid credentials (password)")
		}

		res, err := u.sessionRepository.Create(txCtx,
			session.CreateSessionModel{
				UserID:   user.ID,
				Username: r.Username,
				Device:   r.Device,
			})
		if err != nil {
			return err
		}

		token = res
		return nil
	})

	if err != nil {
		err_v := err.Error()
		return CreateSessionResponse{Token: nil, Error: &err_v}
	}

	return CreateSessionResponse{Token: &token, Error: nil}
}
