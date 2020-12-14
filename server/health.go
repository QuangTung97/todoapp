package server

import (
	"context"
	"todoapp-rpc/rpc/health/v1"
	"todoapp/lib/dblib"
	"todoapp/pkg/errors"
)

// HealthServer for health check
type HealthServer struct {
	health.UnimplementedHealthServiceServer
}

var _ = dblib.NewQuery(`
SELECT * FROM tung
`)

var _ = dblib.NewQuery(`
SELECT * FROM testing WHERE id = ?
`)

var _ = dblib.NewQuery(`
DELETE FROM testing WHERE id IN (?)
`)

var _ = dblib.NewNamedQuery(`
INSERT INTO testing (id, version, value)
VALUES (:id, :version, :value)
`)

// Live ...
func (s *HealthServer) Live(context.Context, *health.LiveRequest) (*health.LiveResponse, error) {
	return nil, errors.General.UnknownValue.Err()
}

// Ready ...
func (s *HealthServer) Ready(context.Context, *health.ReadyRequest) (*health.ReadyResponse, error) {
	return nil, errors.General.UnknownValue.WithMin(223).WithMaxValue("Quang Tung").Err()
}
