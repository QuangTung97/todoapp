package util

import (
	"todoapp/pkg/errors"
	"todoapp/todoapp/model"
	"todoapp/todoapp/types"
)

// UpdateTodoActions ...
type UpdateTodoActions struct {
	DeletedItems  []model.TodoItemID
	UpdatedItems  []model.TodoItemSave
	InsertedItems []model.TodoItemSave
}

// ComputeUpdateTodoActions ...
func ComputeUpdateTodoActions(
	todoID model.TodoID,
	dbItems []model.TodoItem,
	inputItems []types.SaveTodoItem,
) (UpdateTodoActions, error) {
	inputSet := make(map[model.TodoItemID]struct{})

	var inserted []model.TodoItemSave
	for _, item := range inputItems {
		if item.ID == 0 {
			inserted = append(inserted, model.TodoItemSave{
				TodoID: todoID,
				Name:   item.Name,
			})
		} else {
			inputSet[item.ID] = struct{}{}
		}
	}

	dbSet := make(map[model.TodoItemID]struct{})
	for _, item := range dbItems {
		dbSet[item.ID] = struct{}{}
	}

	var deleted []model.TodoItemID
	for _, item := range dbItems {
		_, existed := inputSet[item.ID]
		if !existed {
			deleted = append(deleted, item.ID)
		}
	}

	var updated []model.TodoItemSave
	for _, item := range inputItems {
		if item.ID == 0 {
			continue
		}

		_, existed := dbSet[item.ID]
		if !existed {
			return UpdateTodoActions{}, errors.Todo.NotFoundTodoItem.Err()
		}
		updated = append(updated, model.TodoItemSave{
			ID:   item.ID,
			Name: item.Name,
		})
	}

	return UpdateTodoActions{
		DeletedItems:  deleted,
		UpdatedItems:  updated,
		InsertedItems: inserted,
	}, nil
}
