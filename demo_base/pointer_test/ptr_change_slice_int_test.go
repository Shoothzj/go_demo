package pointer

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestChangeSliceVal(t *testing.T) {
	var intValue = [3]int{1, 9, 9}
	unsafePtr := unsafe.Pointer(&intValue[0])
	uintPtr := uintptr(unsafePtr) + 2*unsafe.Sizeof(intValue[0])
	*(*int)(unsafe.Pointer(uintPtr)) = 6
	fmt.Println(intValue)
}
