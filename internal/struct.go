package internal

import (
	"bytes"
	"fmt"
	"go/ast"
	"reflect"
	"strings"

	"github.com/anqur/procstruct/edsl"
)

type field struct {
	Name   string
	Typ    reflect.Type
	RawTyp ast.Expr
	Tags   []edsl.Tag

	tagCache reflect.StructTag
}

func (f *field) Tag() reflect.StructTag {
	if len(f.Tags) == 0 {
		return ""
	}
	if f.tagCache == "" {
		var tags []string
		for _, tag := range f.Tags {
			tags = append(tags, tag.String())
		}
		s := strings.Join(tags, " ")
		f.tagCache = reflect.StructTag(s)
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
		var typ string
		if rawTyp := field.RawTyp; rawTyp != nil {
			typ = FormatTypeExpr(rawTyp)
		} else {
			typ = FormatType(field.Typ)
		}
		if typ == "" {
			panic(fmt.Errorf(
				"field %q of type %q has no name",
				field.Name,
				s.Name,
			))
		}
		line := []string{field.Name, typ}
		if tag := string(field.Tag()); tag != "" {
			line = append(line, fmt.Sprintf("`%s`", tag))
		}
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

func (s Structer) RawTypedField(
	name string,
	rawType ast.Expr,
	tags ...edsl.Tag,
) edsl.Structer {
	s.Fields = append(s.Fields, &field{
		Name:   name,
		RawTyp: rawType,
		Tags:   tags,
	})
	return s
}

func (s Structer) Of(val interface{}) edsl.Structer {
	typ := DerefStructType(val)
	var ret edsl.Structer = s
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		ret = ret.Field(
			field.Name,
			field.Type,
			ParseTag(string(field.Tag))...,
		)
	}
	return ret
}

func (s Structer) ForEach(
	fn func(name string, typ reflect.Type, tags []edsl.Tag),
) {
	for _, f := range s.Fields {
		fn(f.Name, f.Typ, f.Tags)
	}
}

func (s Structer) FieldNames() (ret []string) {
	for _, field := range s.Fields {
		ret = append(ret, field.Name)
	}
	return
}

func (s Structer) TagKeys(name string) (ret []string) {
	for _, field := range s.Fields {
		for _, tag := range field.Tags {
			if tag.Name() != name {
				continue
			}
			ret = append(ret, tag.FirstKey())
		}
	}
	return
}

func (s Structer) TagValues(name, key string) (ret []string) {
	for _, field := range s.Fields {
		for _, tag := range field.Tags {
			if tag.Name() != name {
				continue
			}
			ret = append(ret, tag.Value(key))
		}
	}
	return
}

func (s Structer) Type() reflect.Type {
	var reflFields []reflect.StructField
	for _, field := range s.Fields {
		if field.RawTyp != nil {
			panic(fmt.Errorf(
				"creating rtype from raw-typed field %q not supported",
				field.Name,
			))
		}
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
