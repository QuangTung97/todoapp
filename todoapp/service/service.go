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

func saveTodoTx(
	ctx context.Context,
	input types.SaveTodoInput,

	getTodo func(ctx context.Context, id model.TodoID) (model.NullTodo, error),
	getTodoItems func(ctx context.Context, todoID model.TodoID) ([]model.TodoItem, error),

	insertTodo func(ctx context.Context, save model.TodoSave) (model.TodoID, error),
	updateTodo func(ctx context.Context, save model.TodoSave) error,

	deleteItems func(ctx context.Context, itemIDs []model.TodoItemID) error,
	updateItem func(ctx context.Context, item model.TodoItemSave) error,
	insertItem func(ctx context.Context, item model.TodoItemSave) (model.TodoItemID, error),

	insertEvent func(ctx context.Context, event model.Event) (model.EventID, error),
) (model.TodoID, error) {
	if input.ID == 0 {
		id, err := insertTodo(ctx, model.TodoSave{
			Name: input.Name,
		})
		if err != nil {
			return 0, err
		}

		for _, item := range input.Items {
			_, err := insertItem(ctx, model.TodoItemSave{
				TodoID: id,
				Name:   item.Name,
			})
			if err != nil {
				return 0, err
			}
		}

		event := types.Event{
			Data: &todoapp_rpc.Event{
				Type: todoapp_rpc.EventType_EVENT_TYPE_TODO_SAVE,
				TodoSave: &todoapp_rpc.EventTodoSave{
					Id:   uint64(id),
					Name: input.Name,
				},
			},
		}

		_, err = insertEvent(ctx, event.ToModel())
		if err != nil {
			return 0, err
		}

		return id, nil
	}

	nullTodo, err := getTodo(ctx, input.ID)
	if err != nil {
		return 0, err
	}
	if !nullTodo.Valid {
		return 0, errors.Todo.NotFoundTodo.Err()
	}

	items, err := getTodoItems(ctx, input.ID)
	if err != nil {
		return 0, err
	}

	err = updateTodo(ctx, model.TodoSave{
		ID:   input.ID,
		Name: input.Name,
	})
	if err != nil {
		return 0, err
	}

	actions, err := util.ComputeUpdateTodoActions(input.ID, items, input.Items)
	if err != nil {
		return 0, err
	}

	err = deleteItems(ctx, actions.DeletedItems)
	if err != nil {
		return 0, err
	}

	for _, item := range actions.UpdatedItems {
		err := updateItem(ctx, item)
		if err != nil {
			return 0, err
		}
	}

	for _, item := range actions.InsertedItems {
		_, err := insertItem(ctx, item)
		if err != nil {
			return 0, err
		}
	}

	event := types.Event{
		Data: &todoapp_rpc.Event{
			Type: todoapp_rpc.EventType_EVENT_TYPE_TODO_SAVE,
			TodoSave: &todoapp_rpc.EventTodoSave{
				Id:   uint64(input.ID),
				Name: input.Name,
			},
		},
	}

	_, err = insertEvent(ctx, event.ToModel())
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
			tx.GetTodo, tx.GetTodoItemsByTodoID,
			tx.InsertTodo, tx.UpdateTodo,
			tx.DeleteTodoItems, tx.UpdateTodoITem, tx.InsertTodoItem,
			tx.ToEventRepository().InsertEvent,
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
