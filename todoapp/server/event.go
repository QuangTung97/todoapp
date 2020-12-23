package server

import (
	"context"
	todoapp_rpc "todoapp-rpc/rpc/todoapp/v1"
	"todoapp/todoapp/event/core"
)

// EventServer ...
type EventServer struct {
	todoapp_rpc.UnimplementedEventServiceServer
	core *core.Core
}

// NewEventServer ...
func NewEventServer(core *core.Core) *EventServer {
	return &EventServer{
		core: core,
	}
}

// Signal ...
func (s *EventServer) Signal(context.Context, *todoapp_rpc.SignalRequest,
) (*todoapp_rpc.SignalResponse, error) {
	s.core.Signal()
	return &todoapp_rpc.SignalResponse{}, nil
}
