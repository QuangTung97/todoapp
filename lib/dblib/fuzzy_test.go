package dblib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFuzzyMatchNormal(t *testing.T) {
	table := []struct {
		name       string
		input      []registeredQuery
		filter     string
		expected   []registeredQuery
		highlights []string
	}{
		{
			name: "normal",
			input: []registeredQuery{
				{
					query: "some string test",
				},
				{
					query: "do nothing",
				},
				{
					query: "something testing",
				},
			},
			filter: "so test",
			expected: []registeredQuery{
				{
					query: "something testing",
				},
				{
					query: "some string test",
				},
			},
			highlights: []string{
				"'s`'o`mething' `'t`'e`'s`'t`ing",
				"'s`'o`me' `string 't`'e`'s`'t`",
			},
		},
	}

	for _, e := range table {
		t.Run(e.name, func(t *testing.T) {
			result, highlights := fuzzyMatchNormal(e.input, e.filter, "'", "`")
			assert.Equal(t, e.expected, result)
			assert.Equal(t, e.highlights, highlights)
		})
	}
}
