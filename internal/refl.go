package internal

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
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
	case reflect.Map:
		return fmt.Sprintf(
			"map[%s]%s",
			FormatType(typ.Key()),
			FormatType(typ.Elem()),
		)
	}
	return typ.Name()
}

func FormatTypeExpr(expr ast.Expr) string {
	buf := new(bytes.Buffer)
	if err := printer.Fprint(buf, token.NewFileSet(), expr); err != nil {
		panic(err)
	}
	return buf.String()
}
