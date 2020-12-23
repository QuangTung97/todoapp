package service

import (
	"context"
	"testing"
	"todoapp/pkg/errors"
	types_mocks "todoapp/todoapp/mocks"
	"todoapp/todoapp/model"
	"todoapp/todoapp/types"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func mockTransact(ctrl *gomock.Controller, repo *types_mocks.MockRepository,
) *types_mocks.MockTxnRepository {
	mockTxn := types_mocks.NewMockTxnRepository(ctrl)

	repo.EXPECT().Transact(gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, fn func(repository types.TxnRepository) error) error {
			return fn(mockTxn)
		})
	return mockTxn
}

func TestService_SaveTodo_Insert(t *testing.T) {

	table := []struct {
		name  string
		input types.SaveTodoInput

		expectedCall func(tx *types_mocks.MockTxnRepository)
		expectedErr  error
		expectedID   model.TodoID
	}{
		{
			name: "empty items",
			input: types.SaveTodoInput{
				Items: nil,
			},
			expectedErr: errors.Todo.InvalidArgumentEmptyItems.Err(),
		},
		{
			name: "empty items",
			input: types.SaveTodoInput{
				Name: "Todo Item",
				Items: []types.SaveTodoItem{
					{Name: "Item 1"},
				},
			},
			expectedCall: func(tx *types_mocks.MockTxnRepository) {
				inserted := model.TodoSave{
					Name: "Todo Item",
				}
				id := model.TodoID(100)
				tx.EXPECT().InsertTodo(gomock.Any(), inserted).Return(id, nil)

				item := model.TodoItemSave{
					TodoID: 100,
					Name:   "Item 1",
				}
				idItem := model.TodoItemID(22)
				tx.EXPECT().InsertTodoItem(gomock.Any(), item).Return(idItem, nil)
			},
			expectedErr: nil,
			expectedID:  100,
		},
	}

	for _, e := range table {
		t.Run(e.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := types_mocks.NewMockRepository(ctrl)
			tx := mockTransact(ctrl, repo)

			if e.expectedCall != nil {
				// txn.EXPECT().GetTodo(gomock.Any(), model.TodoID(0)).Return(model.NullTodo{}, nil)
				e.expectedCall(tx)
			}

			s := NewService(repo, nil)

			ctx := context.Background()
			id, err := s.SaveTodo(ctx, e.input)

			assert.Equal(t, e.expectedErr, err)
			assert.Equal(t, e.expectedID, id)
		})
	}

}
