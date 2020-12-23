package service

import (
	"context"
	todoapp_rpc "todoapp-rpc/rpc/todoapp/v1"
	"todoapp/pkg/errors"
	"todoapp/todoapp/model"
	"todoapp/todoapp/types"
	"todoapp/todoapp/util"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
)

// Service ...
type Service struct {
	repo   types.Repository
	client types.EventClient
}

var _ types.Service = &Service{}

// NewService creates a service
func NewService(repo types.Repository, client types.EventClient) *Service {
	return &Service{
		repo:   repo,
		client: client,
	}
}

func insertTodo(ctx context.Context, tx types.TxnRepository, input types.SaveTodoInput,
) (model.TodoID, error) {
	if len(input.Items) == 0 {
		return 0, errors.Todo.InvalidArgumentEmptyItems.Err()
	}

	id, err := tx.InsertTodo(ctx, model.TodoSave{
		Name: input.Name,
	})
	if err != nil {
		return 0, err
	}

	for _, item := range input.Items {
		_, err := tx.InsertTodoItem(ctx, model.TodoItemSave{
			TodoID: id,
			Name:   item.Name,
		})
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}

func addEvent(ctx context.Context, tx types.TxnRepository, todoID model.TodoID, name string) error {
	event := types.Event{
		Data: &todoapp_rpc.Event{
			Type: todoapp_rpc.EventType_EVENT_TYPE_TODO_SAVE,
			TodoSave: &todoapp_rpc.EventTodoSave{
				Id:   uint64(todoID),
				Name: name,
			},
		},
	}

	eventTx := tx.ToEventRepository()
	_, err := eventTx.InsertEvent(ctx, event.ToModel())
	return err
}

// SaveTodo ...
func (s *Service) SaveTodo(ctx context.Context, input types.SaveTodoInput) (model.TodoID, error) {
	todoID := input.ID
	err := s.repo.Transact(ctx, func(tx types.TxnRepository) error {
		if input.ID == 0 {
			id, err := insertTodo(ctx, tx, input)
			if err != nil {
				return err
			}
			todoID = id
			return addEvent(ctx, tx, todoID, input.Name)
		}

		nullTodo, err := tx.GetTodo(ctx, input.ID)
		if err != nil {
			return err
		}
		if !nullTodo.Valid {
			return errors.Todo.NotFoundTodo.Err()
		}

		items, err := tx.GetTodoItemsByTodoID(ctx, input.ID)
		if err != nil {
			return err
		}

		actions, err := util.ComputeUpdateTodoActions(input.ID, items, input.Items)
		if err != nil {
			return err
		}

		err = tx.UpdateTodo(ctx, model.TodoSave{
			ID:   input.ID,
			Name: input.Name,
		})
		if err != nil {
			return err
		}

		ctxzap.Extract(ctx).Info("DeletedIDs", zap.Any("todoItemIDs", actions.DeletedItems))

		err = tx.DeleteTodoItems(ctx, actions.DeletedItems)
		if err != nil {
			return err
		}

		for _, item := range actions.UpdatedItems {
			err := tx.UpdateTodoITem(ctx, item)
			if err != nil {
				return err
			}
		}

		for _, item := range actions.InsertedItems {
			_, err := tx.InsertTodoItem(ctx, item)
			if err != nil {
				return err
			}
		}

		return addEvent(ctx, tx, todoID, input.Name)
	})
	if err != nil {
		return 0, err
	}

	s.client.Signal(ctx)

	return todoID, nil
}
