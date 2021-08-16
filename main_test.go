package procstruct_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/anqur/procstruct"
)

func TestNew(t *testing.T) {
	s1 := procstruct.
		Struct("Foo").
		Field(
			"Data",
			reflect.Int,
			procstruct.
				Tag().Comma("json").
				Key("name").
				Key("omitempty").
				Nil(),
		)
	fmt.Println(s1)
}
