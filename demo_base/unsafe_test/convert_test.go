package unsafe_test

import (
	"testing"
	"unsafe"
)

func TestConvert(t *testing.T) {
	// -1 二进制表达 11111111
	var i int8 = -1
	// -1 二进制表达 11111111 11111111
	var j = int16(i)
	println(i, j)
	var k uint8 = *(*uint8)(unsafe.Pointer(&i))
	println(k)
	// 255 is the uint8 value for the binary 11111111
}
