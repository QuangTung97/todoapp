package service_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"todoapp/pkg/errors"
	types_mocks "todoapp/todoapp/mocks"
	"todoapp/todoapp/model"
	"todoapp/todoapp/service"
	"todoapp/todoapp/types"
)

func TestService_SaveTodo_Update_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := types_mocks.NewMockRepository(ctrl)
	mockClient := types_mocks.NewMockEventClient(ctrl)
	mockTx := types_mocks.NewMockTxnRepository(ctrl)
	mockEventRepo := types_mocks.NewMockEventTxnRepository(ctrl)

	mockTx.EXPECT().ToEventRepository().Return(mockEventRepo)
	service.GetTodoHelper(mockTx, 123, model.NullTodo{}, errors.General.InternalErrorAccessingDatabase.Err())

	mockRepo.EXPECT().Transact(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, fn func(tx types.TxnRepository) error) error {
			return fn(mockTx)
		})

	s := service.NewService(mockRepo, mockClient)
	id, err := s.SaveTodo(context.Background(), types.SaveTodoInput{
		ID:   123,
		Name: "some save todo",
		Items: []types.SaveTodoItem{
			{
				ID:   12,
				Name: "some item",
			},
		},
	})
	assert.Equal(t, errors.General.InternalErrorAccessingDatabase.Err(), err)
	assert.Equal(t, model.TodoID(0), id)
}

func TestService_SaveTodo_Insert_OK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := types_mocks.NewMockRepository(ctrl)
	mockClient := types_mocks.NewMockEventClient(ctrl)
	mockTx := types_mocks.NewMockTxnRepository(ctrl)
	mockEventRepo := types_mocks.NewMockEventTxnRepository(ctrl)

	input := types.SaveTodoInput{
		Name: "some save todo",
		Items: []types.SaveTodoItem{
			{
				Name: "some item",
			},
		},
	}

	mockTx.EXPECT().ToEventRepository().Return(mockEventRepo)
	service.InsertTodoHelper(mockTx, model.TodoSave{
		Name: "some save todo",
	}, 555, nil)
	service.InsertItemHelper(mockTx, model.TodoItemSave{
		TodoID: 555,
		Name:   "some item",
	}, 33, nil)

	newInput := input
	newInput.ID = 555
	service.InsertEventHelper(mockEventRepo, service.BuildTodoSaveEvent(newInput), 88, nil)

	mockClient.EXPECT().Signal(gomock.Any())

	mockRepo.EXPECT().Transact(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, fn func(tx types.TxnRepository) error) error {
			return fn(mockTx)
		})

	s := service.NewService(mockRepo, mockClient)
	id, err := s.SaveTodo(context.Background(), input)
	assert.Equal(t, nil, err)
	assert.Equal(t, model.TodoID(555), id)
}
