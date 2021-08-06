package demo_base

import (
	"fmt"
	"testing"
)

func TestAppend(t *testing.T) {
	result := make([]string, 0)
	result1 := append(result, "a")
	result2 := append(result, "b")
	result3 := append(result1, "c")
	result4 := append(result1, "a")
	fmt.Println(result)
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
}
