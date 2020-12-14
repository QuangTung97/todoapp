package log

import (
	"reflect"
)

const maskedValue = "****"

func shallowCopyStruct(e interface{}) (interface{}, bool) {
	v := reflect.ValueOf(e)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, false
	}

	newVal := reflect.New(v.Type())
	ptr := newVal.Elem()

	n := v.NumField()
	for i := 0; i < n; i++ {
		if ptr.Field(i).CanSet() {
			ptr.Field(i).Set(v.Field(i))
		}
	}
	return newVal.Interface(), true
}

func existInList(v string, list []string) bool {
	for _, e := range list {
		if e == v {
			return true
		}
	}
	return false
}

func hideStringFields(e interface{}, fields ...string) (interface{}, bool) {
	e, ok := shallowCopyStruct(e)
	if !ok {
		return nil, false
	}

	ptrVal := reflect.ValueOf(e)
	v := ptrVal.Elem()
	t := v.Type()

	n := t.NumField()
	for i := 0; i < n; i++ {
		name := t.Field(i).Name
		if existInList(name, fields) {
			field := v.Field(i)
			if field.Kind() != reflect.String {
				return nil, false
			}
			field.SetString(maskedValue)
		}
	}

	return ptrVal.Interface(), true
}
