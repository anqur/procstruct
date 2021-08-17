package internal

import (
	"bytes"
	"fmt"
	"go/format"

	"github.com/anqur/procstruct/edsl"
)

type Filer struct {
	PkgName string

	header  string
	structs []edsl.Structer
}

func (f Filer) String() string {
	buf := bytes.NewBufferString("package ")
	buf.WriteString(f.PkgName)
	for _, st := range f.structs {
		buf.WriteString("\n\n")
		buf.WriteString(st.String())
	}
	data, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}
	return string(data)
}

func (f Filer) Header(text string) edsl.Filer {
	if f.header == "" {
		f.header = fmt.Sprintf("Package %s %s", f.PkgName, text)
	}
	return f
}

func (f Filer) Structs(structs ...edsl.Structer) edsl.Filer {
	f.structs = append(f.structs, structs...)
	return f
}
