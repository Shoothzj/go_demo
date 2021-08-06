package demo_base

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type ReflectBase struct {
	xx string
}

type ReflectTest struct {
	id      string
	version int
	ReflectBase
}

func TestReflect(t *testing.T) {
	typeOf := reflect.TypeOf(&ReflectTest{id: "12"})
	fmt.Println(typeOf.Elem().String())
	value := reflect.New(typeOf.Elem())
	pointer := unsafe.Pointer(value.Pointer())
	aux := (*ReflectTest)(pointer)
	fmt.Println(aux)
	fmt.Println(pointer)
}
