package edsl

import (
	"fmt"
	"reflect"
)

type Structer interface {
	fmt.Stringer

	Field(name string, kind reflect.Kind, tags ...Tag) Structer

	Type() reflect.Type
	Value() reflect.Value
}
