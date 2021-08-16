package internal

import (
	"reflect"

	"github.com/anqur/procstruct/edsl"
)

type Structer struct {
	Name string
}

func (s Structer) String() string {
	panic("implement me")
}

func (s Structer) Field(
	name string,
	kind reflect.Kind,
	tags ...edsl.Tag,
) edsl.Structer {
	panic("implement me")
}

func (s Structer) Type() reflect.Type {
	panic("implement me")
}

func (s Structer) Value() reflect.Value {
	panic("implement me")
}
