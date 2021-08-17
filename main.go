package procstruct

import (
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

func Of(structs ...interface{}) edsl.Structer {
	return nil
}
