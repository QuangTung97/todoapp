package errors

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestError_Error(t *testing.T) {
	e := &Error{
		RPCStatus: 3,
		Code:      "0322",
		Message:   "Some error",
		Details: map[string]interface{}{
			"max": int64(100),
		},
	}
	s := e.Error()
	expected := "rpc status: InvalidArgument, code: 0322, message: Some error, details: map[max:100]"
	assert.Equal(t, expected, s)
}

func TestError_WithDetail(t *testing.T) {
	e0 := &Error{
		RPCStatus: 3,
		Code:      "0322",
		Message:   "Some error",
		Details: map[string]interface{}{
			"max": int64(100),
		},
	}

	expected := "rpc status: InvalidArgument, code: 0322, message: Some error, details: map[max:100]"
	assert.Equal(t, expected, e0.Error())

	e1 := e0.WithDetail("max", int64(200))
	expected = "rpc status: InvalidArgument, code: 0322, message: Some error, details: map[max:200]"
	assert.Equal(t, expected, e1.Error())
}

func TestError_ToRPCError_Normal(t *testing.T) {
	e := &Error{
		RPCStatus: 3,
		Code:      "0322",
		Message:   "Some error",
		Details: map[string]interface{}{
			"max": int64(100),
		},
	}

	err := e.ToRPCError()
	e1, ok := FromRPCError(err)

	assert.True(t, ok)
	assert.Equal(t, e, e1)
}

func TestError_ToRPCError_Details(t *testing.T) {
	table := []struct {
		name           string
		Details        map[string]interface{}
		expected       string
		notConvertBack bool
	}{
		{
			name:     "ok",
			expected: "rpc error: code = InvalidArgument desc = Some test error",
		},
		{
			name: "ok bool",
			Details: map[string]interface{}{
				"isBool": true,
			},
			expected: "rpc error: code = InvalidArgument desc = Some test error",
		},
		{
			name: "ok time.Time",
			Details: map[string]interface{}{
				"startedAt": time.Now(),
			},
			expected: "rpc error: code = InvalidArgument desc = Some test error",
		},
		{
			name: "unrecognized type",
			Details: map[string]interface{}{
				"num": 200,
			},
			expected:       "unrecognized error detail type",
			notConvertBack: true,
		},
	}

	for _, e := range table {
		t.Run(e.name, func(t *testing.T) {
			err := &Error{
				RPCStatus: 3,
				Code:      "0302",
				Message:   "Some test error",
				Details:   e.Details,
			}
			rpcErr := err.ToRPCError()
			assert.NotNil(t, rpcErr)
			if rpcErr != nil {
				assert.Equal(t, e.expected, rpcErr.Error())
			}

			if !e.notConvertBack {
				_, ok := FromRPCError(rpcErr)
				assert.True(t, ok)
			}
		})
	}
}
