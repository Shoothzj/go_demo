package demo_leetcode

import (
	"fmt"
	"testing"
)

func Test_partitionCase1(t *testing.T) {
	result := partition("aab")
	fmt.Println(result)
}

func Test_partitionCase2(t *testing.T) {
	result := partition("a")
	fmt.Println(result)
}

func Test_partitionCase3(t *testing.T) {
	result := partition("ababbbabbaba")
	fmt.Println(result)
}
