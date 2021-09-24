package internal

import (
	"fmt"
	"reflect"
)

func DerefStructType(val interface{}) reflect.Type {
	v := reflect.ValueOf(val)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() == reflect.Struct {
		return v.Type()
	}
	panic(fmt.Errorf("expected a struct, found %q", v.Kind()))
}
