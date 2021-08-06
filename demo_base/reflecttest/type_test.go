package reflecttest

import (
	"fmt"
	"reflect"
	"testing"
)

func TestType(t *testing.T) {
	var a = turbo{}
	types := reflect.TypeOf(a)
	fmt.Println(types.Name(), types.Kind())
}
