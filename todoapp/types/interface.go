//go:generate mockgen -destination=../mocks/repository.go -package=types_mocks . Repository
//go:generate mockgen -destination=../mocks/txn_repository.go -package=types_mocks . TxnRepository
//go:generate mockgen -destination=../mocks/event_txn_repository.go -package=types_mocks . EventTxnRepository
//go:generate mockgen -destination=../mocks/event_client.go -package=types_mocks . EventClient

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
		// For Todos
		GetTodo(ctx context.Context, id model.TodoID) (model.NullTodo, error)
		InsertTodo(ctx context.Context, save model.TodoSave) (model.TodoID, error)
		UpdateTodo(ctx context.Context, save model.TodoSave) error

		// For Todo Items
		GetTodoItemsByTodoID(ctx context.Context, todoID model.TodoID) ([]model.TodoItem, error)
		DeleteTodoItems(ctx context.Context, todoItemIDs []model.TodoItemID) error
		InsertTodoItem(ctx context.Context, save model.TodoItemSave) (model.TodoItemID, error)
		UpdateTodoITem(ctx context.Context, save model.TodoItemSave) error

		ToEventRepository() EventTxnRepository
	}

	// EventTxnRepository ...
	EventTxnRepository interface {
		InsertEvent(ctx context.Context, event model.Event) (model.EventID, error)
	}

	// EventClient ...
	EventClient interface {
		Signal(ctx context.Context)
	}
)
