package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	todoapp_rpc "todoapp-rpc/rpc/todoapp/v1"
	"todoapp/pkg/errors"
	types_mocks "todoapp/todoapp/mocks"
	"todoapp/todoapp/model"
	"todoapp/todoapp/types"
)

func TestTodoSaveTx_Insert_OK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tx := types_mocks.NewMockTxnRepository(ctrl)
	eventTx := types_mocks.NewMockEventTxnRepository(ctrl)

	tx.EXPECT().InsertTodo(gomock.Any(), model.TodoSave{
		Name: "some insert todo",
	}).Return(model.TodoID(322), nil)

	gomock.InOrder(
		tx.EXPECT().InsertTodoItem(gomock.Any(),
			model.TodoItemSave{
				TodoID: 322,
				Name:   "item insert 1",
			},
		).Return(model.TodoItemID(20), nil),
		tx.EXPECT().InsertTodoItem(gomock.Any(),
			model.TodoItemSave{
				TodoID: 322,
				Name:   "item insert 2",
			},
		).Return(model.TodoItemID(20), nil),
	)

	eventTx.EXPECT().InsertEvent(gomock.Any(), types.Event{
		Data: &todoapp_rpc.Event{
			Type: todoapp_rpc.EventType_EVENT_TYPE_TODO_SAVE,
			TodoSave: &todoapp_rpc.EventTodoSave{
				Id: 322,
				Name: "some insert todo",
			},
		},
	}.ToModel()).Return(model.EventID(55), nil)

	id, err := saveTodoTx(
		context.Background(), types.SaveTodoInput{
			Name: "some insert todo",
			Items: []types.SaveTodoItem{
				{Name: "item insert 1"},
				{Name: "item insert 2"},
			},
		},
		tx, eventTx,
	)

	assert.Equal(t, nil, err)
	assert.Equal(t, model.TodoID(322), id)
}

