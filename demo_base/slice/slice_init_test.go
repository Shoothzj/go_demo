package slice

import (
	"fmt"
	"testing"
)

func Test_SliceInit1(t *testing.T) {
	aux := make([]int, 5)
	aux[0] = 1
	aux[1] = 2
	aux[2] = 3
	aux[3] = 4
	aux[4] = 5
	fmt.Println(aux)
}

func Test_SliceInit2(t *testing.T) {
	aux := [5]int{1, 2, 3, 4, 5}
	fmt.Println(aux)
}

func Test_SliceInit3(t *testing.T) {
	aux := [...]int{1, 2, 3, 4, 5}
	fmt.Println(aux)
}

func Test_SliceInit4(t *testing.T) {
	aux := [...]int{1: 2, 3: 4, 5}
	fmt.Println(aux)
}
