package errors

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
)

// WrapDBError wraps database errors
func WrapDBError(ctx context.Context, err error) error {
	if err != nil {
		ctxzap.Extract(ctx).WithOptions(zap.AddCallerSkip(1)).
			Error("Error accessing database", zap.Error(err))
		return General.InternalErrorAccessingDatabase.Err()
	}
	return nil
}
