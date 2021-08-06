package demo_base

import (
	"fmt"
	"testing"
)

func TestNilSliceAppend(t *testing.T) {
	var slice []int = nil
	ints := append(slice, 1)
	fmt.Println(ints)
}

func TestMakeSlice(t *testing.T) {
	slice := make([]int, 10, 15)
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
}

func TestSliceLength(t *testing.T) {
	strings := make([]int32, 0)
	strings = append(strings, 1)
	strings = append(strings, 2)
	fmt.Println(len(strings))
	fmt.Println(strings)
}
