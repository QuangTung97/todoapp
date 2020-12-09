package generate

import "sort"

// ErrorInfo error config
type ErrorInfo struct {
	RPCStatus uint32            `yaml:"rpcStatus"`
	Code      string            `yaml:"code"`
	Message   string            `yaml:"message"`
	Details   map[string]string `yaml:"details"`
}

// ErrorMap map of errors
type ErrorMap map[string]ErrorInfo

// Detail Entry

type detailEntry struct {
	field     string
	fieldType string
}

type sortDetailEntry []detailEntry

var _ sort.Interface = sortDetailEntry(nil)

func (s sortDetailEntry) Len() int {
	return len(s)
}

func (s sortDetailEntry) Less(i, j int) bool {
	return s[i].field < s[j].field
}

func (s sortDetailEntry) Swap(i, j int) {
	s[j], s[i] = s[i], s[j]
}

// Error Entry

type errorEntry struct {
	name string
	info ErrorInfo
}

type sortErrorEntry []errorEntry

var _ sort.Interface = sortErrorEntry(nil)

func (s sortErrorEntry) Len() int {
	return len(s)
}

func (s sortErrorEntry) Less(i, j int) bool {
	return s[i].name < s[j].name
}

func (s sortErrorEntry) Swap(i, j int) {
	s[j], s[i] = s[i], s[j]
}

// Error Tag Entry

type errorTagEntry struct {
	name   string
	errors []errorEntry
}

type sortErrorTagEntry []errorTagEntry

var _ sort.Interface = sortErrorTagEntry(nil)

func (s sortErrorTagEntry) Len() int {
	return len(s)
}

func (s sortErrorTagEntry) Less(i, j int) bool {
	return s[i].name < s[j].name
}

func (s sortErrorTagEntry) Swap(i, j int) {
	s[j], s[i] = s[i], s[j]
}

// Functions

func errorMapToList(m ErrorMap) []errorEntry {
	result := make([]errorEntry, 0, len(m))
	for name, info := range m {
		result = append(result, errorEntry{
			name: name,
			info: info,
		})
	}
	sort.Sort(sortErrorEntry(result))
	return result
}

func tagMapToList(tags map[string]ErrorMap) []errorTagEntry {
	result := make([]errorTagEntry, 0, len(tags))
	for name, m := range tags {
		result = append(result, errorTagEntry{
			name:   name,
			errors: errorMapToList(m),
		})
	}
	sort.Sort(sortErrorTagEntry(result))
	return result
}
