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
	imports []string
}

func (f Filer) String() string {
	buf := new(bytes.Buffer)

	if f.header != "" {
		buf.WriteString("// ")
		buf.WriteString(f.header)
		buf.WriteByte('\n')
	}

	buf.WriteString("package ")
	buf.WriteString(f.PkgName)

	if len(f.imports) > 0 {
		buf.WriteString("\n\nimport (\n")
		for _, pkg := range f.imports {
			buf.WriteByte('\t')
			buf.WriteString(fmt.Sprintf("%q", pkg))
			buf.WriteByte('\n')
		}
		buf.WriteString(")\n")
	}

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

func (f Filer) Imports(packages ...string) edsl.Filer {
	f.imports = append(f.imports, packages...)
	return f
}

func (f Filer) Structs(structs ...edsl.Structer) edsl.Filer {
	f.structs = append(f.structs, structs...)
	return f
}
