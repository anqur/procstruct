package internal

import (
	"github.com/anqur/procstruct/edsl"
)

type tagStyles map[string]edsl.TagStyle

var (
	TagStyles       = make(tagStyles)
	TagStyleDefault edsl.TagStyle
)

func (m tagStyles) Get(name string) edsl.TagStyle {
	if s, ok := m[name]; ok {
		return s
	}
	return TagStyleDefault
}
