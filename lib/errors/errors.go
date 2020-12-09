package errors

import (
	"context"
	stderrors "errors"
	"fmt"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"github.com/gogo/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const codeDetailField = "code"

// Error for domain error
type Error struct {
	RPCStatus uint32
	Code      string
	Message   string
	Details   map[string]interface{}
}

func (e *Error) Error() string {
	return fmt.Sprintf(
		"rpc status: %v, code: %s, message: %s, details: %v",
		codes.Code(e.RPCStatus), e.Code, e.Message, e.Details,
	)
}

var _ error = &Error{}

func fieldValueToDetail(field string, value interface{}) (proto.Message, error) {
	switch v := value.(type) {
	case bool:
		return &ErrorDetailBool{Field: field, Value: v}, nil

	case string:
		return &ErrorDetailString{Field: field, Value: v}, nil

	case int64:
		return &ErrorDetailInt64{Field: field, Value: v}, nil

	case float64:
		return &ErrorDetailDouble{Field: field, Value: v}, nil

	case time.Time:
		t, _ := types.TimestampProto(v)
		return &ErrorDetailTimestamp{Field: field, Value: t}, nil

	default:
		return nil, stderrors.New("unrecognized error detail type")
	}
}

func detailToFieldValue(detail interface{}) (string, interface{}, error) {
	switch d := detail.(type) {
	case *ErrorDetailBool:
		return d.Field, d.Value, nil

	case *ErrorDetailString:
		return d.Field, d.Value, nil

	case *ErrorDetailInt64:
		return d.Field, d.Value, nil

	case *ErrorDetailDouble:
		return d.Field, d.Value, nil

	case *ErrorDetailTimestamp:
		t, err := types.TimestampFromProto(d.Value)
		if err != nil {
			return "", nil, err
		}
		return d.Field, t, nil

	default:
		return "", nil, stderrors.New("unrecognized interface type")
	}
}

// ToRPCError converts Error to grpc status error
func (e *Error) ToRPCError() error {
	st := status.New(codes.Code(e.RPCStatus), e.Message)

	var details []proto.Message

	details = append(details, &ErrorDetailString{
		Field: codeDetailField,
		Value: e.Code,
	})

	for field, value := range e.Details {
		detail, err := fieldValueToDetail(field, value)
		if err != nil {
			return err
		}
		details = append(details, detail)
	}

	st, err := st.WithDetails(details...)
	if err != nil {
		return err
	}

	return st.Err()
}

// FromRPCStatus converts status to Error
func FromRPCStatus(st *status.Status) (*Error, bool) {
	details := st.Details()

	if len(details) == 0 {
		return nil, false
	}

	code, ok := details[0].(*ErrorDetailString)
	if !ok {
		return nil, false
	}
	if code.Field != codeDetailField {
		return nil, false
	}

	detailMap := make(map[string]interface{})
	for _, detail := range details[1:] {
		field, value, err := detailToFieldValue(detail)
		if err != nil {
			return nil, false
		}
		detailMap[field] = value
	}

	return &Error{
		RPCStatus: uint32(st.Code()),
		Code:      code.Value,
		Message:   st.Message(),
		Details:   detailMap,
	}, true
}

// UnaryServerInterceptor converts domain error to grpc status error
func UnaryServerInterceptor(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (interface{}, error) {
	resp, err := handler(ctx, req)
	if err != nil {
		var domainErr *Error
		if stderrors.As(err, &domainErr) {
			return nil, domainErr.ToRPCError()
		}

		st, ok := status.FromError(err)
		if ok {
			return nil, st.Err()
		}

		st = status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}
	return resp, nil
}
