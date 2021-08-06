package demo_leetcode

import (
	"fmt"
	"testing"
)

func TestThirdMaxCase1(t *testing.T) {
	max := thirdMax(strToIntegerArray("[3,2,1]"))
	fmt.Println(max)
}

func TestThirdMaxCase2(t *testing.T) {
	max := thirdMax(strToIntegerArray("[1,2]"))
	fmt.Println(max)
}

func TestThirdMaxCase3(t *testing.T) {
	max := thirdMax(strToIntegerArray("[2,2,3,1]"))
	fmt.Println(max)
}
