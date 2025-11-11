package txmanager

import (
	"context"
	"database/sql"
	"withoutforget/cider/internal/dependencies"
)

type ContextKey string

const TxKey ContextKey = "tx"

type ITx interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type TxManager struct {
	db *sql.DB
}

func NewTxManager(deps *dependencies.Dependencies) *TxManager {
	return &TxManager{
		db: deps.Postgres,
	}
}

func (tm *TxManager) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := tm.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	ctx = context.WithValue(ctx, TxKey, tx)

	if err := fn(ctx); err != nil {
		return err
	}

	return tx.Commit()
}
