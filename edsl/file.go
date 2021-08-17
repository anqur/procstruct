package edsl

import "fmt"

type Filer interface {
	fmt.Stringer

	Header(text string) Filer
	Structs(structs ...Structer) Filer
}
