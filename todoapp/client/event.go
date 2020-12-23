package client

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	todoapp_rpc "todoapp-rpc/rpc/todoapp/v1"
	"todoapp/todoapp/types"
)

// EventClient ...
type EventClient struct {
	conn *grpc.ClientConn
}

var _ types.EventClient = &EventClient{}

// NewEventClient ...
func NewEventClient(conn *grpc.ClientConn) *EventClient {
	return &EventClient{
		conn: conn,
	}
}

// Signal ...
func (c *EventClient) Signal(ctx context.Context) {
	client := todoapp_rpc.NewEventServiceClient(c.conn)
	_, err := client.Signal(context.Background(), &todoapp_rpc.SignalRequest{})
	if err != nil {
		ctxzap.Extract(ctx).Error("client.Signal", zap.Error(err))
	}
}
