package server

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"todoapp-rpc/rpc/health/v1"
)

// Root struct for whole app
type Root struct {
	db *sqlx.DB

	health *HealthServer
}

// NewRoot initializes gRPC servers
func NewRoot(db *sqlx.DB) *Root {
	return &Root{
		db:     db,
		health: &HealthServer{},
	}
}

// Register register gRPC & gateway servers
func (r *Root) Register(ctx context.Context, server *grpc.Server, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
	health.RegisterHealthServiceServer(server, r.health)
	if err := health.RegisterHealthServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		panic(err)
	}
}
