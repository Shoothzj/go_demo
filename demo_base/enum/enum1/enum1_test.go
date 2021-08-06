package enum1

import (
	"fmt"
	"testing"
)

type Hero int

const (
	IRONMAN   Hero = 0
	SPIDERMAN Hero = 1
	BATMAN    Hero = 2
)

func TestName(t *testing.T) {
	ironman := IRONMAN
	fmt.Println(ironman == SPIDERMAN)
	fmt.Println(ironman == IRONMAN)
}
