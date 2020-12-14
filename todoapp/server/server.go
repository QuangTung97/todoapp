package server

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	todoapp_rpc "todoapp-rpc/rpc/todoapp/v1"
)

// Server for todoapp server
type Server struct {
	todoapp_rpc.UnimplementedTodoServiceServer
}

var _ todoapp_rpc.TodoServiceServer = &Server{}

// NewServer creates a todoapp server
func NewServer() *Server {
	return &Server{}
}

// List for listing
func (s *Server) List(ctx context.Context, req *todoapp_rpc.TodoListRequest,
) (*todoapp_rpc.TodoListResponse, error) {
	return &todoapp_rpc.TodoListResponse{
		Todos: []*todoapp_rpc.TodoData{
			{
				Id:        2,
				Name:      "Quang Tung",
				CreatedAt: ptypes.TimestampNow(),
			},
		},
	}, nil
}
