package server

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

// Root
type Root struct {
	db *sqlx.DB
}

// NewRoot initializes gRPC servers
func NewRoot(db *sqlx.DB) *Root {
	return &Root{
		db: db,
	}
}

// Register register gRPC & gateway servers
func (r *Root) Register(ctx context.Context, server *grpc.Server, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
}
