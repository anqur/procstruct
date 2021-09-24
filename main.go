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

func RegisterTagStyle(name string, style edsl.TagStyle) {
	internal.TagStyles[name] = style
}

func SetDefaultTagStyle(style edsl.TagStyle) {
	internal.TagStyleDefault = style
}

func Of(val interface{}) edsl.Structer {
	return Struct(internal.DerefStructType(val).Name()).Of(val)
}

func init() {
	SetDefaultTagStyle(edsl.TagStyleComma)

	RegisterTagStyle("uri", edsl.TagStyleComma)
	RegisterTagStyle("form", edsl.TagStyleComma)
	RegisterTagStyle("json", edsl.TagStyleComma)
	RegisterTagStyle("header", edsl.TagStyleComma)
	RegisterTagStyle("xml", edsl.TagStyleComma)
	RegisterTagStyle("yaml", edsl.TagStyleComma)
	RegisterTagStyle("toml", edsl.TagStyleComma)

	RegisterTagStyle("binding", edsl.TagStyleCommaEqSpace)
	RegisterTagStyle("validate", edsl.TagStyleCommaEqSpace)

	RegisterTagStyle("gorm", edsl.TagStyleSemiColon)
}
