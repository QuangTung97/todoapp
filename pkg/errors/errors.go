// Code generated by bin/errors. DO NOT EDIT.
package errors

import (
	liberrors "todoapp/lib/errors"
)

import (
	"time"
)

// ErrGeneralUnknown ...
type ErrGeneralUnknown liberrors.Error

// NewErrGeneralUnknown ...
func NewErrGeneralUnknown() *ErrGeneralUnknown {
	return &ErrGeneralUnknown{
		RPCStatus: 2,
		Code:      "0200",
		Message:   "Unknown",
	}
}

// ErrGeneralUnknownValue ...
type ErrGeneralUnknownValue liberrors.Error

// NewErrGeneralUnknownValue ...
func NewErrGeneralUnknownValue() *ErrGeneralUnknownValue {
	return &ErrGeneralUnknownValue{
		RPCStatus: 2,
		Code:      "0201",
		Message:   "Unknown value",
	}
}

// WithMin ...
func (e *ErrGeneralUnknownValue) WithMin(value int64) *ErrGeneralUnknownValue {
	err := (*liberrors.Error)(e)
	return (*ErrGeneralUnknownValue)(err.WithDetail("min", value))
}

// GeneralTag ...
type GeneralTag struct {
	Unknown      *ErrGeneralUnknown
	UnknownValue *ErrGeneralUnknownValue
}

// General ...
var General = &GeneralTag{
	Unknown:      NewErrGeneralUnknown(),
	UnknownValue: NewErrGeneralUnknownValue(),
}

// ErrTodoDeadlineExceededCommand ...
type ErrTodoDeadlineExceededCommand liberrors.Error

// NewErrTodoDeadlineExceededCommand ...
func NewErrTodoDeadlineExceededCommand() *ErrTodoDeadlineExceededCommand {
	return &ErrTodoDeadlineExceededCommand{
		RPCStatus: 4,
		Code:      "04",
		Message:   "Deadline exceeded",
	}
}

// WithMsg ...
func (e *ErrTodoDeadlineExceededCommand) WithMsg(value string) *ErrTodoDeadlineExceededCommand {
	err := (*liberrors.Error)(e)
	return (*ErrTodoDeadlineExceededCommand)(err.WithDetail("msg", value))
}

// WithStartedAt ...
func (e *ErrTodoDeadlineExceededCommand) WithStartedAt(value time.Time) *ErrTodoDeadlineExceededCommand {
	err := (*liberrors.Error)(e)
	return (*ErrTodoDeadlineExceededCommand)(err.WithDetail("startedAt", value))
}

// ErrTodoInvalidArgument ...
type ErrTodoInvalidArgument liberrors.Error

// NewErrTodoInvalidArgument ...
func NewErrTodoInvalidArgument() *ErrTodoInvalidArgument {
	return &ErrTodoInvalidArgument{
		RPCStatus: 3,
		Code:      "03",
		Message:   "Not found todo",
	}
}

// TodoTag ...
type TodoTag struct {
	DeadlineExceededCommand *ErrTodoDeadlineExceededCommand
	InvalidArgument         *ErrTodoInvalidArgument
}

// Todo ...
var Todo = &TodoTag{
	DeadlineExceededCommand: NewErrTodoDeadlineExceededCommand(),
	InvalidArgument:         NewErrTodoInvalidArgument(),
}
