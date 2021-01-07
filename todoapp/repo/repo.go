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
		return errors.WrapDBError(ctx, err)
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

var insertTodoQuery = dblib.NewNamedQuery(`
INSERT INTO todos (name) VALUES (:name)
`)

// InsertTodo ...
func (r *TxnRepository) InsertTodo(ctx context.Context, save model.Todo) (model.TodoID, error) {
	res, err := r.tx.NamedExecContext(ctx, insertTodoQuery, save)
	if err != nil {
		return 0, errors.WrapDBError(ctx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, errors.WrapDBError(ctx, err)
	}
	return model.TodoID(id), nil
}

var updateTodoQuery = dblib.NewNamedQuery(`
UPDATE todos
SET name = :name, updated_at = CURRENT_TIMESTAMP
WHERE id = :id
`)

// UpdateTodo ...
func (r *TxnRepository) UpdateTodo(ctx context.Context, save model.Todo) error {
	_, err := r.tx.NamedExecContext(ctx, updateTodoQuery, save)
	if err != nil {
		return errors.WrapDBError(ctx, err)
	}
	return nil
}

var getTodoItemsByTodoIDQuery = dblib.NewQuery(`
SELECT id, todo_id, name FROM todo_items WHERE todo_id = ?
`)

// GetTodoItemsByTodoID ...
func (r *TxnRepository) GetTodoItemsByTodoID(ctx context.Context, todoID model.TodoID,
) ([]model.TodoItem, error) {
	var result []model.TodoItem
	err := r.tx.SelectContext(ctx, &result, getTodoItemsByTodoIDQuery, todoID)
	if err != nil {
		return nil, errors.WrapDBError(ctx, err)
	}
	return result, nil
}

var deleteTodoItemsQuery = dblib.NewQuery(`
DELETE FROM todo_items WHERE id IN (?)
`)

// DeleteTodoItems ...
func (r *TxnRepository) DeleteTodoItems(ctx context.Context, todoItemIDs []model.TodoItemID) error {
	if len(todoItemIDs) == 0 {
		return nil
	}

	query, args, err := sqlx.In(deleteTodoItemsQuery, todoItemIDs)
	if err != nil {
		return errors.WrapDBError(ctx, err)
	}

	_, err = r.tx.ExecContext(ctx, query, args...)
	return errors.WrapDBError(ctx, err)
}

var insertTodoItemQuery = dblib.NewNamedQuery(`
INSERT INTO todo_items (todo_id, name) VALUES (:todo_id, :name)
`)

// InsertTodoItem ...
func (r *TxnRepository) InsertTodoItem(ctx context.Context, save model.TodoItem,
) (model.TodoItemID, error) {
	res, err := r.tx.NamedExecContext(ctx, insertTodoItemQuery, save)
	if err != nil {
		return 0, errors.WrapDBError(ctx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, errors.WrapDBError(ctx, err)
	}
	return model.TodoItemID(id), nil
}

var updateTodoItemQuery = dblib.NewNamedQuery(`
UPDATE todo_items SET name = :name WHERE id = :id
`)

// UpdateTodoITem ...
func (r *TxnRepository) UpdateTodoITem(ctx context.Context, save model.TodoItem) error {
	_, err := r.tx.NamedExecContext(ctx, updateTodoItemQuery, save)
	if err != nil {
		return errors.WrapDBError(ctx, err)
	}
	return nil
}

// ToEventRepository ...
func (r *TxnRepository) ToEventRepository() types.EventTxnRepository {
	return NewEventTxnRepository(r.tx)
}
