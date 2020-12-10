package server

import (
	"context"
	"todoapp-rpc/rpc/health/v1"
	"todoapp/pkg/errors"
)

// HealthServer for health check
type HealthServer struct {
	health.UnimplementedHealthServiceServer
}

// Live ...
func (s *HealthServer) Live(context.Context, *health.LiveRequest) (*health.LiveResponse, error) {
	return nil, errors.General.UnknownValue.Err()
}

// Ready ...
func (s *HealthServer) Ready(context.Context, *health.ReadyRequest) (*health.ReadyResponse, error) {
	return nil, errors.General.UnknownValue.WithMin(223).Err()
}
