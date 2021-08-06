package reflecttest

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPointer(t *testing.T) {
	var a = &turbo{}
	types := reflect.TypeOf(a)
	if types.Kind() == reflect.Ptr {
		elem := types.Elem()
		fmt.Println(elem.Name(), elem.Kind())
	}
}
