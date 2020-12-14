package repo

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"todoapp/lib/dblib"
	"todoapp/pkg/errors"
	"todoapp/todoapp/model"
	"todoapp/todoapp/types"
)

type (
	// Repository ...
	Repository struct {
		db *sqlx.DB
	}

	// TxnRepository ...
	TxnRepository struct {
		tx *sqlx.Tx
	}
)

var _ types.Repository = &Repository{}

var _ types.TxnRepository = &TxnRepository{}

// NewRepository ...
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Transact ...
func (r *Repository) Transact(ctx context.Context, fn func(tx types.TxnRepository) error) error {
	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	err = fn(&TxnRepository{tx: tx})
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return errors.WrapDBError(ctx, err)
	}
	return nil
}

var getTodoQuery = dblib.NewQuery(`
SELECT id, name FROM todos
WHERE id = ? FOR UPDATE
`)

// GetTodo ...
func (r *TxnRepository) GetTodo(ctx context.Context, id model.TodoID) (model.NullTodo, error) {
	var todo model.Todo

	err := r.tx.GetContext(ctx, &todo, getTodoQuery, id)
	if err == sql.ErrNoRows {
		return model.NullTodo{}, nil
	}
	if err != nil {
		return model.NullTodo{}, errors.WrapDBError(ctx, err)
	}

	return model.NullTodo{Valid: true, Todo: todo}, nil
}
