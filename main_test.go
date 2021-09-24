package procstruct_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/anqur/procstruct"
	"github.com/anqur/procstruct/edsl"
)

func ExampleStruct() {
	s := procstruct.Struct("Foo").
		Field("Data", reflect.TypeOf(0))
	fmt.Println(s)
	// Output:
	// type Foo struct {
	// 	Data int
	// }
}

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
			Tag().SemiColon("gorm").
			Key("not null").
			Entry("column", "float").
			Nil().
			Nil())

	fmt.Println(s)
	s.ForEach(func(name string, typ reflect.Type, tags []edsl.Tag) {
		fmt.Println(name, typ, tags)
	})
	fmt.Println(s.FieldNames())
	fmt.Println(s.TagKeys("json"))
	fmt.Println(s.TagValues("gorm", "column"))
	fmt.Println(s.Interface())
}

func TestFile(t *testing.T) {
	f := procstruct.File("foo").
		Header("DO NOT EDIT!").
		Structs(
			procstruct.Struct("Foo").
				Field("A", reflect.TypeOf(0)).
				Field("B", reflect.TypeOf("")),
			procstruct.Struct("Bar").
				Field("A", reflect.TypeOf(float64(0))).
				Field("B", reflect.TypeOf(false), procstruct.
					Tag().Comma("json")),
		)

	fmt.Println(f)
}

type Data struct {
	Total int `json:"total,,,omitempty,,," validate:"required,oneof=1 2,,," gorm:"column:total;not null;;;;"`
}

func TestOf(t *testing.T) {
	s := procstruct.File("foo").
		Structs(
			procstruct.Of(&Data{}).
				Field("Results", reflect.TypeOf(nil), procstruct.
					Tag().Comma("json").Key("results")),
		)

	fmt.Println(s)
}

func TestStructerOf(t *testing.T) {
	s := procstruct.Struct("Item").
		Field("Name", reflect.TypeOf("")).
		Of(&Data{})
	fmt.Println(s)
}
