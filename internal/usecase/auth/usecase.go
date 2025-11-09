package auth

import (
	"withoutforget/cider/internal/infra/dependencies"
	"withoutforget/cider/internal/infra/repository/session"
)

type AuthUsecase struct {
	session_repository *session.SessionRepository
}

func NewAuthUsecase(deps *dependencies.Dependencies) *AuthUsecase {
	return &AuthUsecase{
		session_repository: session.NewSessionRepository(deps),
	}
}
