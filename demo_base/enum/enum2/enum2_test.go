package enum2

import (
	"fmt"
	"testing"
)

type Hero int

const (
	IRONMAN Hero = iota
	SPIDERMAN
	BATMAN
)

func TestName(t *testing.T) {
	ironman := IRONMAN
	fmt.Println(ironman == SPIDERMAN)
	fmt.Println(ironman == IRONMAN)
}
