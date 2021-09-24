package edsl

import (
	"fmt"
	"reflect"
)

type Structer interface {
	fmt.Stringer

	Field(name string, typ reflect.Type, tags ...Tag) Structer
	Of(val interface{}) Structer

	ForEach(fn func(name string, typ reflect.Type, tags []Tag))
	FieldNames() []string
	TagKeys(name string) []string
	TagValues(name, key string) []string

	Type() reflect.Type
	Value() reflect.Value
	Interface() interface{}
}
