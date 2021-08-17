package internal

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/anqur/procstruct/edsl"
)

type field struct {
	Name string
	Typ  reflect.Type
	Tags []edsl.Tag

	tagCache reflect.StructTag
}

func (f *field) Tag() reflect.StructTag {
	if f.tagCache == "" {
		var tags []string
		for _, tag := range f.Tags {
			tags = append(tags, tag.String())
		}
		f.tagCache = reflect.StructTag(strings.Join(tags, " "))
	}
	return f.tagCache
}

type Structer struct {
	Name   string
	Fields []*field
}

func (s Structer) String() string {
	buf := bytes.NewBufferString("type ")
	buf.WriteString(s.Name)
	buf.WriteString(" struct {\n")
	for _, field := range s.Fields {
		typ := field.Typ.Name()
		if typ == "" {
			panic(fmt.Errorf("field %q of type %q has no name", field.Name, s.Name))
		}
		line := []string{field.Name, typ, string(field.Tag())}
		buf.WriteByte('\t')
		buf.WriteString(strings.Join(line, " "))
		buf.WriteByte('\n')
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s Structer) Field(
	name string,
	typ reflect.Type,
	tags ...edsl.Tag,
) edsl.Structer {
	s.Fields = append(s.Fields, &field{
		Name: name,
		Typ:  typ,
		Tags: tags,
	})
	return s
}

func (s Structer) Type() reflect.Type {
	var reflFields []reflect.StructField
	for _, field := range s.Fields {
		reflFields = append(reflFields, reflect.StructField{
			Name: field.Name,
			Type: field.Typ,
			Tag:  field.Tag(),
		})
	}
	return reflect.StructOf(reflFields)
}

func (s Structer) Value() reflect.Value {
	return reflect.New(s.Type())
}

func (s Structer) Interface() interface{} {
	return s.Value().Interface()
}
