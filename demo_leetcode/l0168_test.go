package demo_leetcode

import (
	"fmt"
	"testing"
)

func Test_convertToTitleCase1(t *testing.T) {
	title := convertToTitle(1)
	fmt.Println(title)
}

func Test_convertToTitleCase2(t *testing.T) {
	title := convertToTitle(701)
	fmt.Println(title)
}

func Test_convertToTitleCase3(t *testing.T) {
	title := convertToTitle(2147483647)
	fmt.Println(title)
}
