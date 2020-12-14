package model

// TodoID ...
type TodoID uint32

// TodoSave for saving todo
type TodoSave struct {
	ID   TodoID `db:"id"`
	Name string `db:"name"`
}

// TodoItemID ...
type TodoItemID uint32

// TodoItemInsert for inserting todo item
type TodoItemInsert struct {
	TodoID TodoID `db:"todo_id"`
	Name   string `db:"name"`
}

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
