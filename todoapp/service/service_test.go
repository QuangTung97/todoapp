package service

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"todoapp/pkg/errors"
	types_mocks "todoapp/todoapp/mocks"
	"todoapp/todoapp/model"
	"todoapp/todoapp/types"
)

func GetTodoHelper(
	tx *types_mocks.MockTxnRepository,
	id model.TodoID,
	nullTodo model.NullTodo, err error,
) *gomock.Call {
	return tx.EXPECT().GetTodo(gomock.Any(), id).Return(nullTodo, err)
}

func GetTodoItemsHelper(
	tx *types_mocks.MockTxnRepository,
	todoID model.TodoID,
	items []model.TodoItem, err error,
) *gomock.Call {
	return tx.EXPECT().GetTodoItemsByTodoID(gomock.Any(), todoID).Return(items, err)
}

func UpdateTodoHelper(
	tx *types_mocks.MockTxnRepository,
	save model.Todo, err error,
) *gomock.Call {
	return tx.EXPECT().UpdateTodo(gomock.Any(), save).Return(err)
}

func DeleteItemsHelper(
	tx *types_mocks.MockTxnRepository,
	itemIDs []model.TodoItemID, err error,
) *gomock.Call {
	return tx.EXPECT().DeleteTodoItems(gomock.Any(), itemIDs).Return(err)
}

func UpdateItemHelper(
	tx *types_mocks.MockTxnRepository,
	item model.TodoItem, err error,
) *gomock.Call {
	return tx.EXPECT().UpdateTodoITem(gomock.Any(), item).Return(err)
}

func InsertItemHelper(
	tx *types_mocks.MockTxnRepository,
	item model.TodoItem, id model.TodoItemID, err error,
) *gomock.Call {
	return tx.EXPECT().InsertTodoItem(gomock.Any(), item).Return(id, err)
}

func InsertTodoHelper(
	tx *types_mocks.MockTxnRepository,
	todo model.Todo, id model.TodoID, err error,
) *gomock.Call {
	return tx.EXPECT().InsertTodo(gomock.Any(), todo).Return(id, err)
}

