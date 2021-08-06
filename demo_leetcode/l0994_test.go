package demo_leetcode

import (
	"fmt"
	"testing"
)

func TestOrangesRottingCase0(t *testing.T) {
	dimensionIntegerArray := strToTwoDimensionIntegerArray("[[0]]")
	rotting := orangesRotting(dimensionIntegerArray)
	fmt.Println(rotting)
}

func TestOrangesRottingCase1(t *testing.T) {
	dimensionIntegerArray := strToTwoDimensionIntegerArray("[[2,1,1],[1,1,0],[0,1,1]]")
	rotting := orangesRotting(dimensionIntegerArray)
	fmt.Println(rotting)
}

func TestOrangesRottingCase2(t *testing.T) {
	dimensionIntegerArray := strToTwoDimensionIntegerArray("[[1,2]]")
	rotting := orangesRotting(dimensionIntegerArray)
	fmt.Println(rotting)
}
