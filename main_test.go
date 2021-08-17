package procstruct_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/anqur/procstruct"
)

func TestStruct(t *testing.T) {
	s := procstruct.
		Struct("Foo").
		Field("Num", reflect.TypeOf(0), procstruct.
			Tag().Comma("json").
			Key("data").
			Key("omitempty").
			Nil().
			Nil()).
		Field("Str", reflect.TypeOf(""), procstruct.
			Tag().CommaEqSpace("binding").
			Key("required").
			Entry("oneof", "todo", "pending", "done").
			Nil().
			Nil()).
		Field("Float", reflect.TypeOf(float64(0)), procstruct.
			Tag().SemiComma("gorm").
			Key("not null").
			Entry("column", "float").
			Nil().
			Nil())
	fmt.Println(s)
}
