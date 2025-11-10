package auth

import (
	"withoutforget/cider/internal/dependencies"
	"withoutforget/cider/internal/infra/repository/session"
	"withoutforget/cider/internal/infra/repository/txmanager"
	"withoutforget/cider/internal/infra/repository/user"
	"withoutforget/cider/internal/provider"
)

type AuthUsecase struct {
	sessionRepository *session.SessionRepository
	txManager         *txmanager.TxManager
	userRepository    *user.UserRepository
	hasher            *provider.PasswordHasher
}

func NewAuthUsecase(deps *dependencies.Dependencies) *AuthUsecase {
	return &AuthUsecase{
		sessionRepository: session.NewSessionRepository(deps),
		txManager:         txmanager.NewTxManager(deps),
		userRepository:    user.NewUserRepository(deps),
		hasher:            provider.NewPasswordHasher(),
	}
}
