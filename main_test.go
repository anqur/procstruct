package procstruct_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/anqur/procstruct"
)

func TestNew(t *testing.T) {
	s := procstruct.
		Struct("Foo").
		Field("Data", reflect.Int)
	fmt.Println(s)
}
