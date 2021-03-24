package internal

import "reflect"

func ParseType(v interface{}) reflect.Type {
	t := reflect.TypeOf(v)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}
