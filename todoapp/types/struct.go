package types

import "todoapp/todoapp/model"

type (
	// SaveTodoInput ...
	SaveTodoInput struct {
		ID    model.TodoID
		Name  string
		Items []model.TodoItem
	}
)
