package demo_leetcode

import (
	"fmt"
	"testing"
)

func Test_validMountainArrayCase1(t *testing.T) {
	array := validMountainArray(strToIntegerArray("[0,3,2,1]"))
	fmt.Println(array)
}

func Test_validMountainArrayCase2(t *testing.T) {
	array := validMountainArray(strToIntegerArray("[1,1,1,1,1,1,1,2,1]"))
	fmt.Println(array)
}
