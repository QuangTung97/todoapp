package generate

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateError(t *testing.T) {
	table := []struct {
		name      string
		errorName string
		info      ErrorInfo
		err       error
	}{
		{
			name: "empty rpc status",
			err:  fmt.Errorf("invalid rpc status '0'"),
		},
		{
			name: "exceed rpc status",
			info: ErrorInfo{
				RPCStatus: 17,
			},
			err: fmt.Errorf("invalid rpc status '17'"),
		},
		{
			name: "code prefix",
			info: ErrorInfo{
				RPCStatus: 1,
			},
			err: fmt.Errorf("code must be prefix with '01'"),
		},
		{
			name: "code prefix",
			info: ErrorInfo{
				RPCStatus: 16,
			},
			err: fmt.Errorf("code must be prefix with '16'"),
		},
		{
			name: "name prefix",
			info: ErrorInfo{
				RPCStatus: 16,
				Code:      "1600",
			},
			err: fmt.Errorf("error name must prefix with 'unauthenticated'"),
		},
		{
			name:      "ok",
			errorName: "unauthenticatedPasswordIncorrect",
			info: ErrorInfo{
				RPCStatus: 16,
				Code:      "1600",
			},
			err: nil,
		},
	}

	for _, e := range table {
		t.Run(e.name, func(t *testing.T) {
			err := validateError(e.errorName, e.info)
			assert.Equal(t, e.err, err)
		})
	}
}

func TestValidate(t *testing.T) {
	table := []struct {
		name string
		tags map[string]ErrorMap
		err  error
	}{
		{
			name: "error name prefix",
			tags: map[string]ErrorMap{
				"auth": {
					"unauthenticated": {
						RPCStatus: 14,
						Code:      "1403",
						Message:   "Some error",
					},
				},
			},
			err: fmt.Errorf("error name must prefix with 'unavailable'"),
		},
		{
			name: "duplicated",
			tags: map[string]ErrorMap{
				"auth": {
					"unknownValue": {
						RPCStatus: 2,
						Code:      "0203",
						Message:   "Some error",
					},
				},
				"general": {
					"unknownField": {
						RPCStatus: 2,
						Code:      "0203",
						Message:   "Unknown field",
					},
				},
			},
			err: fmt.Errorf("code '0203' duplicated"),
		},
		{
			name: "ok",
			tags: map[string]ErrorMap{
				"auth": {
					"unknownValue": {
						RPCStatus: 2,
						Code:      "0203",
						Message:   "Some error",
					},
				},
				"general": {
					"unknownField": {
						RPCStatus: 2,
						Code:      "0204",
						Message:   "Unknown field",
					},
				},
			},
			err: nil,
		},
	}

	for _, e := range table {
		t.Run(e.name, func(t *testing.T) {
			err := Validate(e.tags)
			assert.Equal(t, e.err, err)
		})
	}
}

func TestFindNextErrorCode(t *testing.T) {
	table := []struct {
		name      string
		rpcStatus uint32
		codes     []int
		expected  string
	}{
		{

			name:      "empty",
			rpcStatus: 1,
			expected:  "0100",
		},
		{
			name:      "two",
			rpcStatus: 2,
			codes: []int{
				0, 1,
			},
			expected: "0202",
		},
		{
			name:      "duplicated",
			rpcStatus: 2,
			codes: []int{
				0, 1, 2, 2,
			},
			expected: "0203",
		},
		{
			name:      "duplicated",
			rpcStatus: 2,
			codes: []int{
				0, 1, 3, 4, 4, 5, 2, 2,
			},
			expected: "0206",
		},
		{
			name:      "missing middle",
			rpcStatus: 2,
			codes: []int{
				0, 1, 4, 2, 5,
			},
			expected: "0203",
		},
	}

	for _, e := range table {
		t.Run(e.name, func(t *testing.T) {
			code := findNextErrorCode(e.rpcStatus, e.codes)
			assert.Equal(t, e.expected, code)
		})
	}
}
