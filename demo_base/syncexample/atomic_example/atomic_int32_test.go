package atomic_example

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func Test_AtomicInt32(t *testing.T) {
	var x int32 = 0
	fmt.Println(atomic.AddInt32(&x, 1))
	fmt.Println(atomic.AddInt32(&x, -1))
}
