package server

import (
	"context"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	health_rpc "todoapp-rpc/rpc/health/v1"
	todoapp_rpc "todoapp-rpc/rpc/todoapp/v1"
	common_server "todoapp/common/server"
	"todoapp/config"
	"todoapp/lib/errors"
	"todoapp/lib/log"
	"todoapp/lib/mysql"
	todoapp_server "todoapp/todoapp/server"
)

// Root struct for whole app
type Root struct {
	conf   config.Config
	db     *sqlx.DB
	logger *zap.Logger

	health        *common_server.HealthServer
	todoappServer *todoapp_server.Server
}

// NewRoot initializes gRPC servers
func NewRoot(conf config.Config) *Root {
	logger := log.NewLogger(conf.Log)
	db := mysql.MustConnect(conf.MySQL)

	conn, err := grpc.Dial(conf.Event.GRPC.String(), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	todoappServer := todoapp_server.InitServer(db, conn)

	return &Root{
		conf:   conf,
		db:     db,
		logger: logger,

		health:        &common_server.HealthServer{},
		todoappServer: todoappServer,
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
	health_rpc.RegisterHealthServiceServer(server, r.health)
	if err := health_rpc.RegisterHealthServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		panic(err)
	}

	todoapp_rpc.RegisterTodoServiceServer(server, r.todoappServer)
	if err := todoapp_rpc.RegisterTodoServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
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
