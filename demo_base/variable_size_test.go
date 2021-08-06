package demo_base

import (
	"fmt"
	"testing"
	"time"
	"unsafe"
)

func TestSize(t *testing.T) {
	{
		var x byte = 0
		fmt.Println("byte size is ", unsafe.Sizeof(x))
	}
	{
		var x int32 = 0
		fmt.Println("int32 size is ", unsafe.Sizeof(x))
	}
	{
		var x int = 0
		fmt.Println("int size is ", unsafe.Sizeof(x))
	}
	{
		var x int64 = 0
		fmt.Println("int64 size is ", unsafe.Sizeof(x))
	}
	{
		var x float32 = 0
		fmt.Println("float32 size is ", unsafe.Sizeof(x))
	}
	{
		var x float64 = 0
		fmt.Println("float64 size is ", unsafe.Sizeof(x))
	}
	{
		t := time.Now()
		fmt.Println("time size is ", unsafe.Sizeof(t))
	}
}
