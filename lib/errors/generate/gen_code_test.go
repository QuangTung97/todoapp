package generate

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGenerateNewErrorFunc(t *testing.T) {
	var buf bytes.Buffer
	err := generateNewErrorFunc("ErrGeneralNotFound", ErrorInfo{
		RPCStatus: 13,
		Code:      "1303",
		Message:   "Not found",
	}, &buf)

	expected := `
// NewErrGeneralNotFound ...
func NewErrGeneralNotFound() *ErrGeneralNotFound {
	return &ErrGeneralNotFound{
		RPCStatus: 13,
		Code:      "1303",
		Message:   "Not found",
	}
}
`
	expected = strings.TrimSpace(expected)

	assert.Nil(t, err)
	assert.Equal(t, expected, buf.String())
}

func TestGenerateWithMethod(t *testing.T) {
	var buf bytes.Buffer
	err := generateWithMethod("ErrGeneralUnknown", "total", "int64", &buf)
	assert.Nil(t, err)

	expected := `
// WithTotal ...
func (e *ErrGeneralUnknown) WithTotal(value int64) *ErrGeneralUnknown {
	err := (*liberrors.Error)(e)
	return (*ErrGeneralUnknown)(err.WithDetail("total", value))
}
`
	expected = strings.TrimSpace(expected)
	assert.Equal(t, expected, buf.String())
}

func TestGenerateErrorEntry(t *testing.T) {
	table := []struct {
		name     string
		input    errorEntry
		tagName  string
		expected string
	}{
		{
			name: "simple",
			input: errorEntry{
				name: "notFound",
				info: ErrorInfo{
					RPCStatus: 11,
					Code:      "1101",
					Message:   "Not found",
				},
			},
			tagName: "General",
			expected: `
// ErrGeneralNotFound ...
type ErrGeneralNotFound liberrors.Error

// NewErrGeneralNotFound ...
func NewErrGeneralNotFound() *ErrGeneralNotFound {
	return &ErrGeneralNotFound{
		RPCStatus: 11,
		Code:      "1101",
		Message:   "Not found",
	}
}

// Err ...
func (e *ErrGeneralNotFound) Err() error {
	return (*liberrors.Error)(e)
}
`,
		},
		{
			name: "with-details",
			input: errorEntry{
				name: "unknown",
				info: ErrorInfo{
					RPCStatus: 3,
					Code:      "0302",
					Message:   "Unknown",
					Details: map[string]string{
						"total":       "int64",
						"startedTime": "time.Time",
					},
				},
			},
			tagName: "General",
			expected: `
// ErrGeneralUnknown ...
type ErrGeneralUnknown liberrors.Error

// NewErrGeneralUnknown ...
func NewErrGeneralUnknown() *ErrGeneralUnknown {
	return &ErrGeneralUnknown{
		RPCStatus: 3,
		Code:      "0302",
		Message:   "Unknown",
	}
}

// Err ...
func (e *ErrGeneralUnknown) Err() error {
	return (*liberrors.Error)(e)
}

// WithStartedTime ...
func (e *ErrGeneralUnknown) WithStartedTime(value time.Time) *ErrGeneralUnknown {
	err := (*liberrors.Error)(e)
	return (*ErrGeneralUnknown)(err.WithDetail("startedTime", value))
}

// WithTotal ...
func (e *ErrGeneralUnknown) WithTotal(value int64) *ErrGeneralUnknown {
	err := (*liberrors.Error)(e)
	return (*ErrGeneralUnknown)(err.WithDetail("total", value))
}
`,
		},
	}

	for _, e := range table {
		t.Run(e.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := generateErrorEntry(e.tagName, e.input, &buf)
			assert.Nil(t, err)

			e.expected = strings.TrimSpace(e.expected)
			assert.Equal(t, e.expected, buf.String())
		})
	}
}

func TestGenerateErrorTag(t *testing.T) {
	table := []struct {
		name     string
		input    errorTagEntry
		expected string
	}{
		{
			name: "simple",
			input: errorTagEntry{
				name: "auth",
				errors: []errorEntry{
					{
						name: "notFound",
						info: ErrorInfo{
							RPCStatus: 9,
							Code:      "0903",
							Message:   "Not found",
						},
					},
				},
			},
			expected: `
// ErrAuthNotFound ...
type ErrAuthNotFound liberrors.Error

// NewErrAuthNotFound ...
func NewErrAuthNotFound() *ErrAuthNotFound {
	return &ErrAuthNotFound{
		RPCStatus: 9,
		Code:      "0903",
		Message:   "Not found",
	}
}

// Err ...
func (e *ErrAuthNotFound) Err() error {
	return (*liberrors.Error)(e)
}

// AuthTag ...
type AuthTag struct {
	NotFound *ErrAuthNotFound
}

// Auth ...
var Auth = &AuthTag{
	NotFound: NewErrAuthNotFound(),
}
`,
		},
	}
	for _, e := range table {
		t.Run(e.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := generateErrorTag(e.input, &buf)
			assert.Nil(t, err)
			e.expected = strings.TrimSpace(e.expected)
			assert.Equal(t, e.expected, buf.String())
		})
	}
}