func InsertEventHelper(
	tx *types_mocks.MockEventTxnRepository,
	event model.Event, id model.EventID, err error,
) *gomock.Call {
	return tx.EXPECT().InsertEvent(gomock.Any(), event).Return(id, err)
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
				GetTodoHelper(tx, 11, model.NullTodo{}, nil)
			},
			expectedErr: errors.Todo.NotFoundTodo.Err(),
		},
		{
			name: "get-todo-with-error",
			input: types.SaveTodoInput{
				ID: 11,
			},
			expectCall: func(e testCase, tx *types_mocks.MockTxnRepository, eventTx *types_mocks.MockEventTxnRepository) {
				GetTodoHelper(tx, 11, model.NullTodo{}, errors.General.InternalErrorAccessingDatabase.Err())
			},
			expectedErr: errors.General.InternalErrorAccessingDatabase.Err(),
		},
		{
			name: "get-todo-items-with-error",
			input: types.SaveTodoInput{
				ID: 11,
			},
			expectCall: func(e testCase, tx *types_mocks.MockTxnRepository, eventTx *types_mocks.MockEventTxnRepository) {
				nullTodo := model.NullTodo{
					Valid: true,
					Todo: model.Todo{
						ID:   11,
						Name: "Test todo",
					},
				}
				GetTodoHelper(tx, 11, nullTodo, nil)
				GetTodoItemsHelper(tx, 11, nil, errors.General.InternalErrorAccessingDatabase.Err())
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
				nullTodo := model.NullTodo{
					Valid: true,
					Todo: model.Todo{
						ID:   11,
						Name: "Test todo",
					},
				}
				GetTodoHelper(tx, 11, nullTodo, nil)
				GetTodoItemsHelper(tx, 11, nil, nil)
				UpdateTodoHelper(tx, model.Todo{
					ID:   11,
					Name: "new todo",
				}, errors.General.InternalErrorAccessingDatabase.Err())
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
				nullTodo := model.NullTodo{
					Valid: true,
					Todo: model.Todo{
						ID:   11,
						Name: "Test todo",
					},
				}
				GetTodoHelper(tx, 11, nullTodo, nil)
				GetTodoItemsHelper(tx, 11, []model.TodoItem{
					{ID: 33},
					{ID: 44},
				}, nil)
				UpdateTodoHelper(tx, model.Todo{
					ID:   11,
					Name: "new todo",
				}, nil)

				DeleteItemsHelper(tx, []model.TodoItemID{33, 44}, errors.General.InternalErrorAccessingDatabase.Err())
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
				nullTodo := model.NullTodo{
					Valid: true,
					Todo: model.Todo{
						ID:   11,
						Name: "Test todo",
					},
				}
				GetTodoHelper(tx, 11, nullTodo, nil)
				GetTodoItemsHelper(tx, 11, []model.TodoItem{
					{ID: 33},
					{ID: 44},
				}, nil)
				UpdateTodoHelper(tx, model.Todo{
					ID:   11,
					Name: "new todo",
				}, nil)

				DeleteItemsHelper(tx, []model.TodoItemID{33, 44}, nil)

				InsertEventHelper(eventTx, BuildTodoSaveEvent(e.input), 21,
					errors.General.InternalErrorAccessingDatabase.Err())
			},
			expectedErr: errors.General.InternalErrorAccessingDatabase.Err(),
		},
		{
			name: "compute-actions-error",
			input: types.SaveTodoInput{
				ID:   11,
				Name: "new todo",
				Items: []model.TodoItem{
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
				nullTodo := model.NullTodo{
					Valid: true,
					Todo: model.Todo{
						ID:   11,
						Name: "Test todo",
					},
				}
				GetTodoHelper(tx, 11, nullTodo, nil)
				GetTodoItemsHelper(tx, 11, []model.TodoItem{
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
				Items: []model.TodoItem{
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
				nullTodo := model.NullTodo{
					Valid: true,
					Todo: model.Todo{
						ID:   11,
						Name: "Test todo",
					},
				}
				GetTodoHelper(tx, 11, nullTodo, nil)
				GetTodoItemsHelper(tx, 11, []model.TodoItem{
					{ID: 33},
					{ID: 44},
					{ID: 55},
				}, nil)
				UpdateTodoHelper(tx, model.Todo{
					ID:   11,
					Name: "new todo",
				}, nil)

				DeleteItemsHelper(tx, []model.TodoItemID{33}, nil)

				UpdateItemHelper(tx, model.TodoItem{
					ID:   44,
					Name: "new item 1",
				}, errors.General.InternalErrorAccessingDatabase.Err())
			},
			expectedErr: errors.General.InternalErrorAccessingDatabase.Err(),
		},
		{
			name: "insert-items-error",
			input: types.SaveTodoInput{
				ID:   11,
				Name: "new todo",
				Items: []model.TodoItem{
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
				nullTodo := model.NullTodo{
					Valid: true,
					Todo: model.Todo{
						ID:   11,
						Name: "Test todo",
					},
				}
				GetTodoHelper(tx, 11, nullTodo, nil)
				GetTodoItemsHelper(tx, 11, []model.TodoItem{
					{ID: 33},
					{ID: 44},
					{ID: 55},
				}, nil)
				UpdateTodoHelper(tx, model.Todo{
					ID:   11,
					Name: "new todo",
				}, nil)

				DeleteItemsHelper(tx, []model.TodoItemID{33}, nil)

				gomock.InOrder(
					UpdateItemHelper(tx, model.TodoItem{
						ID:   44,
						Name: "new item 1",
					}, nil),
					UpdateItemHelper(tx, model.TodoItem{
						ID:   55,
						Name: "new item 2",
					}, nil),
				)
				InsertItemHelper(tx, model.TodoItem{
					TodoID: 11,
					Name:   "new item 3",
				}, 0, errors.General.InternalErrorAccessingDatabase.Err())
			},
			expectedErr: errors.General.InternalErrorAccessingDatabase.Err(),
		},
		{
			name: "update-ok",
			input: types.SaveTodoInput{
				ID:   11,
				Name: "new todo",
				Items: []model.TodoItem{
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
				nullTodo := model.NullTodo{
					Valid: true,
					Todo: model.Todo{
						ID:   11,
						Name: "Test todo",
					},
				}
				GetTodoHelper(tx, 11, nullTodo, nil)
				GetTodoItemsHelper(tx, 11, []model.TodoItem{
					{ID: 33},
					{ID: 44},
					{ID: 55},
				}, nil)
				UpdateTodoHelper(tx, model.Todo{
					ID:   11,
					Name: "new todo",
				}, nil)

				DeleteItemsHelper(tx, []model.TodoItemID{33}, nil)

				gomock.InOrder(
					UpdateItemHelper(tx, model.TodoItem{
						ID:   44,
						Name: "new item 1",
					}, nil),
					UpdateItemHelper(tx, model.TodoItem{
						ID:   55,
						Name: "new item 2",
					}, nil),
				)
				InsertItemHelper(tx, model.TodoItem{
					TodoID: 11,
					Name:   "new item 3",
				}, 0, nil)

				InsertEventHelper(eventTx, BuildTodoSaveEvent(e.input), 31, nil)
			},
			expectedErr: nil,
			expectedID:  11,
		},
		{
			name: "insert-todo-error",
			input: types.SaveTodoInput{
				ID:   0,
				Name: "new todo create",
				Items: []model.TodoItem{
					{
						Name: "new item 4",
					},
					{
						Name: "new item 5",
					},
				},
			},
			expectCall: func(e testCase, tx *types_mocks.MockTxnRepository, eventTx *types_mocks.MockEventTxnRepository) {
				InsertTodoHelper(tx, model.Todo{
					Name: "new todo create",
				}, 0, errors.General.InternalErrorAccessingDatabase.Err())
			},
			expectedErr: errors.General.InternalErrorAccessingDatabase.Err(),
		},
		{
			name: "insert-todo-insert-items-error",
			input: types.SaveTodoInput{
				ID:   0,
				Name: "new todo create",
				Items: []model.TodoItem{
					{
						Name: "new item 4",
					},
					{
						Name: "new item 5",
					},
				},
			},
			expectCall: func(e testCase, tx *types_mocks.MockTxnRepository, eventTx *types_mocks.MockEventTxnRepository) {
				InsertTodoHelper(tx, model.Todo{
					Name: "new todo create",
				}, 55, nil)

				InsertItemHelper(tx, model.TodoItem{
					TodoID: 55,
					Name:   "new item 4",
				}, 1, errors.General.InternalErrorAccessingDatabase.Err())
			},
			expectedErr: errors.General.InternalErrorAccessingDatabase.Err(),
		},
		{
			name: "insert-todo-add-event-error",
			input: types.SaveTodoInput{
				ID:   0,
				Name: "new todo create",
				Items: []model.TodoItem{
					{
						Name: "new item 4",
					},
					{
						Name: "new item 5",
					},
				},
			},
			expectCall: func(e testCase, tx *types_mocks.MockTxnRepository, eventTx *types_mocks.MockEventTxnRepository) {
				InsertTodoHelper(tx, model.Todo{
					Name: "new todo create",
				}, 55, nil)

				gomock.InOrder(
					InsertItemHelper(tx, model.TodoItem{
						TodoID: 55,
						Name:   "new item 4",
					}, 1, nil),
					InsertItemHelper(tx, model.TodoItem{
						TodoID: 55,
						Name:   "new item 5",
					}, 2, nil),
				)

				newInput := e.input
				newInput.ID = 55
				InsertEventHelper(eventTx, BuildTodoSaveEvent(newInput), 31,
					errors.General.InternalErrorAccessingDatabase.Err())
			},
			expectedErr: errors.General.InternalErrorAccessingDatabase.Err(),
		},
		{
			name: "insert-ok",
			input: types.SaveTodoInput{
				ID:   0,
				Name: "new todo create",
				Items: []model.TodoItem{
					{
						Name: "new item 4",
					},
					{
						Name: "new item 5",
					},
				},
			},
			expectCall: func(e testCase, tx *types_mocks.MockTxnRepository, eventTx *types_mocks.MockEventTxnRepository) {
				InsertTodoHelper(tx, model.Todo{
					Name: "new todo create",
				}, 55, nil)

				gomock.InOrder(
					InsertItemHelper(tx, model.TodoItem{
						TodoID: 55,
						Name:   "new item 4",
					}, 1, nil),
					InsertItemHelper(tx, model.TodoItem{
						TodoID: 55,
						Name:   "new item 5",
					}, 2, nil),
				)

				newInput := e.input
				newInput.ID = 55
				InsertEventHelper(eventTx, BuildTodoSaveEvent(newInput), 31, nil)
			},
			expectedErr: nil,
			expectedID:  55,
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
