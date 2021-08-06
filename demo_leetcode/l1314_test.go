package demo_leetcode

import (
	"fmt"
	"testing"
)

func Test_matrixBlockSumCase1(t *testing.T) {
	result := matrixBlockSum(strToTwoDimensionIntegerArray("[[1,2,3],[4,5,6],[7,8,9]]"), 1)
	fmt.Println(result)
}
