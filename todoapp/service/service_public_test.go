package service_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	todoapp_rpc "todoapp-rpc/rpc/todoapp/v1"
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
	mockTx.EXPECT().GetTodo(gomock.Any(), model.TodoID(123)).
		Return(model.NullTodo{}, errors.General.InternalErrorAccessingDatabase.Err())

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

	mockTx.EXPECT().ToEventRepository().Return(mockEventRepo)
	mockTx.EXPECT().InsertTodo(gomock.Any(), model.TodoSave{
		Name: "some save todo",
	}).Return(model.TodoID(555), nil)
	mockTx.EXPECT().InsertTodoItem(gomock.Any(), model.TodoItemSave{
		TodoID: 555,
		Name:   "some item",
	}).Return(model.TodoItemID(33), nil)

	mockEventRepo.EXPECT().InsertEvent(gomock.Any(), types.Event{
		Data: &todoapp_rpc.Event{
			Type: todoapp_rpc.EventType_EVENT_TYPE_TODO_SAVE,
			TodoSave: &todoapp_rpc.EventTodoSave{
				Id:   555,
				Name: "some save todo",
			},
		},
	}.ToModel()).Return(model.EventID(88), nil)

	mockClient.EXPECT().Signal(gomock.Any())

	mockRepo.EXPECT().Transact(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, fn func(tx types.TxnRepository) error) error {
			return fn(mockTx)
		})

	s := service.NewService(mockRepo, mockClient)
	id, err := s.SaveTodo(context.Background(), types.SaveTodoInput{
		Name: "some save todo",
		Items: []types.SaveTodoItem{
			{
				Name: "some item",
			},
		},
	})
	assert.Equal(t, nil, err)
	assert.Equal(t, model.TodoID(555), id)
}
