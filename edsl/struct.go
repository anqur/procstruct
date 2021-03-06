package edsl

import (
	"fmt"
	"go/ast"
	"reflect"
)

type Structer interface {
	fmt.Stringer

	Header(text string) Structer

	Field(name string, typ reflect.Type, comment string, tags ...Tag) Structer
	RawTypedField(name string, rawType ast.Expr, comment string, tags ...Tag) Structer
	Of(val interface{}) Structer

	ForEach(fn func(name string, typ reflect.Type, tags []Tag))
	FieldNames() []string
	TagKeys(name string) []string
	TagValues(name, key string) []string

	Type() reflect.Type
	Value() reflect.Value
	Interface() interface{}
}
