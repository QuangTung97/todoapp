package server

import (
	"context"
	health_rpc "todoapp-rpc/rpc/health/v1"
)

// HealthServer for health check
type HealthServer struct {
	health_rpc.UnimplementedHealthServiceServer
}

// Live ...
func (s *HealthServer) Live(context.Context, *health_rpc.LiveRequest) (*health_rpc.LiveResponse, error) {
	return &health_rpc.LiveResponse{}, nil
}

// Ready ...
func (s *HealthServer) Ready(context.Context, *health_rpc.ReadyRequest) (*health_rpc.ReadyResponse, error) {
	return &health_rpc.ReadyResponse{}, nil
}
