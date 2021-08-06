package demo_reedsolomon

import (
	"fmt"
	"github.com/klauspost/reedsolomon"
	"testing"
)

func TestReedSolomon(t *testing.T) {
	enc, err := reedsolomon.New(10, 3)
	if err != nil {
		panic(err)
	}
	fmt.Println(enc)
}
