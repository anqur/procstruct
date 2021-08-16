package edsl

import (
	"fmt"
	"reflect"
)

type Struct interface {
	fmt.Stringer

	Field(name string, kind reflect.Kind, tags ...Tag) Struct

	Type() reflect.Type
	Value() reflect.Value
}
