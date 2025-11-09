package auth

import "withoutforget/cider/internal/infra/repository/session"

type AuthUsecase struct {
	session_repository *session.SessionRepository
}

func NewAuthUsecase(session_repository *session.SessionRepository) *AuthUsecase {
	return &AuthUsecase{
		session_repository: session_repository,
	}
}
