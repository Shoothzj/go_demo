package timer

import (
	"fmt"
	"testing"
	"time"
)

func TestAfterFunc(t *testing.T) {
	time.AfterFunc(time.Second, func() {
		fmt.Println("log")
	})
}
