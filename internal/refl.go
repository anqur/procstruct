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

func FormatType(typ reflect.Type) string {
	if typ == nil {
		return "interface{}"
	}
	switch typ.Kind() {
	case reflect.Ptr:
		return "*" + FormatType(typ.Elem())
	case reflect.Slice:
		return "[]" + FormatType(typ.Elem())
	}
	return typ.Name()
}
