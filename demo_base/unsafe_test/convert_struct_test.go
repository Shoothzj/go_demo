package unsafe

import (
	"testing"
	"unsafe"
)

type A struct {
	A int8
	B string
	C float32
}

type B struct {
	D int8
	E string
	F float32
}

func TestStructConvert(t *testing.T) {
	a := A{A: 1, B: `foo`, C: 1.23}
	b := *(*B)(unsafe.Pointer(&a))
	println(b.D, b.E, b.F) // 1 foo 1.23
}
