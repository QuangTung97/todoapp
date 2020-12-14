//go:generate mockgen -destination=../mocks/repository.go -package=types_mocks . Repository
//go:generate mockgen -destination=../mocks/txn_repository.go -package=types_mocks . TxnRepository

package types

import (
	"context"
	"todoapp/todoapp/model"
)

type (
	// Service ...
	Service interface {
		SaveTodo(ctx context.Context, input SaveTodoInput) (model.TodoID, error)
	}

	// Repository ...
	Repository interface {
		Transact(ctx context.Context, fn func(tx TxnRepository) error) error
	}

	// TxnRepository ...
	TxnRepository interface {
		GetTodo(ctx context.Context, id model.TodoID) (model.NullTodo, error)
	}
)
