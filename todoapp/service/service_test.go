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

func mockTransact(ctrl *gomock.Controller, repo *types_mocks.MockRepository,
) *types_mocks.MockTxnRepository {
	mockTxn := types_mocks.NewMockTxnRepository(ctrl)

	repo.EXPECT().Transact(gomock.Any(), gomock.Any()).DoAndReturn(
		func(ctx context.Context, fn func(repository types.TxnRepository) error) error {
			return fn(mockTxn)
		})
	return mockTxn
}

func TestService_SaveTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := types_mocks.NewMockRepository(ctrl)
	txn := mockTransact(ctrl, repo)
	txn.EXPECT().GetTodo(gomock.Any(), model.TodoID(0)).Return(model.NullTodo{}, nil)

	s := NewService(repo)

	ctx := context.Background()
	input := types.SaveTodoInput{}
	_, err := s.SaveTodo(ctx, input)
	assert.Equal(t, errors.Todo.NotFoundTodo.Err(), err)
}
