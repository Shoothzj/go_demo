package demo_base

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

const invalidUtf8 = "ab\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

func TestJudgeUtf8(t *testing.T) {
	byteArray := []byte(invalidUtf8)
	index := 0
	for {
		decodeLastRune, size := utf8.DecodeRune(byteArray[index:])
		if decodeLastRune == utf8.RuneError {
			fmt.Println("rune error")
			break
		}
		fmt.Println("valid rune is ", decodeLastRune)
		index += size
	}
}
