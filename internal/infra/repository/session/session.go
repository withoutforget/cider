package session

import (
	"context"
	"encoding/json"
	"errors"
	"time"
	"withoutforget/cider/internal/config"
	"withoutforget/cider/internal/dependencies"
	"withoutforget/cider/internal/provider"

	"github.com/redis/go-redis/v9"
)

type SessionModel struct {
	UserID    int64     `json:"user_id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`
	Device    string    `json:"device"`
}

type CreateSessionModel struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Device   string `json:"device"`
}

type SessionRepository struct {
	r        *redis.Client
	datetime *provider.DatetimeProvider
	token    *provider.TokenProvider
	cfg      *config.Session
}

func NewSessionRepository(deps *dependencies.Dependencies) *SessionRepository {
	return &SessionRepository{
		r:        deps.Redis,
		datetime: provider.NewDatetimeProvider(),
		token:    provider.NewTokenProvider(),
		cfg:      deps.Config.Session,
	}
}

func (s *SessionRepository) Create(ctx context.Context, model CreateSessionModel) (string, error) {
	current_time := s.datetime.Now()

	session_model := SessionModel{
		UserID:    model.UserID,
		Username:  model.Username,
		CreatedAt: current_time,
		ExpiredAt: current_time.Add(time.Duration(s.cfg.Timeout)),
		Device:    model.Device,
	}

	bytes, err := json.Marshal(session_model)
	if err != nil {
		return "", err
	}

	token := s.token.Provide(bytes)

	if token == "" {
		return "", errors.New("error during generation of the token")
	}

	res := s.r.SetEx(
		ctx,
		"session_"+token,
		bytes,
		time.Duration(s.cfg.Timeout)*time.Second,
	)

	if res.Err() != nil {
		return "", res.Err()
	}

	return token, nil

}

func (s *SessionRepository) Validate(ctx context.Context, token string) (*SessionModel, error) {
	res := s.r.Get(
		ctx,
		"session_"+token,
	)

	if res.Err() != nil {
		return nil, res.Err()
	}

	data := res.Val()
	var model SessionModel
	err := json.Unmarshal([]byte(data), &model)
	if err != nil {
		return nil, err
	}

	current_time := s.datetime.Now()

	if current_time.Compare(model.ExpiredAt) == 1 {
		return &model, nil
	}

	return nil, errors.New("session expired")
}

func (s *SessionRepository) Revoke(ctx context.Context, token string) error {
	res := s.r.Del(
		ctx,
		"session_"+token,
	)

	if res.Err() != nil {
		return res.Err()
	}

	return nil
}
