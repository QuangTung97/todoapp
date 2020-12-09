package errors

import (
	liberrors "todoapp/lib/errors"
)

// ErrGeneralFailedPrecondition ...
type ErrGeneralFailedPrecondition liberrors.Error

// NewErrGeneralFailedPrecondition ...
func NewErrGeneralFailedPrecondition() *ErrGeneralFailedPrecondition {
	return &ErrGeneralFailedPrecondition{
		RPCStatus: 2,
		Code:      "02",
		Message:   "Failed precondition",
	}
}

// WithMin ...
func (e *ErrGeneralFailedPrecondition) WithMin(value int64) *ErrGeneralFailedPrecondition {
	err := (*liberrors.Error)(e)
	return (*ErrGeneralFailedPrecondition)(err.WithDetail("min", value))
}

// ErrGeneralNotFound ...
type ErrGeneralNotFound liberrors.Error

// NewErrGeneralNotFound ...
func NewErrGeneralNotFound() *ErrGeneralNotFound {
	return &ErrGeneralNotFound{
		RPCStatus: 1,
		Code:      "01",
		Message:   "Not found",
	}
}

// GeneralTag ...
type GeneralTag struct {
	FailedPrecondition *ErrGeneralFailedPrecondition
	NotFound           *ErrGeneralNotFound
}

// General ...
var General = &GeneralTag{
	FailedPrecondition: NewErrGeneralFailedPrecondition(),
	NotFound:           NewErrGeneralNotFound(),
}

// ErrTodoInvalidInputCreate ...
type ErrTodoInvalidInputCreate liberrors.Error

// NewErrTodoInvalidInputCreate ...
func NewErrTodoInvalidInputCreate() *ErrTodoInvalidInputCreate {
	return &ErrTodoInvalidInputCreate{
		RPCStatus: 4,
		Code:      "04",
		Message:   "Invalid create input",
	}
}

// WithMsg ...
func (e *ErrTodoInvalidInputCreate) WithMsg(value string) *ErrTodoInvalidInputCreate {
	err := (*liberrors.Error)(e)
	return (*ErrTodoInvalidInputCreate)(err.WithDetail("msg", value))
}

// ErrTodoNotFoundTodo ...
type ErrTodoNotFoundTodo liberrors.Error

// NewErrTodoNotFoundTodo ...
func NewErrTodoNotFoundTodo() *ErrTodoNotFoundTodo {
	return &ErrTodoNotFoundTodo{
		RPCStatus: 1,
		Code:      "02",
		Message:   "Not found todo",
	}
}

// TodoTag ...
type TodoTag struct {
	InvalidInputCreate *ErrTodoInvalidInputCreate
	NotFoundTodo       *ErrTodoNotFoundTodo
}

// Todo ...
var Todo = &TodoTag{
	InvalidInputCreate: NewErrTodoInvalidInputCreate(),
	NotFoundTodo:       NewErrTodoNotFoundTodo(),
}
