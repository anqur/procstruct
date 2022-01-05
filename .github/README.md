# procstruct

> *Procedural structs.*

Define a Go struct in a procedural way, a.k.a. eDSL library for Go structs.

```bash
$ go get github.com/anqur/procstruct
```

## Examples

Basic example, just build a struct:

```go
package main

import (
	"fmt"
	"reflect"

	"github.com/anqur/procstruct"
)

func main() {
	s := procstruct.
		Struct("Foo").
		Field("Data", reflect.TypeOf(0), "Data of Foo.")
	fmt.Println(s)
	// Output:
	// type Foo struct {
	//  // Data of Foo.
	//  Data int
	// }
}
```

More complicated example, if you want to achieve some meta-programming in your project, you could create a file called
`tools.go` and make a build constraint `//+build tools` to avoid its compilation. Then put a `go:generate` to run the
tool script.

For instance, you have this struct:

```go
package item

type Item struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}
```

And you have some service to sort these fields (e.g. *name, ascending order* or *price, descending order*), with some
field validation, then you may need this struct:

```go
package item

type ItemSorting struct {
	Key   string `json:"key" validate:"required,oneof=name price"`
	Order string `json:"order" validate:"required,oneof=asc desc"`
}
```

We could see that the type `ItemSorting` could be easily generated from type `Item`, by knowing that what kind of fields
and related tags our real business model `Item` already has, so with this tool you could conveniently create your own
meta-programming/code-generation scripts for automating your business models. For this example, please check out
the [examples](../examples) directory.

## License

MIT
