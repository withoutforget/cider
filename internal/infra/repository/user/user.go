package user

import (
	"context"
	"database/sql"
	"errors"
	"withoutforget/cider/internal/dependencies"
	"withoutforget/cider/internal/infra/repository/txmanager"
)

type UserModel struct {
	ID           int64  `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
}

type CreateUserModel struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
}

type UserRepository struct {
	postgres *sql.DB
}

func NewUserRepository(deps *dependencies.Dependencies) *UserRepository {
	return &UserRepository{
		postgres: deps.Postgres,
	}
}

func (r *UserRepository) getTx(ctx context.Context) (txmanager.ITx, error) {
	if tx, ok := ctx.Value(txmanager.TxKey).(*sql.Tx); ok {
		return tx, nil
	}
	return r.postgres, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, model CreateUserModel) (*UserModel, error) {
	tx, err := r.getTx(ctx)

	if err != nil {
		return nil, err
	}

	query := `
		INSERT INTO users
		(username, password_hash)
		VALUES ($1, $2)
		RETURNING id, username, password_hash;
	`

	res, err := tx.QueryContext(ctx, query, model.Username, model.PasswordHash)

	if err != nil {
		return nil, err
	}

	if !res.Next() {
		return nil, errors.New("cannot get returning value")
	}

	var ret UserModel

	if err := res.Scan(
		&ret.ID,
		&ret.Username,
		&ret.PasswordHash,
	); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*UserModel, error) {
	tx, err := r.getTx(ctx)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT id, username, password_hash
		FROM users
		WHERE username = $1
	`

	row := tx.QueryRowContext(ctx, query, username)

	var user UserModel
	if err := row.Scan(&user.ID, &user.Username, &user.PasswordHash); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
