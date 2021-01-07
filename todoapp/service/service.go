package service

import (
	"context"
	todoapp_rpc "todoapp-rpc/rpc/todoapp/v1"
	"todoapp/pkg/errors"
	"todoapp/todoapp/model"
	"todoapp/todoapp/types"
	"todoapp/todoapp/util"
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

// BuildTodoSaveEvent ...
func BuildTodoSaveEvent(input types.SaveTodoInput) model.Event {
	return types.Event{
		Data: &todoapp_rpc.Event{
			Type: todoapp_rpc.EventType_EVENT_TYPE_TODO_SAVE,
			TodoSave: &todoapp_rpc.EventTodoSave{
				Id:   uint64(input.ID),
				Name: input.Name,
			},
		},
	}.ToModel()
}

func saveTodoTx(
	ctx context.Context, input types.SaveTodoInput,
	tx types.TxnRepository, eventTx types.EventTxnRepository,
) (model.TodoID, error) {
	if input.ID == 0 {
		id, err := tx.InsertTodo(ctx, model.Todo{
			Name: input.Name,
		})
		if err != nil {
			return 0, err
		}

		for _, item := range input.Items {
			item.TodoID = id
			_, err := tx.InsertTodoItem(ctx, item)
			if err != nil {
				return 0, err
			}
		}

		input.ID = id
		_, err = eventTx.InsertEvent(ctx, BuildTodoSaveEvent(input))
		if err != nil {
			return 0, err
		}

		return id, nil
	}

	nullTodo, err := tx.GetTodo(ctx, input.ID)
	if err != nil {
		return 0, err
	}
	if !nullTodo.Valid {
		return 0, errors.Todo.NotFoundTodo.Err()
	}

	items, err := tx.GetTodoItemsByTodoID(ctx, input.ID)
	if err != nil {
		return 0, err
	}

	actions, err := util.ComputeUpdateTodoActions(input.ID, items, input.Items)
	if err != nil {
		return 0, err
	}

	err = tx.UpdateTodo(ctx, model.Todo{
		ID:   input.ID,
		Name: input.Name,
	})
	if err != nil {
		return 0, err
	}

	err = tx.DeleteTodoItems(ctx, actions.DeletedItems)
	if err != nil {
		return 0, err
	}

	for _, item := range actions.UpdatedItems {
		err := tx.UpdateTodoITem(ctx, item)
		if err != nil {
			return 0, err
		}
	}

	for _, item := range actions.InsertedItems {
		_, err := tx.InsertTodoItem(ctx, item)
		if err != nil {
			return 0, err
		}
	}

	_, err = eventTx.InsertEvent(ctx, BuildTodoSaveEvent(input))
	if err != nil {
		return 0, err
	}

	return input.ID, nil
}

// SaveTodo ...
func (s *Service) SaveTodo(ctx context.Context, input types.SaveTodoInput) (model.TodoID, error) {
	todoID := input.ID
	err := s.repo.Transact(ctx, func(tx types.TxnRepository) error {
		id, err := saveTodoTx(
			ctx, input,
			tx, tx.ToEventRepository(),
		)
		if err != nil {
			return err
		}

		todoID = id
		return nil
	})
	if err != nil {
		return 0, err
	}

	s.client.Signal(ctx)

	return todoID, nil
}
