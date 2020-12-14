package types

import "todoapp/todoapp/model"

type (
	// SaveTodoItem for todo items
	SaveTodoItem struct {
		ID   model.TodoItemID
		Name string
	}

	// SaveTodoInput ...
	SaveTodoInput struct {
		ID    model.TodoID
		Name  string
		Items []SaveTodoItem
	}
)