func TestTodoSaveTx_Error(t *testing.T) {
	type testCase struct {
		name  string
		input types.SaveTodoInput

		expectCall func(e testCase, tx *types_mocks.MockTxnRepository, eventTx *types_mocks.MockEventTxnRepository)

		expectedErr error
		expectedID  model.TodoID
	}

	table := []testCase{
		{
			name: "not-found-todo",
			input: types.SaveTodoInput{
				ID: 11,
			},
			expectCall: func(e testCase, tx *types_mocks.MockTxnRepository, eventTx *types_mocks.MockEventTxnRepository) {
				tx.EXPECT().GetTodo(gomock.Any(), gomock.Any()).Return(model.NullTodo{}, nil)
			},
			expectedErr: errors.Todo.NotFoundTodo.Err(),
		},
		{
			name: "get-todo-with-error",
			input: types.SaveTodoInput{
				ID: 11,
			},
			expectCall: func(e testCase, tx *types_mocks.MockTxnRepository, eventTx *types_mocks.MockEventTxnRepository) {
				tx.EXPECT().GetTodo(gomock.Any(), gomock.Any()).
					Return(model.NullTodo{}, errors.General.InternalErrorAccessingDatabase.Err())
			},
			expectedErr: errors.General.InternalErrorAccessingDatabase.Err(),
		},
		{
			name: "get-todo-items-with-error",
			input: types.SaveTodoInput{
				ID: 11,
			},
			expectCall: func(e testCase, tx *types_mocks.MockTxnRepository, eventTx *types_mocks.MockEventTxnRepository) {
				tx.EXPECT().GetTodo(gomock.Any(), gomock.Any()).
					Return(model.NullTodo{
						Valid: true,
						Todo: model.Todo{
							ID:   11,
							Name: "Test todo",
						},
					}, nil)

				tx.EXPECT().GetTodoItemsByTodoID(gomock.Any(), gomock.Any()).
					Return(nil, errors.General.InternalErrorAccessingDatabase.Err())
			},
			expectedErr: errors.General.InternalErrorAccessingDatabase.Err(),
		},
		{
			name: "update-todo-error",
			input: types.SaveTodoInput{
				ID:   11,
				Name: "new todo",
			},
			expectCall: func(e testCase, tx *types_mocks.MockTxnRepository, eventTx *types_mocks.MockEventTxnRepository) {
				tx.EXPECT().GetTodo(gomock.Any(), gomock.Any()).
					Return(model.NullTodo{
						Valid: true,
						Todo: model.Todo{
							ID:   11,
							Name: "Test todo",
						},
					}, nil)

				tx.EXPECT().GetTodoItemsByTodoID(gomock.Any(), gomock.Any()).
					Return(nil, nil)

				tx.EXPECT().UpdateTodo(gomock.Any(), gomock.Any()).
					Return(errors.General.InternalErrorAccessingDatabase.Err())
			},
			expectedErr: errors.General.InternalErrorAccessingDatabase.Err(),
		},
		{
			name: "delete-items-error",
			input: types.SaveTodoInput{
				ID:   11,
				Name: "new todo",
			},
			expectCall: func(e testCase, tx *types_mocks.MockTxnRepository, eventTx *types_mocks.MockEventTxnRepository) {
				tx.EXPECT().GetTodo(gomock.Any(), gomock.Any()).
					Return(model.NullTodo{
						Valid: true,
						Todo: model.Todo{
							ID:   11,
							Name: "Test todo",
						},
					}, nil)

				tx.EXPECT().GetTodoItemsByTodoID(gomock.Any(), gomock.Any()).
					Return([]model.TodoItem{
						{ID: 33},
						{ID: 44},
					}, nil)

				tx.EXPECT().UpdateTodo(gomock.Any(), gomock.Any()).
					Return(nil)

				tx.EXPECT().DeleteTodoItems(gomock.Any(), gomock.Any()).
					Return(errors.General.InternalErrorAccessingDatabase.Err())
			},

			expectedErr: errors.General.InternalErrorAccessingDatabase.Err(),
		},
		{
			name: "add-event-error",
			input: types.SaveTodoInput{
				ID:   11,
				Name: "new todo",
			},
			expectCall: func(e testCase, tx *types_mocks.MockTxnRepository, eventTx *types_mocks.MockEventTxnRepository) {
				tx.EXPECT().GetTodo(gomock.Any(), gomock.Any()).
					Return(model.NullTodo{
						Valid: true,
						Todo: model.Todo{
							ID:   11,
							Name: "Test todo",
						},
					}, nil)

				tx.EXPECT().GetTodoItemsByTodoID(gomock.Any(), gomock.Any()).
					Return([]model.TodoItem{
						{ID: 33},
						{ID: 44},
					}, nil)

				tx.EXPECT().UpdateTodo(gomock.Any(), gomock.Any()).
					Return(nil)

				tx.EXPECT().DeleteTodoItems(gomock.Any(), gomock.Any()).
					Return(nil)

				eventTx.EXPECT().InsertEvent(gomock.Any(), gomock.Any()).
					Return(model.EventID(0), errors.General.InternalErrorAccessingDatabase.Err())
			},
			expectedErr: errors.General.InternalErrorAccessingDatabase.Err(),
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
			expectCall: func(e testCase, tx *types_mocks.MockTxnRepository, eventTx *types_mocks.MockEventTxnRepository) {
				tx.EXPECT().GetTodo(gomock.Any(), gomock.Any()).
					Return(model.NullTodo{
						Valid: true,
						Todo: model.Todo{
							ID:   11,
							Name: "Test todo",
						},
					}, nil)

				tx.EXPECT().GetTodoItemsByTodoID(gomock.Any(), gomock.Any()).
					Return([]model.TodoItem{
						{ID: 33},
						{ID: 44},
					}, nil)
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
			expectCall: func(e testCase, tx *types_mocks.MockTxnRepository, eventTx *types_mocks.MockEventTxnRepository) {
				tx.EXPECT().GetTodo(gomock.Any(), gomock.Any()).
					Return(model.NullTodo{
						Valid: true,
						Todo: model.Todo{
							ID:   11,
							Name: "Test todo",
						},
					}, nil)

				tx.EXPECT().GetTodoItemsByTodoID(gomock.Any(), gomock.Any()).
					Return([]model.TodoItem{
						{ID: 33},
						{ID: 44},
						{ID: 55},
					}, nil)

				tx.EXPECT().UpdateTodo(gomock.Any(), gomock.Any()).
					Return(nil)

				tx.EXPECT().DeleteTodoItems(gomock.Any(), gomock.Any()).
					Return(nil)

				tx.EXPECT().UpdateTodoITem(gomock.Any(), gomock.Any()).Return(errors.General.InternalErrorAccessingDatabase.Err())
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
			expectCall: func(e testCase, tx *types_mocks.MockTxnRepository, eventTx *types_mocks.MockEventTxnRepository) {
				tx.EXPECT().GetTodo(gomock.Any(), gomock.Any()).
					Return(model.NullTodo{
						Valid: true,
						Todo: model.Todo{
							ID:   11,
							Name: "Test todo",
						},
					}, nil)

				tx.EXPECT().GetTodoItemsByTodoID(gomock.Any(), gomock.Any()).
					Return([]model.TodoItem{
						{ID: 33},
						{ID: 44},
						{ID: 55},
					}, nil)

				tx.EXPECT().UpdateTodo(gomock.Any(), gomock.Any()).
					Return(nil)

				tx.EXPECT().DeleteTodoItems(gomock.Any(), gomock.Any()).
					Return(nil)

				tx.EXPECT().UpdateTodoITem(gomock.Any(), gomock.Any()).Return(nil).Times(2)

				tx.EXPECT().InsertTodoItem(gomock.Any(), gomock.Any()).
					Return(model.TodoItemID(0), errors.General.InternalErrorAccessingDatabase.Err())
			},
			expectedErr: errors.General.InternalErrorAccessingDatabase.Err(),
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
			expectCall: func(e testCase, tx *types_mocks.MockTxnRepository, eventTx *types_mocks.MockEventTxnRepository) {
				tx.EXPECT().InsertTodo(gomock.Any(), gomock.Any()).
					Return(model.TodoID(0), errors.General.InternalErrorAccessingDatabase.Err())
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
			expectCall: func(e testCase, tx *types_mocks.MockTxnRepository, eventTx *types_mocks.MockEventTxnRepository) {
				tx.EXPECT().InsertTodo(gomock.Any(), gomock.Any()).
					Return(model.TodoID(0), nil)

				tx.EXPECT().InsertTodoItem(gomock.Any(), gomock.Any()).
					Return(model.TodoItemID(0), errors.General.InternalErrorAccessingDatabase.Err())
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
			expectCall: func(e testCase, tx *types_mocks.MockTxnRepository, eventTx *types_mocks.MockEventTxnRepository) {
				tx.EXPECT().InsertTodo(gomock.Any(), gomock.Any()).
					Return(model.TodoID(0), nil)

				tx.EXPECT().InsertTodoItem(gomock.Any(), gomock.Any()).
					Return(model.TodoItemID(0), nil).Times(2)

				eventTx.EXPECT().InsertEvent(gomock.Any(), gomock.Any()).
					Return(model.EventID(0), errors.General.InternalErrorAccessingDatabase.Err())
			},
			expectedErr: errors.General.InternalErrorAccessingDatabase.Err(),
		},
	}

	for _, e := range table {
		t.Run(e.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			tx := types_mocks.NewMockTxnRepository(ctrl)
			eventTx := types_mocks.NewMockEventTxnRepository(ctrl)

			e.expectCall(e, tx, eventTx)

			id, err := saveTodoTx(
				context.Background(), e.input,
				tx, eventTx,
			)

			assert.Equal(t, e.expectedErr, err)
			assert.Equal(t, e.expectedID, id)
		})
	}
}
