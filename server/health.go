package server

import (
	"context"
	"todoapp-rpc/rpc/health/v1"
)

// HealthServer for health check
type HealthServer struct {
	health.UnimplementedHealthServiceServer
}

// Live ...
func (s *HealthServer) Live(context.Context, *health.LiveRequest) (*health.LiveResponse, error) {
	return &health.LiveResponse{}, nil
}

// Ready ...
func (s *HealthServer) Ready(context.Context, *health.ReadyRequest) (*health.ReadyResponse, error) {
	return &health.ReadyResponse{}, nil
}
