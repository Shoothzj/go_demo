package demo_leetcode

import (
	"fmt"
	"testing"
)

func TestCanPlaceFlowersCase1(t *testing.T) {
	canPlaceFlowers := canPlaceFlowers(strToIntegerArray("[1,0,0,0,0,1]"), 2)
	fmt.Println(canPlaceFlowers)
}

func TestCanPlaceFlowersCase2(t *testing.T) {
	canPlaceFlowers := canPlaceFlowers(strToIntegerArray("[1,0,1,0,1,0,1]"), 1)
	fmt.Println(canPlaceFlowers)
}

func TestCanPlaceFlowersCase3(t *testing.T) {
	canPlaceFlowers := canPlaceFlowers(strToIntegerArray("[1,0,0,0,1,0,0]"), 2)
	fmt.Println(canPlaceFlowers)
}

func TestCanPlaceFlowersCase4(t *testing.T) {
	canPlaceFlowers := canPlaceFlowers(strToIntegerArray("[0,1,0]"), 1)
	if !canPlaceFlowers {
		t.Error("error")
	}
}
