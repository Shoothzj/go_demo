package slice

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	a := []int{1, 2, 3, 4}
	result := make([]int, 0)
	result = append(result, a[0:]...)
	fmt.Println(result)
}
