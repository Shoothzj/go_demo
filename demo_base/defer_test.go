package demo_base

import (
	"fmt"
	"testing"
)

func TestDeferLoop(t *testing.T) {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

func TestDeferResultVal(t *testing.T) {
	fmt.Println(deferResultVal())
}

func deferResultVal() (i int) {
	defer func() { i++ }()
	return 1
}

func TestPanicWhereDefer(t *testing.T) {
	panicWhereDefer()
}

func panicWhereDefer() {
	defer func() {
		panic("666")
	}()
	fmt.Println("abc")
	panicWhereDefer2()
}

func panicWhereDefer2() {
	defer func() {
		panic("555")
	}()
	panicWhereDefer3()
}

func panicWhereDefer3() {
	defer func() {
		panic("444")
	}()
	fmt.Println("abc")
}
