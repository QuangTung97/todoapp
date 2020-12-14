package log

import (
	"bytes"
	"context"
	"fmt"
	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"path"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"go.uber.org/zap/zapcore"
)

var (
	// SystemField is used in every log statement made through grpc_zap. Can be overwritten before any initialization code.
	SystemField = zap.String("system", "grpc")

	// ServerField is used in every server-side log statement made through grpc_zap.Can be overwritten before initialization.
	ServerField = zap.String("span.kind", "server")
)

var (
	// JSONPbMarshaller is the marshaller used for serializing protobuf messages.
	// If needed, this variable can be reassigned with a different marshaller with the same Marshal() signature.
	JSONPbMarshaller grpc_logging.JsonPbMarshaler = &jsonpb.Marshaler{}
)

// PayloadUnaryServerInterceptor returns a new unary server interceptors that logs the payloads of requests.
//
// This *only* works when placed *after* the `grpc_zap.UnaryServerInterceptor`. However, the logging can be done to a
// separate instance of the logger.
//
// *maskedFields* *only* support fields of type string
func PayloadUnaryServerInterceptor(logger *zap.Logger, decider grpc_logging.ServerPayloadLoggingDecider, maskedFields ...string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if !decider(ctx, info.FullMethod, info.Server) {
			return handler(ctx, req)
		}

		maskedReq, ok := hideStringFields(req, maskedFields...)
		if !ok {
			return handler(ctx, req)
		}

		// Use the provided zap.Logger for logging but use the fields from context.
		logEntry := logger.With(append(serverCallFields(info.FullMethod), ctxzap.TagsToFields(ctx)...)...)
		logProtoMessageAsJSON(logEntry, maskedReq, "grpc.request.content", "server request payload logged as grpc.request.content field")
		resp, err := handler(ctx, req)
		if err == nil {
			maskedResp, ok := hideStringFields(resp, maskedFields...)
			if ok {
				logProtoMessageAsJSON(logEntry, maskedResp, "grpc.response.content", "server response payload logged as grpc.response.content field")
			}
		}
		return resp, err
	}
}

// PayloadStreamServerInterceptor returns a new server server interceptors that logs the payloads of requests.
//
// This *only* works when placed *after* the `grpc_zap.StreamServerInterceptor`. However, the logging can be done to a
// separate instance of the logger.
func PayloadStreamServerInterceptor(logger *zap.Logger, decider grpc_logging.ServerPayloadLoggingDecider) grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if !decider(stream.Context(), info.FullMethod, srv) {
			return handler(srv, stream)
		}
		logEntry := logger.With(append(serverCallFields(info.FullMethod), ctxzap.TagsToFields(stream.Context())...)...)
		newStream := &loggingServerStream{ServerStream: stream, logger: logEntry}
		return handler(srv, newStream)
	}
}

type loggingServerStream struct {
	grpc.ServerStream
	logger *zap.Logger
}

func (l *loggingServerStream) SendMsg(m interface{}) error {
	err := l.ServerStream.SendMsg(m)
	if err == nil {
		logProtoMessageAsJSON(l.logger, m, "grpc.response.content", "server response payload logged as grpc.response.content field")
	}
	return err
}

func (l *loggingServerStream) RecvMsg(m interface{}) error {
	err := l.ServerStream.RecvMsg(m)
	if err == nil {
		logProtoMessageAsJSON(l.logger, m, "grpc.request.content", "server request payload logged as grpc.request.content field")
	}
	return err
}

func logProtoMessageAsJSON(logger *zap.Logger, pbMsg interface{}, key string, msg string) {
	if p, ok := pbMsg.(proto.Message); ok {
		logger.Check(zapcore.InfoLevel, msg).Write(zap.Object(key, &jsonpbObjectMarshaler{pb: p}))
	}
}

type jsonpbObjectMarshaler struct {
	pb proto.Message
}

func (j *jsonpbObjectMarshaler) MarshalLogObject(e zapcore.ObjectEncoder) error {
	// ZAP jsonEncoder deals with AddReflect by using json.MarshalObject. The same thing applies for consoleEncoder.
	return e.AddReflected("msg", j)
}

func (j *jsonpbObjectMarshaler) MarshalJSON() ([]byte, error) {
	b := &bytes.Buffer{}
	if err := JSONPbMarshaller.Marshal(b, j.pb); err != nil {
		return nil, fmt.Errorf("jsonpb serializer failed: %v", err)
	}
	return b.Bytes(), nil
}

func serverCallFields(fullMethodString string) []zapcore.Field {
	service := path.Dir(fullMethodString)[1:]
	method := path.Base(fullMethodString)
	return []zapcore.Field{
		SystemField,
		ServerField,
		zap.String("grpc.service", service),
		zap.String("grpc.method", method),
	}
}
