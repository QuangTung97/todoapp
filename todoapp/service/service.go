package service

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
	"todoapp/pkg/errors"
	"todoapp/todoapp/model"
	"todoapp/todoapp/types"
)

// Service ...
type Service struct {
	repo types.Repository
}

var _ types.Service = &Service{}

// NewService creates a service
func NewService(repo types.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// SaveTodo ...
func (s *Service) SaveTodo(ctx context.Context, input types.SaveTodoInput) (model.TodoID, error) {
	err := s.repo.Transact(ctx, func(tx types.TxnRepository) error {
		nullTodo, err := tx.GetTodo(ctx, input.ID)
		if err != nil {
			return err
		}

		if !nullTodo.Valid {
			return errors.Todo.NotFoundTodo.Err()
		}

		ctxzap.Extract(ctx).Info("Log todo", zap.Any("todo", nullTodo.Todo))

		return nil
	})
	if err != nil {
		return 0, err
	}

	return input.ID, nil
}
