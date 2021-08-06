package demo_base

import (
	"fmt"
	"testing"
)

func TestAppendWithoutMake(t *testing.T) {
	var array []string
	strings := append(array, "1111")
	fmt.Println(strings)
}
