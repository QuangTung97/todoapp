package generate

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"sort"
	"strings"
)

func generateNewErrorFunc(errName string, info ErrorInfo, writer io.Writer) error {
	code := `
// New%s ...
func New%s() *%s {
	return &%s{
		RPCStatus: %d,
		Code:      %q,
		Message:   %q,
	}
}
`
	code = strings.TrimSpace(code)
	code = fmt.Sprintf(code, errName, errName, errName, errName, info.RPCStatus, info.Code, info.Message)
	_, err := writer.Write([]byte(code))
	return err
}

func generateWithMethod(errName string, field string, fieldType string, writer io.Writer) error {
	method := "With" + strings.Title(field)

	code := `
// %s ...
func (e *%s) %s(value %s) *%s {
	err := (*liberrors.Error)(e)
	return (*%s)(err.WithDetail(%q, value))
}
`
	code = fmt.Sprintf(code, method, errName, method, fieldType, errName, errName, field)
	code = strings.TrimSpace(code)
	_, err := writer.Write([]byte(code))
	return err
}

func generateErrorEntry(tagName string, e errorEntry, writer io.Writer) error {
	name := strings.Title(e.name)
	name = "Err" + tagName + name
	typeCode := fmt.Sprintf("// %s ...\ntype %s liberrors.Error", name, name)
	_, err := writer.Write([]byte(typeCode))
	if err != nil {
		return err
	}

	_, err = writer.Write([]byte("\n\n"))
	if err != nil {
		return err
	}

	err = generateNewErrorFunc(name, e.info, writer)
	if err != nil {
		return err
	}

	details := make([]detailEntry, 0, len(e.info.Details))
	for field, fieldType := range e.info.Details {
		details = append(details, detailEntry{
			field:     field,
			fieldType: fieldType,
		})
	}
	sort.Sort(sortDetailEntry(details))

	for _, d := range details {
		_, err := writer.Write([]byte("\n\n"))
		if err != nil {
			return err
		}

		err = generateWithMethod(name, d.field, d.fieldType, writer)
		if err != nil {
			return err
		}
	}

	return nil
}

func generateErrorTag(tag errorTagEntry, writer io.Writer) error {
	tagName := strings.Title(tag.name)
	for _, errEntry := range tag.errors {
		err := generateErrorEntry(tagName, errEntry, writer)
		if err != nil {
			return err
		}

		_, err = writer.Write([]byte("\n\n"))
		if err != nil {
			return err
		}
	}

	structCode := fmt.Sprintf("// %sTag ...\ntype %sTag struct {", tagName, tagName)
	_, err := writer.Write([]byte(structCode))
	if err != nil {
		return err
	}

	for _, errEntry := range tag.errors {
		name := strings.Title(errEntry.name)
		errName := "Err" + tagName + name

		field := fmt.Sprintf("\n\t%s *%s", name, errName)
		_, err := writer.Write([]byte(field))
		if err != nil {
			return err
		}
	}

	_, err = writer.Write([]byte("\n}\n\n"))
	if err != nil {
		return err
	}

	initCode := fmt.Sprintf("// %s ...\nvar %s = &%sTag{", tagName, tagName, tagName)
	_, err = writer.Write([]byte(initCode))
	if err != nil {
		return err
	}

	for _, errEntry := range tag.errors {
		name := strings.Title(errEntry.name)
		errName := "Err" + tagName + name

		field := fmt.Sprintf("\n\t%s: New%s(),", name, errName)
		_, err := writer.Write([]byte(field))
		if err != nil {
			return err
		}
	}

	_, err = writer.Write([]byte("\n}"))
	if err != nil {
		return err
	}

	return nil
}

// Generate generates file from yml
func Generate(tags map[string]ErrorMap, output io.Writer) error {
	var buf bytes.Buffer
	_, err := buf.Write([]byte(imports))
	if err != nil {
		return err
	}

	tagList := tagMapToList(tags)
	for _, tag := range tagList {
		_, err = buf.Write([]byte("\n\n"))
		if err != nil {
			return err
		}

		err := generateErrorTag(tag, &buf)
		if err != nil {
			return err
		}
	}

	_, err = buf.Write([]byte("\n"))
	if err != nil {
		return err
	}

	result, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	_, err = output.Write(result)
	return err
}
