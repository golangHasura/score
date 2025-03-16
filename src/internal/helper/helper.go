package helper

import "reflect"

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	// Note: Arrays are not nilable. Calling IsNil on an array value will panic.
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Chan, reflect.Slice, reflect.Func:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}
