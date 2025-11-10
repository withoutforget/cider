package auth

import (
	"context"
	"withoutforget/cider/internal/infra/repository/user"
)

type RegisterUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterUserResponse struct {
	UserID *int64  `json:"user_id,omitempty"`
	Error  *string `json:"error,omitempty"`
}

func (u *AuthUsecase) RegisterUser(ctx context.Context, r RegisterUserRequest) RegisterUserResponse {
	var userID int64

	err := u.txManager.WithTransaction(ctx, func(txCtx context.Context) error {
		hashPassword, err := u.hasher.HashPassword(r.Password)
		if err != nil {
			return err
		}

		userModel, err := u.userRepository.CreateUser(txCtx, user.CreateUserModel{
			Username:     r.Username,
			PasswordHash: hashPassword,
		})
		if err != nil {
			return err
		}

		userID = userModel.ID
		return nil
	})

	if err != nil {
		err_v := err.Error()
		return RegisterUserResponse{UserID: nil, Error: &err_v}
	}

	return RegisterUserResponse{UserID: &userID, Error: nil}
}
