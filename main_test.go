package procstruct_test

import (
	"fmt"
	"go/ast"
	"reflect"
	"testing"

	"github.com/anqur/procstruct"
	"github.com/anqur/procstruct/edsl"
)

func ExampleStruct() {
	s := procstruct.Struct("Foo").
		Field("Data", reflect.TypeOf(0), "Data of Foo.")
	fmt.Println(s)
	// Output:
	// type Foo struct {
	// 	// Data of Foo.
	// 	Data int
	// }
}

func TestStruct(t *testing.T) {
	s := procstruct.
		Struct("Foo").
		Field(
			"Num",
			reflect.TypeOf(0),
			"Num",
			procstruct.Tag().
				Comma("json").
				Key("data").
				Key("omitempty").
				Nil().
				Nil(),
		).
		Field(
			"Str",
			reflect.PtrTo(reflect.TypeOf("")),
			"Str",
			procstruct.Tag().
				CommaEqSpace("binding").
				Key("required").
				Entry("oneof", "todo", "pending", "done").
				Nil().
				Nil(),
		).
		Field(
			"Float",
			reflect.SliceOf(reflect.TypeOf(float64(0))),
			"Float",
			procstruct.Tag().
				SemiColon("gorm").
				Key("not null").
				Entry("column", "float").
				Nil().
				Nil(),
		)

	fmt.Println(s)
	s.ForEach(func(name string, typ reflect.Type, tags []edsl.Tag) {
		fmt.Println(name, typ, tags)
	})
	fmt.Println(s.FieldNames())
	fmt.Println(s.TagKeys("json"))
	fmt.Println(s.TagValues("gorm", "column"))
	fmt.Println(s.Interface())
}

func TestRawTypedField(t *testing.T) {
	s := procstruct.Struct("Foo").
		RawTypedField(
			"Ctx",
			&ast.StarExpr{
				X: &ast.SelectorExpr{
					X:   ast.NewIdent("foo"),
					Sel: ast.NewIdent("Bar"),
				},
			},
			"Ctx is the context",
		)

	defer func() {
		if r := recover(); r == nil {
			t.Fatal(s)
		}
	}()

	fmt.Println(s)
	fmt.Println(s.Interface())
}

func TestFile(t *testing.T) {
	f := procstruct.File("foo").
		Header("DO NOT EDIT!").
		Imports(
			"go/parser",
		).
		Structs(
			procstruct.Struct("Foo").
				Field("A", reflect.TypeOf(0), "A").
				Field("B", reflect.TypeOf(""), "B"),
			procstruct.Struct("Bar").
				Field("A", reflect.TypeOf(float64(0)), "A").
				Field(
					"B",
					reflect.TypeOf(false),
					"B",
					procstruct.Tag().Comma("json"),
				),
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
				Field(
					"Results",
					reflect.TypeOf(nil),
					"Results",
					procstruct.Tag().Comma("json").Key("results"),
				),
		)

	fmt.Println(s)
}

func TestStructerOf(t *testing.T) {
	s := procstruct.Struct("Item").
		Field("Name", reflect.TypeOf(""), "Name").
		Of(&Data{})
	fmt.Println(s)
}
