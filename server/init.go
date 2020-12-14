package server

import (
	"context"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"todoapp-rpc/rpc/health/v1"
	"todoapp/config"
	"todoapp/lib/errors"
	"todoapp/lib/log"
	"todoapp/lib/mysql"
)

// Root struct for whole app
type Root struct {
	conf   config.Config
	db     *sqlx.DB
	logger *zap.Logger

	health *HealthServer
}

// NewRoot initializes gRPC servers
func NewRoot(conf config.Config) *Root {
	logger := log.NewLogger(conf.Log)
	db := mysql.MustConnect(conf.MySQL)

	return &Root{
		conf:   conf,
		db:     db,
		logger: logger,
		health: &HealthServer{},
	}
}

func deciderAllMethods(ctx context.Context, fullMethodName string, servingObject interface{}) bool {
	return true
}

// UnaryInterceptor creates unary server interceptor
func (r *Root) UnaryInterceptor() grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		grpc_ctxtags.UnaryServerInterceptor(),
		grpc_prometheus.UnaryServerInterceptor,
		grpc_zap.UnaryServerInterceptor(r.logger),
		log.PayloadUnaryServerInterceptor(r.logger, deciderAllMethods, r.conf.Log.MaskedFields...),
		grpc_recovery.UnaryServerInterceptor(),
		errors.UnaryServerInterceptor,
	)
}

// StreamInterceptor creates stream server interceptor
func (r *Root) StreamInterceptor() grpc.ServerOption {
	return grpc.ChainStreamInterceptor(
		grpc_ctxtags.StreamServerInterceptor(),
		grpc_prometheus.StreamServerInterceptor,
		grpc_zap.StreamServerInterceptor(r.logger),
		grpc_recovery.StreamServerInterceptor(),
	)
}

// Register register gRPC & gateway servers
func (r *Root) Register(ctx context.Context, server *grpc.Server, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
	health.RegisterHealthServiceServer(server, r.health)
	if err := health.RegisterHealthServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		panic(err)
	}
}

// Shutdown for graceful shutdown
func (r *Root) Shutdown() {
	if err := r.db.Close(); err != nil {
		panic(err)
	}

	r.logger.Info("Graceful shutdown completed")
	_ = r.logger.Sync()
}
