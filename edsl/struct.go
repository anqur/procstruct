package edsl

import (
	"fmt"
	"reflect"
)

type Structer interface {
	fmt.Stringer

	Field(name string, typ reflect.Type, tags ...Tag) Structer

	FieldNames() []string
	TagKeys(name string) []string
	TagValues(name, key string) []string

	Type() reflect.Type
	Value() reflect.Value
	Interface() interface{}
}
