package pointer

import (
	"fmt"
	"testing"
)

func Test_Pointer(t *testing.T) {
	// p0指向int的零值
	p0 := new(int)
	// hex表达的地址
	fmt.Println(p0)
	// 0
	fmt.Println(*p0)

	// x是p0指向值的拷贝
	x := *p0
	// 都取得x的地址
	// x, *p1, *p2 的值相同
	p1, p2 := &x, &x
	// true
	fmt.Println(p1 == p2)
	// false
	fmt.Println(p0 == p1)
	// p3 和 p0 也存储相同的地址
	p3 := &*p0
	fmt.Println(p0 == p3)
	*p0, *p1 = 123, 789
	// 789 789 123
	fmt.Println(*p2, x, *p3)

	// int, int
	fmt.Printf("%T, %T \n", *p0, x)
	// *int, *int
	fmt.Printf("%T, %T \n", p0, p1)
}
