package procstruct

import (
	"fmt"
	"reflect"

	"github.com/anqur/procstruct/edsl"
	"github.com/anqur/procstruct/internal"
)

func File(pkg string) edsl.Filer {
	return internal.Filer{PkgName: pkg}
}

func Struct(name string) edsl.Structer {
	return internal.Structer{Name: name}
}

func Tag() edsl.Tagger {
	return internal.Tagger{}
}

func Of(val interface{}) edsl.Structer {
	v := reflect.ValueOf(val)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		panic(fmt.Errorf("expected a struct, found %q", v.Kind()))
	}
	typ := v.Type()
	s := Struct(typ.Name())
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		s = s.Field(field.Name, field.Type, Tag().AsIs(string(field.Tag)))
	}
	return s
}
