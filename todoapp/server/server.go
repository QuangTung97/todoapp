package server

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	todoapp_rpc "todoapp-rpc/rpc/todoapp/v1"
	"todoapp/todoapp/types"
)

// Server for todoapp server
type Server struct {
	todoapp_rpc.UnimplementedTodoServiceServer

	service types.Service
}

var _ todoapp_rpc.TodoServiceServer = &Server{}

// NewServer creates a todoapp server
func NewServer(service types.Service) *Server {
	return &Server{
		service: service,
	}
}

// Save ...
func (s *Server) Save(ctx context.Context, req *todoapp_rpc.TodoSaveRequest,
) (*todoapp_rpc.TodoSaveResponse, error) {
	input, err := transformSaveRequest(req)
	if err != nil {
		return nil, err
	}

	id, err := s.service.SaveTodo(ctx, input)
	if err != nil {
		return nil, err
	}

	return &todoapp_rpc.TodoSaveResponse{
		Id: int64(id),
	}, nil
}

// List for listing
func (s *Server) List(ctx context.Context, req *todoapp_rpc.TodoListRequest,
) (*todoapp_rpc.TodoListResponse, error) {
	return &todoapp_rpc.TodoListResponse{
		Todos: []*todoapp_rpc.TodoData{
			{
				Id:        100,
				Name:      "Quang Tung",
				CreatedAt: ptypes.TimestampNow(),
			},
		},
	}, nil
}
