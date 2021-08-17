package edsl

import (
	"fmt"
	"reflect"
)

type Structer interface {
	fmt.Stringer

	Field(name string, typ reflect.Type, tags ...Tag) Structer

	Type() reflect.Type
	Value() reflect.Value
	Interface() interface{}
}
