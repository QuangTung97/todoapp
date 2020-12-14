package log

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHideStringFields(t *testing.T) {
	type exam struct {
		Name   string
		APIKey string
		Age    int
	}

	type inner struct {
		Password string
	}

	type nestedExam struct {
		APIKey string
		Age    int
		Inner  *inner
	}

	table := []struct {
		name     string
		value    interface{}
		expected interface{}
		fields   []string
		ok       bool
	}{
		{
			name:     "int",
			value:    int(100),
			expected: nil,
			ok:       false,
		},
		{
			name: "struct",
			value: exam{
				Name:   "some name",
				APIKey: "12345",
				Age:    12,
			},
			expected: &exam{
				Name:   "some name",
				APIKey: "12345",
				Age:    12,
			},
			ok: true,
		},
		{
			name: "pointer",
			value: &exam{
				Name:   "some name",
				APIKey: "12345",
				Age:    12,
			},
			expected: &exam{
				Name:   "some name",
				APIKey: "12345",
				Age:    12,
			},
			ok: true,
		},
		{
			name: "hide-name",
			value: &exam{
				Name:   "some name",
				APIKey: "12345",
				Age:    12,
			},
			fields: []string{"Name"},
			expected: &exam{
				Name:   maskedValue,
				APIKey: "12345",
				Age:    12,
			},
			ok: true,
		},
		{
			name: "hide-name-and-api-key",
			value: &exam{
				Name:   "some name",
				APIKey: "12345",
				Age:    12,
			},
			fields: []string{"Name", "APIKey"},
			expected: &exam{
				Name:   maskedValue,
				APIKey: maskedValue,
				Age:    12,
			},
			ok: true,
		},
		{
			name: "hide-age",
			value: &exam{
				Name:   "some name",
				APIKey: "12345",
				Age:    12,
			},
			fields:   []string{"Age"},
			expected: nil,
			ok:       false,
		},
		{
			name: "hide-age",
			value: &nestedExam{
				APIKey: "12345",
				Age:    12,
				Inner: &inner{
					Password: "aabbcc",
				},
			},
			fields: []string{"APIKey", "Password"},
			expected: &nestedExam{
				APIKey: maskedValue,
				Age:    12,
				Inner: &inner{
					Password: "aabbcc",
				},
			},
			ok: true,
		},
	}

	for _, e := range table {
		t.Run(e.name, func(t *testing.T) {
			result, ok := hideStringFields(e.value, e.fields...)
			assert.Equal(t, e.ok, ok)
			assert.Equal(t, e.expected, result)
		})
	}
}

func BenchmarkHideStringFields(b *testing.B) {
	type exam struct {
		APIKey string
		Name   string
		Age    int
	}
	for n := 0; n < b.N; n++ {
		e := &exam{
			APIKey: "1233333333abdfsafddfaa",
			Name:   "some name",
			Age:    12,
		}
		hideStringFields(e, "APIKey")
	}
}
