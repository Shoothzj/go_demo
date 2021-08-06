package demo_base

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	var xx uint32 = 0x80000000
	fmt.Println(int32(xx))
}
