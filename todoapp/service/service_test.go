package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	todoapp_rpc "todoapp-rpc/rpc/todoapp/v1"
	"todoapp/pkg/errors"
	"todoapp/todoapp/model"
	"todoapp/todoapp/types"
)

func TestTodoSaveTx(t *testing.T) {
	type getTodoOutput struct {
		nullTodo model.NullTodo
		err      error
	}

	type getTodoItemsOutput struct {
		items []model.TodoItem
		err   error
	}

	type insertTodoOutput struct {
		todoID model.TodoID
		err    error
	}

	type addEventOutput struct {
		eventID model.EventID
		err     error
	}

	table := []struct {
		name  string
		input types.SaveTodoInput

		getTodoOutput      getTodoOutput
		getTodoItemsOutput getTodoItemsOutput
		insertTodoOutput   insertTodoOutput
		updateTodoOutput   error
		deleteItemsOutput  error
		updateItemOutputs  []error
		createItemOutputs  []error
		addEventOutput     addEventOutput

		expectedGetTodoInput      model.TodoID
		expectedGetTodoItemsInput model.TodoID
		expectedInsertInput       model.TodoSave
		expectedUpdateTodoInput   model.TodoSave
		expectedDeleteItemsInput  []model.TodoItemID
		expectedUpdateItemInputs  []model.TodoItemSave
		expectedCreateItemInputs  []model.TodoItemSave
		expectedAddEventInput     model.Event

		expectedErr error
		expectedID  model.TodoID
	}{
		{
			name: "not-found-todo",
			input: types.SaveTodoInput{
				ID: 11,
			},
			getTodoOutput: getTodoOutput{
				nullTodo: model.NullTodo{Valid: false},
			},
			expectedGetTodoInput: 11,
			expectedErr:          errors.Todo.NotFoundTodo.Err(),
		},
		{
			name: "get-todo-with-error",
			input: types.SaveTodoInput{
				ID: 11,
			},
			getTodoOutput: getTodoOutput{
				err: errors.General.InternalErrorAccessingDatabase.Err(),
			},
			expectedGetTodoInput: 11,
			expectedErr:          errors.General.InternalErrorAccessingDatabase.Err(),
		},
		{
			name: "get-todo-items-with-error",
			input: types.SaveTodoInput{
				ID: 11,
			},
			getTodoOutput: getTodoOutput{
				nullTodo: model.NullTodo{
					Valid: true,
					Todo: model.Todo{
						ID:   11,
						Name: "Test todo",
					},
				},
			},
			getTodoItemsOutput: getTodoItemsOutput{
				err: errors.General.InternalErrorAccessingDatabase.Err(),
			},
			expectedGetTodoInput:      11,
			expectedGetTodoItemsInput: 11,
			expectedErr:               errors.General.InternalErrorAccessingDatabase.Err(),
		},
		{
			name: "update-todo-error",
			input: types.SaveTodoInput{
				ID:   11,
				Name: "new todo",
			},
			getTodoOutput: getTodoOutput{
				nullTodo: model.NullTodo{
					Valid: true,
					Todo: model.Todo{
						ID:   11,
						Name: "Test todo",
					},
				},
			},
			getTodoItemsOutput: getTodoItemsOutput{
				err: nil,
			},
			updateTodoOutput: errors.General.InternalErrorAccessingDatabase.Err(),

			expectedGetTodoInput:      11,
			expectedGetTodoItemsInput: 11,
			expectedUpdateTodoInput: model.TodoSave{
				ID:   11,
				Name: "new todo",
			},
			expectedErr: errors.General.InternalErrorAccessingDatabase.Err(),
		},
		{
			name: "delete-items-error",
			input: types.SaveTodoInput{
				ID:   11,
				Name: "new todo",
			},
			getTodoOutput: getTodoOutput{
				nullTodo: model.NullTodo{
					Valid: true,
					Todo: model.Todo{
						ID:   11,
						Name: "test todo",
					},
				},
			},
			getTodoItemsOutput: getTodoItemsOutput{
				items: []model.TodoItem{
					{ID: 33},
					{ID: 44},
				},
				err: nil,
			},
			updateTodoOutput:  nil,
			deleteItemsOutput: errors.General.InternalErrorAccessingDatabase.Err(),

			expectedGetTodoInput:      11,
			expectedGetTodoItemsInput: 11,
			expectedUpdateTodoInput: model.TodoSave{
				ID:   11,
				Name: "new todo",
			},
			expectedDeleteItemsInput: []model.TodoItemID{33, 44},

			expectedErr: errors.General.InternalErrorAccessingDatabase.Err(),
		},
		{
			name: "add-event-error",
			input: types.SaveTodoInput{
				ID:   11,
				Name: "new todo",
			},
			getTodoOutput: getTodoOutput{
				nullTodo: model.NullTodo{
					Valid: true,
					Todo: model.Todo{
						ID:   11,
						Name: "test todo",
					},
				},
			},
			getTodoItemsOutput: getTodoItemsOutput{
				items: []model.TodoItem{
					{ID: 33},
					{ID: 44},
				},
				err: nil,
			},
			updateTodoOutput:  nil,
			deleteItemsOutput: nil,
			addEventOutput: addEventOutput{
				err: errors.General.InternalErrorAccessingDatabase.Err(),
			},

			expectedGetTodoInput:      11,
			expectedGetTodoItemsInput: 11,
			expectedUpdateTodoInput: model.TodoSave{
				ID:   11,
				Name: "new todo",
			},
			expectedDeleteItemsInput: []model.TodoItemID{33, 44},
			expectedAddEventInput: types.Event{
				ID:       0,
				Sequence: 0,
				Data: &todoapp_rpc.Event{
					Type: todoapp_rpc.EventType_EVENT_TYPE_TODO_SAVE,
					TodoSave: &todoapp_rpc.EventTodoSave{
						Id:   11,
						Name: "new todo",
					},
				},
			}.ToModel(),

			expectedErr: errors.General.InternalErrorAccessingDatabase.Err(),
		},
		{
			name: "ok-no-update-insert-item",
			input: types.SaveTodoInput{
				ID:   11,
				Name: "new todo",
			},
			getTodoOutput: getTodoOutput{
				nullTodo: model.NullTodo{
					Valid: true,
					Todo: model.Todo{
						ID:   11,
						Name: "test todo",
					},
				},
			},
			getTodoItemsOutput: getTodoItemsOutput{
				items: []model.TodoItem{
					{ID: 33},
					{ID: 44},
				},
				err: nil,
			},
			updateTodoOutput:  nil,
			deleteItemsOutput: nil,
			addEventOutput:    addEventOutput{},

			expectedGetTodoInput:      11,
			expectedGetTodoItemsInput: 11,
			expectedUpdateTodoInput: model.TodoSave{
				ID:   11,
				Name: "new todo",
			},
			expectedDeleteItemsInput: []model.TodoItemID{33, 44},
			expectedAddEventInput: types.Event{
				ID:       0,
				Sequence: 0,
				Data: &todoapp_rpc.Event{
					Type: todoapp_rpc.EventType_EVENT_TYPE_TODO_SAVE,
					TodoSave: &todoapp_rpc.EventTodoSave{
						Id:   11,
						Name: "new todo",
					},
				},
			}.ToModel(),

			expectedErr: nil,
			expectedID:  11,
		},
		{
			name: "compute-actions-error",
			input: types.SaveTodoInput{
				ID:   11,
				Name: "new todo",
				Items: []types.SaveTodoItem{
					{
						ID:   44,
						Name: "new item 1",
					},
					{
						ID:   55,
						Name: "new item 2",
					},
				},
			},
			getTodoOutput: getTodoOutput{
				nullTodo: model.NullTodo{
					Valid: true,
					Todo: model.Todo{
						ID:   11,
						Name: "test todo",
					},
				},
			},
			getTodoItemsOutput: getTodoItemsOutput{
				items: []model.TodoItem{
					{ID: 33},
					{ID: 44},
				},
				err: nil,
			},
			updateTodoOutput:  nil,
			deleteItemsOutput: nil,
			addEventOutput:    addEventOutput{},

			expectedGetTodoInput:      11,
			expectedGetTodoItemsInput: 11,
			expectedUpdateTodoInput: model.TodoSave{
				ID:   11,
				Name: "new todo",
			},

			expectedErr: errors.Todo.NotFoundTodoItem.Err(),
		},
		{
			name: "update-items-error",
			input: types.SaveTodoInput{
				ID:   11,
				Name: "new todo",
				Items: []types.SaveTodoItem{
					{
						ID:   44,
						Name: "new item 1",
					},
					{
						ID:   55,
						Name: "new item 2",
					},
				},
			},
			getTodoOutput: getTodoOutput{
				nullTodo: model.NullTodo{
					Valid: true,
					Todo: model.Todo{
						ID:   11,
						Name: "test todo",
					},
				},
			},
			getTodoItemsOutput: getTodoItemsOutput{
				items: []model.TodoItem{
					{ID: 33},
					{ID: 44},
					{ID: 55},
				},
				err: nil,
			},
			updateTodoOutput:  nil,
			deleteItemsOutput: nil,
			updateItemOutputs: []error{nil, errors.General.InternalErrorAccessingDatabase.Err()},
			addEventOutput:    addEventOutput{},

			expectedGetTodoInput:      11,
			expectedGetTodoItemsInput: 11,
			expectedDeleteItemsInput:  []model.TodoItemID{33},
			expectedUpdateItemInputs: []model.TodoItemSave{
				{
					ID:   44,
					Name: "new item 1",
				},
				{
					ID:   55,
					Name: "new item 2",
				},
			},
			expectedUpdateTodoInput: model.TodoSave{
				ID:   11,
				Name: "new todo",
			},

			expectedErr: errors.General.InternalErrorAccessingDatabase.Err(),
		},
		{
			name: "insert-items-error",
			input: types.SaveTodoInput{
				ID:   11,
				Name: "new todo",
				Items: []types.SaveTodoItem{
					{
						ID:   44,
						Name: "new item 1",
					},
					{
						ID:   55,
						Name: "new item 2",
					},
					{
						Name: "new item 3",
					},
				},
			},
			getTodoOutput: getTodoOutput{
				nullTodo: model.NullTodo{
					Valid: true,
					Todo: model.Todo{
						ID:   11,
						Name: "test todo",
					},
				},
			},
			getTodoItemsOutput: getTodoItemsOutput{
				items: []model.TodoItem{
					{ID: 33},
					{ID: 44},
					{ID: 55},
				},
				err: nil,
			},
			updateTodoOutput:  nil,
			deleteItemsOutput: nil,
			updateItemOutputs: []error{nil, nil},
			createItemOutputs: []error{errors.General.InternalErrorAccessingDatabase.Err()},
			addEventOutput:    addEventOutput{},

			expectedGetTodoInput:      11,
			expectedGetTodoItemsInput: 11,
			expectedDeleteItemsInput:  []model.TodoItemID{33},
			expectedUpdateItemInputs: []model.TodoItemSave{
				{
					ID:   44,
					Name: "new item 1",
				},
				{
					ID:   55,
					Name: "new item 2",
				},
			},
			expectedCreateItemInputs: []model.TodoItemSave{
				{
					TodoID: 11,
					Name:   "new item 3",
				},
			},
			expectedUpdateTodoInput: model.TodoSave{
				ID:   11,
				Name: "new todo",
			},

			expectedErr: errors.General.InternalErrorAccessingDatabase.Err(),
		},
		{
			name: "insert-items-ok",
			input: types.SaveTodoInput{
				ID:   11,
				Name: "new todo",
				Items: []types.SaveTodoItem{
					{
						ID:   44,
						Name: "new item 1",
					},
					{
						ID:   55,
						Name: "new item 2",
					},
					{
						Name: "new item 3",
					},
				},
			},
			getTodoOutput: getTodoOutput{
				nullTodo: model.NullTodo{
					Valid: true,
					Todo: model.Todo{
						ID:   11,
						Name: "test todo",
					},
				},
			},
			getTodoItemsOutput: getTodoItemsOutput{
				items: []model.TodoItem{
					{ID: 33},
					{ID: 44},
					{ID: 55},
				},
				err: nil,
			},
			updateTodoOutput:  nil,
			deleteItemsOutput: nil,
			updateItemOutputs: []error{nil, nil},
			createItemOutputs: []error{nil},
			addEventOutput:    addEventOutput{},

			expectedGetTodoInput:      11,
			expectedGetTodoItemsInput: 11,
			expectedDeleteItemsInput:  []model.TodoItemID{33},
			expectedUpdateItemInputs: []model.TodoItemSave{
				{
					ID:   44,
					Name: "new item 1",
				},
				{
					ID:   55,
					Name: "new item 2",
				},
			},
			expectedCreateItemInputs: []model.TodoItemSave{
				{
					TodoID: 11,
					Name:   "new item 3",
				},
			},
			expectedUpdateTodoInput: model.TodoSave{
				ID:   11,
				Name: "new todo",
			},
			expectedAddEventInput: types.Event{
				ID:       0,
				Sequence: 0,
				Data: &todoapp_rpc.Event{
					Type: todoapp_rpc.EventType_EVENT_TYPE_TODO_SAVE,
					TodoSave: &todoapp_rpc.EventTodoSave{
						Id:   11,
						Name: "new todo",
					},
				},
			}.ToModel(),

			expectedErr: nil,
			expectedID:  11,
		},
		{
			name: "insert-todo-ok",
			input: types.SaveTodoInput{
				ID:   0,
				Name: "new todo create",
				Items: []types.SaveTodoItem{
					{
						Name: "new item 4",
					},
					{
						Name: "new item 5",
					},
				},
			},

			insertTodoOutput: insertTodoOutput{
				todoID: 123,
				err:    nil,
			},
			createItemOutputs: []error{nil, nil},
			addEventOutput:    addEventOutput{},

			expectedInsertInput: model.TodoSave{
				Name: "new todo create",
			},
			expectedCreateItemInputs: []model.TodoItemSave{
				{
					TodoID: 123,
					Name:   "new item 4",
				},
				{
					TodoID: 123,
					Name:   "new item 5",
				},
			},
			expectedAddEventInput: types.Event{
				ID:       0,
				Sequence: 0,
				Data: &todoapp_rpc.Event{
					Type: todoapp_rpc.EventType_EVENT_TYPE_TODO_SAVE,
					TodoSave: &todoapp_rpc.EventTodoSave{
						Id:   123,
						Name: "new todo create",
					},
				},
			}.ToModel(),

			expectedErr: nil,
			expectedID:  123,
		},
		{
			name: "insert-todo-error",
			input: types.SaveTodoInput{
				ID:   0,
				Name: "new todo create",
				Items: []types.SaveTodoItem{
					{
						Name: "new item 4",
					},
					{
						Name: "new item 5",
					},
				},
			},

			insertTodoOutput: insertTodoOutput{
				err: errors.General.InternalErrorAccessingDatabase.Err(),
			},
			createItemOutputs: []error{nil, nil},
			addEventOutput:    addEventOutput{},

			expectedInsertInput: model.TodoSave{
				Name: "new todo create",
			},

			expectedErr: errors.General.InternalErrorAccessingDatabase.Err(),
		},
		{
			name: "insert-todo-insert-items-error",
			input: types.SaveTodoInput{
				ID:   0,
				Name: "new todo create",
				Items: []types.SaveTodoItem{
					{
						Name: "new item 4",
					},
					{
						Name: "new item 5",
					},
				},
			},

			insertTodoOutput: insertTodoOutput{
				todoID: 233,
				err:    nil,
			},
			createItemOutputs: []error{nil, errors.General.InternalErrorAccessingDatabase.Err()},
			addEventOutput:    addEventOutput{},

			expectedInsertInput: model.TodoSave{
				Name: "new todo create",
			},
			expectedCreateItemInputs: []model.TodoItemSave{
				{
					TodoID: 233,
					Name:   "new item 4",
				},
				{
					TodoID: 233,
					Name:   "new item 5",
				},
			},

			expectedErr: errors.General.InternalErrorAccessingDatabase.Err(),
		},
		{
			name: "insert-todo-add-event-error",
			input: types.SaveTodoInput{
				ID:   0,
				Name: "new todo create",
				Items: []types.SaveTodoItem{
					{
						Name: "new item 4",
					},
					{
						Name: "new item 5",
					},
				},
			},

			insertTodoOutput: insertTodoOutput{
				todoID: 233,
				err:    nil,
			},
			createItemOutputs: []error{nil, nil},
			addEventOutput: addEventOutput{
				err: errors.General.InternalErrorAccessingDatabase.Err(),
			},

			expectedInsertInput: model.TodoSave{
				Name: "new todo create",
			},
			expectedCreateItemInputs: []model.TodoItemSave{
				{
					TodoID: 233,
					Name:   "new item 4",
				},
				{
					TodoID: 233,
					Name:   "new item 5",
				},
			},
			expectedAddEventInput: types.Event{
				ID:       0,
				Sequence: 0,
				Data: &todoapp_rpc.Event{
					Type: todoapp_rpc.EventType_EVENT_TYPE_TODO_SAVE,
					TodoSave: &todoapp_rpc.EventTodoSave{
						Id:   233,
						Name: "new todo create",
					},
				},
			}.ToModel(),

			expectedErr: errors.General.InternalErrorAccessingDatabase.Err(),
		},
	}

	for _, e := range table {
		t.Run(e.name, func(t *testing.T) {
			var getTodoInput model.TodoID
			var getTodoItemsInput model.TodoID
			var insertTodoInput model.TodoSave
			var updateTodoInput model.TodoSave
			var deleteItemsInput []model.TodoItemID
			var updateItemInputs []model.TodoItemSave
			var createItemInputs []model.TodoItemSave
			var addEventInput model.Event

			id, err := saveTodoTx(
				context.Background(), e.input,
				func(ctx context.Context, id model.TodoID) (model.NullTodo, error) {
					getTodoInput = id
					return e.getTodoOutput.nullTodo, e.getTodoOutput.err
				},
				func(ctx context.Context, todoID model.TodoID) ([]model.TodoItem, error) {
					getTodoItemsInput = todoID
					return e.getTodoItemsOutput.items, e.getTodoItemsOutput.err
				},
				func(ctx context.Context, save model.TodoSave) (model.TodoID, error) {
					insertTodoInput = save
					return e.insertTodoOutput.todoID, e.insertTodoOutput.err
				},
				func(ctx context.Context, save model.TodoSave) error {
					updateTodoInput = save
					return e.updateTodoOutput
				},
				func(ctx context.Context, itemIDs []model.TodoItemID) error {
					deleteItemsInput = itemIDs
					return e.deleteItemsOutput
				},
				func(ctx context.Context, item model.TodoItemSave) error {
					index := len(updateItemInputs)
					updateItemInputs = append(updateItemInputs, item)
					return e.updateItemOutputs[index]
				},
				func(ctx context.Context, item model.TodoItemSave) (model.TodoItemID, error) {
					index := len(createItemInputs)
					createItemInputs = append(createItemInputs, item)
					return 0, e.createItemOutputs[index]
				},
				func(ctx context.Context, event model.Event) (model.EventID, error) {
					addEventInput = event
					return e.addEventOutput.eventID, e.addEventOutput.err
				},
			)

			assert.Equal(t, e.expectedErr, err)
			assert.Equal(t, e.expectedID, id)

			assert.Equal(t, e.expectedGetTodoInput, getTodoInput)
			assert.Equal(t, e.expectedGetTodoItemsInput, getTodoItemsInput)
			assert.Equal(t, e.expectedInsertInput, insertTodoInput)
			assert.Equal(t, e.expectedUpdateTodoInput, updateTodoInput)
			assert.Equal(t, e.expectedDeleteItemsInput, deleteItemsInput)
			assert.Equal(t, e.expectedUpdateItemInputs, updateItemInputs)
			assert.Equal(t, e.expectedCreateItemInputs, createItemInputs)
			assert.Equal(t, e.expectedAddEventInput, addEventInput)
		})
	}
}
