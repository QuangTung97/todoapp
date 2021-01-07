package server

import (
	todoapp_rpc "todoapp-rpc/rpc/todoapp/v1"
	"todoapp/todoapp/model"
	"todoapp/todoapp/types"
)

func transformTodoItems(items []*todoapp_rpc.TodoItem) []model.TodoItem {
	result := make([]model.TodoItem, 0, len(items))
	for _, i := range items {
		result = append(result, model.TodoItem{
			ID:   model.TodoItemID(i.Id),
			Name: i.Name,
		})
	}
	return result
}

func transformSaveRequest(req *todoapp_rpc.TodoSaveRequest) (types.SaveTodoInput, error) {
	return types.SaveTodoInput{
		ID:    model.TodoID(req.Id),
		Name:  req.Name,
		Items: transformTodoItems(req.Items),
	}, nil
}
