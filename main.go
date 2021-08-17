package procstruct

import (
	"github.com/anqur/procstruct/edsl"
	"github.com/anqur/procstruct/internal"
)

func File(pkg string) edsl.Filer {
	return nil
}

func Struct(name string) edsl.Structer {
	return internal.Structer{Name: name}
}

func Tag() edsl.Tagger {
	return nil
}

func Of(structs ...interface{}) edsl.Structer {
	return nil
}
