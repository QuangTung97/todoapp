package model

//=====================
// ID Definitions
//=====================

// TodoID ...
type TodoID uint32

// TodoItemID ...
type TodoItemID uint32

//=====================
// Save Models
//=====================

// TodoSave for saving todo
type TodoSave struct {
	ID   TodoID `db:"id"`
	Name string `db:"name"`
}

// TodoItemSave for saving todo item
type TodoItemSave struct {
	ID     TodoItemID `db:"id"`
	TodoID TodoID     `db:"todo_id"`
	Name   string     `db:"name"`
}

//=====================
// Get Models
//=====================

// Todo ...
type Todo struct {
	ID   TodoID `db:"id"`
	Name string `db:"name"`
}

// NullTodo ...
type NullTodo struct {
	Valid bool
	Todo  Todo
}

// TodoItem ...
type TodoItem struct {
	ID     TodoItemID `db:"id"`
	TodoID TodoID     `db:"todo_id"`
	Name   string     `db:"name"`
}

// NullTodoItem ...
type NullTodoItem struct {
	Valid bool
	Item  TodoItem
}
