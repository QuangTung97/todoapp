package dblib

import (
	"github.com/sahilm/fuzzy"
	"strings"
)

func contains(num int, list []int) bool {
	for _, i := range list {
		if num == i {
			return true
		}
	}
	return false
}

type sourceRegisteredQuery []registeredQuery

var _ fuzzy.Source = sourceRegisteredQuery(nil)

func (s sourceRegisteredQuery) String(i int) string {
	return s[i].query
}

func (s sourceRegisteredQuery) Len() int {
	return len(s)
}

func fuzzyMatchNormal(list []registeredQuery, filter string, colorHighlight string, colorNone string,
) ([]registeredQuery, []string) {
	matches := fuzzy.FindFrom(filter, sourceRegisteredQuery(list))

	result := make([]registeredQuery, 0, len(matches))
	highlights := make([]string, 0, len(matches))

	var builder strings.Builder
	for _, m := range matches {
		result = append(result, list[m.Index])

		builder.Reset()
		for i := 0; i < len(m.Str); i++ {
			ch := string(m.Str[i])
			if contains(i, m.MatchedIndexes) {
				builder.WriteString(colorHighlight)
				builder.WriteString(ch)
				builder.WriteString(colorNone)

			} else {
				builder.WriteString(ch)
			}
		}
		highlights = append(highlights, builder.String())
	}

	return result, highlights
}

type sourceRegisteredNamedQuery []registeredNamedQuery

var _ fuzzy.Source = sourceRegisteredNamedQuery(nil)

func (s sourceRegisteredNamedQuery) String(i int) string {
	return s[i].query
}

func (s sourceRegisteredNamedQuery) Len() int {
	return len(s)
}

func fuzzyMatchNamed(list []registeredNamedQuery, filter string, colorHighlight string, colorNone string,
) ([]registeredNamedQuery, []string) {
	matches := fuzzy.FindFrom(filter, sourceRegisteredNamedQuery(list))

	result := make([]registeredNamedQuery, 0, len(matches))
	highlights := make([]string, 0, len(matches))

	var builder strings.Builder
	for _, m := range matches {
		result = append(result, list[m.Index])

		builder.Reset()
		for i := 0; i < len(m.Str); i++ {
			ch := string(m.Str[i])
			if contains(i, m.MatchedIndexes) {
				builder.WriteString(colorHighlight)
				builder.WriteString(ch)
				builder.WriteString(colorNone)

			} else {
				builder.WriteString(ch)
			}
		}
		highlights = append(highlights, builder.String())
	}

	return result, highlights
}
