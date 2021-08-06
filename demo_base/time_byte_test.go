package demo_base

import (
	"fmt"
	"testing"
	"time"
	"unsafe"
)

func TestTimeActualBytes(t *testing.T) {
	now := time.Now()
	fmt.Println(unsafe.Sizeof(now))
}

func TestTimeToBytes(t *testing.T) {
	now := time.Now()
	bytes := []byte(now.Format("2006-01-02 15:04:05.000000000"))
	s := string(bytes)
	parse, err := time.Parse("2006-01-02 15:04:05.000000000", s)
	if err != nil {
		panic(err)
	}
	fmt.Println(parse)
}
