package demo_leetcode

import (
	"fmt"
	"testing"
)

func TestStrToTwoDimensionIntegerArray(t *testing.T) {
	str := "[[2,1,1],[1,1,0],[0,1,1]]"
	integerArray := strToTwoDimensionIntegerArray(str)
	fmt.Println(integerArray)
}

func TestStrToTwoDimensionIntegerArrayCase2(t *testing.T) {
	str := "[[1,2,3],[4,5,6],[7,8,9]]"
	integerArray := strToTwoDimensionIntegerArray(str)
	fmt.Println(integerArray)
}
