package atomic_example

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func Test_AtomicUint32(t *testing.T) {
	var x uint32 = 0
	fmt.Println(atomic.AddUint32(&x, 1))
}
